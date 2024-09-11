# \AddressAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressBlobs**](AddressAPI.md#AddressBlobs) | **Get** /address/{hash}/blobs | Get blobs pushed by address
[**AddressDelegations**](AddressAPI.md#AddressDelegations) | **Get** /address/{hash}/delegations | Get delegations made by address
[**AddressGrantee**](AddressAPI.md#AddressGrantee) | **Get** /address/{hash}/granters | Get grants where address is grantee
[**AddressGrants**](AddressAPI.md#AddressGrants) | **Get** /address/{hash}/grants | Get grants made by address
[**AddressMessages**](AddressAPI.md#AddressMessages) | **Get** /address/{hash}/messages | Get address messages
[**AddressRedelegations**](AddressAPI.md#AddressRedelegations) | **Get** /address/{hash}/redelegations | Get redelegations made by address
[**AddressStats**](AddressAPI.md#AddressStats) | **Get** /address/{hash}/stats/{name}/{timeframe} | Get address stats
[**AddressTransactions**](AddressAPI.md#AddressTransactions) | **Get** /address/{hash}/txs | Get address transactions
[**AddressUndelegations**](AddressAPI.md#AddressUndelegations) | **Get** /address/{hash}/undelegations | Get undelegations made by address
[**AddressVesting**](AddressAPI.md#AddressVesting) | **Get** /address/{hash}/vestings | Get vesting for address
[**GetAddress**](AddressAPI.md#GetAddress) | **Get** /address/{hash} | Get address info
[**GetAddressCount**](AddressAPI.md#GetAddressCount) | **Get** /address/count | Get count of addresses in network
[**ListAddress**](AddressAPI.md#ListAddress) | **Get** /address | List address info



## AddressBlobs

> []ResponsesBlobLog AddressBlobs(ctx, hash).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Joins(joins).Execute()

Get blobs pushed by address



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
	hash := "hash_example" // string | Hash
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)
	sortBy := "sortBy_example" // string | Sort field. If it's empty internal id is used (optional)
	joins := true // bool | Flag indicating whether entities of transaction and namespace should be attached or not. Default: true (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressBlobs(context.Background(), hash).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Joins(joins).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressBlobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlobs`: []ResponsesBlobLog
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressBlobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **sortBy** | **string** | Sort field. If it&#39;s empty internal id is used | 
 **joins** | **bool** | Flag indicating whether entities of transaction and namespace should be attached or not. Default: true | 

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


## AddressDelegations

> []ResponsesDelegation AddressDelegations(ctx, hash).Limit(limit).Offset(offset).ShowZero(showZero).Execute()

Get delegations made by address



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
	hash := "hash_example" // string | Hash
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	showZero := true // bool | Show zero delegations (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressDelegations(context.Background(), hash).Limit(limit).Offset(offset).ShowZero(showZero).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressDelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressDelegations`: []ResponsesDelegation
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressDelegations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressDelegationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **showZero** | **bool** | Show zero delegations | 

### Return type

[**[]ResponsesDelegation**](ResponsesDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressGrantee

> []ResponsesGrant AddressGrantee(ctx, hash).Limit(limit).Offset(offset).Execute()

Get grants where address is grantee



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
	hash := "hash_example" // string | Hash
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressGrantee(context.Background(), hash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressGrantee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressGrantee`: []ResponsesGrant
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressGrantee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressGranteeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesGrant**](ResponsesGrant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressGrants

> []ResponsesGrant AddressGrants(ctx, hash).Limit(limit).Offset(offset).Execute()

Get grants made by address



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
	hash := "hash_example" // string | Hash
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressGrants(context.Background(), hash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressGrants`: []ResponsesGrant
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesGrant**](ResponsesGrant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressMessages

> []ResponsesMessageForAddress AddressMessages(ctx, hash).Limit(limit).Offset(offset).Sort(sort).MsgType(msgType).Execute()

Get address messages



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
	hash := "hash_example" // string | Hash
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order (optional)
	msgType := "msgType_example" // string | Comma-separated message types list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressMessages(context.Background(), hash).Limit(limit).Offset(offset).Sort(sort).MsgType(msgType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressMessages`: []ResponsesMessageForAddress
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order | 
 **msgType** | **string** | Comma-separated message types list | 

### Return type

[**[]ResponsesMessageForAddress**](ResponsesMessageForAddress.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressRedelegations

> []ResponsesRedelegation AddressRedelegations(ctx, hash).Limit(limit).Offset(offset).Execute()

Get redelegations made by address



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
	hash := "hash_example" // string | Hash
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressRedelegations(context.Background(), hash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressRedelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressRedelegations`: []ResponsesRedelegation
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressRedelegations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressRedelegationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesRedelegation**](ResponsesRedelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressStats

> []ResponsesHistogramItem AddressStats(ctx, hash, name, timeframe).From(from).To(to).Execute()

Get address stats



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
	hash := "hash_example" // string | Hash
	name := "name_example" // string | Series name
	timeframe := "timeframe_example" // string | Timeframe
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressStats(context.Background(), hash, name, timeframe).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressStats`: []ResponsesHistogramItem
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 
**name** | **string** | Series name | 
**timeframe** | **string** | Timeframe | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressStatsRequest struct via the builder pattern


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


## AddressTransactions

> []ResponsesTx AddressTransactions(ctx, hash).Limit(limit).Offset(offset).Sort(sort).Status(status).MsgType(msgType).From(from).To(to).Height(height).Execute()

Get address transactions



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
	hash := "hash_example" // string | Hash
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order (optional)
	status := "status_example" // string | Comma-separated status list (optional)
	msgType := "msgType_example" // string | Comma-separated message types list (optional)
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)
	height := int32(56) // int32 | Block number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressTransactions(context.Background(), hash).Limit(limit).Offset(offset).Sort(sort).Status(status).MsgType(msgType).From(from).To(to).Height(height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressTransactions`: []ResponsesTx
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order | 
 **status** | **string** | Comma-separated status list | 
 **msgType** | **string** | Comma-separated message types list | 
 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 
 **height** | **int32** | Block number | 

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


## AddressUndelegations

> []ResponsesUndelegation AddressUndelegations(ctx, hash).Limit(limit).Offset(offset).Execute()

Get undelegations made by address



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
	hash := "hash_example" // string | Hash
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressUndelegations(context.Background(), hash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressUndelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressUndelegations`: []ResponsesUndelegation
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressUndelegations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressUndelegationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesUndelegation**](ResponsesUndelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressVesting

> []ResponsesVesting AddressVesting(ctx, hash).Limit(limit).Offset(offset).ShowEnded(showEnded).Execute()

Get vesting for address



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
	hash := "hash_example" // string | Hash
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	showEnded := true // bool | Show finished vestings delegations (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressVesting(context.Background(), hash).Limit(limit).Offset(offset).ShowEnded(showEnded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressVesting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressVesting`: []ResponsesVesting
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressVesting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressVestingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **showEnded** | **bool** | Show finished vestings delegations | 

### Return type

[**[]ResponsesVesting**](ResponsesVesting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddress

> ResponsesAddress GetAddress(ctx, hash).Execute()

Get address info



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
	hash := "hash_example" // string | Hash

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.GetAddress(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.GetAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddress`: ResponsesAddress
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.GetAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesAddress**](ResponsesAddress.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressCount

> int32 GetAddressCount(ctx).Execute()

Get count of addresses in network



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
	resp, r, err := apiClient.AddressAPI.GetAddressCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.GetAddressCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.GetAddressCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressCountRequest struct via the builder pattern


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


## ListAddress

> []ResponsesAddress ListAddress(ctx).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()

List address info



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
	sortBy := "sortBy_example" // string | Sort field (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.ListAddress(context.Background()).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.ListAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAddress`: []ResponsesAddress
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.ListAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order | 
 **sortBy** | **string** | Sort field | 

### Return type

[**[]ResponsesAddress**](ResponsesAddress.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

