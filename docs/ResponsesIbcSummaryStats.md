# ResponsesIbcSummaryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusiestChannel** | Pointer to [**ResponsesBusiestChannel**](ResponsesBusiestChannel.md) |  | [optional] 
**LargestTransfer** | Pointer to [**ResponsesIbcTransfer**](ResponsesIbcTransfer.md) |  | [optional] 

## Methods

### NewResponsesIbcSummaryStats

`func NewResponsesIbcSummaryStats() *ResponsesIbcSummaryStats`

NewResponsesIbcSummaryStats instantiates a new ResponsesIbcSummaryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesIbcSummaryStatsWithDefaults

`func NewResponsesIbcSummaryStatsWithDefaults() *ResponsesIbcSummaryStats`

NewResponsesIbcSummaryStatsWithDefaults instantiates a new ResponsesIbcSummaryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiestChannel

`func (o *ResponsesIbcSummaryStats) GetBusiestChannel() ResponsesBusiestChannel`

GetBusiestChannel returns the BusiestChannel field if non-nil, zero value otherwise.

### GetBusiestChannelOk

`func (o *ResponsesIbcSummaryStats) GetBusiestChannelOk() (*ResponsesBusiestChannel, bool)`

GetBusiestChannelOk returns a tuple with the BusiestChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiestChannel

`func (o *ResponsesIbcSummaryStats) SetBusiestChannel(v ResponsesBusiestChannel)`

SetBusiestChannel sets BusiestChannel field to given value.

### HasBusiestChannel

`func (o *ResponsesIbcSummaryStats) HasBusiestChannel() bool`

HasBusiestChannel returns a boolean if a field has been set.

### GetLargestTransfer

`func (o *ResponsesIbcSummaryStats) GetLargestTransfer() ResponsesIbcTransfer`

GetLargestTransfer returns the LargestTransfer field if non-nil, zero value otherwise.

### GetLargestTransferOk

`func (o *ResponsesIbcSummaryStats) GetLargestTransferOk() (*ResponsesIbcTransfer, bool)`

GetLargestTransferOk returns a tuple with the LargestTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestTransfer

`func (o *ResponsesIbcSummaryStats) SetLargestTransfer(v ResponsesIbcTransfer)`

SetLargestTransfer sets LargestTransfer field to given value.

### HasLargestTransfer

`func (o *ResponsesIbcSummaryStats) HasLargestTransfer() bool`

HasLargestTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


