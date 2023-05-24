# AdminAddBalanceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**Balance** | Pointer to **int64** |  | [optional] 

## Methods

### NewAdminAddBalanceCommand

`func NewAdminAddBalanceCommand(customerId string, ) *AdminAddBalanceCommand`

NewAdminAddBalanceCommand instantiates a new AdminAddBalanceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAddBalanceCommandWithDefaults

`func NewAdminAddBalanceCommandWithDefaults() *AdminAddBalanceCommand`

NewAdminAddBalanceCommandWithDefaults instantiates a new AdminAddBalanceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *AdminAddBalanceCommand) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AdminAddBalanceCommand) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AdminAddBalanceCommand) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetBalance

`func (o *AdminAddBalanceCommand) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AdminAddBalanceCommand) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AdminAddBalanceCommand) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AdminAddBalanceCommand) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


