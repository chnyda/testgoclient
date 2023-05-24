# ProjectChartDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succeeded** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 
**Pending** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 
**Updating** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 

## Methods

### NewProjectChartDto

`func NewProjectChartDto() *ProjectChartDto`

NewProjectChartDto instantiates a new ProjectChartDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectChartDtoWithDefaults

`func NewProjectChartDtoWithDefaults() *ProjectChartDto`

NewProjectChartDtoWithDefaults instantiates a new ProjectChartDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucceeded

`func (o *ProjectChartDto) GetSucceeded() []ProjectCommonRecordDto`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *ProjectChartDto) GetSucceededOk() (*[]ProjectCommonRecordDto, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *ProjectChartDto) SetSucceeded(v []ProjectCommonRecordDto)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *ProjectChartDto) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *ProjectChartDto) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *ProjectChartDto) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetPending

`func (o *ProjectChartDto) GetPending() []ProjectCommonRecordDto`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ProjectChartDto) GetPendingOk() (*[]ProjectCommonRecordDto, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ProjectChartDto) SetPending(v []ProjectCommonRecordDto)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ProjectChartDto) HasPending() bool`

HasPending returns a boolean if a field has been set.

### SetPendingNil

`func (o *ProjectChartDto) SetPendingNil(b bool)`

 SetPendingNil sets the value for Pending to be an explicit nil

### UnsetPending
`func (o *ProjectChartDto) UnsetPending()`

UnsetPending ensures that no value is present for Pending, not even an explicit nil
### GetUpdating

`func (o *ProjectChartDto) GetUpdating() []ProjectCommonRecordDto`

GetUpdating returns the Updating field if non-nil, zero value otherwise.

### GetUpdatingOk

`func (o *ProjectChartDto) GetUpdatingOk() (*[]ProjectCommonRecordDto, bool)`

GetUpdatingOk returns a tuple with the Updating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdating

`func (o *ProjectChartDto) SetUpdating(v []ProjectCommonRecordDto)`

SetUpdating sets Updating field to given value.

### HasUpdating

`func (o *ProjectChartDto) HasUpdating() bool`

HasUpdating returns a boolean if a field has been set.

### SetUpdatingNil

`func (o *ProjectChartDto) SetUpdatingNil(b bool)`

 SetUpdatingNil sets the value for Updating to be an explicit nil

### UnsetUpdating
`func (o *ProjectChartDto) UnsetUpdating()`

UnsetUpdating ensures that no value is present for Updating, not even an explicit nil
### GetTotalCount

`func (o *ProjectChartDto) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProjectChartDto) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProjectChartDto) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ProjectChartDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetFailed

`func (o *ProjectChartDto) GetFailed() []ProjectCommonRecordDto`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ProjectChartDto) GetFailedOk() (*[]ProjectCommonRecordDto, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ProjectChartDto) SetFailed(v []ProjectCommonRecordDto)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ProjectChartDto) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *ProjectChartDto) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *ProjectChartDto) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


