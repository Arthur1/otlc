# otlc

"otlc" is a command line tool that allows you to easily post metrics by OTLP. It acts as a simple exporter and helps you testing for the OTLP endpoint.

## Install

Currently we support macOS and Linux. You may be able to run otlc on Windows by using docker or go install.

### Homebrew

```sh
brew install Arthur1/tap/otlc
```

### Docker

- [`ghcr.io/arthur1/otlc`](https://github.com/Arthur1/otlc/pkgs/container/otlc)

### go install

```sh
go install github.com/Arthur1/otlc/cmd/otlc@latest
```

## Run

### post metrics

```console
$ export OTEL_EXPORTER_OTLP_ENDPOINT="otlp.mackerelio.com:4317"
$ export OTEL_EXPORTER_OTLP_HEADERS="Mackerel-Api-Key=***your_api_key***"
$ otlc metrics post --name awesome_gauge --attrs hoge=poyo,fuga=1 123.45
exported.
```

```
Usage: otlc metrics post --otlp-endpoint=STRING --name=STRING <data-point-value> [flags]

post a metric datapoint

Arguments:
  <data-point-value>    datapoint value

Flags:
  -h, --help                             Show context-sensitive help.
  -v, --version                          print version and quit

      --otlp-endpoint=STRING             OTLP endpoint ($OTEL_EXPORTER_OTLP_ENDPOINT,
                                         $OTEL_EXPORTER_OTLP_METRICS_ENDPOINT)
      --otlp-headers=KEY=VALUE,...       OTLP headers ($OTEL_EXPORTER_OTLP_HEADERS, $OTEL_EXPORTER_OTLP_METRICS_HEADERS)
      --otlp-protocol="grpc"             OTLP protocol ($OTEL_EXPORTER_OTLP_PROTOCOL)
      --otlp-insecure                    disable secure connection (required for such as localhost)
  -n, --name=STRING                      metric name
  -t, --type="gauge"                     metric type
  -d, --description=STRING               metric description
  -u, --unit="1"                         metric unit
      --resource-attrs=KEY=VALUE,...     resource attributes
      --datapoint-attrs=KEY=VALUE,...    datapoint attributes
      --timestamp=INT-64                 datapoint timestamp (unix seconds)
```
