# upsmonitor-prometheus

Metrics Exporter for CyberPower [PowerPannel](https://www.cyberpowersystems.com/products/software/power-panel-personal/#platform-section)

Get's status from `pwrstat -status`

See `example/metrics` for some example metrics.

Runs metrics at `http://localhost:9090/metrics`

Systemd service can be found at `example/upsmonitor.service`