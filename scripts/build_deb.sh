#!/bin/bash

BASE_PATH=$(git rev-parse --show-toplevel)

cd $BASE_PATH

latest_tag=$(git describe --abbrev=0)
commits_since=$(git rev-list --count HEAD "^${latest_tag}")

if git describe --exact-match >/dev/null 2>&1; then
    deb_version=${latest_tag}
else
    # Increment tag and add a ~dev
    major="${latest_tag%.*}"
    patch="${latest_tag##*.}"
    deb_version=${major}.$((patch+1))~dev-${commits_since}
fi

go build

docker build -t fpm -f fpm.Dockerfile .
docker run -ti -v $PWD:/packaging fpm fpm --verbose -s dir -t deb -n opensirene -v $deb_version \
  --description "French company database based on French government open data" \
  opensirene=/usr/bin/opensirene \
  conf-example.yml=/usr/share/opensirene/conf-example.yml \
  systemd.service=/lib/systemd/system/opensirene.service

sha256sum *.deb > SHA256SUMS
md5sum *.deb > MD5SUMS
