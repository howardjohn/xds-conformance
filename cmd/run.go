package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/envoyproxy/xds-conformance/pkg/conformance"
	"github.com/envoyproxy/xds-conformance/pkg/runner"
)

var testArgs conformance.TestInput
var rawLabelSelector []string

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVar(&testArgs.Address, "server.address", testArgs.Address, "Address of the XDS server under test")
	runCmd.Flags().StringSliceVarP(&rawLabelSelector, "selector", "l", rawLabelSelector, "Label selector for tests to use")

}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run XDS conformance tests",
	Long:  "Run XDS conformance tests",
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateInput(); err != nil {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), "invalid input: %v\n", err)
			os.Exit(2)
		}
		results := runner.RunChecks(nil, testArgs)
		displayResults(results, cmd.OutOrStdout())
		if countErrors(results) > 0 {
			os.Exit(1)
		}
	},
}

func validateInput() error {
	return nil
}

func infoString(i string) string {
	if i == "" {
		return ""
	}
	return ". " + i
}

func displayResults(results []conformance.TestResult, stdout io.Writer) {
	print := func(format string, args ...interface{}) {
		_, _ = fmt.Fprintf(stdout, format+"\n", args...)
	}
	print("Processed %d tests. %d skipped, %d passed, %d failed", len(results), countSkipped(results), countSuccess(results), countErrors(results))
	for _, result := range results {
		if result.Skipped {
			print("  %s %s skipped%s", color.YellowString("-"), result.Name, infoString(result.Information))
			continue
		}
		if result.Error != nil {

			print("  %s %s failed in %v: %s%s", color.RedString("✘"), result.Name,
				result.Duration.Truncate(time.Millisecond), result.Error, infoString(result.Information))
			continue
		}

		print("  %s %s passed in %v%s", color.GreenString("✔"), result.Name,
			result.Duration.Truncate(time.Millisecond), infoString(result.Information))
		continue
	}
}

func countErrors(results []conformance.TestResult) int {
	res := 0
	for _, r := range results {
		if r.Error != nil && !r.Skipped {
			res++
		}
	}
	return res
}

func countSuccess(results []conformance.TestResult) int {
	res := 0
	for _, r := range results {
		if r.Error == nil && !r.Skipped {
			res++
		}
	}
	return res
}

func countSkipped(results []conformance.TestResult) int {
	res := 0
	for _, r := range results {
		if r.Skipped {
			res++
		}
	}
	return res
}
