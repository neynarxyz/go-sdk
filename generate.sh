#!/bin/bash
set -eux -o pipefail

export GO_POST_PROCESS_FILE='gofmt -w'

rm -rf ./generated/api/*
npx --yes @openapitools/openapi-generator-cli \
    generate \
    --git-user-id neynarxyz \
    --git-repo-id go-sdk/generated \
    -g go \
    --config ./src/api/openapi-generator-config.json \
    -i ./src/OAS/src/v2/spec.yaml \
    -o ./generated/api \
    --openapi-normalizer 'SIMPLIFY_ONEOF_ANYOF=false' \
    --inline-schema-options 'SKIP_SCHEMA_REUSE=true' \
    --enable-post-process-file
(
    cd ./generated/api
    go get github.com/stretchr/testify/assert
    go get golang.org/x/net/context
    go build
    go test ./...
)

rm -rf ./generated/hub-api/*
npx --yes @openapitools/openapi-generator-cli \
    generate \
    --git-user-id neynarxyz \
    --git-repo-id go-sdk/generated \
    -g go \
    --config ./src/hub-api/openapi-generator-config.json \
    -i ./src/OAS/src/hub-rest-api/spec.yaml \
    -o ./generated/hub-api \
    --openapi-normalizer 'SIMPLIFY_ONEOF_ANYOF=false' \
    --inline-schema-options 'SKIP_SCHEMA_REUSE=true' \
    --type-mappings='file=[]byte' \
    --enable-post-process-file
(
    cd ./generated/hub-api
    go get github.com/stretchr/testify/assert
    go get golang.org/x/net/context
    go build
    go test ./...
)