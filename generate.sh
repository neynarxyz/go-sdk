#!/bin/bash -eu -o pipefail

export GO_POST_PROCESS_FILE='gofmt -w'

rm -rf ./api
npx @openapitools/openapi-generator-cli generate --enable-post-process-file --git-user-id neynarxyz --git-repo-id go-sdk -i ./oas/src/v2/spec.yaml -g go -o ./api --config ./openapi-generator-config.api.json --openapi-normalizer 'SIMPLIFY_ONEOF_ANYOF=false'

rm -rf ./hub-api
npx @openapitools/openapi-generator-cli generate --enable-post-process-file --git-user-id neynarxyz --git-repo-id go-sdk -i ./oas/src/hub-rest-api/spec.yaml -g go -o ./hub-api --config ./openapi-generator-config.hub-api.json
