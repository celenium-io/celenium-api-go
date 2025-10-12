# ResponsesMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedProposalsCount** | Pointer to **int32** |  | [optional] 
**BlockMissedCount** | Pointer to **int32** |  | [optional] 
**BlockMissedMetric** | Pointer to **string** |  | [optional] 
**CommissionMetric** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**MaxChangeRate** | Pointer to **string** |  | [optional] 
**MaxRate** | Pointer to **string** |  | [optional] 
**Moniker** | Pointer to **string** |  | [optional] 
**OperationTimeMetric** | Pointer to **string** |  | [optional] 
**SelfDelegationAmount** | Pointer to **string** |  | [optional] 
**SelfDelegationMetric** | Pointer to **string** |  | [optional] 
**Stake** | Pointer to **string** |  | [optional] 
**VotesCount** | Pointer to **int32** |  | [optional] 
**VotesMetric** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesMetrics

`func NewResponsesMetrics() *ResponsesMetrics`

NewResponsesMetrics instantiates a new ResponsesMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesMetricsWithDefaults

`func NewResponsesMetricsWithDefaults() *ResponsesMetrics`

NewResponsesMetricsWithDefaults instantiates a new ResponsesMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedProposalsCount

`func (o *ResponsesMetrics) GetAppliedProposalsCount() int32`

GetAppliedProposalsCount returns the AppliedProposalsCount field if non-nil, zero value otherwise.

### GetAppliedProposalsCountOk

`func (o *ResponsesMetrics) GetAppliedProposalsCountOk() (*int32, bool)`

GetAppliedProposalsCountOk returns a tuple with the AppliedProposalsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedProposalsCount

`func (o *ResponsesMetrics) SetAppliedProposalsCount(v int32)`

SetAppliedProposalsCount sets AppliedProposalsCount field to given value.

### HasAppliedProposalsCount

`func (o *ResponsesMetrics) HasAppliedProposalsCount() bool`

HasAppliedProposalsCount returns a boolean if a field has been set.

### GetBlockMissedCount

`func (o *ResponsesMetrics) GetBlockMissedCount() int32`

GetBlockMissedCount returns the BlockMissedCount field if non-nil, zero value otherwise.

### GetBlockMissedCountOk

`func (o *ResponsesMetrics) GetBlockMissedCountOk() (*int32, bool)`

GetBlockMissedCountOk returns a tuple with the BlockMissedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockMissedCount

`func (o *ResponsesMetrics) SetBlockMissedCount(v int32)`

SetBlockMissedCount sets BlockMissedCount field to given value.

### HasBlockMissedCount

`func (o *ResponsesMetrics) HasBlockMissedCount() bool`

HasBlockMissedCount returns a boolean if a field has been set.

### GetBlockMissedMetric

`func (o *ResponsesMetrics) GetBlockMissedMetric() string`

GetBlockMissedMetric returns the BlockMissedMetric field if non-nil, zero value otherwise.

### GetBlockMissedMetricOk

`func (o *ResponsesMetrics) GetBlockMissedMetricOk() (*string, bool)`

GetBlockMissedMetricOk returns a tuple with the BlockMissedMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockMissedMetric

`func (o *ResponsesMetrics) SetBlockMissedMetric(v string)`

SetBlockMissedMetric sets BlockMissedMetric field to given value.

### HasBlockMissedMetric

`func (o *ResponsesMetrics) HasBlockMissedMetric() bool`

HasBlockMissedMetric returns a boolean if a field has been set.

### GetCommissionMetric

`func (o *ResponsesMetrics) GetCommissionMetric() string`

GetCommissionMetric returns the CommissionMetric field if non-nil, zero value otherwise.

### GetCommissionMetricOk

`func (o *ResponsesMetrics) GetCommissionMetricOk() (*string, bool)`

GetCommissionMetricOk returns a tuple with the CommissionMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionMetric

`func (o *ResponsesMetrics) SetCommissionMetric(v string)`

SetCommissionMetric sets CommissionMetric field to given value.

### HasCommissionMetric

`func (o *ResponsesMetrics) HasCommissionMetric() bool`

HasCommissionMetric returns a boolean if a field has been set.

### GetCreationTime

`func (o *ResponsesMetrics) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ResponsesMetrics) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ResponsesMetrics) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ResponsesMetrics) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *ResponsesMetrics) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesMetrics) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesMetrics) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesMetrics) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxChangeRate

`func (o *ResponsesMetrics) GetMaxChangeRate() string`

GetMaxChangeRate returns the MaxChangeRate field if non-nil, zero value otherwise.

### GetMaxChangeRateOk

`func (o *ResponsesMetrics) GetMaxChangeRateOk() (*string, bool)`

GetMaxChangeRateOk returns a tuple with the MaxChangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChangeRate

`func (o *ResponsesMetrics) SetMaxChangeRate(v string)`

SetMaxChangeRate sets MaxChangeRate field to given value.

### HasMaxChangeRate

`func (o *ResponsesMetrics) HasMaxChangeRate() bool`

HasMaxChangeRate returns a boolean if a field has been set.

### GetMaxRate

`func (o *ResponsesMetrics) GetMaxRate() string`

GetMaxRate returns the MaxRate field if non-nil, zero value otherwise.

### GetMaxRateOk

`func (o *ResponsesMetrics) GetMaxRateOk() (*string, bool)`

GetMaxRateOk returns a tuple with the MaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRate

`func (o *ResponsesMetrics) SetMaxRate(v string)`

SetMaxRate sets MaxRate field to given value.

### HasMaxRate

`func (o *ResponsesMetrics) HasMaxRate() bool`

HasMaxRate returns a boolean if a field has been set.

### GetMoniker

`func (o *ResponsesMetrics) GetMoniker() string`

GetMoniker returns the Moniker field if non-nil, zero value otherwise.

### GetMonikerOk

`func (o *ResponsesMetrics) GetMonikerOk() (*string, bool)`

GetMonikerOk returns a tuple with the Moniker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoniker

`func (o *ResponsesMetrics) SetMoniker(v string)`

SetMoniker sets Moniker field to given value.

### HasMoniker

`func (o *ResponsesMetrics) HasMoniker() bool`

HasMoniker returns a boolean if a field has been set.

### GetOperationTimeMetric

`func (o *ResponsesMetrics) GetOperationTimeMetric() string`

GetOperationTimeMetric returns the OperationTimeMetric field if non-nil, zero value otherwise.

### GetOperationTimeMetricOk

`func (o *ResponsesMetrics) GetOperationTimeMetricOk() (*string, bool)`

GetOperationTimeMetricOk returns a tuple with the OperationTimeMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTimeMetric

`func (o *ResponsesMetrics) SetOperationTimeMetric(v string)`

SetOperationTimeMetric sets OperationTimeMetric field to given value.

### HasOperationTimeMetric

`func (o *ResponsesMetrics) HasOperationTimeMetric() bool`

HasOperationTimeMetric returns a boolean if a field has been set.

### GetSelfDelegationAmount

`func (o *ResponsesMetrics) GetSelfDelegationAmount() string`

GetSelfDelegationAmount returns the SelfDelegationAmount field if non-nil, zero value otherwise.

### GetSelfDelegationAmountOk

`func (o *ResponsesMetrics) GetSelfDelegationAmountOk() (*string, bool)`

GetSelfDelegationAmountOk returns a tuple with the SelfDelegationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfDelegationAmount

`func (o *ResponsesMetrics) SetSelfDelegationAmount(v string)`

SetSelfDelegationAmount sets SelfDelegationAmount field to given value.

### HasSelfDelegationAmount

`func (o *ResponsesMetrics) HasSelfDelegationAmount() bool`

HasSelfDelegationAmount returns a boolean if a field has been set.

### GetSelfDelegationMetric

`func (o *ResponsesMetrics) GetSelfDelegationMetric() string`

GetSelfDelegationMetric returns the SelfDelegationMetric field if non-nil, zero value otherwise.

### GetSelfDelegationMetricOk

`func (o *ResponsesMetrics) GetSelfDelegationMetricOk() (*string, bool)`

GetSelfDelegationMetricOk returns a tuple with the SelfDelegationMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfDelegationMetric

`func (o *ResponsesMetrics) SetSelfDelegationMetric(v string)`

SetSelfDelegationMetric sets SelfDelegationMetric field to given value.

### HasSelfDelegationMetric

`func (o *ResponsesMetrics) HasSelfDelegationMetric() bool`

HasSelfDelegationMetric returns a boolean if a field has been set.

### GetStake

`func (o *ResponsesMetrics) GetStake() string`

GetStake returns the Stake field if non-nil, zero value otherwise.

### GetStakeOk

`func (o *ResponsesMetrics) GetStakeOk() (*string, bool)`

GetStakeOk returns a tuple with the Stake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStake

`func (o *ResponsesMetrics) SetStake(v string)`

SetStake sets Stake field to given value.

### HasStake

`func (o *ResponsesMetrics) HasStake() bool`

HasStake returns a boolean if a field has been set.

### GetVotesCount

`func (o *ResponsesMetrics) GetVotesCount() int32`

GetVotesCount returns the VotesCount field if non-nil, zero value otherwise.

### GetVotesCountOk

`func (o *ResponsesMetrics) GetVotesCountOk() (*int32, bool)`

GetVotesCountOk returns a tuple with the VotesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesCount

`func (o *ResponsesMetrics) SetVotesCount(v int32)`

SetVotesCount sets VotesCount field to given value.

### HasVotesCount

`func (o *ResponsesMetrics) HasVotesCount() bool`

HasVotesCount returns a boolean if a field has been set.

### GetVotesMetric

`func (o *ResponsesMetrics) GetVotesMetric() string`

GetVotesMetric returns the VotesMetric field if non-nil, zero value otherwise.

### GetVotesMetricOk

`func (o *ResponsesMetrics) GetVotesMetricOk() (*string, bool)`

GetVotesMetricOk returns a tuple with the VotesMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesMetric

`func (o *ResponsesMetrics) SetVotesMetric(v string)`

SetVotesMetric sets VotesMetric field to given value.

### HasVotesMetric

`func (o *ResponsesMetrics) HasVotesMetric() bool`

HasVotesMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


