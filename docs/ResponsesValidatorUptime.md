# ResponsesValidatorUptime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | Pointer to [**[]ResponsesSignedBlocks**](ResponsesSignedBlocks.md) |  | [optional] 
**Uptime** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesValidatorUptime

`func NewResponsesValidatorUptime() *ResponsesValidatorUptime`

NewResponsesValidatorUptime instantiates a new ResponsesValidatorUptime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesValidatorUptimeWithDefaults

`func NewResponsesValidatorUptimeWithDefaults() *ResponsesValidatorUptime`

NewResponsesValidatorUptimeWithDefaults instantiates a new ResponsesValidatorUptime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *ResponsesValidatorUptime) GetBlocks() []ResponsesSignedBlocks`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *ResponsesValidatorUptime) GetBlocksOk() (*[]ResponsesSignedBlocks, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *ResponsesValidatorUptime) SetBlocks(v []ResponsesSignedBlocks)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *ResponsesValidatorUptime) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetUptime

`func (o *ResponsesValidatorUptime) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ResponsesValidatorUptime) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ResponsesValidatorUptime) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ResponsesValidatorUptime) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


