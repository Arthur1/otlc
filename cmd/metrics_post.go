package cmd

import (
	"log"

	"github.com/Arthur1/otlc/metrics"
	"github.com/spf13/cobra"
)

func NewMetricsPostCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post",
		Short: "Post a metric",
		Long:  `Post a metric using OTLP. Currently, only gauge is supported.`,
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				log.Fatalln(err)
			}
			value, err := cmd.Flags().GetFloat64("value")
			if err != nil {
				log.Fatalln(err)
			}
			description, err := cmd.Flags().GetString("description")
			if err != nil {
				log.Fatalln(err)
			}
			attributes, err := cmd.Flags().GetStringToString("attributes")
			if err != nil {
				log.Fatalln(err)
			}

			p := metrics.NewPoster(config.Endpoint, config.Headers)
			if err := p.Post(&metrics.PostParams{
				Name:           name,
				Description:    description,
				DataPointAttrs: attributes,
				DataPointValue: value,
			}); err != nil {
				log.Fatalln(err)
			}
		},
	}

	cmd.Flags().Float64P("value", "v", 0, "metric value")
	cmd.Flags().StringP("type", "t", "gauge", "metric value type")
	cmd.Flags().StringP("name", "n", "", "metric name")
	cmd.Flags().StringP("description", "d", "", "metric description")
	cmd.Flags().StringToStringP("attributes", "a", nil, "metric datapoint attributes. format: key1=value1,key2=value2")
	cmd.MarkFlagRequired("value")
	cmd.MarkFlagRequired("name")

	return cmd
}
