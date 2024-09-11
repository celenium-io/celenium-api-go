# ResponsesSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **map[string]interface{}** | Search result. Can be one of folowwing types: Block, Address, Namespace, Tx, Validator, Rollup | [optional] 
**Type** | Pointer to **string** | Result type which is in the result. Can be &#39;block&#39;, &#39;address&#39;, &#39;namespace&#39;, &#39;tx&#39;, &#39;validator&#39;, &#39;rollup&#39; | [optional] 

## Methods

### NewResponsesSearchItem

`func NewResponsesSearchItem() *ResponsesSearchItem`

NewResponsesSearchItem instantiates a new ResponsesSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesSearchItemWithDefaults

`func NewResponsesSearchItemWithDefaults() *ResponsesSearchItem`

NewResponsesSearchItemWithDefaults instantiates a new ResponsesSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ResponsesSearchItem) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResponsesSearchItem) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResponsesSearchItem) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *ResponsesSearchItem) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetType

`func (o *ResponsesSearchItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesSearchItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesSearchItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesSearchItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


