package checks

import (
	"context"
	"fmt"
	"time"

	"github.com/envoyproxy/xds-conformance/pkg/conformance"
	"github.com/envoyproxy/xds-conformance/pkg/label"
	"github.com/envoyproxy/xds-conformance/pkg/xds"
)

var ConnectionCheck = RegisterCheck(conformance.Check{
	Name:        "Server Connection",
	Description: "Clients must be able to connect to the server under test.",
	Labels:      []label.Instance{label.Server, label.XdsV3},
	Timeout:     time.Second * 5,
	Run: func(ctx context.Context, runner conformance.TestReporter, input conformance.TestInput) {
		if input.Address == "" {
			runner.Error(fmt.Errorf("invalid address provided %q", input.Address))
			return
		}
		c, err := xds.ConnectAds(ctx, input.Address)
		if err != nil {
			runner.Error(connectionFailure(input.Address))
			return
		}
		defer c.Cleanup()
		runner.Pass(fmt.Sprintf("Connection established to %q.", input.Address))
	},
})
