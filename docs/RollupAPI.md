# \RollupAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRollup**](RollupAPI.md#GetRollup) | **Get** /rollup/{id} | Get rollup info
[**GetRollupAllSeries**](RollupAPI.md#GetRollupAllSeries) | **Get** /rollup/stats/series | Get series for all rollups
[**GetRollupBlobs**](RollupAPI.md#GetRollupBlobs) | **Get** /rollup/{id}/blobs | Get rollup blobs
[**GetRollupBySlug**](RollupAPI.md#GetRollupBySlug) | **Get** /rollup/slug/{slug} | Get rollup by slug
[**GetRollupDistribution**](RollupAPI.md#GetRollupDistribution) | **Get** /rollup/{id}/distribution/{name}/{timeframe} | Get rollup distribution
[**GetRollupNamespaces**](RollupAPI.md#GetRollupNamespaces) | **Get** /rollup/{id}/namespaces | Get rollup namespaces info
[**GetRollupStats**](RollupAPI.md#GetRollupStats) | **Get** /rollup/{id}/stats/{name}/{timeframe} | Get rollup stats
[**GetRollupsCount**](RollupAPI.md#GetRollupsCount) | **Get** /rollup/count | Get count of rollups in network
[**ListRollup**](RollupAPI.md#ListRollup) | **Get** /rollup | List rollups info
[**ListRollup24h**](RollupAPI.md#ListRollup24h) | **Get** /rollup/day | List rollups info with stats by previous 24 hours
[**RollupExport**](RollupAPI.md#RollupExport) | **Get** /rollup/{id}/export | Export rollup blobs



## GetRollup

> ResponsesRollup GetRollup(ctx, id).Execute()

Get rollup info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {
	id := int32(56) // int32 | Internal identity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.GetRollup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.GetRollup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRollup`: ResponsesRollup
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.GetRollup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRollupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesRollup**](ResponsesRollup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRollupAllSeries

> []ResponsesRollupAllSeriesItem GetRollupAllSeries(ctx).Execute()

Get series for all rollups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.GetRollupAllSeries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.GetRollupAllSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRollupAllSeries`: []ResponsesRollupAllSeriesItem
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.GetRollupAllSeries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRollupAllSeriesRequest struct via the builder pattern


### Return type

[**[]ResponsesRollupAllSeriesItem**](ResponsesRollupAllSeriesItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRollupBlobs

> []ResponsesBlobLog GetRollupBlobs(ctx, id).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Joins(joins).Execute()

Get rollup blobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {
	id := int32(56) // int32 | Internal identity
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)
	sortBy := "sortBy_example" // string | Sort field. If it's empty internal id is used (optional)
	joins := true // bool | Flag indicating whether entities of transaction and signer should be attached or not. Default: true (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.GetRollupBlobs(context.Background(), id).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Joins(joins).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.GetRollupBlobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRollupBlobs`: []ResponsesBlobLog
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.GetRollupBlobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRollupBlobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **sortBy** | **string** | Sort field. If it&#39;s empty internal id is used | 
 **joins** | **bool** | Flag indicating whether entities of transaction and signer should be attached or not. Default: true | 

### Return type

[**[]ResponsesBlobLog**](ResponsesBlobLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRollupBySlug

> ResponsesRollup GetRollupBySlug(ctx, slug).Execute()

Get rollup by slug



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {
	slug := "slug_example" // string | Slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.GetRollupBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.GetRollupBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRollupBySlug`: ResponsesRollup
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.GetRollupBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRollupBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesRollup**](ResponsesRollup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRollupDistribution

> []ResponsesDistributionItem GetRollupDistribution(ctx, id, name, timeframe).Execute()

Get rollup distribution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {
	id := int32(56) // int32 | Internal identity
	name := "name_example" // string | Series name
	timeframe := "timeframe_example" // string | Timeframe

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.GetRollupDistribution(context.Background(), id, name, timeframe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.GetRollupDistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRollupDistribution`: []ResponsesDistributionItem
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.GetRollupDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 
**name** | **string** | Series name | 
**timeframe** | **string** | Timeframe | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRollupDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ResponsesDistributionItem**](ResponsesDistributionItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRollupNamespaces

> []ResponsesNamespace GetRollupNamespaces(ctx, id).Limit(limit).Offset(offset).Execute()

Get rollup namespaces info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {
	id := int32(56) // int32 | Internal identity
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.GetRollupNamespaces(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.GetRollupNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRollupNamespaces`: []ResponsesNamespace
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.GetRollupNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRollupNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesNamespace**](ResponsesNamespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRollupStats

> []ResponsesHistogramItem GetRollupStats(ctx, id, name, timeframe).From(from).To(to).Execute()

Get rollup stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {
	id := int32(56) // int32 | Internal identity
	name := "name_example" // string | Series name
	timeframe := "timeframe_example" // string | Timeframe
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.GetRollupStats(context.Background(), id, name, timeframe).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.GetRollupStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRollupStats`: []ResponsesHistogramItem
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.GetRollupStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 
**name** | **string** | Series name | 
**timeframe** | **string** | Timeframe | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRollupStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 

### Return type

[**[]ResponsesHistogramItem**](ResponsesHistogramItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRollupsCount

> int32 GetRollupsCount(ctx).Execute()

Get count of rollups in network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.GetRollupsCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.GetRollupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRollupsCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.GetRollupsCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRollupsCountRequest struct via the builder pattern


### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRollup

> []ResponsesRollupWithStats ListRollup(ctx).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()

List rollups info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)
	sortBy := "sortBy_example" // string | Sort field. Default: size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.ListRollup(context.Background()).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.ListRollup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRollup`: []ResponsesRollupWithStats
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.ListRollup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRollupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **sortBy** | **string** | Sort field. Default: size | 

### Return type

[**[]ResponsesRollupWithStats**](ResponsesRollupWithStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRollup24h

> []ResponsesRollupWithDayStats ListRollup24h(ctx).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()

List rollups info with stats by previous 24 hours



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)
	sortBy := "sortBy_example" // string | Sort field. Default: mb_price (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RollupAPI.ListRollup24h(context.Background()).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.ListRollup24h``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRollup24h`: []ResponsesRollupWithDayStats
	fmt.Fprintf(os.Stdout, "Response from `RollupAPI.ListRollup24h`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRollup24hRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **sortBy** | **string** | Sort field. Default: mb_price | 

### Return type

[**[]ResponsesRollupWithDayStats**](ResponsesRollupWithDayStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollupExport

> RollupExport(ctx, id).From(from).To(to).Execute()

Export rollup blobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/celenium-io/celenium-api-go"
)

func main() {
	id := int32(56) // int32 | Internal identity
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RollupAPI.RollupExport(context.Background(), id).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RollupAPI.RollupExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollupExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

