# ResponsesZkISM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creator** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Groth16Vkey** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MerkleTreeAddress** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateMembershipVkey** | Pointer to **string** |  | [optional] 
**StateTransitionVkey** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxHash** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesZkISM

`func NewResponsesZkISM() *ResponsesZkISM`

NewResponsesZkISM instantiates a new ResponsesZkISM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesZkISMWithDefaults

`func NewResponsesZkISMWithDefaults() *ResponsesZkISM`

NewResponsesZkISMWithDefaults instantiates a new ResponsesZkISM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreator

`func (o *ResponsesZkISM) GetCreator() ResponsesShortAddress`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ResponsesZkISM) GetCreatorOk() (*ResponsesShortAddress, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ResponsesZkISM) SetCreator(v ResponsesShortAddress)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ResponsesZkISM) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetExternalId

`func (o *ResponsesZkISM) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ResponsesZkISM) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ResponsesZkISM) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ResponsesZkISM) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGroth16Vkey

`func (o *ResponsesZkISM) GetGroth16Vkey() string`

GetGroth16Vkey returns the Groth16Vkey field if non-nil, zero value otherwise.

### GetGroth16VkeyOk

`func (o *ResponsesZkISM) GetGroth16VkeyOk() (*string, bool)`

GetGroth16VkeyOk returns a tuple with the Groth16Vkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroth16Vkey

`func (o *ResponsesZkISM) SetGroth16Vkey(v string)`

SetGroth16Vkey sets Groth16Vkey field to given value.

### HasGroth16Vkey

`func (o *ResponsesZkISM) HasGroth16Vkey() bool`

HasGroth16Vkey returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesZkISM) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesZkISM) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesZkISM) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesZkISM) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesZkISM) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesZkISM) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesZkISM) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesZkISM) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerkleTreeAddress

`func (o *ResponsesZkISM) GetMerkleTreeAddress() string`

GetMerkleTreeAddress returns the MerkleTreeAddress field if non-nil, zero value otherwise.

### GetMerkleTreeAddressOk

`func (o *ResponsesZkISM) GetMerkleTreeAddressOk() (*string, bool)`

GetMerkleTreeAddressOk returns a tuple with the MerkleTreeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleTreeAddress

`func (o *ResponsesZkISM) SetMerkleTreeAddress(v string)`

SetMerkleTreeAddress sets MerkleTreeAddress field to given value.

### HasMerkleTreeAddress

`func (o *ResponsesZkISM) HasMerkleTreeAddress() bool`

HasMerkleTreeAddress returns a boolean if a field has been set.

### GetState

`func (o *ResponsesZkISM) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponsesZkISM) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponsesZkISM) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponsesZkISM) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMembershipVkey

`func (o *ResponsesZkISM) GetStateMembershipVkey() string`

GetStateMembershipVkey returns the StateMembershipVkey field if non-nil, zero value otherwise.

### GetStateMembershipVkeyOk

`func (o *ResponsesZkISM) GetStateMembershipVkeyOk() (*string, bool)`

GetStateMembershipVkeyOk returns a tuple with the StateMembershipVkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMembershipVkey

`func (o *ResponsesZkISM) SetStateMembershipVkey(v string)`

SetStateMembershipVkey sets StateMembershipVkey field to given value.

### HasStateMembershipVkey

`func (o *ResponsesZkISM) HasStateMembershipVkey() bool`

HasStateMembershipVkey returns a boolean if a field has been set.

### GetStateTransitionVkey

`func (o *ResponsesZkISM) GetStateTransitionVkey() string`

GetStateTransitionVkey returns the StateTransitionVkey field if non-nil, zero value otherwise.

### GetStateTransitionVkeyOk

`func (o *ResponsesZkISM) GetStateTransitionVkeyOk() (*string, bool)`

GetStateTransitionVkeyOk returns a tuple with the StateTransitionVkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTransitionVkey

`func (o *ResponsesZkISM) SetStateTransitionVkey(v string)`

SetStateTransitionVkey sets StateTransitionVkey field to given value.

### HasStateTransitionVkey

`func (o *ResponsesZkISM) HasStateTransitionVkey() bool`

HasStateTransitionVkey returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesZkISM) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesZkISM) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesZkISM) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesZkISM) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesZkISM) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesZkISM) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesZkISM) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesZkISM) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


