package networkfirewall

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/networkfirewall/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func FirewallPolicies() *schema.Table {
	tableName := "aws_networkfirewall_firewall_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_FirewallPolicy.html`,
		Resolver:            fetchFirewallPolicies,
		PreResourceResolver: getFirewallPolicy,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "network-firewall"),
		Transform: transformers.TransformWithStruct(
			&models.FirewallPolicyWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FirewallPolicyArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchFirewallPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input networkfirewall.ListFirewallPoliciesInput
	c := meta.(*client.Client)
	svc := c.Services().Networkfirewall
	p := networkfirewall.NewListFirewallPoliciesPaginator(svc, &input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *networkfirewall.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.FirewallPolicies
	}
	return nil
}

func getFirewallPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Networkfirewall
	metadata := resource.Item.(types.FirewallPolicyMetadata)

	policy, err := svc.DescribeFirewallPolicy(ctx, &networkfirewall.DescribeFirewallPolicyInput{
		FirewallPolicyArn: metadata.Arn,
	}, func(options *networkfirewall.Options) {
		options.Region = c.Region
	})
	if err != nil && !c.IsNotFoundError(err) {
		return err
	}

	resource.Item = &models.FirewallPolicyWrapper{
		FirewallPolicy:         policy.FirewallPolicy,
		FirewallPolicyResponse: policy.FirewallPolicyResponse,
	}
	return nil
}
