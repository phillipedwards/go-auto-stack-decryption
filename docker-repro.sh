#!/bin/bash

set -euox pipefail

# cleanup the old directory
sudo rm -rf ./.pulumi

# shellcheck disable=SC2034
SCRIPT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

docker run \
  --entrypoint /bin/sh \
  --rm  \
  --volume "$PWD":/workspace \
  -w /workspace \
  --tty \
   pulumi/pulumi:3.57.0 \
   -c "go build -buildvcs=false && go run main.go"