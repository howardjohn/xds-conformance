package checks

import (
	"context"
	"fmt"
	"time"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/golang/protobuf/proto"

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

var AdsCheck = RegisterCheck(conformance.Check{
	Name:        "ADS",
	Description: "Can connect over ADS and get a valid response",
	Labels:      []label.Instance{label.Server, label.XdsV3},
	Timeout:     time.Second * 5,
	Run: func(ctx context.Context, runner conformance.TestReporter, input conformance.TestInput) {
		c, err := xds.ConnectAds(ctx, input.Address)
		if err != nil {
			runner.Error(connectionFailure(input.Address))
			return
		}
		defer c.Cleanup()

		if err := c.Send(&discovery.DiscoveryRequest{
			Node:    constructNode(),
			TypeUrl: resource.ClusterType,
		}); err != nil {
			runner.Error(requestFailure(err))
			return
		}
		resp, err := c.Recv()
		if err != nil {
			runner.Error(responseFailure(err))
			return
		}
		if resp.TypeUrl != resource.ClusterType {
			runner.Error(fmt.Errorf("expected type URL %q, got %q", resource.ClusterType, resp.TypeUrl))
		} else {
			runner.Pass(fmt.Sprintf("Response has the correct TypeUrl: %s", resource.ClusterType))
		}
		for i, resource := range resp.Resources {
			cl := &cluster.Cluster{}
			// TODO is this safe? With Any we may not know all types. It would be nice to call .Validate() though..
			if err := proto.Unmarshal(resource.Value, cl); err != nil {
				runner.Error(fmt.Errorf("failed to unmarshal resource %d", i))
				continue
			}
			if err := cl.Validate(); err != nil {
				runner.Error(fmt.Errorf("failed to validate resource %d: %v", i, err))
				continue
			}
		}
		runner.Pass(fmt.Sprintf("Recieved %d valid clusters", len(resp.Resources)))
	},
})
