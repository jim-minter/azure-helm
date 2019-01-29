#!/bin/bash -ex

if [[ $1 -ne "" ]]; then
    echo usage: $0 resourcegroup
    exit 1
fi

if [ -z "$RESOURCEGROUP" ]
then
    export RESOURCEGROUP=$1
fi

rm -rf _data
mkdir -p _data/_out

set -x

if [[ -n "$TEST_IN_PRODUCTION" ]]; then
    TEST_IN_PRODUCTION="-use-prod=true"
else
    go generate ./...
    go run cmd/fakerp/main.go &
fi
if [[ -n "$ADMIN_MANIFEST" ]]; then
    ADMIN_MANIFEST="-admin-manifest=$ADMIN_MANIFEST"
fi

trap 'return_id=$?; set +ex; kill $(lsof -t -i :8080); wait $(lsof -t -i :8080); exit $return_id' EXIT

# default location for ci-secrets
if [[ -f /usr/local/e2e-secrets/azure/secret ]] ;then
    source /usr/local/e2e-secrets/azure/secret
fi

go run cmd/createorupdate/createorupdate.go ${TEST_IN_PRODUCTION:-} ${ADMIN_MANIFEST:-}
