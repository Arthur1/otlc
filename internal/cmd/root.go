package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type VersionInfo struct {
	version string
	commit  string
	date    string
}

var versionInfo VersionInfo

type Config struct {
	Endpoint string            `yaml:"endpoint"`
	Headers  map[string]string `yaml:"headers"`
}

var configFile string
var config Config

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "otlc",
		Short: "otlc: Command Line Tool of OpenTelemetry Protocol",
		Long: `"otlc" is a command line tool that allows you to easily post metrics by OTLP.
It acts as a simple exporter and helps you validate the OTLP endpoint.`,
	}

	cmd.PersistentFlags().StringVar(&configFile, "conf", "otlc.yaml", "config file path")
	cobra.OnInitialize(initConfig)

	cmd.Version = fmt.Sprintf("%s (rev %s)", versionInfo.version, versionInfo.commit)

	cmd.AddCommand(NewMetricsCmd())
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := NewRootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}
}

func SetVersionInfo(version, commit, date string) {
	versionInfo.version = version
	versionInfo.commit = commit
	versionInfo.date = date
}
