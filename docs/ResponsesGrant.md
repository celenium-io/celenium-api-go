# ResponsesGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **string** |  | [optional] 
**Grantee** | Pointer to **string** |  | [optional] 
**Granter** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Params** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**RevokeHeight** | Pointer to **int32** |  | [optional] 
**Revoked** | Pointer to **bool** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesGrant

`func NewResponsesGrant() *ResponsesGrant`

NewResponsesGrant instantiates a new ResponsesGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesGrantWithDefaults

`func NewResponsesGrantWithDefaults() *ResponsesGrant`

NewResponsesGrantWithDefaults instantiates a new ResponsesGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *ResponsesGrant) GetAuthorization() string`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *ResponsesGrant) GetAuthorizationOk() (*string, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *ResponsesGrant) SetAuthorization(v string)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *ResponsesGrant) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetExpiration

`func (o *ResponsesGrant) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ResponsesGrant) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ResponsesGrant) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ResponsesGrant) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetGrantee

`func (o *ResponsesGrant) GetGrantee() string`

GetGrantee returns the Grantee field if non-nil, zero value otherwise.

### GetGranteeOk

`func (o *ResponsesGrant) GetGranteeOk() (*string, bool)`

GetGranteeOk returns a tuple with the Grantee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantee

`func (o *ResponsesGrant) SetGrantee(v string)`

SetGrantee sets Grantee field to given value.

### HasGrantee

`func (o *ResponsesGrant) HasGrantee() bool`

HasGrantee returns a boolean if a field has been set.

### GetGranter

`func (o *ResponsesGrant) GetGranter() string`

GetGranter returns the Granter field if non-nil, zero value otherwise.

### GetGranterOk

`func (o *ResponsesGrant) GetGranterOk() (*string, bool)`

GetGranterOk returns a tuple with the Granter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranter

`func (o *ResponsesGrant) SetGranter(v string)`

SetGranter sets Granter field to given value.

### HasGranter

`func (o *ResponsesGrant) HasGranter() bool`

HasGranter returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesGrant) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesGrant) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesGrant) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesGrant) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetParams

`func (o *ResponsesGrant) GetParams() map[string]map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ResponsesGrant) GetParamsOk() (*map[string]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ResponsesGrant) SetParams(v map[string]map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ResponsesGrant) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetRevokeHeight

`func (o *ResponsesGrant) GetRevokeHeight() int32`

GetRevokeHeight returns the RevokeHeight field if non-nil, zero value otherwise.

### GetRevokeHeightOk

`func (o *ResponsesGrant) GetRevokeHeightOk() (*int32, bool)`

GetRevokeHeightOk returns a tuple with the RevokeHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeHeight

`func (o *ResponsesGrant) SetRevokeHeight(v int32)`

SetRevokeHeight sets RevokeHeight field to given value.

### HasRevokeHeight

`func (o *ResponsesGrant) HasRevokeHeight() bool`

HasRevokeHeight returns a boolean if a field has been set.

### GetRevoked

`func (o *ResponsesGrant) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *ResponsesGrant) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *ResponsesGrant) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *ResponsesGrant) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesGrant) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesGrant) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesGrant) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesGrant) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


