package cli

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"

	"github.com/Arthur1/otlc"
)

type VersionCmd struct{}

func (c *VersionCmd) Run(globals *Globals) error {
	return printVersion()
}

func printVersion() error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	if _, err := fmt.Fprintf(writer, "otlc is a Command-line Tool to Export Telemetry by OTLP (OpenTelemetry Protocol).\n"); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(writer, "Version:\t%s\n", otlc.Version); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(writer, "Go version:\t%s\n", runtime.Version()); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(writer, "Arch:\t%s\n", runtime.GOARCH); err != nil {
		return err
	}
	return writer.Flush()
}
