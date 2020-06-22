package conformance

import (
	"context"
	"time"

	"github.com/envoyproxy/xds-conformance/pkg/label"
)

type TestResult struct {
	Name string

	// Error represents a test failure. If this is present, the test is considered failed
	Error error

	// Skipped records if a test was skipped
	Skipped bool

	// Information provides additional information about a test result. This may include warnings, comments, etc
	Information string

	// Duration records the time it took to complete the tests
	Duration time.Duration
}

// LabelSelector defines a label selection for tests. This is used to filter out supported tests. For example,
// a user developing an XDS server may filter to Match={"server"} to run only server tests.
type LabelSelector struct {
	Match    []string
	NotMatch []string
}

// LabelSelectors is a set of labels, which are "OR"ed together.
type LabelSelectors []LabelSelector

func (l LabelSelectors) Matches(test []label.Instance) bool {
	// TODO implement match
	return true
}

type TestInput struct {
	Address string
}

type Check struct {
	Name        string
	Description string
	// Timeout for the test
	Timeout time.Duration
	Labels  []label.Instance
	Run     func(ctx context.Context, input TestInput) TestResult
}
