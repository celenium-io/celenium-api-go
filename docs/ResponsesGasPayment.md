# ResponsesGasPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**GasAmount** | Pointer to **string** |  | [optional] 
**IgpId** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewResponsesGasPayment

`func NewResponsesGasPayment() *ResponsesGasPayment`

NewResponsesGasPayment instantiates a new ResponsesGasPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesGasPaymentWithDefaults

`func NewResponsesGasPaymentWithDefaults() *ResponsesGasPayment`

NewResponsesGasPaymentWithDefaults instantiates a new ResponsesGasPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ResponsesGasPayment) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResponsesGasPayment) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResponsesGasPayment) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ResponsesGasPayment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetGasAmount

`func (o *ResponsesGasPayment) GetGasAmount() string`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *ResponsesGasPayment) GetGasAmountOk() (*string, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *ResponsesGasPayment) SetGasAmount(v string)`

SetGasAmount sets GasAmount field to given value.

### HasGasAmount

`func (o *ResponsesGasPayment) HasGasAmount() bool`

HasGasAmount returns a boolean if a field has been set.

### GetIgpId

`func (o *ResponsesGasPayment) GetIgpId() *os.File`

GetIgpId returns the IgpId field if non-nil, zero value otherwise.

### GetIgpIdOk

`func (o *ResponsesGasPayment) GetIgpIdOk() (**os.File, bool)`

GetIgpIdOk returns a tuple with the IgpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgpId

`func (o *ResponsesGasPayment) SetIgpId(v *os.File)`

SetIgpId sets IgpId field to given value.

### HasIgpId

`func (o *ResponsesGasPayment) HasIgpId() bool`

HasIgpId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


