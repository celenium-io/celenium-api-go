# ResponsesBlobLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commitment** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Namespace** | Pointer to [**ResponsesNamespace**](ResponsesNamespace.md) |  | [optional] 
**Rollup** | Pointer to [**ResponsesShortRollup**](ResponsesShortRollup.md) |  | [optional] 
**Signer** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Tx** | Pointer to [**ResponsesTx**](ResponsesTx.md) |  | [optional] 

## Methods

### NewResponsesBlobLog

`func NewResponsesBlobLog() *ResponsesBlobLog`

NewResponsesBlobLog instantiates a new ResponsesBlobLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesBlobLogWithDefaults

`func NewResponsesBlobLogWithDefaults() *ResponsesBlobLog`

NewResponsesBlobLogWithDefaults instantiates a new ResponsesBlobLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitment

`func (o *ResponsesBlobLog) GetCommitment() string`

GetCommitment returns the Commitment field if non-nil, zero value otherwise.

### GetCommitmentOk

`func (o *ResponsesBlobLog) GetCommitmentOk() (*string, bool)`

GetCommitmentOk returns a tuple with the Commitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitment

`func (o *ResponsesBlobLog) SetCommitment(v string)`

SetCommitment sets Commitment field to given value.

### HasCommitment

`func (o *ResponsesBlobLog) HasCommitment() bool`

HasCommitment returns a boolean if a field has been set.

### GetContentType

`func (o *ResponsesBlobLog) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ResponsesBlobLog) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ResponsesBlobLog) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ResponsesBlobLog) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesBlobLog) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesBlobLog) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesBlobLog) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesBlobLog) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesBlobLog) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesBlobLog) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesBlobLog) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesBlobLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNamespace

`func (o *ResponsesBlobLog) GetNamespace() ResponsesNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ResponsesBlobLog) GetNamespaceOk() (*ResponsesNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ResponsesBlobLog) SetNamespace(v ResponsesNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ResponsesBlobLog) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRollup

`func (o *ResponsesBlobLog) GetRollup() ResponsesShortRollup`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *ResponsesBlobLog) GetRollupOk() (*ResponsesShortRollup, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *ResponsesBlobLog) SetRollup(v ResponsesShortRollup)`

SetRollup sets Rollup field to given value.

### HasRollup

`func (o *ResponsesBlobLog) HasRollup() bool`

HasRollup returns a boolean if a field has been set.

### GetSigner

`func (o *ResponsesBlobLog) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *ResponsesBlobLog) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *ResponsesBlobLog) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *ResponsesBlobLog) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetSize

`func (o *ResponsesBlobLog) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponsesBlobLog) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponsesBlobLog) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponsesBlobLog) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesBlobLog) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesBlobLog) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesBlobLog) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesBlobLog) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTx

`func (o *ResponsesBlobLog) GetTx() ResponsesTx`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *ResponsesBlobLog) GetTxOk() (*ResponsesTx, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *ResponsesBlobLog) SetTx(v ResponsesTx)`

SetTx sets Tx field to given value.

### HasTx

`func (o *ResponsesBlobLog) HasTx() bool`

HasTx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


