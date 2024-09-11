# \NamespaceAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlob**](NamespaceAPI.md#GetBlob) | **Post** /blob | Get namespace blob by commitment on height
[**GetBlobLogs**](NamespaceAPI.md#GetBlobLogs) | **Get** /namespace/{id}/{version}/blobs | Get blob changes for namespace
[**GetBlobMetadata**](NamespaceAPI.md#GetBlobMetadata) | **Post** /blob/metadata | Get blob metadata by commitment on height
[**GetBlobs**](NamespaceAPI.md#GetBlobs) | **Get** /blob | List all blobs with filters
[**GetNamespace**](NamespaceAPI.md#GetNamespace) | **Get** /namespace/{id} | Get namespace info
[**GetNamespaceActive**](NamespaceAPI.md#GetNamespaceActive) | **Get** /namespace/active | Get last used namespace
[**GetNamespaceBase64**](NamespaceAPI.md#GetNamespaceBase64) | **Get** /namespace_by_hash/{hash} | Get namespace info by base64
[**GetNamespaceBlobs**](NamespaceAPI.md#GetNamespaceBlobs) | **Get** /namespace_by_hash/{hash}/{height} | Get namespace blobs on height
[**GetNamespaceByVersionAndId**](NamespaceAPI.md#GetNamespaceByVersionAndId) | **Get** /namespace/{id}/{version} | Get namespace info by id and version
[**GetNamespaceCount**](NamespaceAPI.md#GetNamespaceCount) | **Get** /namespace/count | Get count of namespaces in network
[**GetNamespaceMessages**](NamespaceAPI.md#GetNamespaceMessages) | **Get** /namespace/{id}/{version}/messages | Get namespace messages by id and version
[**GetNamespaceRollups**](NamespaceAPI.md#GetNamespaceRollups) | **Get** /namespace/{id}/{version}/rollups | List rollups using the namespace
[**ListNamespace**](NamespaceAPI.md#ListNamespace) | **Get** /namespace | List namespace info



## GetBlob

> ResponsesBlob GetBlob(ctx).Commitment(commitment).Execute()

Get namespace blob by commitment on height



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
	commitment := "commitment_example" // string | Blob commitment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetBlob(context.Background()).Commitment(commitment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetBlob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlob`: ResponsesBlob
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetBlob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitment** | **string** | Blob commitment | 

### Return type

[**ResponsesBlob**](ResponsesBlob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobLogs

> []ResponsesBlobLog GetBlobLogs(ctx, id, version).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Commitment(commitment).From(from).To(to).Joins(joins).Signers(signers).Cursor(cursor).Execute()

Get blob changes for namespace



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
	id := "id_example" // string | Namespace id in hexadecimal
	version := int32(56) // int32 | Version of namespace
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)
	sort := "sort_example" // string | Sort order. Default: desc (optional)
	sortBy := "sortBy_example" // string | Sort field. If it's empty internal id is used (optional)
	commitment := "commitment_example" // string | Commitment value in URLbase64 format (optional)
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)
	joins := true // bool | Flag indicating whether entities of rollup, transaction and signer should be attached or not. Default: true (optional)
	signers := "signers_example" // string | Comma-separated celestia addresses (optional)
	cursor := int32(56) // int32 | Last entity id which is used for cursor pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetBlobLogs(context.Background(), id, version).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Commitment(commitment).From(from).To(to).Joins(joins).Signers(signers).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetBlobLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobLogs`: []ResponsesBlobLog
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetBlobLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace id in hexadecimal | 
**version** | **int32** | Version of namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **sortBy** | **string** | Sort field. If it&#39;s empty internal id is used | 
 **commitment** | **string** | Commitment value in URLbase64 format | 
 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 
 **joins** | **bool** | Flag indicating whether entities of rollup, transaction and signer should be attached or not. Default: true | 
 **signers** | **string** | Comma-separated celestia addresses | 
 **cursor** | **int32** | Last entity id which is used for cursor pagination | 

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


## GetBlobMetadata

> ResponsesBlobLog GetBlobMetadata(ctx).Commitment(commitment).Execute()

Get blob metadata by commitment on height



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
	commitment := "commitment_example" // string | Blob commitment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetBlobMetadata(context.Background()).Commitment(commitment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetBlobMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobMetadata`: ResponsesBlobLog
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetBlobMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitment** | **string** | Blob commitment | 

### Return type

[**ResponsesBlobLog**](ResponsesBlobLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobs

> []ResponsesLightBlobLog GetBlobs(ctx).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Commitment(commitment).From(from).To(to).Signers(signers).Namespaces(namespaces).Cursor(cursor).Execute()

List all blobs with filters



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
	sortBy := "sortBy_example" // string | Sort field. If it's empty internal id is used (optional)
	commitment := "commitment_example" // string | Commitment value in URLbase64 format (optional)
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)
	signers := "signers_example" // string | Comma-separated celestia addresses (optional)
	namespaces := "namespaces_example" // string | Comma-separated celestia namespaces (optional)
	cursor := int32(56) // int32 | Last entity id which is used for cursor pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetBlobs(context.Background()).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Commitment(commitment).From(from).To(to).Signers(signers).Namespaces(namespaces).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetBlobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobs`: []ResponsesLightBlobLog
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetBlobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **sortBy** | **string** | Sort field. If it&#39;s empty internal id is used | 
 **commitment** | **string** | Commitment value in URLbase64 format | 
 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 
 **signers** | **string** | Comma-separated celestia addresses | 
 **namespaces** | **string** | Comma-separated celestia namespaces | 
 **cursor** | **int32** | Last entity id which is used for cursor pagination | 

### Return type

[**[]ResponsesLightBlobLog**](ResponsesLightBlobLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespace

> []ResponsesNamespace GetNamespace(ctx, id).Execute()

Get namespace info



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
	id := "id_example" // string | Namespace id in hexadecimal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetNamespace(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespace`: []ResponsesNamespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace id in hexadecimal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ResponsesNamespace**](ResponsesNamespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceActive

> []ResponsesNamespace GetNamespaceActive(ctx).Sort(sort).Execute()

Get last used namespace



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
	sort := "sort_example" // string | Sort field. Default: time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetNamespaceActive(context.Background()).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetNamespaceActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceActive`: []ResponsesNamespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetNamespaceActive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Default: time | 

### Return type

[**[]ResponsesNamespace**](ResponsesNamespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceBase64

> ResponsesNamespace GetNamespaceBase64(ctx, hash).Execute()

Get namespace info by base64



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
	hash := "hash_example" // string | Base64-encoded namespace id and version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetNamespaceBase64(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetNamespaceBase64``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceBase64`: ResponsesNamespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetNamespaceBase64`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Base64-encoded namespace id and version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceBase64Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsesNamespace**](ResponsesNamespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceBlobs

> []ResponsesBlob GetNamespaceBlobs(ctx, hash, height).Execute()

Get namespace blobs on height



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
	hash := "hash_example" // string | Base64-encoded namespace id and version
	height := int32(56) // int32 | Block heigth

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetNamespaceBlobs(context.Background(), hash, height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetNamespaceBlobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceBlobs`: []ResponsesBlob
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetNamespaceBlobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Base64-encoded namespace id and version | 
**height** | **int32** | Block heigth | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceBlobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ResponsesBlob**](ResponsesBlob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceByVersionAndId

> ResponsesNamespace GetNamespaceByVersionAndId(ctx, id, version).Execute()

Get namespace info by id and version



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
	id := "id_example" // string | Namespace id in hexadecimal
	version := int32(56) // int32 | Version of namespace

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetNamespaceByVersionAndId(context.Background(), id, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetNamespaceByVersionAndId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceByVersionAndId`: ResponsesNamespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetNamespaceByVersionAndId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace id in hexadecimal | 
**version** | **int32** | Version of namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceByVersionAndIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponsesNamespace**](ResponsesNamespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceCount

> int32 GetNamespaceCount(ctx).Execute()

Get count of namespaces in network



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
	resp, r, err := apiClient.NamespaceAPI.GetNamespaceCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetNamespaceCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetNamespaceCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceCountRequest struct via the builder pattern


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


## GetNamespaceMessages

> []ResponsesNamespaceMessage GetNamespaceMessages(ctx, id, version).Limit(limit).Offset(offset).Execute()

Get namespace messages by id and version



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
	id := "id_example" // string | Namespace id in hexadecimal
	version := int32(56) // int32 | Version of namespace
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetNamespaceMessages(context.Background(), id, version).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetNamespaceMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceMessages`: []ResponsesNamespaceMessage
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetNamespaceMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace id in hexadecimal | 
**version** | **int32** | Version of namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesNamespaceMessage**](ResponsesNamespaceMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceRollups

> []ResponsesRollup GetNamespaceRollups(ctx, id, version).Limit(limit).Offset(offset).Execute()

List rollups using the namespace



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
	id := "id_example" // string | Namespace id in hexadecimal
	version := int32(56) // int32 | Version of namespace
	limit := int32(56) // int32 | Count of requested entities (optional)
	offset := int32(56) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.GetNamespaceRollups(context.Background(), id, version).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.GetNamespaceRollups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceRollups`: []ResponsesRollup
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.GetNamespaceRollups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace id in hexadecimal | 
**version** | **int32** | Version of namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceRollupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 

### Return type

[**[]ResponsesRollup**](ResponsesRollup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespace

> []ResponsesNamespace ListNamespace(ctx).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()

List namespace info



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
	sortBy := "sortBy_example" // string | Sort field. If it's empty internal id is used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.ListNamespace(context.Background()).Limit(limit).Offset(offset).Sort(sort).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.ListNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespace`: []ResponsesNamespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.ListNamespace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of requested entities | 
 **offset** | **int32** | Offset | 
 **sort** | **string** | Sort order. Default: desc | 
 **sortBy** | **string** | Sort field. If it&#39;s empty internal id is used | 

### Return type

[**[]ResponsesNamespace**](ResponsesNamespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

