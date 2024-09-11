# \GeneralAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConstants**](GeneralAPI.md#GetConstants) | **Get** /constants | Get network constants
[**GetEnums**](GeneralAPI.md#GetEnums) | **Get** /enums | Get celenium enumerators
[**Head**](GeneralAPI.md#Head) | **Get** /head | Get current indexer head



## GetConstants

> ResponsesConstants GetConstants(ctx).Execute()

Get network constants



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
	resp, r, err := apiClient.GeneralAPI.GetConstants(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetConstants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConstants`: ResponsesConstants
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.GetConstants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConstantsRequest struct via the builder pattern


### Return type

[**ResponsesConstants**](ResponsesConstants.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnums

> ResponsesEnums GetEnums(ctx).Execute()

Get celenium enumerators



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
	resp, r, err := apiClient.GeneralAPI.GetEnums(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnums`: ResponsesEnums
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.GetEnums`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnumsRequest struct via the builder pattern


### Return type

[**ResponsesEnums**](ResponsesEnums.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Head

> ResponsesState Head(ctx).Execute()

Get current indexer head



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
	resp, r, err := apiClient.GeneralAPI.Head(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.Head``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Head`: ResponsesState
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.Head`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHeadRequest struct via the builder pattern


### Return type

[**ResponsesState**](ResponsesState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

