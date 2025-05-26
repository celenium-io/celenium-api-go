# \IbcAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIbcChannel**](IbcAPI.md#GetIbcChannel) | **Get** /ibc/channel/{id} | Get ibc channel info
[**GetIbcChannels**](IbcAPI.md#GetIbcChannels) | **Get** /ibc/channel | Get ibc channels info
[**GetIbcClient**](IbcAPI.md#GetIbcClient) | **Get** /ibc/client/{id} | Get ibc client info
[**GetIbcClients**](IbcAPI.md#GetIbcClients) | **Get** /ibc/client | Get ibc clients info
[**GetIbcConn**](IbcAPI.md#GetIbcConn) | **Get** /ibc/connection/{id} | Get ibc connection info
[**GetIbcConns**](IbcAPI.md#GetIbcConns) | **Get** /ibc/connection | Get ibc connections info
[**GetIbcTransfers**](IbcAPI.md#GetIbcTransfers) | **Get** /ibc/transfer | Get ibc transfers info



## GetIbcChannel

> ResponsesIbcChannel GetIbcChannel(ctx, id).Execute()

Get ibc channel info



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
	id := "id_example" // string | IBC channel id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IbcAPI.GetIbcChannel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IbcAPI.GetIbcChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIbcChannel`: ResponsesIbcChannel
	fmt.Fprintf(os.Stdout, "Response from `IbcAPI.GetIbcChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | IBC channel id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIbcChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesIbcChannel**](ResponsesIbcChannel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIbcChannels

> []ResponsesIbcChannel GetIbcChannels(ctx).Limit(limit).Offset(offset).Sort(sort).ClientId(clientId).ConnectionId(connectionId).Status(status).Execute()

Get ibc channels info



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
	clientId := "clientId_example" // string | Client id (optional)
	connectionId := "connectionId_example" // string | Connection id (optional)
	status := "status_example" // string | Channel status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IbcAPI.GetIbcChannels(context.Background()).Limit(limit).Offset(offset).Sort(sort).ClientId(clientId).ConnectionId(connectionId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IbcAPI.GetIbcChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIbcChannels`: []ResponsesIbcChannel
	fmt.Fprintf(os.Stdout, "Response from `IbcAPI.GetIbcChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIbcChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **clientId** | **string** | Client id | 
 **connectionId** | **string** | Connection id | 
 **status** | **string** | Channel status | 

### Return type

[**[]ResponsesIbcChannel**](ResponsesIbcChannel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIbcClient

> ResponsesIbcClient GetIbcClient(ctx, id).Execute()

Get ibc client info



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
	id := "id_example" // string | IBC client id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IbcAPI.GetIbcClient(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IbcAPI.GetIbcClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIbcClient`: ResponsesIbcClient
	fmt.Fprintf(os.Stdout, "Response from `IbcAPI.GetIbcClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | IBC client id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIbcClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesIbcClient**](ResponsesIbcClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIbcClients

> []ResponsesIbcClient GetIbcClients(ctx).Limit(limit).Offset(offset).Sort(sort).Execute()

Get ibc clients info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IbcAPI.GetIbcClients(context.Background()).Limit(limit).Offset(offset).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IbcAPI.GetIbcClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIbcClients`: []ResponsesIbcClient
	fmt.Fprintf(os.Stdout, "Response from `IbcAPI.GetIbcClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIbcClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 

### Return type

[**[]ResponsesIbcClient**](ResponsesIbcClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIbcConn

> ResponsesIbcConnection GetIbcConn(ctx, id).Execute()

Get ibc connection info



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
	id := "id_example" // string | IBC connection id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IbcAPI.GetIbcConn(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IbcAPI.GetIbcConn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIbcConn`: ResponsesIbcConnection
	fmt.Fprintf(os.Stdout, "Response from `IbcAPI.GetIbcConn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | IBC connection id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIbcConnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesIbcConnection**](ResponsesIbcConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIbcConns

> []ResponsesIbcConnection GetIbcConns(ctx).Limit(limit).Offset(offset).Sort(sort).ClientId(clientId).Execute()

Get ibc connections info



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
	clientId := "clientId_example" // string | Client id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IbcAPI.GetIbcConns(context.Background()).Limit(limit).Offset(offset).Sort(sort).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IbcAPI.GetIbcConns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIbcConns`: []ResponsesIbcConnection
	fmt.Fprintf(os.Stdout, "Response from `IbcAPI.GetIbcConns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIbcConnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **clientId** | **string** | Client id | 

### Return type

[**[]ResponsesIbcConnection**](ResponsesIbcConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIbcTransfers

> []ResponsesIbcTransfer GetIbcTransfers(ctx).Limit(limit).Offset(offset).Sort(sort).ClientId(clientId).ConnectionId(connectionId).Status(status).Execute()

Get ibc transfers info



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
	clientId := "clientId_example" // string | Client id (optional)
	connectionId := "connectionId_example" // string | Connection id (optional)
	status := "status_example" // string | Channel status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IbcAPI.GetIbcTransfers(context.Background()).Limit(limit).Offset(offset).Sort(sort).ClientId(clientId).ConnectionId(connectionId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IbcAPI.GetIbcTransfers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIbcTransfers`: []ResponsesIbcTransfer
	fmt.Fprintf(os.Stdout, "Response from `IbcAPI.GetIbcTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIbcTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **clientId** | **string** | Client id | 
 **connectionId** | **string** | Connection id | 
 **status** | **string** | Channel status | 

### Return type

[**[]ResponsesIbcTransfer**](ResponsesIbcTransfer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

