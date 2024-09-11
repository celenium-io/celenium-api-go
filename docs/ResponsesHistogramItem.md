# ResponsesHistogramItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesHistogramItem

`func NewResponsesHistogramItem() *ResponsesHistogramItem`

NewResponsesHistogramItem instantiates a new ResponsesHistogramItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesHistogramItemWithDefaults

`func NewResponsesHistogramItemWithDefaults() *ResponsesHistogramItem`

NewResponsesHistogramItemWithDefaults instantiates a new ResponsesHistogramItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ResponsesHistogramItem) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesHistogramItem) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesHistogramItem) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesHistogramItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValue

`func (o *ResponsesHistogramItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponsesHistogramItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponsesHistogramItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResponsesHistogramItem) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


