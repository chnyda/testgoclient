# ProxmoxFlavorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ProxmoxFlavorData**](ProxmoxFlavorData.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewProxmoxFlavorList

`func NewProxmoxFlavorList() *ProxmoxFlavorList`

NewProxmoxFlavorList instantiates a new ProxmoxFlavorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxmoxFlavorListWithDefaults

`func NewProxmoxFlavorListWithDefaults() *ProxmoxFlavorList`

NewProxmoxFlavorListWithDefaults instantiates a new ProxmoxFlavorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProxmoxFlavorList) GetData() []ProxmoxFlavorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProxmoxFlavorList) GetDataOk() (*[]ProxmoxFlavorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProxmoxFlavorList) SetData(v []ProxmoxFlavorData)`

SetData sets Data field to given value.

### HasData

`func (o *ProxmoxFlavorList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ProxmoxFlavorList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ProxmoxFlavorList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetCount

`func (o *ProxmoxFlavorList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ProxmoxFlavorList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ProxmoxFlavorList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ProxmoxFlavorList) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


