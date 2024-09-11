# ResponsesODSItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **[]int32** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**To** | Pointer to **[]int32** |  | [optional] 
**Type** | Pointer to [**ResponsesNamespaceKind**](ResponsesNamespaceKind.md) |  | [optional] 

## Methods

### NewResponsesODSItem

`func NewResponsesODSItem() *ResponsesODSItem`

NewResponsesODSItem instantiates a new ResponsesODSItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesODSItemWithDefaults

`func NewResponsesODSItemWithDefaults() *ResponsesODSItem`

NewResponsesODSItemWithDefaults instantiates a new ResponsesODSItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *ResponsesODSItem) GetFrom() []int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ResponsesODSItem) GetFromOk() (*[]int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ResponsesODSItem) SetFrom(v []int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ResponsesODSItem) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetNamespace

`func (o *ResponsesODSItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ResponsesODSItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ResponsesODSItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ResponsesODSItem) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTo

`func (o *ResponsesODSItem) GetTo() []int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ResponsesODSItem) GetToOk() (*[]int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ResponsesODSItem) SetTo(v []int32)`

SetTo sets To field to given value.

### HasTo

`func (o *ResponsesODSItem) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetType

`func (o *ResponsesODSItem) GetType() ResponsesNamespaceKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesODSItem) GetTypeOk() (*ResponsesNamespaceKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesODSItem) SetType(v ResponsesNamespaceKind)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesODSItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


