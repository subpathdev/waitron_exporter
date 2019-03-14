# Prometheus metric for Waitron

This is a tool, which can be used to export waitron metrics to [Prometheus](https://prometheus.io/)

## Dependencies
- an running waitron instance

## Command line parameters
| command | parameter |
| ------- | --------- |
| listen  | is the TCP network address |
| waitron | is the url to the waitron server |

## Metrics
| name | value | Description |
| ---- | ----- | ----- |
| waitron\_node | string | name of the nodes |
| waitron\_health | string | output of \<waitron\_server\>/health |
| waitron\_node\_state{node=\<node\_name\>} | string | output of \<waitron\_server\>/status/\<node\_name\> |

## Installation
To install this exporter you have to build this with the following command:
```
go build -o waitron\_exporter main.go
```
To use the service file you have change the user and place it in /etc/systemd/system/.
