# \HyperlaneAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHyperlaneMailbox**](HyperlaneAPI.md#GetHyperlaneMailbox) | **Get** /hyperlane/mailbox/{id} | Get hyperlane mailbox info
[**GetHyperlaneToken**](HyperlaneAPI.md#GetHyperlaneToken) | **Get** /hyperlane/token/{id} | Get hyperlane token info
[**GetHyperlaneTransfer**](HyperlaneAPI.md#GetHyperlaneTransfer) | **Get** /hyperlane/transfer/{id} | Get transfer by id
[**ListHyperlaneDomains**](HyperlaneAPI.md#ListHyperlaneDomains) | **Get** /hyperlane/domains | List hyperlane domains info
[**ListHyperlaneMailbox**](HyperlaneAPI.md#ListHyperlaneMailbox) | **Get** /hyperlane/mailbox | List hyperlane mailboxes info
[**ListHyperlaneTokens**](HyperlaneAPI.md#ListHyperlaneTokens) | **Get** /hyperlane/token | List hyperlane tokens info
[**ListHyperlaneTransfers**](HyperlaneAPI.md#ListHyperlaneTransfers) | **Get** /hyperlane/transfer | List hyperlane transfers info



## GetHyperlaneMailbox

> ResponsesHyperlaneMailbox GetHyperlaneMailbox(ctx, id).Execute()

Get hyperlane mailbox info



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
	id := "id_example" // string | Hyperlane mailbox id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HyperlaneAPI.GetHyperlaneMailbox(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HyperlaneAPI.GetHyperlaneMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHyperlaneMailbox`: ResponsesHyperlaneMailbox
	fmt.Fprintf(os.Stdout, "Response from `HyperlaneAPI.GetHyperlaneMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Hyperlane mailbox id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHyperlaneMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesHyperlaneMailbox**](ResponsesHyperlaneMailbox.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHyperlaneToken

> ResponsesHyperlaneToken GetHyperlaneToken(ctx, id).Execute()

Get hyperlane token info



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
	id := "id_example" // string | Hyperlane token id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HyperlaneAPI.GetHyperlaneToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HyperlaneAPI.GetHyperlaneToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHyperlaneToken`: ResponsesHyperlaneToken
	fmt.Fprintf(os.Stdout, "Response from `HyperlaneAPI.GetHyperlaneToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Hyperlane token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHyperlaneTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesHyperlaneToken**](ResponsesHyperlaneToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHyperlaneTransfer

> ResponsesHyperlaneTransfer GetHyperlaneTransfer(ctx, id).Execute()

Get transfer by id



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
	resp, r, err := apiClient.HyperlaneAPI.GetHyperlaneTransfer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HyperlaneAPI.GetHyperlaneTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHyperlaneTransfer`: ResponsesHyperlaneTransfer
	fmt.Fprintf(os.Stdout, "Response from `HyperlaneAPI.GetHyperlaneTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHyperlaneTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesHyperlaneTransfer**](ResponsesHyperlaneTransfer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHyperlaneDomains

> []ResponsesDomainMetadata ListHyperlaneDomains(ctx).Execute()

List hyperlane domains info



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
	resp, r, err := apiClient.HyperlaneAPI.ListHyperlaneDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HyperlaneAPI.ListHyperlaneDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHyperlaneDomains`: []ResponsesDomainMetadata
	fmt.Fprintf(os.Stdout, "Response from `HyperlaneAPI.ListHyperlaneDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHyperlaneDomainsRequest struct via the builder pattern


### Return type

[**[]ResponsesDomainMetadata**](ResponsesDomainMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHyperlaneMailbox

> []ResponsesHyperlaneMailbox ListHyperlaneMailbox(ctx).Limit(limit).Offset(offset).Execute()

List hyperlane mailboxes info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HyperlaneAPI.ListHyperlaneMailbox(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HyperlaneAPI.ListHyperlaneMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHyperlaneMailbox`: []ResponsesHyperlaneMailbox
	fmt.Fprintf(os.Stdout, "Response from `HyperlaneAPI.ListHyperlaneMailbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHyperlaneMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesHyperlaneMailbox**](ResponsesHyperlaneMailbox.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHyperlaneTokens

> []ResponsesHyperlaneToken ListHyperlaneTokens(ctx).Limit(limit).Offset(offset).Sort(sort).Owner(owner).Mailbox(mailbox).Type_(type_).Execute()

List hyperlane tokens info



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
	owner := "owner_example" // string | Owner celestia address (optional)
	mailbox := "mailbox_example" // string | Mailbox hexademical identity (optional)
	type_ := "type__example" // string | Comma-separated string of tokens type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HyperlaneAPI.ListHyperlaneTokens(context.Background()).Limit(limit).Offset(offset).Sort(sort).Owner(owner).Mailbox(mailbox).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HyperlaneAPI.ListHyperlaneTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHyperlaneTokens`: []ResponsesHyperlaneToken
	fmt.Fprintf(os.Stdout, "Response from `HyperlaneAPI.ListHyperlaneTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHyperlaneTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **owner** | **string** | Owner celestia address | 
 **mailbox** | **string** | Mailbox hexademical identity | 
 **type_** | **string** | Comma-separated string of tokens type | 

### Return type

[**[]ResponsesHyperlaneToken**](ResponsesHyperlaneToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHyperlaneTransfers

> []ResponsesHyperlaneTransfer ListHyperlaneTransfers(ctx).Limit(limit).Offset(offset).Sort(sort).Address(address).Relayer(relayer).Mailbox(mailbox).Token(token).Type_(type_).Domain(domain).Hash(hash).Execute()

List hyperlane transfers info



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
	address := "address_example" // string | Celestia address (optional)
	relayer := "relayer_example" // string | Celestia address of relayer (optional)
	mailbox := "mailbox_example" // string | Mailbox hexademical identity (optional)
	token := "token_example" // string | Token hexademical identity (optional)
	type_ := "type__example" // string | Comma-separated string of transfer type (optional)
	domain := int32(56) // int32 | Domain of counterparty chain (optional)
	hash := "hash_example" // string | Transaction hash in hexadecimal (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HyperlaneAPI.ListHyperlaneTransfers(context.Background()).Limit(limit).Offset(offset).Sort(sort).Address(address).Relayer(relayer).Mailbox(mailbox).Token(token).Type_(type_).Domain(domain).Hash(hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HyperlaneAPI.ListHyperlaneTransfers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHyperlaneTransfers`: []ResponsesHyperlaneTransfer
	fmt.Fprintf(os.Stdout, "Response from `HyperlaneAPI.ListHyperlaneTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHyperlaneTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **address** | **string** | Celestia address | 
 **relayer** | **string** | Celestia address of relayer | 
 **mailbox** | **string** | Mailbox hexademical identity | 
 **token** | **string** | Token hexademical identity | 
 **type_** | **string** | Comma-separated string of transfer type | 
 **domain** | **int32** | Domain of counterparty chain | 
 **hash** | **string** | Transaction hash in hexadecimal | 

### Return type

[**[]ResponsesHyperlaneTransfer**](ResponsesHyperlaneTransfer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

