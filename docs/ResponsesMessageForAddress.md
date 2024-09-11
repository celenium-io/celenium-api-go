# ResponsesMessageForAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvocationType** | Pointer to [**TypesMsgAddressType**](TypesMsgAddressType.md) |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Tx** | Pointer to [**ResponsesTxForAddress**](ResponsesTxForAddress.md) |  | [optional] 
**TxId** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**TypesMsgType**](TypesMsgType.md) |  | [optional] 

## Methods

### NewResponsesMessageForAddress

`func NewResponsesMessageForAddress() *ResponsesMessageForAddress`

NewResponsesMessageForAddress instantiates a new ResponsesMessageForAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesMessageForAddressWithDefaults

`func NewResponsesMessageForAddressWithDefaults() *ResponsesMessageForAddress`

NewResponsesMessageForAddressWithDefaults instantiates a new ResponsesMessageForAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponsesMessageForAddress) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponsesMessageForAddress) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponsesMessageForAddress) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ResponsesMessageForAddress) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesMessageForAddress) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesMessageForAddress) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesMessageForAddress) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesMessageForAddress) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesMessageForAddress) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesMessageForAddress) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesMessageForAddress) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesMessageForAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvocationType

`func (o *ResponsesMessageForAddress) GetInvocationType() TypesMsgAddressType`

GetInvocationType returns the InvocationType field if non-nil, zero value otherwise.

### GetInvocationTypeOk

`func (o *ResponsesMessageForAddress) GetInvocationTypeOk() (*TypesMsgAddressType, bool)`

GetInvocationTypeOk returns a tuple with the InvocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationType

`func (o *ResponsesMessageForAddress) SetInvocationType(v TypesMsgAddressType)`

SetInvocationType sets InvocationType field to given value.

### HasInvocationType

`func (o *ResponsesMessageForAddress) HasInvocationType() bool`

HasInvocationType returns a boolean if a field has been set.

### GetPosition

`func (o *ResponsesMessageForAddress) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResponsesMessageForAddress) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResponsesMessageForAddress) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ResponsesMessageForAddress) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSize

`func (o *ResponsesMessageForAddress) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponsesMessageForAddress) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponsesMessageForAddress) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponsesMessageForAddress) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesMessageForAddress) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesMessageForAddress) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesMessageForAddress) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesMessageForAddress) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTx

`func (o *ResponsesMessageForAddress) GetTx() ResponsesTxForAddress`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *ResponsesMessageForAddress) GetTxOk() (*ResponsesTxForAddress, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *ResponsesMessageForAddress) SetTx(v ResponsesTxForAddress)`

SetTx sets Tx field to given value.

### HasTx

`func (o *ResponsesMessageForAddress) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxId

`func (o *ResponsesMessageForAddress) GetTxId() int64`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ResponsesMessageForAddress) GetTxIdOk() (*int64, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ResponsesMessageForAddress) SetTxId(v int64)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *ResponsesMessageForAddress) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetType

`func (o *ResponsesMessageForAddress) GetType() TypesMsgType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesMessageForAddress) GetTypeOk() (*TypesMsgType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesMessageForAddress) SetType(v TypesMsgType)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesMessageForAddress) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


