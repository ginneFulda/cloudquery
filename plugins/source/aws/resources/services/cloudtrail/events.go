package cloudtrail

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/mitchellh/hashstructure/v2"
)

const tableName = "aws_cloudtrail_events"

func Events() *schema.Table {
	return &schema.Table{
		Name:          tableName,
		Description:   `https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_Event.html`,
		Resolver:      fetchCloudtrailEvents,
		Multiplex:     client.ServiceAccountRegionMultiplexer(tableName, "cloudtrail"),
		Transform:     transformers.TransformWithStruct(&types.Event{}, transformers.WithPrimaryKeys("EventId")),
		IsIncremental: true,
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "cloud_trail_event",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudTrailEvent"),
			},
			{
				Name:     "event_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EventTime"),
				CreationOptions: schema.ColumnCreationOptions{
					IncrementalKey: true,
				},
			},
		},
	}
}

func fetchCloudtrailEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudtrail

	allConfigs := []tableoptions.CustomLookupEventsOpts{{}}
	noTableConfig := true
	if cl.Spec.TableOptions.CloudTrailEvents != nil {
		allConfigs = cl.Spec.TableOptions.CloudTrailEvents.LookupEventsOpts
		noTableConfig = false
	}
	for _, w := range allConfigs {
		le := w.LookupEventsInput

		var backendKey string
		if cl.Backend != nil {
			// Retrieve the last event time from the backend for this table option config.
			// We use a hash of the config as the key, so changing the config will cause a full refresh.
			hash, err := hashstructure.Hash(le, hashstructure.FormatV2, nil)
			if err != nil {
				return err
			}
			backendKey = fmt.Sprintf("%s-%d", cl.ID(), hash)
			if noTableConfig {
				// for backwards-compatibility, default to client id if there is no table config
				backendKey = cl.ID()
			}
			value, err := cl.Backend.Get(ctx, tableName, backendKey)
			if err != nil {
				return fmt.Errorf("failed to retrieve state from backend: %w", err)
			}

			if value != "" {
				date, err := time.Parse(time.RFC3339Nano, value)
				if err != nil {
					return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
				}
				le.StartTime = &date
			}
		}

		var lastEventTime *time.Time
		// var err error
		paginator := cloudtrail.NewLookupEventsPaginator(svc, &le)
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx, func(options *cloudtrail.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- page.Events

			// Retrieve the timestamp from the latest event
			for _, event := range page.Events {
				if lastEventTime == nil {
					lastEventTime = event.EventTime
					continue
				}
				if event.EventTime.After(*lastEventTime) {
					lastEventTime = event.EventTime
				}
			}
		}

		if cl.Backend != nil && lastEventTime != nil {
			err := cl.Backend.Set(ctx, tableName, backendKey, lastEventTime.Format(time.RFC3339Nano))
			if err != nil {
				return fmt.Errorf("failed to save state to backend: %w", err)
			}
		}
	}

	return nil
}
