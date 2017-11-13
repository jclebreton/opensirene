#!/usr/bin/env bash

# Version
source ./scripts/version.sh
deb_version=$(getVersionFromGitTag)

# Build binary
./scripts/build.sh

# Build Debian package
docker build -t fpm -f fpm.Dockerfile .
docker run -ti -v $PWD:/packaging fpm fpm --verbose -s dir -t deb -n opensirene -v $deb_version \
  --description "French company database based on French government open data" \
  opensirene=/usr/bin/opensirene \
  conf-example.yml=/usr/share/opensirene/conf-example.yml \
  systemd.service=/lib/systemd/system/opensirene.service

# Create shasums files
sha256sum *.deb > SHA256SUMS
md5sum *.deb > MD5SUMS
