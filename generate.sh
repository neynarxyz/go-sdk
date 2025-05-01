#!/bin/bash
set -eux -o pipefail
export GO_POST_PROCESS_FILE='gofmt -w'

rm -rf ./api
npx --yes @openapitools/openapi-generator-cli generate --enable-post-process-file --git-user-id neynarxyz --git-repo-id go-sdk -i ./oas/src/v2/spec.yaml -g go -o ./api --config ./openapi-generator-config.api.json --openapi-normalizer 'SIMPLIFY_ONEOF_ANYOF=false'
(
    cd ./api
    go get github.com/stretchr/testify/assert
    go get golang.org/x/net/context
    go build
    go test ./...
)

rm -rf ./hub-api
npx --yes @openapitools/openapi-generator-cli generate --enable-post-process-file --git-user-id neynarxyz --git-repo-id go-sdk -i ./oas/src/hub-rest-api/spec.yaml -g go -o ./hub-api --config ./openapi-generator-config.hub-api.json
(
    cd ./hub-api
    go get github.com/stretchr/testify/assert
    go get golang.org/x/net/context
    go build
    go test ./...
)