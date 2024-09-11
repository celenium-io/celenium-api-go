# \BlockAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockBlobsCount**](BlockAPI.md#BlockBlobsCount) | **Get** /block/{height}/blobs/count | Count of blobs which was pushed by transaction
[**GetBlock**](BlockAPI.md#GetBlock) | **Get** /block/{height} | Get block info
[**GetBlockBlobs**](BlockAPI.md#GetBlockBlobs) | **Get** /block/{height}/blobs | List blobs which was pushed in the block
[**GetBlockCount**](BlockAPI.md#GetBlockCount) | **Get** /block/count | Get count of blocks in network
[**GetBlockEvents**](BlockAPI.md#GetBlockEvents) | **Get** /block/{height}/events | Get events from begin and end of block
[**GetBlockMessages**](BlockAPI.md#GetBlockMessages) | **Get** /block/{height}/messages | Get messages contained in the block
[**GetBlockStats**](BlockAPI.md#GetBlockStats) | **Get** /block/{height}/stats | Get block stats by height
[**ListBlock**](BlockAPI.md#ListBlock) | **Get** /block | List blocks info



## BlockBlobsCount

> int32 BlockBlobsCount(ctx, height).Execute()

Count of blobs which was pushed by transaction



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
	height := int32(56) // int32 | Block height

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.BlockBlobsCount(context.Background(), height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.BlockBlobsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockBlobsCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.BlockBlobsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Block height | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockBlobsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetBlock

> ResponsesBlock GetBlock(ctx, height).Stats(stats).Execute()

Get block info



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
	height := int32(56) // int32 | Block height
	stats := true // bool | Need join stats for block (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.GetBlock(context.Background(), height).Stats(stats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.GetBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlock`: ResponsesBlock
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.GetBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Block height | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stats** | **bool** | Need join stats for block | 

### Return type

[**ResponsesBlock**](ResponsesBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockBlobs

> []ResponsesBlobLog GetBlockBlobs(ctx, height).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()

List blobs which was pushed in the block



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
	height := int32(56) // int32 | Block height
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)
	sortBy := "sortBy_example" // string | Sort field. If it's empty internal id is used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.GetBlockBlobs(context.Background(), height).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.GetBlockBlobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockBlobs`: []ResponsesBlobLog
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.GetBlockBlobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Block height | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockBlobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **sortBy** | **string** | Sort field. If it&#39;s empty internal id is used | 

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


## GetBlockCount

> int32 GetBlockCount(ctx).Execute()

Get count of blocks in network



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
	resp, r, err := apiClient.BlockAPI.GetBlockCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.GetBlockCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.GetBlockCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockCountRequest struct via the builder pattern


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


## GetBlockEvents

> []ResponsesEvent GetBlockEvents(ctx, height).Limit(limit).Offset(offset).Execute()

Get events from begin and end of block



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
	height := int32(56) // int32 | Block height
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.GetBlockEvents(context.Background(), height).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.GetBlockEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockEvents`: []ResponsesEvent
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.GetBlockEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Block height | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesEvent**](ResponsesEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockMessages

> []ResponsesMessage GetBlockMessages(ctx, height).Limit(limit).Offset(offset).MsgType(msgType).ExcludedMsgType(excludedMsgType).Execute()

Get messages contained in the block



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
	height := int32(56) // int32 | Block height
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	msgType := "msgType_example" // string | Comma-separated message types list (optional)
	excludedMsgType := "excludedMsgType_example" // string | Comma-separated message types which should be excluded from list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.GetBlockMessages(context.Background(), height).Limit(limit).Offset(offset).MsgType(msgType).ExcludedMsgType(excludedMsgType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.GetBlockMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockMessages`: []ResponsesMessage
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.GetBlockMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Block height | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **msgType** | **string** | Comma-separated message types list | 
 **excludedMsgType** | **string** | Comma-separated message types which should be excluded from list | 

### Return type

[**[]ResponsesMessage**](ResponsesMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockStats

> ResponsesBlockStats GetBlockStats(ctx, height).Execute()

Get block stats by height



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
	height := int32(56) // int32 | Block height

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.GetBlockStats(context.Background(), height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.GetBlockStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockStats`: ResponsesBlockStats
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.GetBlockStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Block height | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesBlockStats**](ResponsesBlockStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlock

> []ResponsesBlock ListBlock(ctx).Limit(limit).Offset(offset).Sort(sort).Stats(stats).Execute()

List blocks info



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
	sort := "sort_example" // string | Sort order (optional)
	stats := true // bool | Need join stats for block (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.ListBlock(context.Background()).Limit(limit).Offset(offset).Sort(sort).Stats(stats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.ListBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlock`: []ResponsesBlock
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.ListBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order | 
 **stats** | **bool** | Need join stats for block | 

### Return type

[**[]ResponsesBlock**](ResponsesBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

