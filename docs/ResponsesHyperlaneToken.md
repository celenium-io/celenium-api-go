# ResponsesHyperlaneToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denom** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Mailbox** | Pointer to ***os.File** |  | [optional] 
**Owner** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Received** | Pointer to **string** |  | [optional] 
**ReceivedTransfers** | Pointer to **int64** |  | [optional] 
**Sent** | Pointer to **string** |  | [optional] 
**SentTransfers** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TokenId** | Pointer to ***os.File** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesHyperlaneToken

`func NewResponsesHyperlaneToken() *ResponsesHyperlaneToken`

NewResponsesHyperlaneToken instantiates a new ResponsesHyperlaneToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesHyperlaneTokenWithDefaults

`func NewResponsesHyperlaneTokenWithDefaults() *ResponsesHyperlaneToken`

NewResponsesHyperlaneTokenWithDefaults instantiates a new ResponsesHyperlaneToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenom

`func (o *ResponsesHyperlaneToken) GetDenom() string`

GetDenom returns the Denom field if non-nil, zero value otherwise.

### GetDenomOk

`func (o *ResponsesHyperlaneToken) GetDenomOk() (*string, bool)`

GetDenomOk returns a tuple with the Denom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenom

`func (o *ResponsesHyperlaneToken) SetDenom(v string)`

SetDenom sets Denom field to given value.

### HasDenom

`func (o *ResponsesHyperlaneToken) HasDenom() bool`

HasDenom returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesHyperlaneToken) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesHyperlaneToken) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesHyperlaneToken) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesHyperlaneToken) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesHyperlaneToken) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesHyperlaneToken) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesHyperlaneToken) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesHyperlaneToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMailbox

`func (o *ResponsesHyperlaneToken) GetMailbox() *os.File`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *ResponsesHyperlaneToken) GetMailboxOk() (**os.File, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *ResponsesHyperlaneToken) SetMailbox(v *os.File)`

SetMailbox sets Mailbox field to given value.

### HasMailbox

`func (o *ResponsesHyperlaneToken) HasMailbox() bool`

HasMailbox returns a boolean if a field has been set.

### GetOwner

`func (o *ResponsesHyperlaneToken) GetOwner() ResponsesShortAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResponsesHyperlaneToken) GetOwnerOk() (*ResponsesShortAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResponsesHyperlaneToken) SetOwner(v ResponsesShortAddress)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ResponsesHyperlaneToken) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetReceived

`func (o *ResponsesHyperlaneToken) GetReceived() string`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *ResponsesHyperlaneToken) GetReceivedOk() (*string, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *ResponsesHyperlaneToken) SetReceived(v string)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *ResponsesHyperlaneToken) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetReceivedTransfers

`func (o *ResponsesHyperlaneToken) GetReceivedTransfers() int64`

GetReceivedTransfers returns the ReceivedTransfers field if non-nil, zero value otherwise.

### GetReceivedTransfersOk

`func (o *ResponsesHyperlaneToken) GetReceivedTransfersOk() (*int64, bool)`

GetReceivedTransfersOk returns a tuple with the ReceivedTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTransfers

`func (o *ResponsesHyperlaneToken) SetReceivedTransfers(v int64)`

SetReceivedTransfers sets ReceivedTransfers field to given value.

### HasReceivedTransfers

`func (o *ResponsesHyperlaneToken) HasReceivedTransfers() bool`

HasReceivedTransfers returns a boolean if a field has been set.

### GetSent

`func (o *ResponsesHyperlaneToken) GetSent() string`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *ResponsesHyperlaneToken) GetSentOk() (*string, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *ResponsesHyperlaneToken) SetSent(v string)`

SetSent sets Sent field to given value.

### HasSent

`func (o *ResponsesHyperlaneToken) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetSentTransfers

`func (o *ResponsesHyperlaneToken) GetSentTransfers() int64`

GetSentTransfers returns the SentTransfers field if non-nil, zero value otherwise.

### GetSentTransfersOk

`func (o *ResponsesHyperlaneToken) GetSentTransfersOk() (*int64, bool)`

GetSentTransfersOk returns a tuple with the SentTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTransfers

`func (o *ResponsesHyperlaneToken) SetSentTransfers(v int64)`

SetSentTransfers sets SentTransfers field to given value.

### HasSentTransfers

`func (o *ResponsesHyperlaneToken) HasSentTransfers() bool`

HasSentTransfers returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesHyperlaneToken) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesHyperlaneToken) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesHyperlaneToken) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesHyperlaneToken) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTokenId

`func (o *ResponsesHyperlaneToken) GetTokenId() *os.File`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ResponsesHyperlaneToken) GetTokenIdOk() (**os.File, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ResponsesHyperlaneToken) SetTokenId(v *os.File)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ResponsesHyperlaneToken) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesHyperlaneToken) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesHyperlaneToken) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesHyperlaneToken) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesHyperlaneToken) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetType

`func (o *ResponsesHyperlaneToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesHyperlaneToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesHyperlaneToken) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesHyperlaneToken) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


