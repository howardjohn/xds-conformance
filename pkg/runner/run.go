package runner

import (
	"context"
	"sort"
	"time"

	"github.com/envoyproxy/xds-conformance/pkg/conformance"
	"github.com/envoyproxy/xds-conformance/pkg/conformance/checks"
)

func RunChecks(labels conformance.LabelSelectors, input conformance.TestInput) []conformance.TestResult {
	checksToRun := checks.All
	results := make([]conformance.TestResult, 0, len(checksToRun))
	sort.Slice(checksToRun, func(i, j int) bool {
		return checksToRun[i].Name < checksToRun[j].Name
	})
	for _, check := range checksToRun {
		t0 := time.Now()
		result := check.Run(context.TODO(), input)
		if result.Duration == 0 {
			result.Duration = time.Since(t0)
		}
		result.Name = check.Name
		results = append(results, result)
	}
	return results
}
