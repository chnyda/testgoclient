# DuplicateNameCheckerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**Type** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewDuplicateNameCheckerCommand

`func NewDuplicateNameCheckerCommand(type_ string, name string, ) *DuplicateNameCheckerCommand`

NewDuplicateNameCheckerCommand instantiates a new DuplicateNameCheckerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicateNameCheckerCommandWithDefaults

`func NewDuplicateNameCheckerCommandWithDefaults() *DuplicateNameCheckerCommand`

NewDuplicateNameCheckerCommandWithDefaults instantiates a new DuplicateNameCheckerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *DuplicateNameCheckerCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DuplicateNameCheckerCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DuplicateNameCheckerCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DuplicateNameCheckerCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *DuplicateNameCheckerCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *DuplicateNameCheckerCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetType

`func (o *DuplicateNameCheckerCommand) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DuplicateNameCheckerCommand) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DuplicateNameCheckerCommand) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *DuplicateNameCheckerCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DuplicateNameCheckerCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DuplicateNameCheckerCommand) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


