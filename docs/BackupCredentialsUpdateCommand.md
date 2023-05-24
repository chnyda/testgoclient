# BackupCredentialsUpdateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**S3Name** | **string** |  | 
**S3AccessKeyId** | **string** |  | 
**S3SecretKey** | **string** |  | 

## Methods

### NewBackupCredentialsUpdateCommand

`func NewBackupCredentialsUpdateCommand(id int32, s3Name string, s3AccessKeyId string, s3SecretKey string, ) *BackupCredentialsUpdateCommand`

NewBackupCredentialsUpdateCommand instantiates a new BackupCredentialsUpdateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupCredentialsUpdateCommandWithDefaults

`func NewBackupCredentialsUpdateCommandWithDefaults() *BackupCredentialsUpdateCommand`

NewBackupCredentialsUpdateCommandWithDefaults instantiates a new BackupCredentialsUpdateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupCredentialsUpdateCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupCredentialsUpdateCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupCredentialsUpdateCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetS3Name

`func (o *BackupCredentialsUpdateCommand) GetS3Name() string`

GetS3Name returns the S3Name field if non-nil, zero value otherwise.

### GetS3NameOk

`func (o *BackupCredentialsUpdateCommand) GetS3NameOk() (*string, bool)`

GetS3NameOk returns a tuple with the S3Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Name

`func (o *BackupCredentialsUpdateCommand) SetS3Name(v string)`

SetS3Name sets S3Name field to given value.


### GetS3AccessKeyId

`func (o *BackupCredentialsUpdateCommand) GetS3AccessKeyId() string`

GetS3AccessKeyId returns the S3AccessKeyId field if non-nil, zero value otherwise.

### GetS3AccessKeyIdOk

`func (o *BackupCredentialsUpdateCommand) GetS3AccessKeyIdOk() (*string, bool)`

GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessKeyId

`func (o *BackupCredentialsUpdateCommand) SetS3AccessKeyId(v string)`

SetS3AccessKeyId sets S3AccessKeyId field to given value.


### GetS3SecretKey

`func (o *BackupCredentialsUpdateCommand) GetS3SecretKey() string`

GetS3SecretKey returns the S3SecretKey field if non-nil, zero value otherwise.

### GetS3SecretKeyOk

`func (o *BackupCredentialsUpdateCommand) GetS3SecretKeyOk() (*string, bool)`

GetS3SecretKeyOk returns a tuple with the S3SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3SecretKey

`func (o *BackupCredentialsUpdateCommand) SetS3SecretKey(v string)`

SetS3SecretKey sets S3SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


