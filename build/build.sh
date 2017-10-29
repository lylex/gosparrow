#!/bin/bash

set -o errexit
set -o nounset

if [ -z "${VERSION}" ]; then
    echo "VERSION must be set"
    exit 1
fi

HOME=$(dirname "${BASH_SOURCE[0]}")/..
cd $HOME/cmd/gosparrow

go install                                     \
    -installsuffix "static"
