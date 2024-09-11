# \SearchAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search**](SearchAPI.md#Search) | **Get** /search | Search by hash



## Search

> []ResponsesSearchItem Search(ctx).Query(query).Execute()

Search by hash



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
	query := "query_example" // string | Search string

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.Search(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: []ResponsesSearchItem
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Search string | 

### Return type

[**[]ResponsesSearchItem**](ResponsesSearchItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

