# ResponsesHyperlaneMailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultHook** | Pointer to ***os.File** |  | [optional] 
**DefaultIsm** | Pointer to ***os.File** |  | [optional] 
**Domain** | Pointer to **int64** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Mailbox** | Pointer to ***os.File** |  | [optional] 
**Owner** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**ReceivedMessages** | Pointer to **int64** |  | [optional] 
**RequiredHook** | Pointer to ***os.File** |  | [optional] 
**SentMessages** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewResponsesHyperlaneMailbox

`func NewResponsesHyperlaneMailbox() *ResponsesHyperlaneMailbox`

NewResponsesHyperlaneMailbox instantiates a new ResponsesHyperlaneMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesHyperlaneMailboxWithDefaults

`func NewResponsesHyperlaneMailboxWithDefaults() *ResponsesHyperlaneMailbox`

NewResponsesHyperlaneMailboxWithDefaults instantiates a new ResponsesHyperlaneMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultHook

`func (o *ResponsesHyperlaneMailbox) GetDefaultHook() *os.File`

GetDefaultHook returns the DefaultHook field if non-nil, zero value otherwise.

### GetDefaultHookOk

`func (o *ResponsesHyperlaneMailbox) GetDefaultHookOk() (**os.File, bool)`

GetDefaultHookOk returns a tuple with the DefaultHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHook

`func (o *ResponsesHyperlaneMailbox) SetDefaultHook(v *os.File)`

SetDefaultHook sets DefaultHook field to given value.

### HasDefaultHook

`func (o *ResponsesHyperlaneMailbox) HasDefaultHook() bool`

HasDefaultHook returns a boolean if a field has been set.

### GetDefaultIsm

`func (o *ResponsesHyperlaneMailbox) GetDefaultIsm() *os.File`

GetDefaultIsm returns the DefaultIsm field if non-nil, zero value otherwise.

### GetDefaultIsmOk

`func (o *ResponsesHyperlaneMailbox) GetDefaultIsmOk() (**os.File, bool)`

GetDefaultIsmOk returns a tuple with the DefaultIsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIsm

`func (o *ResponsesHyperlaneMailbox) SetDefaultIsm(v *os.File)`

SetDefaultIsm sets DefaultIsm field to given value.

### HasDefaultIsm

`func (o *ResponsesHyperlaneMailbox) HasDefaultIsm() bool`

HasDefaultIsm returns a boolean if a field has been set.

### GetDomain

`func (o *ResponsesHyperlaneMailbox) GetDomain() int64`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResponsesHyperlaneMailbox) GetDomainOk() (*int64, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResponsesHyperlaneMailbox) SetDomain(v int64)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ResponsesHyperlaneMailbox) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesHyperlaneMailbox) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesHyperlaneMailbox) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesHyperlaneMailbox) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesHyperlaneMailbox) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesHyperlaneMailbox) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesHyperlaneMailbox) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesHyperlaneMailbox) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesHyperlaneMailbox) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMailbox

`func (o *ResponsesHyperlaneMailbox) GetMailbox() *os.File`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *ResponsesHyperlaneMailbox) GetMailboxOk() (**os.File, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *ResponsesHyperlaneMailbox) SetMailbox(v *os.File)`

SetMailbox sets Mailbox field to given value.

### HasMailbox

`func (o *ResponsesHyperlaneMailbox) HasMailbox() bool`

HasMailbox returns a boolean if a field has been set.

### GetOwner

`func (o *ResponsesHyperlaneMailbox) GetOwner() ResponsesShortAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResponsesHyperlaneMailbox) GetOwnerOk() (*ResponsesShortAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResponsesHyperlaneMailbox) SetOwner(v ResponsesShortAddress)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ResponsesHyperlaneMailbox) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetReceivedMessages

`func (o *ResponsesHyperlaneMailbox) GetReceivedMessages() int64`

GetReceivedMessages returns the ReceivedMessages field if non-nil, zero value otherwise.

### GetReceivedMessagesOk

`func (o *ResponsesHyperlaneMailbox) GetReceivedMessagesOk() (*int64, bool)`

GetReceivedMessagesOk returns a tuple with the ReceivedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedMessages

`func (o *ResponsesHyperlaneMailbox) SetReceivedMessages(v int64)`

SetReceivedMessages sets ReceivedMessages field to given value.

### HasReceivedMessages

`func (o *ResponsesHyperlaneMailbox) HasReceivedMessages() bool`

HasReceivedMessages returns a boolean if a field has been set.

### GetRequiredHook

`func (o *ResponsesHyperlaneMailbox) GetRequiredHook() *os.File`

GetRequiredHook returns the RequiredHook field if non-nil, zero value otherwise.

### GetRequiredHookOk

`func (o *ResponsesHyperlaneMailbox) GetRequiredHookOk() (**os.File, bool)`

GetRequiredHookOk returns a tuple with the RequiredHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredHook

`func (o *ResponsesHyperlaneMailbox) SetRequiredHook(v *os.File)`

SetRequiredHook sets RequiredHook field to given value.

### HasRequiredHook

`func (o *ResponsesHyperlaneMailbox) HasRequiredHook() bool`

HasRequiredHook returns a boolean if a field has been set.

### GetSentMessages

`func (o *ResponsesHyperlaneMailbox) GetSentMessages() int64`

GetSentMessages returns the SentMessages field if non-nil, zero value otherwise.

### GetSentMessagesOk

`func (o *ResponsesHyperlaneMailbox) GetSentMessagesOk() (*int64, bool)`

GetSentMessagesOk returns a tuple with the SentMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentMessages

`func (o *ResponsesHyperlaneMailbox) SetSentMessages(v int64)`

SetSentMessages sets SentMessages field to given value.

### HasSentMessages

`func (o *ResponsesHyperlaneMailbox) HasSentMessages() bool`

HasSentMessages returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesHyperlaneMailbox) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesHyperlaneMailbox) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesHyperlaneMailbox) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesHyperlaneMailbox) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesHyperlaneMailbox) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesHyperlaneMailbox) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesHyperlaneMailbox) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesHyperlaneMailbox) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


