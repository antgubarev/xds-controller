#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/.." &> /dev/null && pwd )

bash vendor/k8s.io/code-generator/generate-groups.sh \
	"deepcopy,client,informer,lister" \
	github.com/antgubarev/xds-controller/internal/generated \
	github.com/antgubarev/xds-controller/internal/apis \
	proxy.company.com:v1 \
	--output-base=$SCRIPT_ROOT/../../.. \
	--go-header-file=$SCRIPT_ROOT/hack/boilerplate.go.txt

