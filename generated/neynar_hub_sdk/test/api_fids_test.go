/*
Farcaster Hub API

Testing FidsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package neynar_hub_sdk

import (
	"context"
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_hub_sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_neynar_hub_sdk_FidsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FidsAPIService FetchFids", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FidsAPI.FetchFids(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
