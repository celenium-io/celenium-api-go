# \ProposalAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProposal**](ProposalAPI.md#ListProposal) | **Get** /proposal | List proposal info



## ListProposal

> []ResponsesProposal ListProposal(ctx).Limit(limit).Offset(offset).Sort(sort).Proposer(proposer).Status(status).Type_(type_).Execute()

List proposal info



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
	proposer := "proposer_example" // string | Proposer celestia address (optional)
	status := "status_example" // string | Comma-separated proposal status list (optional)
	type_ := "type__example" // string | Comma-separated proposal type list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProposalAPI.ListProposal(context.Background()).Limit(limit).Offset(offset).Sort(sort).Proposer(proposer).Status(status).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProposalAPI.ListProposal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProposal`: []ResponsesProposal
	fmt.Fprintf(os.Stdout, "Response from `ProposalAPI.ListProposal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProposalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **proposer** | **string** | Proposer celestia address | 
 **status** | **string** | Comma-separated proposal status list | 
 **type_** | **string** | Comma-separated proposal type list | 

### Return type

[**[]ResponsesProposal**](ResponsesProposal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

