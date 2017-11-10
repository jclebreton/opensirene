# opensirene [![Build Status](https://travis-ci.org/jclebreton/opensirene.svg?branch=v2)](https://travis-ci.org/jclebreton/opensirene) [![codecov](https://codecov.io/gh/jclebreton/opensirene/branch/master/graph/badge.svg)](https://codecov.io/gh/jclebreton/opensirene)
French company database based on French government open data

## Getting Started

### Build
```
$ dep ensure
$ go run main.go
```

## Configuration

This project can be configured using both a yaml configuration file and
environment variables (for most of the configuration fields). Environment
variables have a higher priority than the configuration file, which means you
can override almost any value of the configuration file using them. 

yml
---
```
logger:
  level: debug
  format: text

server:
  host: 127.0.0.1
  port: 8080
  debug: false
  cors:
    permissive_mode: true
    enabled: true

database:
  user: xx
  password: xx
  name: opensirene
  host: 127.0.0.1
  port: 5432
  sslmode: disable

prometheus:
  prefix: opensirene

crontab:
  download_path: "downloads"
  every_x_hours: 3

```


| Field                       | Type     | Description                                               | Environment Variable | Default        | Example        |
|-----------------------------|----------|-----------------------------------------------------------|----------------------|----------------|----------------|
| logger.level                | string   | Global log level                                          | `LOGLEVEL`           | "info"         | "debug"        |
| logger.format               | string   | Log format (text, json)                                   | `LOGFORMAT`          | "text"         | "json"         |
| server.host                 | string   | Host on which the server will listen                      | `SERVER_HOST`        | "127.0.0.1"    | "127.0.0.1"    |
| server.port                 | int      | Port on which the server will listen                      | `SERVER_PORT`        | 8080           | 8080           |
| server.debug                | bool     | Debug mode                                                | `SERVER_DEBUG`       | false          | true           |
| server.cors.allow_origins   | []string | Array of accepted origins                                 | -                    | -              | -              |
| server.cors.permissive_mode | bool     | Accept every origin and overrides the allow_origins field | `CORS_PERMISSIVE`    | false          | true           |
| database.user               | string   | User used to connect to the DB                            | `DB_USER`            | "sir"          | "sir"          |
| database.password           | string   | Password associated to the user                           | `DB_PASSWORD`        | -              | -              |
| database.host               | string   | Host on which the DB listens                              | `DB_HOST`            | "127.0.0.1"    | "127.0.0.1"    |
| database.port               | int      | Port on which the DB listens                              | `DB_PORT`            | 5432           | 5432           |
| database.name               | string   | Database name to use                                      | `DB_NAME`            | "opensirenedb" | "opensirenedb" |
| database.sslmode            | string   | Use the SSL mode                                          | `DB_SSL_MODE`        | "disable"      | "disable"      |
| prometheus.prefix           | string   | Prefix the prometheus metrics                             | `PROMETHEUS_PREFIX`  | "opensirene"   | "opensirene"   |
| crontab.download_path       | string   | Downloads path                                            | `DOWNLOAD_PATH`      | "downloads"    | "/tmp"         |
| crontab.every_x_hours       | uint64   | Crontab interval (in hours)                               | `EVERY_X_HOURS`      | 3              | 1              |
