#!/usr/bin/env bash

source ./scripts/version.sh
deb_version=$(getVersionFromGitTag)
build_date=$(date --rfc-3339=seconds)
go build -ldflags "-X 'main.version=${deb_version}' -X 'main.buildDate=${build_date}'"
