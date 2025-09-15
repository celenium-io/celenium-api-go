# \ProposalAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProposal**](ProposalAPI.md#GetProposal) | **Get** /proposal/{id} | Get proposal info
[**ListProposal**](ProposalAPI.md#ListProposal) | **Get** /proposal | List proposal info
[**ProposalVotes**](ProposalAPI.md#ProposalVotes) | **Get** /proposal/{id}/votes | Get proposal&#39;s votes



## GetProposal

> ResponsesProposal GetProposal(ctx, id).Execute()

Get proposal info



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
	resp, r, err := apiClient.ProposalAPI.GetProposal(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProposalAPI.GetProposal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProposal`: ResponsesProposal
	fmt.Fprintf(os.Stdout, "Response from `ProposalAPI.GetProposal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProposalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesProposal**](ResponsesProposal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ProposalVotes

> []ResponsesVote ProposalVotes(ctx, id).Option(option).Voter(voter).Limit(limit).Offset(offset).Address(address).Validator(validator).Execute()

Get proposal's votes



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
	option := "option_example" // string | Option
	voter := "voter_example" // string | Voter type
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	address := "address_example" // string | Voter address (optional)
	validator := "validator_example" // string | Voter address (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProposalAPI.ProposalVotes(context.Background(), id).Option(option).Voter(voter).Limit(limit).Offset(offset).Address(address).Validator(validator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProposalAPI.ProposalVotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProposalVotes`: []ResponsesVote
	fmt.Fprintf(os.Stdout, "Response from `ProposalAPI.ProposalVotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Internal identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiProposalVotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **option** | **string** | Option | 
 **voter** | **string** | Voter type | 
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **address** | **string** | Voter address | 
 **validator** | **string** | Voter address | 

### Return type

[**[]ResponsesVote**](ResponsesVote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

