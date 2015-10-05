# watcher
Multitool integration to provide system monitoring and alerting, ala EM7 or Nagios

InfluxDB (http://influxdb.org) will be the time-series database this project needs.

Stats on the machine(s) running the monitor will be self-monitored by Telegraf
(also from http://influxdb.com).

Graphing and dashboards will either depend on InfluxDB's Chrongraf tool, Grafana
(http://grafana.org), or something else (yet to be determined).

SNMP polling of external devices might be done by GoSNMP, depending on whether
it can adequately configured to support multiple datacenters and many OIDs per device.

Alerting might be handled by FlapJack, or might require a special tool to be written.

