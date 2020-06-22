package runner

import (
	"context"
	"fmt"
	"io"
	"sort"
	"testing"

	"github.com/fatih/color"

	"github.com/envoyproxy/xds-conformance/pkg/conformance"
	"github.com/envoyproxy/xds-conformance/pkg/conformance/checks"
)

func RunChecks(w io.Writer, labels conformance.LabelSelectors, input conformance.TestInput) bool {
	checksToRun := checks.All
	sort.Slice(checksToRun, func(i, j int) bool {
		return checksToRun[i].Name < checksToRun[j].Name
	})
	failed := false
	for _, check := range checksToRun {
		_, _ = fmt.Fprintf(w, "  Running: %v\n", check.Name)
		reporter := &PrintReporter{w, false}
		ctx, _ := context.WithTimeout(context.Background(), check.Timeout)
		check.Run(ctx, reporter, input)
		failed = failed || reporter.failed
	}
	return failed
}

func RunGoTest(t *testing.T, labels conformance.LabelSelectors, input conformance.TestInput) {
	checksToRun := checks.All
	sort.Slice(checksToRun, func(i, j int) bool {
		return checksToRun[i].Name < checksToRun[j].Name
	})
	for _, check := range checksToRun {
		t.Run(check.Name, func(t *testing.T) {
			reporter := GoTestReporter{t}
			ctx, _ := context.WithTimeout(context.Background(), check.Timeout)
			// TODO make reporter have Result() which aggregates everything. Or make it private?
			check.Run(ctx, reporter, input)
		})
	}
}

type PrintReporter struct {
	w io.Writer
	failed bool
}

var _ conformance.TestReporter = &PrintReporter{}

func (g *PrintReporter) ReportRunning(check conformance.Check) {
	_, _ = fmt.Fprintf(g.w, "  Running: %v", check.Name)
}

func (g *PrintReporter) Error(err error) {
	_, _ = fmt.Fprintf(g.w, "    %s %s\n", color.RedString("✘"), err.Error())
	g.failed = true
}

func (g *PrintReporter) Pass(msg string) {
	_, _ = fmt.Fprintf(g.w, "    %s %s\n", color.GreenString("✔"), msg)
}

func (g *PrintReporter) Info(msg string) {
	_, _ = fmt.Fprintf(g.w, "    %s %s\n", color.YellowString("-"), msg)
}

type GoTestReporter struct {
	t *testing.T
}

var _ conformance.TestReporter = &GoTestReporter{}

// Go test handles this for us
func (g GoTestReporter) ReportRunning(check conformance.Check) {}

func (g GoTestReporter) Error(err error) {
	g.t.Error(err)
}

func (g GoTestReporter) Pass(msg string) {
	g.t.Logf("PASS: %v", msg)
}

func (g GoTestReporter) Info(msg string) {
	g.t.Log(msg)
}
