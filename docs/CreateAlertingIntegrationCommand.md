# CreateAlertingIntegrationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Token** | Pointer to **NullableString** |  | [optional] 
**AlertingIntegrationType** | [**AlertingIntegrationType**](AlertingIntegrationType.md) |  | 
**AlertingProfileId** | **int32** |  | 

## Methods

### NewCreateAlertingIntegrationCommand

`func NewCreateAlertingIntegrationCommand(url string, alertingIntegrationType AlertingIntegrationType, alertingProfileId int32, ) *CreateAlertingIntegrationCommand`

NewCreateAlertingIntegrationCommand instantiates a new CreateAlertingIntegrationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlertingIntegrationCommandWithDefaults

`func NewCreateAlertingIntegrationCommandWithDefaults() *CreateAlertingIntegrationCommand`

NewCreateAlertingIntegrationCommandWithDefaults instantiates a new CreateAlertingIntegrationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateAlertingIntegrationCommand) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateAlertingIntegrationCommand) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateAlertingIntegrationCommand) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *CreateAlertingIntegrationCommand) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAlertingIntegrationCommand) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAlertingIntegrationCommand) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAlertingIntegrationCommand) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *CreateAlertingIntegrationCommand) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *CreateAlertingIntegrationCommand) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetAlertingIntegrationType

`func (o *CreateAlertingIntegrationCommand) GetAlertingIntegrationType() AlertingIntegrationType`

GetAlertingIntegrationType returns the AlertingIntegrationType field if non-nil, zero value otherwise.

### GetAlertingIntegrationTypeOk

`func (o *CreateAlertingIntegrationCommand) GetAlertingIntegrationTypeOk() (*AlertingIntegrationType, bool)`

GetAlertingIntegrationTypeOk returns a tuple with the AlertingIntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingIntegrationType

`func (o *CreateAlertingIntegrationCommand) SetAlertingIntegrationType(v AlertingIntegrationType)`

SetAlertingIntegrationType sets AlertingIntegrationType field to given value.


### GetAlertingProfileId

`func (o *CreateAlertingIntegrationCommand) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *CreateAlertingIntegrationCommand) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *CreateAlertingIntegrationCommand) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


