# ResponsesConstants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DenomMetadata** | Pointer to [**[]ResponsesDenomMetadata**](ResponsesDenomMetadata.md) |  | [optional] 
**Module** | Pointer to **map[string]map[string]string** |  | [optional] 

## Methods

### NewResponsesConstants

`func NewResponsesConstants() *ResponsesConstants`

NewResponsesConstants instantiates a new ResponsesConstants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesConstantsWithDefaults

`func NewResponsesConstantsWithDefaults() *ResponsesConstants`

NewResponsesConstantsWithDefaults instantiates a new ResponsesConstants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenomMetadata

`func (o *ResponsesConstants) GetDenomMetadata() []ResponsesDenomMetadata`

GetDenomMetadata returns the DenomMetadata field if non-nil, zero value otherwise.

### GetDenomMetadataOk

`func (o *ResponsesConstants) GetDenomMetadataOk() (*[]ResponsesDenomMetadata, bool)`

GetDenomMetadataOk returns a tuple with the DenomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomMetadata

`func (o *ResponsesConstants) SetDenomMetadata(v []ResponsesDenomMetadata)`

SetDenomMetadata sets DenomMetadata field to given value.

### HasDenomMetadata

`func (o *ResponsesConstants) HasDenomMetadata() bool`

HasDenomMetadata returns a boolean if a field has been set.

### GetModule

`func (o *ResponsesConstants) GetModule() map[string]map[string]string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ResponsesConstants) GetModuleOk() (*map[string]map[string]string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ResponsesConstants) SetModule(v map[string]map[string]string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ResponsesConstants) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


