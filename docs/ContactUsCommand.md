# ContactUsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**BusinessEmail** | **string** |  | 
**CompanyName** | **string** |  | 
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactUsCommand

`func NewContactUsCommand(name string, businessEmail string, companyName string, ) *ContactUsCommand`

NewContactUsCommand instantiates a new ContactUsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactUsCommandWithDefaults

`func NewContactUsCommandWithDefaults() *ContactUsCommand`

NewContactUsCommandWithDefaults instantiates a new ContactUsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContactUsCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactUsCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactUsCommand) SetName(v string)`

SetName sets Name field to given value.


### GetBusinessEmail

`func (o *ContactUsCommand) GetBusinessEmail() string`

GetBusinessEmail returns the BusinessEmail field if non-nil, zero value otherwise.

### GetBusinessEmailOk

`func (o *ContactUsCommand) GetBusinessEmailOk() (*string, bool)`

GetBusinessEmailOk returns a tuple with the BusinessEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessEmail

`func (o *ContactUsCommand) SetBusinessEmail(v string)`

SetBusinessEmail sets BusinessEmail field to given value.


### GetCompanyName

`func (o *ContactUsCommand) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ContactUsCommand) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ContactUsCommand) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetComment

`func (o *ContactUsCommand) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ContactUsCommand) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ContactUsCommand) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ContactUsCommand) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ContactUsCommand) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ContactUsCommand) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


