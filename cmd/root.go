package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	VERSION string
)

var rootCmd = &cobra.Command{
	Use:     "xds-conformance",
	Short:   "XDS Conformance test suite",
	Version: VERSION,
	Long: `XDS conformance test suite in Go. Complete documentation is available at https://github.com/envoyproxy/xds-conformance.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
