# ResponsesChainMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockExplorers** | Pointer to [**[]ResponsesBlockExplorer**](ResponsesBlockExplorer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NativeToken** | Pointer to [**ResponsesNativeToken**](ResponsesNativeToken.md) |  | [optional] 

## Methods

### NewResponsesChainMetadata

`func NewResponsesChainMetadata() *ResponsesChainMetadata`

NewResponsesChainMetadata instantiates a new ResponsesChainMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesChainMetadataWithDefaults

`func NewResponsesChainMetadataWithDefaults() *ResponsesChainMetadata`

NewResponsesChainMetadataWithDefaults instantiates a new ResponsesChainMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockExplorers

`func (o *ResponsesChainMetadata) GetBlockExplorers() []ResponsesBlockExplorer`

GetBlockExplorers returns the BlockExplorers field if non-nil, zero value otherwise.

### GetBlockExplorersOk

`func (o *ResponsesChainMetadata) GetBlockExplorersOk() (*[]ResponsesBlockExplorer, bool)`

GetBlockExplorersOk returns a tuple with the BlockExplorers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExplorers

`func (o *ResponsesChainMetadata) SetBlockExplorers(v []ResponsesBlockExplorer)`

SetBlockExplorers sets BlockExplorers field to given value.

### HasBlockExplorers

`func (o *ResponsesChainMetadata) HasBlockExplorers() bool`

HasBlockExplorers returns a boolean if a field has been set.

### GetName

`func (o *ResponsesChainMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsesChainMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsesChainMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponsesChainMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeToken

`func (o *ResponsesChainMetadata) GetNativeToken() ResponsesNativeToken`

GetNativeToken returns the NativeToken field if non-nil, zero value otherwise.

### GetNativeTokenOk

`func (o *ResponsesChainMetadata) GetNativeTokenOk() (*ResponsesNativeToken, bool)`

GetNativeTokenOk returns a tuple with the NativeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeToken

`func (o *ResponsesChainMetadata) SetNativeToken(v ResponsesNativeToken)`

SetNativeToken sets NativeToken field to given value.

### HasNativeToken

`func (o *ResponsesChainMetadata) HasNativeToken() bool`

HasNativeToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


