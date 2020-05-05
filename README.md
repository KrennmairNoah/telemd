go-telemetry
============

A symmetry telemetry library written in go

Build
-----

To build the `telemd` binary using your local go installation run:

    go build -o bin/telemd cmd/telemd/main.go

To build Docker images for local usage (without a go installation) run:

    scripts/docker-build.sh

Usage
-----

### Topic schema

telemd reports telemetry data into (Redis) topics.

The topic schema is as follows. It starts with the keyword `telem`, followed by the hostname that reports the telemetry,
the specific metric being reported, and optionally the subsystem (e.g., a specific network device or disk).

    telem/<hostname>/<metric>[/<subsystem>]

For example, the CPU utilization of CPU core 0 host `rpi0` could be reported as:

    telem/rpi0/cpu/0

Or it may report an aggregate value into

    telem/rpi0/cpu

### Talking back to clients

Clients listen on the topic

    telemcmd/<hostname>

for commands. Currently, telemd supports the following commands:

* `pause` pauses reporting of metrics
* `unpause` unpauses report of metrics

### Telemetry Daemon Parameters

The `telemd` command allows the following parameters via environment variables.

| Variable | Default | Description |
|---|---|---|
| `telemd_node_name`    | `$HOST`       | The node name determines the value for `<hostname>` in the topics |
| `telemd_redis_host`   | `localhost`   | The redis host to connect to |
| `telemd_redis_port`   | `6379`        | The redis port to connect to |
| `telemd_redis_url`    |               | Can be used to specify the redis URL (e.g., `redis://localhost:1234`). Overwrites anything set to `telemd_redis_host`.
| `telemd_net_devices`  | all           | A list of network devices to be monitored, e.g. `wlan0 eth0`. Monitors all devices per default |
| `telemd_disk_devices` | all           | A list of block devices to be monitored, e.g. `sda sdc sdd0`. Monitors all devices per default |
