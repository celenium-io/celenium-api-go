# \TransactionsAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransaction**](TransactionsAPI.md#GetTransaction) | **Get** /tx/{hash} | Get transaction by hash
[**GetTransactionEvents**](TransactionsAPI.md#GetTransactionEvents) | **Get** /tx/{hash}/events | Get transaction events
[**GetTransactionMessages**](TransactionsAPI.md#GetTransactionMessages) | **Get** /tx/{hash}/messages | Get transaction messages
[**GetTransactionsCount**](TransactionsAPI.md#GetTransactionsCount) | **Get** /tx/count | Get count of transactions in network
[**ListGenesisTransactions**](TransactionsAPI.md#ListGenesisTransactions) | **Get** /tx/genesis | List genesis transactions info
[**ListTransactionBlobs**](TransactionsAPI.md#ListTransactionBlobs) | **Get** /tx/{hash}/blobs | List blobs which was pushed by transaction
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /tx | List transactions info
[**TransactionBlobsCount**](TransactionsAPI.md#TransactionBlobsCount) | **Get** /tx/{hash}/blobs/count | Count of blobs which was pushed by transaction



## GetTransaction

> ResponsesTx GetTransaction(ctx, hash).Execute()

Get transaction by hash



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
	hash := "hash_example" // string | Transaction hash in hexadecimal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransaction(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransaction`: ResponsesTx
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Transaction hash in hexadecimal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesTx**](ResponsesTx.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionEvents

> []ResponsesEvent GetTransactionEvents(ctx, hash).Limit(limit).Offset(offset).Execute()

Get transaction events



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
	hash := "hash_example" // string | Transaction hash in hexadecimal
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionEvents(context.Background(), hash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionEvents`: []ResponsesEvent
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Transaction hash in hexadecimal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionEventsRequest struct via the builder pattern


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


## GetTransactionMessages

> []ResponsesMessage GetTransactionMessages(ctx, hash).Limit(limit).Offset(offset).Execute()

Get transaction messages



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
	hash := "hash_example" // string | Transaction hash in hexadecimal
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionMessages(context.Background(), hash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionMessages`: []ResponsesMessage
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Transaction hash in hexadecimal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

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


## GetTransactionsCount

> int32 GetTransactionsCount(ctx).Execute()

Get count of transactions in network



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
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsCountRequest struct via the builder pattern


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


## ListGenesisTransactions

> []ResponsesTx ListGenesisTransactions(ctx).Limit(limit).Offset(offset).Sort(sort).Execute()

List genesis transactions info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListGenesisTransactions(context.Background()).Limit(limit).Offset(offset).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListGenesisTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGenesisTransactions`: []ResponsesTx
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListGenesisTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGenesisTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order | 

### Return type

[**[]ResponsesTx**](ResponsesTx.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionBlobs

> []ResponsesBlobLog ListTransactionBlobs(ctx, hash).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()

List blobs which was pushed by transaction



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
	hash := "hash_example" // string | Transaction hash in hexadecimal
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)
	sortBy := "sortBy_example" // string | Sort field. If it's empty internal id is used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListTransactionBlobs(context.Background(), hash).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactionBlobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactionBlobs`: []ResponsesBlobLog
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactionBlobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Transaction hash in hexadecimal | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionBlobsRequest struct via the builder pattern


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


## ListTransactions

> []ResponsesTx ListTransactions(ctx).Limit(limit).Offset(offset).Sort(sort).Status(status).MsgType(msgType).ExcludedMsgType(excludedMsgType).From(from).To(to).Height(height).Messages(messages).Execute()

List transactions info



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
	status := "status_example" // string | Comma-separated status list (optional)
	msgType := "msgType_example" // string | Comma-separated message types list (optional)
	excludedMsgType := "excludedMsgType_example" // string | Comma-separated message types list which should be excluded (optional)
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)
	height := int32(56) // int32 | Block number (optional)
	messages := true // bool | If true join messages (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListTransactions(context.Background()).Limit(limit).Offset(offset).Sort(sort).Status(status).MsgType(msgType).ExcludedMsgType(excludedMsgType).From(from).To(to).Height(height).Messages(messages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactions`: []ResponsesTx
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order | 
 **status** | **string** | Comma-separated status list | 
 **msgType** | **string** | Comma-separated message types list | 
 **excludedMsgType** | **string** | Comma-separated message types list which should be excluded | 
 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 
 **height** | **int32** | Block number | 
 **messages** | **bool** | If true join messages | 

### Return type

[**[]ResponsesTx**](ResponsesTx.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionBlobsCount

> int32 TransactionBlobsCount(ctx, hash).Execute()

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
	hash := "hash_example" // string | Transaction hash in hexadecimal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.TransactionBlobsCount(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.TransactionBlobsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionBlobsCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.TransactionBlobsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Transaction hash in hexadecimal | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionBlobsCountRequest struct via the builder pattern


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

