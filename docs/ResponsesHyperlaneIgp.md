# ResponsesHyperlaneIgp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | Pointer to [**[]ResponsesHyperlaneIgpConfig**](ResponsesHyperlaneIgpConfig.md) |  | [optional] 
**Denom** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IgpId** | Pointer to ***os.File** |  | [optional] 
**Owner** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewResponsesHyperlaneIgp

`func NewResponsesHyperlaneIgp() *ResponsesHyperlaneIgp`

NewResponsesHyperlaneIgp instantiates a new ResponsesHyperlaneIgp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesHyperlaneIgpWithDefaults

`func NewResponsesHyperlaneIgpWithDefaults() *ResponsesHyperlaneIgp`

NewResponsesHyperlaneIgpWithDefaults instantiates a new ResponsesHyperlaneIgp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *ResponsesHyperlaneIgp) GetConfigs() []ResponsesHyperlaneIgpConfig`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *ResponsesHyperlaneIgp) GetConfigsOk() (*[]ResponsesHyperlaneIgpConfig, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *ResponsesHyperlaneIgp) SetConfigs(v []ResponsesHyperlaneIgpConfig)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *ResponsesHyperlaneIgp) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetDenom

`func (o *ResponsesHyperlaneIgp) GetDenom() string`

GetDenom returns the Denom field if non-nil, zero value otherwise.

### GetDenomOk

`func (o *ResponsesHyperlaneIgp) GetDenomOk() (*string, bool)`

GetDenomOk returns a tuple with the Denom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenom

`func (o *ResponsesHyperlaneIgp) SetDenom(v string)`

SetDenom sets Denom field to given value.

### HasDenom

`func (o *ResponsesHyperlaneIgp) HasDenom() bool`

HasDenom returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesHyperlaneIgp) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesHyperlaneIgp) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesHyperlaneIgp) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesHyperlaneIgp) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesHyperlaneIgp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesHyperlaneIgp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesHyperlaneIgp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesHyperlaneIgp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgpId

`func (o *ResponsesHyperlaneIgp) GetIgpId() *os.File`

GetIgpId returns the IgpId field if non-nil, zero value otherwise.

### GetIgpIdOk

`func (o *ResponsesHyperlaneIgp) GetIgpIdOk() (**os.File, bool)`

GetIgpIdOk returns a tuple with the IgpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgpId

`func (o *ResponsesHyperlaneIgp) SetIgpId(v *os.File)`

SetIgpId sets IgpId field to given value.

### HasIgpId

`func (o *ResponsesHyperlaneIgp) HasIgpId() bool`

HasIgpId returns a boolean if a field has been set.

### GetOwner

`func (o *ResponsesHyperlaneIgp) GetOwner() ResponsesShortAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResponsesHyperlaneIgp) GetOwnerOk() (*ResponsesShortAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResponsesHyperlaneIgp) SetOwner(v ResponsesShortAddress)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ResponsesHyperlaneIgp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesHyperlaneIgp) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesHyperlaneIgp) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesHyperlaneIgp) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesHyperlaneIgp) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


