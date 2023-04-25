# go-devops
[![Go Report Card](https://goreportcard.com/badge/github.com/mo-silent/go-devops)](https://goreportcard.com/report/github.com/mo-silent/go-devops)
[![Go Reference](https://pkg.go.dev/badge/github.com/mo-silent/go-devops.svg)](https://pkg.go.dev/github.com/mo-silent/go-devops)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/mo-silent/go-devops)](https://pkg.go.dev/mod/github.com/mo-silent/go-devops)

This is the [Go](https://go.dev/) Devops tool library is used to encapsulate common tool methods

## Common

Some commonly used method encapsulation, such as http, ssh

## Logging & Monitoring
Encapsulated commonly used log monitoring queries for use in log monitoring tools. 

### Prometheus

The [prometheus directory](https://github.com/mo-silent/go-devops/tree/main/prometheus) includes the push metrics and range query methods of Prometheus.

The [examples prometheus directory](https://github.com/mo-silent/go-devops/tree/main/examples/prometheus) contains simple examples of instrumented code.

### Grafana

The [open-source directory](https://github.com/mo-silent/go-devops/tree/main/grafana/open-source) includes some querying methods of Open-source Grafana.

The [aliyun directory](https://github.com/mo-silent/go-devops/tree/main/grafana/aliyun) includes some querying methods of Aliyun Grafana.

### ~~Dynatrace~~

> disabled

The [dynatrace directory](https://github.com/mo-silent/go-devops/tree/main/dynatrace) contains commonly used query methods for Dynatrace.

## CI & CD
- Gitlab
- Jenkins

## Database
- MongoDB
- Mysql
- Redis
