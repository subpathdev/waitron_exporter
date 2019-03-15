# Prometheus metric for Waitron

This is a tool, which can be used to export waitron metrics to [Prometheus](https://prometheus.io/)

## Dependencies
- a running waitron instance

## Command line parameters
| command | description |
| ------- | --------- |
| help | print the help |
| listen  | is the TCP network address |
| waitron | is the url to the waitron server |

## Metrics
| name | Description |
| ---- | ----- |
| waitron\_health | output of \<waitron\_server\>/health WHERE 1 is OK|
| waitron\_node\_state{node=\<node\_name\>} | output of \<waitron\_server\>/status/\<node\_name\> WHERE 1 is installing |

## Installation
To install this exporter you have to build this with the following command:
```
go build -o waitron\_exporter main.go
```
To use the service file you have change the user and place it in /etc/systemd/system/.

## Configuration example prometheus
This show an example of the prometheus configuration of this exporter:
```
 - job_name: waitron
    scrape_interval: 5s
    metrics_path: /
    static_configs:
    - targets: ['localhost:9999']
```
