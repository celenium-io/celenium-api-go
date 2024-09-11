# ResponsesMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Tx** | Pointer to [**ResponsesTx**](ResponsesTx.md) |  | [optional] 
**TxId** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**TypesMsgType**](TypesMsgType.md) |  | [optional] 

## Methods

### NewResponsesMessage

`func NewResponsesMessage() *ResponsesMessage`

NewResponsesMessage instantiates a new ResponsesMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesMessageWithDefaults

`func NewResponsesMessageWithDefaults() *ResponsesMessage`

NewResponsesMessageWithDefaults instantiates a new ResponsesMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponsesMessage) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponsesMessage) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponsesMessage) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ResponsesMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesMessage) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesMessage) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesMessage) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesMessage) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesMessage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesMessage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesMessage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *ResponsesMessage) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResponsesMessage) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResponsesMessage) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ResponsesMessage) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSize

`func (o *ResponsesMessage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponsesMessage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponsesMessage) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponsesMessage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesMessage) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesMessage) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesMessage) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesMessage) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTx

`func (o *ResponsesMessage) GetTx() ResponsesTx`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *ResponsesMessage) GetTxOk() (*ResponsesTx, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *ResponsesMessage) SetTx(v ResponsesTx)`

SetTx sets Tx field to given value.

### HasTx

`func (o *ResponsesMessage) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxId

`func (o *ResponsesMessage) GetTxId() int64`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ResponsesMessage) GetTxIdOk() (*int64, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ResponsesMessage) SetTxId(v int64)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *ResponsesMessage) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetType

`func (o *ResponsesMessage) GetType() TypesMsgType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesMessage) GetTypeOk() (*TypesMsgType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesMessage) SetType(v TypesMsgType)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesMessage) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


