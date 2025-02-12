/*
Celenium API

Testing StatsAPIService

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

func Test_celenium_StatsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StatsAPIService Stats24hChanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StatsAPI.Stats24hChanges(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsMessagesCount24h", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StatsAPI.StatsMessagesCount24h(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsNamespaceUsage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StatsAPI.StatsNamespaceUsage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsNsSeries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var timeframe string
		var name string

		resp, httpRes, err := apiClient.StatsAPI.StatsNsSeries(context.Background(), id, timeframe, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsPriceCurrent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StatsAPI.StatsPriceCurrent(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsPriceSeries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var timeframe string

		resp, httpRes, err := apiClient.StatsAPI.StatsPriceSeries(context.Background(), timeframe).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsRollup24h", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StatsAPI.StatsRollup24h(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsSeries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var timeframe string
		var name string

		resp, httpRes, err := apiClient.StatsAPI.StatsSeries(context.Background(), timeframe, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsSeriesCumulative", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var timeframe string
		var name string

		resp, httpRes, err := apiClient.StatsAPI.StatsSeriesCumulative(context.Background(), timeframe, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsSquareSize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StatsAPI.StatsSquareSize(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsStakingSeries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var timeframe string
		var name string

		resp, httpRes, err := apiClient.StatsAPI.StatsStakingSeries(context.Background(), id, timeframe, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatsAPIService StatsSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var table string
		var function string

		resp, httpRes, err := apiClient.StatsAPI.StatsSummary(context.Background(), table, function).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
