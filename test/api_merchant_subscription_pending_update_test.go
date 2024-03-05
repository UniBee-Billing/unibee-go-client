/*
OpenAPI UniBee

Testing MerchantSubscriptionPendingUpdateAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_MerchantSubscriptionPendingUpdateAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MerchantSubscriptionPendingUpdateAPIService MerchantSubscriptionPendingUpdateListGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MerchantSubscriptionPendingUpdateAPI.MerchantSubscriptionPendingUpdateListGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MerchantSubscriptionPendingUpdateAPIService MerchantSubscriptionPendingUpdateListPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MerchantSubscriptionPendingUpdateAPI.MerchantSubscriptionPendingUpdateListPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
