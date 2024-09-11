# ResponsesTimeValueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesTimeValueItem

`func NewResponsesTimeValueItem() *ResponsesTimeValueItem`

NewResponsesTimeValueItem instantiates a new ResponsesTimeValueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesTimeValueItemWithDefaults

`func NewResponsesTimeValueItemWithDefaults() *ResponsesTimeValueItem`

NewResponsesTimeValueItemWithDefaults instantiates a new ResponsesTimeValueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ResponsesTimeValueItem) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesTimeValueItem) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesTimeValueItem) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesTimeValueItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValue

`func (o *ResponsesTimeValueItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponsesTimeValueItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponsesTimeValueItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResponsesTimeValueItem) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


