/*
Celenium API

Testing AddressAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package celenium

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func Test_celenium_AddressAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AddressAPIService AddressBlobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.AddressBlobs(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService AddressDelegations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.AddressDelegations(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService AddressGrantee", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.AddressGrantee(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService AddressGrants", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.AddressGrants(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService AddressMessages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.AddressMessages(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService AddressRedelegations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.AddressRedelegations(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService AddressStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string
		var name string
		var timeframe string

		resp, httpRes, err := apiClient.AddressAPI.AddressStats(context.Background(), hash, name, timeframe).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService AddressTransactions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.AddressTransactions(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService AddressUndelegations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.AddressUndelegations(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService AddressVesting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.AddressVesting(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService GetAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.AddressAPI.GetAddress(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService GetAddressCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AddressAPI.GetAddressCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressAPIService ListAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AddressAPI.ListAddress(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
