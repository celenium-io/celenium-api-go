# \ValidatorAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetValidator**](ValidatorAPI.md#GetValidator) | **Get** /validators/{id} | Get validator info
[**GetValidatorBlocks**](ValidatorAPI.md#GetValidatorBlocks) | **Get** /validators/{id}/blocks | Get blocks which was proposed by validator
[**GetValidatorUptime**](ValidatorAPI.md#GetValidatorUptime) | **Get** /validators/{id}/uptime | Get validator&#39;s uptime and history of signed block
[**ListValidator**](ValidatorAPI.md#ListValidator) | **Get** /validators | List validators
[**ValidatorCount**](ValidatorAPI.md#ValidatorCount) | **Get** /validators/count | Get validator&#39;s count by status
[**ValidatorDelegators**](ValidatorAPI.md#ValidatorDelegators) | **Get** /validators/{id}/delegators | Get validator&#39;s delegators
[**ValidatorJails**](ValidatorAPI.md#ValidatorJails) | **Get** /validators/{id}/jails | Get validator&#39;s jails



## GetValidator

> ResponsesValidator GetValidator(ctx, id).Execute()

Get validator info



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
	id := int32(56) // int32 | Internal validator id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorAPI.GetValidator(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorAPI.GetValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidator`: ResponsesValidator
	fmt.Fprintf(os.Stdout, "Response from `ValidatorAPI.GetValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal validator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesValidator**](ResponsesValidator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidatorBlocks

> ResponsesBlock GetValidatorBlocks(ctx, id).Limit(limit).Offset(offset).Execute()

Get blocks which was proposed by validator



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
	id := int32(56) // int32 | Internal validator id
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorAPI.GetValidatorBlocks(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorAPI.GetValidatorBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidatorBlocks`: ResponsesBlock
	fmt.Fprintf(os.Stdout, "Response from `ValidatorAPI.GetValidatorBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal validator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidatorBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

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


## GetValidatorUptime

> ResponsesValidatorUptime GetValidatorUptime(ctx, id).Limit(limit).Execute()

Get validator's uptime and history of signed block



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
	id := int32(56) // int32 | Internal validator id
	limit := int32(56) // int32 | Count of requested blocks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorAPI.GetValidatorUptime(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorAPI.GetValidatorUptime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidatorUptime`: ResponsesValidatorUptime
	fmt.Fprintf(os.Stdout, "Response from `ValidatorAPI.GetValidatorUptime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal validator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidatorUptimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested blocks | 

### Return type

[**ResponsesValidatorUptime**](ResponsesValidatorUptime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListValidator

> []ResponsesValidator ListValidator(ctx).Limit(limit).Offset(offset).Jailed(jailed).Execute()

List validators



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
	jailed := true // bool | Return only jailed validators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorAPI.ListValidator(context.Background()).Limit(limit).Offset(offset).Jailed(jailed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorAPI.ListValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListValidator`: []ResponsesValidator
	fmt.Fprintf(os.Stdout, "Response from `ValidatorAPI.ListValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **jailed** | **bool** | Return only jailed validators | 

### Return type

[**[]ResponsesValidator**](ResponsesValidator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatorCount

> ResponsesValidatorCount ValidatorCount(ctx).Execute()

Get validator's count by status



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
	resp, r, err := apiClient.ValidatorAPI.ValidatorCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorAPI.ValidatorCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatorCount`: ResponsesValidatorCount
	fmt.Fprintf(os.Stdout, "Response from `ValidatorAPI.ValidatorCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiValidatorCountRequest struct via the builder pattern


### Return type

[**ResponsesValidatorCount**](ResponsesValidatorCount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatorDelegators

> []ResponsesDelegation ValidatorDelegators(ctx, id).Limit(limit).Offset(offset).ShowZero(showZero).Execute()

Get validator's delegators



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
	id := int32(56) // int32 | Internal validator id
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	showZero := true // bool | Show zero delegations (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorAPI.ValidatorDelegators(context.Background(), id).Limit(limit).Offset(offset).ShowZero(showZero).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorAPI.ValidatorDelegators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatorDelegators`: []ResponsesDelegation
	fmt.Fprintf(os.Stdout, "Response from `ValidatorAPI.ValidatorDelegators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal validator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidatorDelegatorsRequest struct via the builder pattern


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


## ValidatorJails

> []ResponsesJail ValidatorJails(ctx, id).Limit(limit).Offset(offset).Execute()

Get validator's jails



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
	id := int32(56) // int32 | Internal validator id
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorAPI.ValidatorJails(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorAPI.ValidatorJails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatorJails`: []ResponsesJail
	fmt.Fprintf(os.Stdout, "Response from `ValidatorAPI.ValidatorJails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal validator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidatorJailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesJail**](ResponsesJail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

