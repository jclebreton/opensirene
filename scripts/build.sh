#!/usr/bin/env bash

source ./scripts/version.sh
deb_version=$(getVersionFromGitTag)
go build -ldflags "-X main.version=${deb_version}"
