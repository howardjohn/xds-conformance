package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/envoyproxy/xds-conformance/pkg/conformance/checks"
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringSliceVarP(&rawLabelSelector, "selector", "l", rawLabelSelector, "Label selector for tests to use")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List XDS conformance tests",
	Long:  "List XDS conformance tests",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO filter by labels
		checksToList := checks.All
		for _, check := range checksToList {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "- %s (%s)\n", check.Description, check.Name)
		}
	},
}
