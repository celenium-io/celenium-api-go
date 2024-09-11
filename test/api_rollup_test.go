/*
Celenium API

Testing RollupAPIService

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

func Test_celenium_RollupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RollupAPIService GetRollup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.RollupAPI.GetRollup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService GetRollupAllSeries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RollupAPI.GetRollupAllSeries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService GetRollupBlobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.RollupAPI.GetRollupBlobs(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService GetRollupBySlug", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var slug string

		resp, httpRes, err := apiClient.RollupAPI.GetRollupBySlug(context.Background(), slug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService GetRollupDistribution", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32
		var name string
		var timeframe string

		resp, httpRes, err := apiClient.RollupAPI.GetRollupDistribution(context.Background(), id, name, timeframe).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService GetRollupNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.RollupAPI.GetRollupNamespaces(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService GetRollupStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32
		var name string
		var timeframe string

		resp, httpRes, err := apiClient.RollupAPI.GetRollupStats(context.Background(), id, name, timeframe).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService GetRollupsCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RollupAPI.GetRollupsCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService ListRollup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RollupAPI.ListRollup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService ListRollup24h", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RollupAPI.ListRollup24h(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RollupAPIService RollupExport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.RollupAPI.RollupExport(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
