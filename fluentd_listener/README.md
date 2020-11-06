# fluentd-listener

This middleware puts the log records that were received from Fluentd into Inflion Event. This is designed for use along with Fluentd's http output plugin.

## Installation

To install this middleware, check out this repository.

## Usage

Just run `go run` like below:

```sh
go run main.go --project $YOUR_PROJECT
```

### Options

* `--project` specifies the project name. Required.
* `--endpoint` is the address of Inflion gRPC endpoint. (Default: `127.0.0.1:50051`)
* `--listen` is the address this middleware listens to HTTP requests. (Default `0.0.0.0:8000`)

## Fluentd Settings

To use this middleware along with Fluentd, you have to make Fluentd forward logs to the HTTP endpoint of this middleware.

### Minimal settings

The simplest setting is below, all logs will be sent to Inflion:

```sh
# my_fluentd.conf
<match nginx.access>
  @type http
  endpoint http://127.0.0.1:8000/
  <buffer>
    flush_interval 1s
  </buffer>
</match>
```

The `endpoint` parameter is the URL of this middleware. The `flush_interval` parameter in the `<buffer>` directive specifies how often Fluentd sends logs. We recommend this value as small as possible so that Inflion can acknowledge problems from real-time logs as soon as possible.

### Filtering logs

To filter the logs to send, use Fluentd's Filter Plugin. For example, if you would like to send only POST request logs from all the access logs, add a `<filter>` directive to the settings like below:

```sh
# my_fluentd.conf
<filter nginx.access>
 @type grep
  <regexp>
    key method
    pattern /^POST$/
  </regexp>
</filter>

<match nginx.access>
  @type http
  endpoint http://127.0.0.1:8000/
  <buffer>
    flush_interval 1s
  </buffer>
</match>
```

To know the full usage of the Filter Plugin, see [the Fluentd's documentation](https://docs.fluentd.org/filter).

## Build

```sh
go build \
    -ldflags "-X main.name=fluentd-listener" \
    -ldflags "-X main.version=1.0.0" \
    -o fluentd-listener \
    main.go
```
