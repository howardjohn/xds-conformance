package checks

import (
	"context"
	"fmt"
	"time"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"

	"github.com/envoyproxy/xds-conformance/pkg/conformance"
	"github.com/envoyproxy/xds-conformance/pkg/label"
	"github.com/envoyproxy/xds-conformance/pkg/xds"
)

// TODO make this a pass in option
func constructNode() *core.Node {
	return &core.Node{
		Id: "sidecar~1.1.1.1~id~domain",
	}
}

var AdsClusterCheck = RegisterCheck(conformance.Check{
	Name:        "ADS Clusters",
	Description: "Can connect over ADS and get a cluster valid response",
	Labels:      []label.Instance{label.Server, label.XdsV3},
	Timeout:     time.Second * 5,
	Run: func(ctx context.Context, runner conformance.TestReporter, input conformance.TestInput) {
		checkAdsForType(ctx, runner, input, resource.ClusterType)
	},
})

var AdsListenersCheck = RegisterCheck(conformance.Check{
	Name:        "ADS Listeners",
	Description: "Can connect over ADS and get a listener valid response",
	Labels:      []label.Instance{label.Server, label.XdsV3},
	Timeout:     time.Second * 5,
	Run: func(ctx context.Context, runner conformance.TestReporter, input conformance.TestInput) {
		checkAdsForType(ctx, runner, input, resource.ListenerType)
	},
})

// TODO pass resource names
var AdsRoutesCheck = RegisterCheck(conformance.Check{
	Name:        "ADS Routes",
	Description: "Can connect over ADS and get a route valid response",
	Labels:      []label.Instance{label.Server, label.XdsV3},
	Timeout:     time.Second * 5,
	Run: func(ctx context.Context, runner conformance.TestReporter, input conformance.TestInput) {
		checkAdsForType(ctx, runner, input, resource.RouteType)
	},
})

var AdsEndpointsCheck = RegisterCheck(conformance.Check{
	Name:        "ADS Endpoints",
	Description: "Can connect over ADS and get a endpoint valid response",
	Labels:      []label.Instance{label.Server, label.XdsV3},
	Timeout:     time.Second * 5,
	Run: func(ctx context.Context, runner conformance.TestReporter, input conformance.TestInput) {
		checkAdsForType(ctx, runner, input, resource.EndpointType)
	},
})

var AdsSecretCheck = RegisterCheck(conformance.Check{
	Name:        "ADS Secrets",
	Description: "Can connect over ADS and get a secret valid response",
	Labels:      []label.Instance{label.Server, label.XdsV3},
	Timeout:     time.Second * 5,
	Run: func(ctx context.Context, runner conformance.TestReporter, input conformance.TestInput) {
		checkAdsForType(ctx, runner, input, resource.SecretType)
	},
})

func checkAdsForType(ctx context.Context, runner conformance.TestReporter, input conformance.TestInput, typeUrl string) {
	c, err := xds.ConnectAds(ctx, input.Address)
	if err != nil {
		runner.Error(connectionFailure(input.Address))
		return
	}
	defer c.Cleanup()

	if err := c.Send(&discovery.DiscoveryRequest{
		Node:    constructNode(),
		TypeUrl: typeUrl,
	}); err != nil {
		runner.Error(requestFailure(err))
		return
	}
	resp, err := c.Recv()
	if err != nil {
		runner.Error(responseFailure(err))
		return
	}
	if resp.TypeUrl != typeUrl {
		runner.Error(fmt.Errorf("expected type URL %q, got %q", typeUrl, resp.TypeUrl))
	} else {
		runner.Pass(fmt.Sprintf("Response has the correct TypeUrl: %s", typeUrl))
	}
	if resp.Nonce == "" {
		runner.Error(fmt.Errorf("expected a nonce in discovery response, found none"))
	}
	if resp.VersionInfo == "" {
		runner.Error(fmt.Errorf("expected version info in discovery response, found none"))
	}
	runner.Pass(fmt.Sprintf("Recieved %d valid clusters", len(resp.Resources)))
}
