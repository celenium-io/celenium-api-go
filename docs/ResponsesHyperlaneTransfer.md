# ResponsesHyperlaneTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Counterparty** | Pointer to [**ResponsesHyperlaneCounterparty**](ResponsesHyperlaneCounterparty.md) |  | [optional] 
**Denom** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Mailbox** | Pointer to ***os.File** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **int64** |  | [optional] 
**Received** | Pointer to **string** |  | [optional] 
**Relayer** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TokenId** | Pointer to ***os.File** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewResponsesHyperlaneTransfer

`func NewResponsesHyperlaneTransfer() *ResponsesHyperlaneTransfer`

NewResponsesHyperlaneTransfer instantiates a new ResponsesHyperlaneTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesHyperlaneTransferWithDefaults

`func NewResponsesHyperlaneTransferWithDefaults() *ResponsesHyperlaneTransfer`

NewResponsesHyperlaneTransferWithDefaults instantiates a new ResponsesHyperlaneTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ResponsesHyperlaneTransfer) GetAddress() ResponsesShortAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ResponsesHyperlaneTransfer) GetAddressOk() (*ResponsesShortAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ResponsesHyperlaneTransfer) SetAddress(v ResponsesShortAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ResponsesHyperlaneTransfer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBody

`func (o *ResponsesHyperlaneTransfer) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ResponsesHyperlaneTransfer) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ResponsesHyperlaneTransfer) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ResponsesHyperlaneTransfer) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCounterparty

`func (o *ResponsesHyperlaneTransfer) GetCounterparty() ResponsesHyperlaneCounterparty`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *ResponsesHyperlaneTransfer) GetCounterpartyOk() (*ResponsesHyperlaneCounterparty, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *ResponsesHyperlaneTransfer) SetCounterparty(v ResponsesHyperlaneCounterparty)`

SetCounterparty sets Counterparty field to given value.

### HasCounterparty

`func (o *ResponsesHyperlaneTransfer) HasCounterparty() bool`

HasCounterparty returns a boolean if a field has been set.

### GetDenom

`func (o *ResponsesHyperlaneTransfer) GetDenom() string`

GetDenom returns the Denom field if non-nil, zero value otherwise.

### GetDenomOk

`func (o *ResponsesHyperlaneTransfer) GetDenomOk() (*string, bool)`

GetDenomOk returns a tuple with the Denom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenom

`func (o *ResponsesHyperlaneTransfer) SetDenom(v string)`

SetDenom sets Denom field to given value.

### HasDenom

`func (o *ResponsesHyperlaneTransfer) HasDenom() bool`

HasDenom returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesHyperlaneTransfer) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesHyperlaneTransfer) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesHyperlaneTransfer) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesHyperlaneTransfer) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesHyperlaneTransfer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesHyperlaneTransfer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesHyperlaneTransfer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesHyperlaneTransfer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMailbox

`func (o *ResponsesHyperlaneTransfer) GetMailbox() *os.File`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *ResponsesHyperlaneTransfer) GetMailboxOk() (**os.File, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *ResponsesHyperlaneTransfer) SetMailbox(v *os.File)`

SetMailbox sets Mailbox field to given value.

### HasMailbox

`func (o *ResponsesHyperlaneTransfer) HasMailbox() bool`

HasMailbox returns a boolean if a field has been set.

### GetMetadata

`func (o *ResponsesHyperlaneTransfer) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResponsesHyperlaneTransfer) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResponsesHyperlaneTransfer) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResponsesHyperlaneTransfer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNonce

`func (o *ResponsesHyperlaneTransfer) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ResponsesHyperlaneTransfer) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ResponsesHyperlaneTransfer) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *ResponsesHyperlaneTransfer) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetReceived

`func (o *ResponsesHyperlaneTransfer) GetReceived() string`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *ResponsesHyperlaneTransfer) GetReceivedOk() (*string, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *ResponsesHyperlaneTransfer) SetReceived(v string)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *ResponsesHyperlaneTransfer) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetRelayer

`func (o *ResponsesHyperlaneTransfer) GetRelayer() ResponsesShortAddress`

GetRelayer returns the Relayer field if non-nil, zero value otherwise.

### GetRelayerOk

`func (o *ResponsesHyperlaneTransfer) GetRelayerOk() (*ResponsesShortAddress, bool)`

GetRelayerOk returns a tuple with the Relayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayer

`func (o *ResponsesHyperlaneTransfer) SetRelayer(v ResponsesShortAddress)`

SetRelayer sets Relayer field to given value.

### HasRelayer

`func (o *ResponsesHyperlaneTransfer) HasRelayer() bool`

HasRelayer returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesHyperlaneTransfer) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesHyperlaneTransfer) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesHyperlaneTransfer) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesHyperlaneTransfer) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTokenId

`func (o *ResponsesHyperlaneTransfer) GetTokenId() *os.File`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ResponsesHyperlaneTransfer) GetTokenIdOk() (**os.File, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ResponsesHyperlaneTransfer) SetTokenId(v *os.File)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ResponsesHyperlaneTransfer) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesHyperlaneTransfer) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesHyperlaneTransfer) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesHyperlaneTransfer) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesHyperlaneTransfer) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetType

`func (o *ResponsesHyperlaneTransfer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesHyperlaneTransfer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesHyperlaneTransfer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesHyperlaneTransfer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *ResponsesHyperlaneTransfer) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponsesHyperlaneTransfer) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponsesHyperlaneTransfer) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponsesHyperlaneTransfer) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


