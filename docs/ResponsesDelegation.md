# ResponsesDelegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Delegator** | Pointer to **string** |  | [optional] 
**Validator** | Pointer to [**ResponsesShortValidator**](ResponsesShortValidator.md) |  | [optional] 

## Methods

### NewResponsesDelegation

`func NewResponsesDelegation() *ResponsesDelegation`

NewResponsesDelegation instantiates a new ResponsesDelegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesDelegationWithDefaults

`func NewResponsesDelegationWithDefaults() *ResponsesDelegation`

NewResponsesDelegationWithDefaults instantiates a new ResponsesDelegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ResponsesDelegation) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResponsesDelegation) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResponsesDelegation) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ResponsesDelegation) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDelegator

`func (o *ResponsesDelegation) GetDelegator() string`

GetDelegator returns the Delegator field if non-nil, zero value otherwise.

### GetDelegatorOk

`func (o *ResponsesDelegation) GetDelegatorOk() (*string, bool)`

GetDelegatorOk returns a tuple with the Delegator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegator

`func (o *ResponsesDelegation) SetDelegator(v string)`

SetDelegator sets Delegator field to given value.

### HasDelegator

`func (o *ResponsesDelegation) HasDelegator() bool`

HasDelegator returns a boolean if a field has been set.

### GetValidator

`func (o *ResponsesDelegation) GetValidator() ResponsesShortValidator`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *ResponsesDelegation) GetValidatorOk() (*ResponsesShortValidator, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *ResponsesDelegation) SetValidator(v ResponsesShortValidator)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *ResponsesDelegation) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


