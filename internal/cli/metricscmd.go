package cli

type MetricsCmd struct {
	Post MetricsPostCmd `cmd:"post" help:"post a metric datapoint"`
}
