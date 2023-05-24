# CreateDnsServerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**AccessProfileId** | **int32** |  | 

## Methods

### NewCreateDnsServerCommand

`func NewCreateDnsServerCommand(address string, accessProfileId int32, ) *CreateDnsServerCommand`

NewCreateDnsServerCommand instantiates a new CreateDnsServerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsServerCommandWithDefaults

`func NewCreateDnsServerCommandWithDefaults() *CreateDnsServerCommand`

NewCreateDnsServerCommandWithDefaults instantiates a new CreateDnsServerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateDnsServerCommand) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateDnsServerCommand) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateDnsServerCommand) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAccessProfileId

`func (o *CreateDnsServerCommand) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *CreateDnsServerCommand) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *CreateDnsServerCommand) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


