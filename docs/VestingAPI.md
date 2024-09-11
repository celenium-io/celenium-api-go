# \VestingAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVestingPeriods**](VestingAPI.md#GetVestingPeriods) | **Get** /vesting/{id}/periods | Periods vesting periods by id



## GetVestingPeriods

> ResponsesVestingPeriod GetVestingPeriods(ctx, id).Limit(limit).Offset(offset).Execute()

Periods vesting periods by id



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
	resp, r, err := apiClient.VestingAPI.GetVestingPeriods(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VestingAPI.GetVestingPeriods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVestingPeriods`: ResponsesVestingPeriod
	fmt.Fprintf(os.Stdout, "Response from `VestingAPI.GetVestingPeriods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVestingPeriodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**ResponsesVestingPeriod**](ResponsesVestingPeriod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

