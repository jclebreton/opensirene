# OpenSirene [![Build Status](https://travis-ci.org/jclebreton/opensirene.svg?branch=v2)](https://travis-ci.org/jclebreton/opensirene) [![codecov](https://codecov.io/gh/jclebreton/opensirene/branch/master/graph/badge.svg)](https://codecov.io/gh/jclebreton/opensirene)

OpenSirene  is a REST micro-service to build and access a database of French companies
provided by the [French government's open data](https://www.data.gouv.fr/fr/datasets/base-sirene-des-entreprises-et-de-leurs-etablissements-siren-siret/).

This micro service will run a crontab to downloads automatically update files
and save them to a `PostgreSQL` database.

## API contract 

API endpoints are defined using a [Swagger file](swagger.yaml) at the root of
the Git repository.

* Search endpoints:
    * GET /api/v1/siret/{siret_id}
        * Retrieve one company from its SIRET identifier
    * GET /api/v1/siren/{siren_id}
        * Retrieve the list of establishments of a company from its SIREN identifier
* Monitoring endpoints:
    * GET /admin/ping
        * For monitoring purpose
    * GET /admin/history
        * Retrieve the list of update files stored in database


## Setup
The micro service needs 10Gb of free space to manipulates update files and
the database server approximately the same to store them.

### With Docker Compose

For development only, it will start two containers: the database (PostgresSQL)
and the micro-service.
```
$ docker-compose up
$ curl localhost:8080/ping
```

### With Debian package (for production environment)

```sh
# Download and install
$ LATEST_VERSION="1.2.1"
$ wget https://github.com/jclebreton/opensirene/releases/download/${LATEST_VERSION}/opensirene_${LATEST_VERSION}_amd64.deb
$ wget https://github.com/jclebreton/opensirene/releases/download/${LATEST_VERSION}/SHA256SUMS
$ sha256sum -c SHA256SUMS
$ dpkg -i opensirene_${LATEST_VERSION}_amd64.deb

# Create the configuration file and edit it with your own configuration
$ mkdir /etc/opensirene/
$ cp /usr/share/opensirene/conf-example.yml /etc/opensirene/conf.yml
$ vi /etc/opensirene/conf.yml

# Start service
$ systemctl start opensirene
$ systemctl daemon-reload # if needed

# Check if the service is up
$ systemctl status opensirene
$ curl 127.0.0.1:8080/ping
```

To redirect logs to syslog, edit `/lib/systemd/system/opensirene.service` and add those lines to the end:
```bash
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=opensirene
```

### Build and run

```sh
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure
$ go run main.go --config=conf-example.yml
```

### Configuration

This project can be configured using both a yaml configuration file and
environment variables (for most of the configuration fields). Environment
variables have a higher priority than the configuration file, which means you
can override almost any value of the configuration file using them. 

#### Example
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
  prefix:
    api: "/api/v1"
    admin: "/admin"

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

#### Description

| Field                       | Type     | Description                                               | Environment Variable | Default        | Example        |
|-----------------------------|----------|-----------------------------------------------------------|----------------------|----------------|----------------|
| logger.level                | string   | Global log level                                          | `LOGLEVEL`           | "info"         | "debug"        |
| logger.format               | string   | Log format (text, json)                                   | `LOGFORMAT`          | "text"         | "json"         |
| server.host                 | string   | Host on which the server will listen                      | `SERVER_HOST`        | "127.0.0.1"    | "127.0.0.1"    |
| server.port                 | int      | Port on which the server will listen                      | `SERVER_PORT`        | 8080           | 8080           |
| server.debug                | bool     | Debug mode                                                | `SERVER_DEBUG`       | false          | true           |
| server.cors.allow_origins   | []string | Array of accepted origins                                 | -                    | -              | -              |
| server.cors.permissive_mode | bool     | Accept every origin and overrides the allow_origins field | `CORS_PERMISSIVE`    | false          | true           |
| server.prefix.api           | string   | API prefix URL for clients                                | -                    | "/api/v1/"     | "/api/v1/"     |
| server.prefix.admin         | string   | API prefix URL for monitoring purpose                     | -                    | "/admin/"      | "/admin/"      |
| database.user               | string   | User used to connect to the DB                            | `DB_USER`            | "sir"          | "sir"          |
| database.password           | string   | Password associated to the user                           | `DB_PASSWORD`        | -              | -              |
| database.host               | string   | Host on which the DB listens                              | `DB_HOST`            | "127.0.0.1"    | "127.0.0.1"    |
| database.port               | int      | Port on which the DB listens                              | `DB_PORT`            | 5432           | 5432           |
| database.name               | string   | Database name to use                                      | `DB_NAME`            | "opensirenedb" | "opensirenedb" |
| database.sslmode            | string   | Use the SSL mode                                          | `DB_SSL_MODE`        | "disable"      | "disable"      |
| prometheus.prefix           | string   | Prefix the prometheus metrics                             | `PROMETHEUS_PREFIX`  | "opensirene"   | "opensirene"   |
| crontab.download_path       | string   | Downloads path                                            | `DOWNLOAD_PATH`      | "downloads"    | "/tmp"         |
| crontab.every_x_hours       | uint64   | Crontab interval (in hours)                               | `EVERY_X_HOURS`      | 3              | 1              |
