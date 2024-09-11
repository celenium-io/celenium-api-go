# ResponsesNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobsCount** | Pointer to **int32** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastHeight** | Pointer to **int64** |  | [optional] 
**LastMessageTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NamespaceId** | Pointer to ***os.File** |  | [optional] 
**PfbCount** | Pointer to **int32** |  | [optional] 
**Reserved** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponsesNamespace

`func NewResponsesNamespace() *ResponsesNamespace`

NewResponsesNamespace instantiates a new ResponsesNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesNamespaceWithDefaults

`func NewResponsesNamespaceWithDefaults() *ResponsesNamespace`

NewResponsesNamespaceWithDefaults instantiates a new ResponsesNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobsCount

`func (o *ResponsesNamespace) GetBlobsCount() int32`

GetBlobsCount returns the BlobsCount field if non-nil, zero value otherwise.

### GetBlobsCountOk

`func (o *ResponsesNamespace) GetBlobsCountOk() (*int32, bool)`

GetBlobsCountOk returns a tuple with the BlobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsCount

`func (o *ResponsesNamespace) SetBlobsCount(v int32)`

SetBlobsCount sets BlobsCount field to given value.

### HasBlobsCount

`func (o *ResponsesNamespace) HasBlobsCount() bool`

HasBlobsCount returns a boolean if a field has been set.

### GetHash

`func (o *ResponsesNamespace) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ResponsesNamespace) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ResponsesNamespace) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ResponsesNamespace) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ResponsesNamespace) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesNamespace) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesNamespace) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesNamespace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastHeight

`func (o *ResponsesNamespace) GetLastHeight() int64`

GetLastHeight returns the LastHeight field if non-nil, zero value otherwise.

### GetLastHeightOk

`func (o *ResponsesNamespace) GetLastHeightOk() (*int64, bool)`

GetLastHeightOk returns a tuple with the LastHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeight

`func (o *ResponsesNamespace) SetLastHeight(v int64)`

SetLastHeight sets LastHeight field to given value.

### HasLastHeight

`func (o *ResponsesNamespace) HasLastHeight() bool`

HasLastHeight returns a boolean if a field has been set.

### GetLastMessageTime

`func (o *ResponsesNamespace) GetLastMessageTime() time.Time`

GetLastMessageTime returns the LastMessageTime field if non-nil, zero value otherwise.

### GetLastMessageTimeOk

`func (o *ResponsesNamespace) GetLastMessageTimeOk() (*time.Time, bool)`

GetLastMessageTimeOk returns a tuple with the LastMessageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageTime

`func (o *ResponsesNamespace) SetLastMessageTime(v time.Time)`

SetLastMessageTime sets LastMessageTime field to given value.

### HasLastMessageTime

`func (o *ResponsesNamespace) HasLastMessageTime() bool`

HasLastMessageTime returns a boolean if a field has been set.

### GetName

`func (o *ResponsesNamespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsesNamespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsesNamespace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponsesNamespace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceId

`func (o *ResponsesNamespace) GetNamespaceId() *os.File`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *ResponsesNamespace) GetNamespaceIdOk() (**os.File, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *ResponsesNamespace) SetNamespaceId(v *os.File)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *ResponsesNamespace) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetPfbCount

`func (o *ResponsesNamespace) GetPfbCount() int32`

GetPfbCount returns the PfbCount field if non-nil, zero value otherwise.

### GetPfbCountOk

`func (o *ResponsesNamespace) GetPfbCountOk() (*int32, bool)`

GetPfbCountOk returns a tuple with the PfbCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfbCount

`func (o *ResponsesNamespace) SetPfbCount(v int32)`

SetPfbCount sets PfbCount field to given value.

### HasPfbCount

`func (o *ResponsesNamespace) HasPfbCount() bool`

HasPfbCount returns a boolean if a field has been set.

### GetReserved

`func (o *ResponsesNamespace) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *ResponsesNamespace) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *ResponsesNamespace) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *ResponsesNamespace) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetSize

`func (o *ResponsesNamespace) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponsesNamespace) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponsesNamespace) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponsesNamespace) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVersion

`func (o *ResponsesNamespace) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponsesNamespace) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponsesNamespace) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponsesNamespace) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


