#!/usr/bin/env bash

function getVersionFromGitTag()
{
    BASE_PATH=$(git rev-parse --show-toplevel)
    cd $BASE_PATH
    latest_tag=$(git describe --abbrev=0)
    if git describe --exact-match >/dev/null 2>&1; then
        deb_version=${latest_tag}
    else
        # Increment tag and add a ~dev
        major="${latest_tag%.*}"
        patch="${latest_tag##*.}"
        commits_since=$(git rev-list --count HEAD "^${latest_tag}")
        deb_version=${major}.$((patch+1))~dev-${commits_since}
    fi
    echo $deb_version
}
