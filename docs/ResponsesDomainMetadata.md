# ResponsesDomainMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockExplorers** | Pointer to [**[]ResponsesBlockExplorer**](ResponsesBlockExplorer.md) |  | [optional] 
**Domain** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NativeToken** | Pointer to [**ResponsesNativeToken**](ResponsesNativeToken.md) |  | [optional] 

## Methods

### NewResponsesDomainMetadata

`func NewResponsesDomainMetadata() *ResponsesDomainMetadata`

NewResponsesDomainMetadata instantiates a new ResponsesDomainMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesDomainMetadataWithDefaults

`func NewResponsesDomainMetadataWithDefaults() *ResponsesDomainMetadata`

NewResponsesDomainMetadataWithDefaults instantiates a new ResponsesDomainMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockExplorers

`func (o *ResponsesDomainMetadata) GetBlockExplorers() []ResponsesBlockExplorer`

GetBlockExplorers returns the BlockExplorers field if non-nil, zero value otherwise.

### GetBlockExplorersOk

`func (o *ResponsesDomainMetadata) GetBlockExplorersOk() (*[]ResponsesBlockExplorer, bool)`

GetBlockExplorersOk returns a tuple with the BlockExplorers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExplorers

`func (o *ResponsesDomainMetadata) SetBlockExplorers(v []ResponsesBlockExplorer)`

SetBlockExplorers sets BlockExplorers field to given value.

### HasBlockExplorers

`func (o *ResponsesDomainMetadata) HasBlockExplorers() bool`

HasBlockExplorers returns a boolean if a field has been set.

### GetDomain

`func (o *ResponsesDomainMetadata) GetDomain() int32`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResponsesDomainMetadata) GetDomainOk() (*int32, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResponsesDomainMetadata) SetDomain(v int32)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ResponsesDomainMetadata) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *ResponsesDomainMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsesDomainMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsesDomainMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponsesDomainMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeToken

`func (o *ResponsesDomainMetadata) GetNativeToken() ResponsesNativeToken`

GetNativeToken returns the NativeToken field if non-nil, zero value otherwise.

### GetNativeTokenOk

`func (o *ResponsesDomainMetadata) GetNativeTokenOk() (*ResponsesNativeToken, bool)`

GetNativeTokenOk returns a tuple with the NativeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeToken

`func (o *ResponsesDomainMetadata) SetNativeToken(v ResponsesNativeToken)`

SetNativeToken sets NativeToken field to given value.

### HasNativeToken

`func (o *ResponsesDomainMetadata) HasNativeToken() bool`

HasNativeToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


