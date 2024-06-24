package cli

import (
	"github.com/alecthomas/kong"
)

type Globals struct {
	Version VersionFlag `name:"version" short:"v" help:"print version and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error {
	return nil
}
func (v VersionFlag) IsBool() bool {
	return true
}
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	printVersion()
	app.Exit(0)
	return nil
}

var cli struct {
	Globals
	Metrics MetricsCmd `cmd:"metrics" help:"command for OpenTelemetry Metrics"`
	Version VersionCmd `cmd:"version" help:"print version information"`
}

type Cli struct{}

func (c *Cli) Run() {
	kctx := kong.Parse(&cli,
		kong.Name("otlc"),
		kong.Description("otlc is a command line tool that allows you to easily post metrics by OTLP. It acts as a simple exporter and helps you testing for the OTLP endpoint."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
	)
	err := kctx.Run(&cli.Globals)
	kctx.FatalIfErrorf(err)
}
