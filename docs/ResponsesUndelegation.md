# ResponsesUndelegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**CompletionTime** | Pointer to **string** |  | [optional] 
**Delegator** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Validator** | Pointer to [**ResponsesShortValidator**](ResponsesShortValidator.md) |  | [optional] 

## Methods

### NewResponsesUndelegation

`func NewResponsesUndelegation() *ResponsesUndelegation`

NewResponsesUndelegation instantiates a new ResponsesUndelegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesUndelegationWithDefaults

`func NewResponsesUndelegationWithDefaults() *ResponsesUndelegation`

NewResponsesUndelegationWithDefaults instantiates a new ResponsesUndelegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ResponsesUndelegation) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResponsesUndelegation) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResponsesUndelegation) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ResponsesUndelegation) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCompletionTime

`func (o *ResponsesUndelegation) GetCompletionTime() string`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *ResponsesUndelegation) GetCompletionTimeOk() (*string, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *ResponsesUndelegation) SetCompletionTime(v string)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *ResponsesUndelegation) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetDelegator

`func (o *ResponsesUndelegation) GetDelegator() string`

GetDelegator returns the Delegator field if non-nil, zero value otherwise.

### GetDelegatorOk

`func (o *ResponsesUndelegation) GetDelegatorOk() (*string, bool)`

GetDelegatorOk returns a tuple with the Delegator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegator

`func (o *ResponsesUndelegation) SetDelegator(v string)`

SetDelegator sets Delegator field to given value.

### HasDelegator

`func (o *ResponsesUndelegation) HasDelegator() bool`

HasDelegator returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesUndelegation) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesUndelegation) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesUndelegation) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesUndelegation) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesUndelegation) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesUndelegation) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesUndelegation) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesUndelegation) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValidator

`func (o *ResponsesUndelegation) GetValidator() ResponsesShortValidator`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *ResponsesUndelegation) GetValidatorOk() (*ResponsesShortValidator, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *ResponsesUndelegation) SetValidator(v ResponsesShortValidator)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *ResponsesUndelegation) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


