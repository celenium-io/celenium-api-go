# \GasAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GasEstimateForPfb**](GasAPI.md#GasEstimateForPfb) | **Get** /gas/estimate_for_pfb | Get estimated gas for pay for blob
[**GasPrice**](GasAPI.md#GasPrice) | **Get** /gas/price | Get estimated gas price
[**GasPricePriority**](GasAPI.md#GasPricePriority) | **Get** /gas/price/{priority} | Get estimated gas price with priority filter



## GasEstimateForPfb

> int32 GasEstimateForPfb(ctx).Sizes(sizes).Execute()

Get estimated gas for pay for blob



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
	sizes := "sizes_example" // string | Comma-separated array of blob sizes

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GasAPI.GasEstimateForPfb(context.Background()).Sizes(sizes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GasAPI.GasEstimateForPfb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GasEstimateForPfb`: int32
	fmt.Fprintf(os.Stdout, "Response from `GasAPI.GasEstimateForPfb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGasEstimateForPfbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sizes** | **string** | Comma-separated array of blob sizes | 

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


## GasPrice

> ResponsesGasPrice GasPrice(ctx).Execute()

Get estimated gas price



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
	resp, r, err := apiClient.GasAPI.GasPrice(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GasAPI.GasPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GasPrice`: ResponsesGasPrice
	fmt.Fprintf(os.Stdout, "Response from `GasAPI.GasPrice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGasPriceRequest struct via the builder pattern


### Return type

[**ResponsesGasPrice**](ResponsesGasPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GasPricePriority

> string GasPricePriority(ctx, priority).Execute()

Get estimated gas price with priority filter



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
	priority := "priority_example" // string | Priority

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GasAPI.GasPricePriority(context.Background(), priority).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GasAPI.GasPricePriority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GasPricePriority`: string
	fmt.Fprintf(os.Stdout, "Response from `GasAPI.GasPricePriority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priority** | **string** | Priority | 

### Other Parameters

Other parameters are passed through a pointer to a apiGasPricePriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

