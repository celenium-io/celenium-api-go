# \SignalAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSignal**](SignalAPI.md#ListSignal) | **Get** /signal | List signals
[**ListUpgrades**](SignalAPI.md#ListUpgrades) | **Get** /signal/upgrade | List upgrades



## ListSignal

> []ResponsesSignalVersion ListSignal(ctx).Version(version).ValidatorId(validatorId).TxHash(txHash).Limit(limit).Offset(offset).From(from).To(to).Sort(sort).Execute()

List signals



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
	version := int32(56) // int32 | Version (optional)
	validatorId := int32(56) // int32 | Validator internal id (optional)
	txHash := "txHash_example" // string | Transaction hash (optional)
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalAPI.ListSignal(context.Background()).Version(version).ValidatorId(validatorId).TxHash(txHash).Limit(limit).Offset(offset).From(from).To(to).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalAPI.ListSignal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSignal`: []ResponsesSignalVersion
	fmt.Fprintf(os.Stdout, "Response from `SignalAPI.ListSignal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSignalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **int32** | Version | 
 **validatorId** | **int32** | Validator internal id | 
 **txHash** | **string** | Transaction hash | 
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 
 **sort** | **string** | Sort order. Default: desc | 

### Return type

[**[]ResponsesSignalVersion**](ResponsesSignalVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUpgrades

> []ResponsesUpgrade ListUpgrades(ctx).Height(height).TxHash(txHash).Signer(signer).Limit(limit).Offset(offset).Sort(sort).Execute()

List upgrades



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
	height := int32(56) // int32 | Number of block (optional)
	txHash := "txHash_example" // string | Transaction hash (optional)
	signer := "signer_example" // string | Signer address (optional)
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalAPI.ListUpgrades(context.Background()).Height(height).TxHash(txHash).Signer(signer).Limit(limit).Offset(offset).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalAPI.ListUpgrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUpgrades`: []ResponsesUpgrade
	fmt.Fprintf(os.Stdout, "Response from `SignalAPI.ListUpgrades`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **int32** | Number of block | 
 **txHash** | **string** | Transaction hash | 
 **signer** | **string** | Signer address | 
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 

### Return type

[**[]ResponsesUpgrade**](ResponsesUpgrade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

