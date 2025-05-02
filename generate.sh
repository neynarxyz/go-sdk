#!/bin/bash
set -eux -o pipefail
export GO_POST_PROCESS_FILE='gofmt -w'

rm -rf ./api
npx --yes @openapitools/openapi-generator-cli \
    generate \
    --enable-post-process-file \
    --git-user-id neynarxyz \
    --git-repo-id go-sdk \
    -g go \
    -i ./oas/src/v2/spec.yaml \
    -o ./api \
    --config ./openapi-generator-config.api.json \
    --openapi-normalizer 'SIMPLIFY_ONEOF_ANYOF=false' \
    --inline-schema-options 'SKIP_SCHEMA_REUSE=true'
(
    cd ./api
    go get github.com/stretchr/testify/assert
    go get golang.org/x/net/context
    go build
    go test ./...
)

rm -rf ./hub-api
npx --yes @openapitools/openapi-generator-cli \
    generate \
    --enable-post-process-file \
    --git-user-id neynarxyz \
    --git-repo-id go-sdk \
    -g go \
    -i ./oas/src/hub-rest-api/spec.yaml \
    -o ./hub-api \
    --config ./openapi-generator-config.hub-api.json
(
    cd ./hub-api
    go get github.com/stretchr/testify/assert
    go get golang.org/x/net/context
    go build
    go test ./...
)