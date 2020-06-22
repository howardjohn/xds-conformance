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
	Run: func(ctx context.Context, input conformance.TestInput) conformance.TestResult {
		if input.Address == "" {
			return conformance.TestResult{Error: fmt.Errorf("invalid address provided %q", input.Address)}
		}
		c, err := xds.ConnectAds(ctx, input.Address)
		if err != nil {
			return conformance.TestResult{
				Error:       fmt.Errorf("failed to conntect to %q", input.Address),
				Information: "Verify the XDS server is running and is reachable at the provided address. If TLS is required by the server, ensure TLS configuration is set.",
			}
		}
		defer c.Cleanup()
		return conformance.TestResult{
			Information: fmt.Sprintf("Connection established to %q.", input.Address),
		}
	},
})

