package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func VirtualNetworkGateways() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_virtual_network_gateways",
		Resolver:    fetchVirtualNetworkGateways,
		Description: "https://learn.microsoft.com/en-us/rest/api/network-gateway/virtual-network-gateways/list?tabs=HTTP#virtualnetworkgateway",
		Multiplex:   client.SubscriptionResourceGroupMultiplexRegisteredNamespace("azure_network_virtual_network_gateways", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.VirtualNetworkGateway{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualNetworkGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewVirtualNetworkGatewaysClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}

	pager := svc.NewListPager(cl.ResourceGroup, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
