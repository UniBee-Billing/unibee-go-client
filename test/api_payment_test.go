/*
OpenAPI UniBee

Testing PaymentService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package unibee

import (
	"context"
	"fmt"
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_unibee_PaymentService(t *testing.T) {
	openapiclient.ApiKey = "EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM"
	openapiclient.Host = "http://api.unibee.top"
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PaymentService PaymentCancelPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Payment.PaymentCancelPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentService PaymentCapturePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Payment.PaymentCapturePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentService PaymentDetailGet", func(t *testing.T) {

		resp, httpRes, err := apiClient.Payment.PaymentDetailGet(context.Background()).PaymentId("pay20240312RushNwJLr4rlk4P").Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentService PaymentListGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Payment.PaymentListGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentService PaymentMethodListGet", func(t *testing.T) {

		resp, httpRes, err := apiClient.Payment.PaymentMethodListGet(context.Background()).GatewayId(29).PaymentId("pay20240312RushNwJLr4rlk4P").Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentService PaymentNewPost", func(t *testing.T) {

		resp, httpRes, err := apiClient.Payment.PaymentNewPost(context.Background()).UnibeeApiMerchantPaymentNewReq(openapiclient.UnibeeApiMerchantPaymentNewReq{
			CountryCode:       openapiclient.String("CH"),
			TotalAmount:       openapiclient.Int64(100),
			Currency:          openapiclient.String("USDT"),
			Email:             openapiclient.String("jack.fu@wowow.io"),
			ExternalPaymentId: openapiclient.String(uuid.New().String()), // your paymentId
			ExternalUserId:    openapiclient.String("1709272139"),
			GatewayId:         29,
			Items:             nil, // without items
			Metadata:          &map[string]string{"key1": "value1", "key2": "value2"},
			RedirectUrl:       openapiclient.String("http://user.unibee.top/paymentResult"),
			GasPayer:          openapiclient.String("user"),
		}).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		fmt.Printf("crypto payment url:%s\n", *resp.Data.Link)

	})

	t.Run("Test PaymentService PaymentRefundCancelPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Payment.PaymentRefundCancelPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentService PaymentRefundDetailGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Payment.PaymentRefundDetailGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentService PaymentRefundListGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Payment.PaymentRefundListGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentService PaymentRefundNewPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Payment.PaymentRefundNewPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentService PaymentTimelineListGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Payment.PaymentTimelineListGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
