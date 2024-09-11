# ResponsesTxForAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to ***os.File** |  | [optional] 
**MessageTypes** | Pointer to [**[]TypesMsgType**](TypesMsgType.md) |  | [optional] 
**MessagesCount** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to [**GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus**](GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus.md) |  | [optional] 

## Methods

### NewResponsesTxForAddress

`func NewResponsesTxForAddress() *ResponsesTxForAddress`

NewResponsesTxForAddress instantiates a new ResponsesTxForAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesTxForAddressWithDefaults

`func NewResponsesTxForAddressWithDefaults() *ResponsesTxForAddress`

NewResponsesTxForAddressWithDefaults instantiates a new ResponsesTxForAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *ResponsesTxForAddress) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ResponsesTxForAddress) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ResponsesTxForAddress) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *ResponsesTxForAddress) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetHash

`func (o *ResponsesTxForAddress) GetHash() *os.File`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ResponsesTxForAddress) GetHashOk() (**os.File, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ResponsesTxForAddress) SetHash(v *os.File)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ResponsesTxForAddress) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMessageTypes

`func (o *ResponsesTxForAddress) GetMessageTypes() []TypesMsgType`

GetMessageTypes returns the MessageTypes field if non-nil, zero value otherwise.

### GetMessageTypesOk

`func (o *ResponsesTxForAddress) GetMessageTypesOk() (*[]TypesMsgType, bool)`

GetMessageTypesOk returns a tuple with the MessageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTypes

`func (o *ResponsesTxForAddress) SetMessageTypes(v []TypesMsgType)`

SetMessageTypes sets MessageTypes field to given value.

### HasMessageTypes

`func (o *ResponsesTxForAddress) HasMessageTypes() bool`

HasMessageTypes returns a boolean if a field has been set.

### GetMessagesCount

`func (o *ResponsesTxForAddress) GetMessagesCount() int64`

GetMessagesCount returns the MessagesCount field if non-nil, zero value otherwise.

### GetMessagesCountOk

`func (o *ResponsesTxForAddress) GetMessagesCountOk() (*int64, bool)`

GetMessagesCountOk returns a tuple with the MessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesCount

`func (o *ResponsesTxForAddress) SetMessagesCount(v int64)`

SetMessagesCount sets MessagesCount field to given value.

### HasMessagesCount

`func (o *ResponsesTxForAddress) HasMessagesCount() bool`

HasMessagesCount returns a boolean if a field has been set.

### GetStatus

`func (o *ResponsesTxForAddress) GetStatus() GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesTxForAddress) GetStatusOk() (*GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesTxForAddress) SetStatus(v GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesTxForAddress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


