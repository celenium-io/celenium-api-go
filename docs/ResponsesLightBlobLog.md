# ResponsesLightBlobLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commitment** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Signer** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewResponsesLightBlobLog

`func NewResponsesLightBlobLog() *ResponsesLightBlobLog`

NewResponsesLightBlobLog instantiates a new ResponsesLightBlobLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesLightBlobLogWithDefaults

`func NewResponsesLightBlobLogWithDefaults() *ResponsesLightBlobLog`

NewResponsesLightBlobLogWithDefaults instantiates a new ResponsesLightBlobLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitment

`func (o *ResponsesLightBlobLog) GetCommitment() string`

GetCommitment returns the Commitment field if non-nil, zero value otherwise.

### GetCommitmentOk

`func (o *ResponsesLightBlobLog) GetCommitmentOk() (*string, bool)`

GetCommitmentOk returns a tuple with the Commitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitment

`func (o *ResponsesLightBlobLog) SetCommitment(v string)`

SetCommitment sets Commitment field to given value.

### HasCommitment

`func (o *ResponsesLightBlobLog) HasCommitment() bool`

HasCommitment returns a boolean if a field has been set.

### GetContentType

`func (o *ResponsesLightBlobLog) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ResponsesLightBlobLog) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ResponsesLightBlobLog) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ResponsesLightBlobLog) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesLightBlobLog) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesLightBlobLog) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesLightBlobLog) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesLightBlobLog) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesLightBlobLog) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesLightBlobLog) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesLightBlobLog) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesLightBlobLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNamespace

`func (o *ResponsesLightBlobLog) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ResponsesLightBlobLog) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ResponsesLightBlobLog) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ResponsesLightBlobLog) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSigner

`func (o *ResponsesLightBlobLog) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *ResponsesLightBlobLog) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *ResponsesLightBlobLog) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *ResponsesLightBlobLog) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetSize

`func (o *ResponsesLightBlobLog) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponsesLightBlobLog) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponsesLightBlobLog) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponsesLightBlobLog) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesLightBlobLog) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesLightBlobLog) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesLightBlobLog) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesLightBlobLog) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesLightBlobLog) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesLightBlobLog) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesLightBlobLog) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesLightBlobLog) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


