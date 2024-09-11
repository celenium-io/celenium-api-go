# ResponsesNamespaceMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Namespace** | Pointer to [**ResponsesNamespace**](ResponsesNamespace.md) |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Tx** | Pointer to [**ResponsesTx**](ResponsesTx.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesNamespaceMessage

`func NewResponsesNamespaceMessage() *ResponsesNamespaceMessage`

NewResponsesNamespaceMessage instantiates a new ResponsesNamespaceMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesNamespaceMessageWithDefaults

`func NewResponsesNamespaceMessageWithDefaults() *ResponsesNamespaceMessage`

NewResponsesNamespaceMessageWithDefaults instantiates a new ResponsesNamespaceMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponsesNamespaceMessage) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponsesNamespaceMessage) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponsesNamespaceMessage) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ResponsesNamespaceMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesNamespaceMessage) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesNamespaceMessage) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesNamespaceMessage) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesNamespaceMessage) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesNamespaceMessage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesNamespaceMessage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesNamespaceMessage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesNamespaceMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNamespace

`func (o *ResponsesNamespaceMessage) GetNamespace() ResponsesNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ResponsesNamespaceMessage) GetNamespaceOk() (*ResponsesNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ResponsesNamespaceMessage) SetNamespace(v ResponsesNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ResponsesNamespaceMessage) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPosition

`func (o *ResponsesNamespaceMessage) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResponsesNamespaceMessage) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResponsesNamespaceMessage) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ResponsesNamespaceMessage) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesNamespaceMessage) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesNamespaceMessage) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesNamespaceMessage) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesNamespaceMessage) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTx

`func (o *ResponsesNamespaceMessage) GetTx() ResponsesTx`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *ResponsesNamespaceMessage) GetTxOk() (*ResponsesTx, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *ResponsesNamespaceMessage) SetTx(v ResponsesTx)`

SetTx sets Tx field to given value.

### HasTx

`func (o *ResponsesNamespaceMessage) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetType

`func (o *ResponsesNamespaceMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesNamespaceMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesNamespaceMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesNamespaceMessage) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


