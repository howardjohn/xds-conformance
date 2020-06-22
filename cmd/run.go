package cmd

import (
	"fmt"
	"os"

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
		if failed := runner.RunChecks(cmd.OutOrStdout(), nil, testArgs); failed {
				os.Exit(1)
		}
	},
}

func validateInput() error {
	return nil
}
