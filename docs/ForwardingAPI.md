# \ForwardingAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetForwarding**](ForwardingAPI.md#GetForwarding) | **Get** /forwarding/{id} | Get forwarding event by ID
[**ListForwarding**](ForwardingAPI.md#ListForwarding) | **Get** /forwarding | List forwarding events



## GetForwarding

> ResponsesForwarding GetForwarding(ctx, id).Execute()

Get forwarding event by ID



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
	id := int32(56) // int32 | Internal forwarding event ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForwardingAPI.GetForwarding(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForwardingAPI.GetForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForwarding`: ResponsesForwarding
	fmt.Fprintf(os.Stdout, "Response from `ForwardingAPI.GetForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal forwarding event ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesForwarding**](ResponsesForwarding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListForwarding

> []ResponsesForwarding ListForwarding(ctx).Limit(limit).Offset(offset).Sort(sort).TxHash(txHash).Address(address).Height(height).From(from).To(to).Execute()

List forwarding events



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
	offset := int32(56) // int32 | Offset for pagination (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)
	txHash := "txHash_example" // string | Filter by transaction hash (hex) (optional)
	address := "address_example" // string | Filter by Celestia address (optional)
	height := int32(56) // int32 | Filter by block height (optional)
	from := int32(56) // int32 | Filter by start time (Unix timestamp) (optional)
	to := int32(56) // int32 | Filter by end time (Unix timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForwardingAPI.ListForwarding(context.Background()).Limit(limit).Offset(offset).Sort(sort).TxHash(txHash).Address(address).Height(height).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForwardingAPI.ListForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListForwarding`: []ResponsesForwarding
	fmt.Fprintf(os.Stdout, "Response from `ForwardingAPI.ListForwarding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset for pagination | 
 **sort** | **string** | Sort order. Default: desc | 
 **txHash** | **string** | Filter by transaction hash (hex) | 
 **address** | **string** | Filter by Celestia address | 
 **height** | **int32** | Filter by block height | 
 **from** | **int32** | Filter by start time (Unix timestamp) | 
 **to** | **int32** | Filter by end time (Unix timestamp) | 

### Return type

[**[]ResponsesForwarding**](ResponsesForwarding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

