# ResponsesRelayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** |  | [optional] 
**Contact** | Pointer to [**ResponsesContact**](ResponsesContact.md) |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesRelayer

`func NewResponsesRelayer() *ResponsesRelayer`

NewResponsesRelayer instantiates a new ResponsesRelayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesRelayerWithDefaults

`func NewResponsesRelayerWithDefaults() *ResponsesRelayer`

NewResponsesRelayerWithDefaults instantiates a new ResponsesRelayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ResponsesRelayer) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ResponsesRelayer) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ResponsesRelayer) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ResponsesRelayer) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetContact

`func (o *ResponsesRelayer) GetContact() ResponsesContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ResponsesRelayer) GetContactOk() (*ResponsesContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ResponsesRelayer) SetContact(v ResponsesContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ResponsesRelayer) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLogo

`func (o *ResponsesRelayer) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ResponsesRelayer) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ResponsesRelayer) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ResponsesRelayer) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *ResponsesRelayer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsesRelayer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsesRelayer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponsesRelayer) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


