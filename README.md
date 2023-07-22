# otlc

"otlc" is a command line tool that allows you to easily post metrics by OTLP. It acts as a simple exporter and helps you validate the OTLP endpoint.

## installation

WIP

## run

### post metrics

```sh
otlc metrics post --conf ./otlc.yaml --name awesome_gauge --value 123.45 \
--resource-attrs service.name=otlc --datapoint-attrs hoge=poyo,fuga=1
```
