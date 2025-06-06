#!/bin/bash
set -eux -o pipefail

# This script generates the Neynar Go SDK from the OpenAPI specification.

export GO_POST_PROCESS_FILE='gofmt -w'

# Delete all directories except test
cd ./generated/neynar_sdk
rm -rf ./*.* api/ docs/
cd ../..
npx --yes @openapitools/openapi-generator-cli \
    generate \
    --git-user-id neynarxyz \
    --git-repo-id go-sdk/generated \
    -g go \
    --additional-properties=packageVersion="$SDK_VERSION" \
    --config ./src/api/openapi-generator-config.json \
    -i ./src/OAS/src/api/spec.yaml \
    -o ./generated/neynar_sdk \
    --openapi-normalizer 'SIMPLIFY_ONEOF_ANYOF=false' \
    --inline-schema-options 'SKIP_SCHEMA_REUSE=true' \
    --ignore-file-override=./.openapi-generator-ignore \
    --enable-post-process-file
(
    cd ./generated/neynar_sdk
    go get github.com/stretchr/testify/assert
    go get golang.org/x/net/context
    go fmt
    go mod tidy
    go build
    # go test ./...
)

# Delete all directories except test
cd ./generated/neynar_hub_sdk
rm -rf ./*.* api/ docs/
cd ../..
npx --yes @openapitools/openapi-generator-cli \
    generate \
    --git-user-id neynarxyz \
    --git-repo-id go-sdk/generated \
    -g go \
    --additional-properties=packageVersion="$SDK_VERSION" \
    --config ./src/hub-api/openapi-generator-config.json \
    -i ./src/OAS/src/hub-api/spec.yaml \
    -o ./generated/neynar_hub_sdk \
    --openapi-normalizer 'SIMPLIFY_ONEOF_ANYOF=false' \
    --inline-schema-options 'SKIP_SCHEMA_REUSE=true' \
    --ignore-file-override=./.openapi-generator-ignore \
    --type-mappings='file=[]byte' \
    --enable-post-process-file
(
    cd ./generated/neynar_hub_sdk
    go get github.com/stretchr/testify/assert
    go get golang.org/x/net/context
    go fmt
    go mod tidy
    go build
    # go test ./...
)