# ResponsesEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxId** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**TypesEventType**](TypesEventType.md) |  | [optional] 

## Methods

### NewResponsesEvent

`func NewResponsesEvent() *ResponsesEvent`

NewResponsesEvent instantiates a new ResponsesEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesEventWithDefaults

`func NewResponsesEventWithDefaults() *ResponsesEvent`

NewResponsesEventWithDefaults instantiates a new ResponsesEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponsesEvent) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponsesEvent) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponsesEvent) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ResponsesEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesEvent) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesEvent) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesEvent) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesEvent) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesEvent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *ResponsesEvent) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResponsesEvent) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResponsesEvent) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ResponsesEvent) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxId

`func (o *ResponsesEvent) GetTxId() int64`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ResponsesEvent) GetTxIdOk() (*int64, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ResponsesEvent) SetTxId(v int64)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *ResponsesEvent) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetType

`func (o *ResponsesEvent) GetType() TypesEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesEvent) GetTypeOk() (*TypesEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesEvent) SetType(v TypesEventType)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


