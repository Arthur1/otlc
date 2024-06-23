package cmd

import (
	"github.com/spf13/cobra"
)

func NewMetricsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "metrics",
		Short: "Command for OpenTelemetry Metrics",
	}
	cmd.AddCommand(NewMetricsPostCmd())

	return cmd
}
