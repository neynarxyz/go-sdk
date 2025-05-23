/*
Farcaster Hub API

Testing MessageAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package neynar_hub_sdk

import (
	"context"
	"encoding/hex"
	"testing"

	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_hub_sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_hub_MessageAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MessageAPIService PublishMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MessageAPI.PublishMessage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessageAPIService ValidateMessage", func(t *testing.T) {
		s := "0a2c080310b88d1f1888cbf64020013a1d0801121908b26c1214f72f82fc3bc5f15d89b25edbe1faa2f4859ee2ac12147eb48982d4578a113b91a0573f406bc25baccb26180122406fc13bc380c78c8636de35939a35f0ef9b5be7754b952dd1ee0a2738f729f01f2ed53bcb398760db4277822a58a3afd71afa05820561c3b163e55cb014d15b0f2801322039ee42d1f77ced51a364a0a9925e7dd6ade2daf029fcb87a0e31f8061567e222"
		data, err := hex.DecodeString(s)
		if err != nil {
			panic(err)
		}
		// Add x-api-key header
		configuration.AddDefaultHeader("x-api-key", "NEYNAR_API_DOCS")
		resp, httpRes, err := apiClient.MessageAPI.ValidateMessage(context.Background()).Body(data).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}
