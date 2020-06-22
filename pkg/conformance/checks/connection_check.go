package checks

import (
	"context"
	"fmt"

	"github.com/envoyproxy/xds-conformance/pkg/conformance"
	"github.com/envoyproxy/xds-conformance/pkg/label"
)

var ConnectionCheck = RegisterCheck(conformance.Check{
	Name:        "Connection",
	Description: "Clients must be able to connect to the server under test.",
	Labels:      []label.Instance{label.Server},
	Run: func(ctx context.Context, input conformance.TestInput) conformance.TestResult {
		return conformance.TestResult{}
	},
})

// Always fails
var FailCheck = RegisterCheck(conformance.Check{
	Name:        "Fail Check",
	Description: "For testing only, should be removed before release",
	Labels:      []label.Instance{label.Server},
	Run: func(ctx context.Context, input conformance.TestInput) conformance.TestResult {
		return conformance.TestResult{
			Error: fmt.Errorf("this test always fails"),
		}
	},
})

// Always skipped
var SkipCheck = RegisterCheck(conformance.Check{
	Name:        "Skip Check",
	Description: "For testing only, should be removed before release",
	Labels:      []label.Instance{label.Server},
	Run: func(ctx context.Context, input conformance.TestInput) conformance.TestResult {
		return conformance.TestResult{Skipped: true, Information: "Skipped because this test is always skipped."}
	},
})
