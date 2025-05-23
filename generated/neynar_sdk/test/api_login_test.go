/*
Farcaster API V2

Testing LoginAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package neynar_sdk

import (
	"context"
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_neynar_sdk_LoginAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LoginAPIService FetchNonce", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.LoginAPI.FetchNonce(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
