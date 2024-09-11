# ResponsesRollupWithStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobsCount** | Pointer to **int32** |  | [optional] 
**BlobsCountPct** | Pointer to **float32** |  | [optional] 
**Bridge** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Explorer** | Pointer to **string** |  | [optional] 
**Fee** | Pointer to **string** |  | [optional] 
**FeePct** | Pointer to **float32** |  | [optional] 
**FirstMessageTime** | Pointer to **time.Time** |  | [optional] 
**Github** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**L2Beat** | Pointer to **string** |  | [optional] 
**LastMessageTime** | Pointer to **time.Time** |  | [optional] 
**Links** | Pointer to **[]string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**SizePct** | Pointer to **float32** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesRollupWithStats

`func NewResponsesRollupWithStats() *ResponsesRollupWithStats`

NewResponsesRollupWithStats instantiates a new ResponsesRollupWithStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesRollupWithStatsWithDefaults

`func NewResponsesRollupWithStatsWithDefaults() *ResponsesRollupWithStats`

NewResponsesRollupWithStatsWithDefaults instantiates a new ResponsesRollupWithStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobsCount

`func (o *ResponsesRollupWithStats) GetBlobsCount() int32`

GetBlobsCount returns the BlobsCount field if non-nil, zero value otherwise.

### GetBlobsCountOk

`func (o *ResponsesRollupWithStats) GetBlobsCountOk() (*int32, bool)`

GetBlobsCountOk returns a tuple with the BlobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsCount

`func (o *ResponsesRollupWithStats) SetBlobsCount(v int32)`

SetBlobsCount sets BlobsCount field to given value.

### HasBlobsCount

`func (o *ResponsesRollupWithStats) HasBlobsCount() bool`

HasBlobsCount returns a boolean if a field has been set.

### GetBlobsCountPct

`func (o *ResponsesRollupWithStats) GetBlobsCountPct() float32`

GetBlobsCountPct returns the BlobsCountPct field if non-nil, zero value otherwise.

### GetBlobsCountPctOk

`func (o *ResponsesRollupWithStats) GetBlobsCountPctOk() (*float32, bool)`

GetBlobsCountPctOk returns a tuple with the BlobsCountPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsCountPct

`func (o *ResponsesRollupWithStats) SetBlobsCountPct(v float32)`

SetBlobsCountPct sets BlobsCountPct field to given value.

### HasBlobsCountPct

`func (o *ResponsesRollupWithStats) HasBlobsCountPct() bool`

HasBlobsCountPct returns a boolean if a field has been set.

### GetBridge

`func (o *ResponsesRollupWithStats) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *ResponsesRollupWithStats) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *ResponsesRollupWithStats) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *ResponsesRollupWithStats) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetDescription

`func (o *ResponsesRollupWithStats) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponsesRollupWithStats) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponsesRollupWithStats) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponsesRollupWithStats) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExplorer

`func (o *ResponsesRollupWithStats) GetExplorer() string`

GetExplorer returns the Explorer field if non-nil, zero value otherwise.

### GetExplorerOk

`func (o *ResponsesRollupWithStats) GetExplorerOk() (*string, bool)`

GetExplorerOk returns a tuple with the Explorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorer

`func (o *ResponsesRollupWithStats) SetExplorer(v string)`

SetExplorer sets Explorer field to given value.

### HasExplorer

`func (o *ResponsesRollupWithStats) HasExplorer() bool`

HasExplorer returns a boolean if a field has been set.

### GetFee

`func (o *ResponsesRollupWithStats) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ResponsesRollupWithStats) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ResponsesRollupWithStats) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *ResponsesRollupWithStats) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeePct

`func (o *ResponsesRollupWithStats) GetFeePct() float32`

GetFeePct returns the FeePct field if non-nil, zero value otherwise.

### GetFeePctOk

`func (o *ResponsesRollupWithStats) GetFeePctOk() (*float32, bool)`

GetFeePctOk returns a tuple with the FeePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePct

`func (o *ResponsesRollupWithStats) SetFeePct(v float32)`

SetFeePct sets FeePct field to given value.

### HasFeePct

`func (o *ResponsesRollupWithStats) HasFeePct() bool`

HasFeePct returns a boolean if a field has been set.

### GetFirstMessageTime

`func (o *ResponsesRollupWithStats) GetFirstMessageTime() time.Time`

GetFirstMessageTime returns the FirstMessageTime field if non-nil, zero value otherwise.

### GetFirstMessageTimeOk

`func (o *ResponsesRollupWithStats) GetFirstMessageTimeOk() (*time.Time, bool)`

GetFirstMessageTimeOk returns a tuple with the FirstMessageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessageTime

`func (o *ResponsesRollupWithStats) SetFirstMessageTime(v time.Time)`

SetFirstMessageTime sets FirstMessageTime field to given value.

### HasFirstMessageTime

`func (o *ResponsesRollupWithStats) HasFirstMessageTime() bool`

HasFirstMessageTime returns a boolean if a field has been set.

### GetGithub

`func (o *ResponsesRollupWithStats) GetGithub() string`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *ResponsesRollupWithStats) GetGithubOk() (*string, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *ResponsesRollupWithStats) SetGithub(v string)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *ResponsesRollupWithStats) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetId

`func (o *ResponsesRollupWithStats) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesRollupWithStats) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesRollupWithStats) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesRollupWithStats) HasId() bool`

HasId returns a boolean if a field has been set.

### GetL2Beat

`func (o *ResponsesRollupWithStats) GetL2Beat() string`

GetL2Beat returns the L2Beat field if non-nil, zero value otherwise.

### GetL2BeatOk

`func (o *ResponsesRollupWithStats) GetL2BeatOk() (*string, bool)`

GetL2BeatOk returns a tuple with the L2Beat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Beat

`func (o *ResponsesRollupWithStats) SetL2Beat(v string)`

SetL2Beat sets L2Beat field to given value.

### HasL2Beat

`func (o *ResponsesRollupWithStats) HasL2Beat() bool`

HasL2Beat returns a boolean if a field has been set.

### GetLastMessageTime

`func (o *ResponsesRollupWithStats) GetLastMessageTime() time.Time`

GetLastMessageTime returns the LastMessageTime field if non-nil, zero value otherwise.

### GetLastMessageTimeOk

`func (o *ResponsesRollupWithStats) GetLastMessageTimeOk() (*time.Time, bool)`

GetLastMessageTimeOk returns a tuple with the LastMessageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageTime

`func (o *ResponsesRollupWithStats) SetLastMessageTime(v time.Time)`

SetLastMessageTime sets LastMessageTime field to given value.

### HasLastMessageTime

`func (o *ResponsesRollupWithStats) HasLastMessageTime() bool`

HasLastMessageTime returns a boolean if a field has been set.

### GetLinks

`func (o *ResponsesRollupWithStats) GetLinks() []string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResponsesRollupWithStats) GetLinksOk() (*[]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResponsesRollupWithStats) SetLinks(v []string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResponsesRollupWithStats) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetLogo

`func (o *ResponsesRollupWithStats) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ResponsesRollupWithStats) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ResponsesRollupWithStats) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ResponsesRollupWithStats) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *ResponsesRollupWithStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsesRollupWithStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsesRollupWithStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponsesRollupWithStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ResponsesRollupWithStats) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponsesRollupWithStats) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponsesRollupWithStats) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponsesRollupWithStats) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizePct

`func (o *ResponsesRollupWithStats) GetSizePct() float32`

GetSizePct returns the SizePct field if non-nil, zero value otherwise.

### GetSizePctOk

`func (o *ResponsesRollupWithStats) GetSizePctOk() (*float32, bool)`

GetSizePctOk returns a tuple with the SizePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizePct

`func (o *ResponsesRollupWithStats) SetSizePct(v float32)`

SetSizePct sets SizePct field to given value.

### HasSizePct

`func (o *ResponsesRollupWithStats) HasSizePct() bool`

HasSizePct returns a boolean if a field has been set.

### GetSlug

`func (o *ResponsesRollupWithStats) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ResponsesRollupWithStats) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ResponsesRollupWithStats) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ResponsesRollupWithStats) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStack

`func (o *ResponsesRollupWithStats) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *ResponsesRollupWithStats) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *ResponsesRollupWithStats) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *ResponsesRollupWithStats) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetTwitter

`func (o *ResponsesRollupWithStats) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *ResponsesRollupWithStats) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *ResponsesRollupWithStats) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *ResponsesRollupWithStats) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetWebsite

`func (o *ResponsesRollupWithStats) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ResponsesRollupWithStats) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ResponsesRollupWithStats) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ResponsesRollupWithStats) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


