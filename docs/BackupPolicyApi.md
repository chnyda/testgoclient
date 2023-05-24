# \BackupPolicyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupByName**](BackupPolicyApi.md#BackupByName) | **Get** /api/v1/backup/{projectId}/{name} | 
[**BackupClearProject**](BackupPolicyApi.md#BackupClearProject) | **Post** /api/v1/backup/clear/project | 
[**BackupCreate**](BackupPolicyApi.md#BackupCreate) | **Post** /api/v1/backup/create | 
[**BackupDeleteBackup**](BackupPolicyApi.md#BackupDeleteBackup) | **Post** /api/v1/backup/delete/backup | 
[**BackupDeleteBackupLocation**](BackupPolicyApi.md#BackupDeleteBackupLocation) | **Post** /api/v1/backup/delete/location | 
[**BackupDeleteRestore**](BackupPolicyApi.md#BackupDeleteRestore) | **Post** /api/v1/backup/delete/restore | 
[**BackupDeleteSchedule**](BackupPolicyApi.md#BackupDeleteSchedule) | **Post** /api/v1/backup/delete/schedule | 
[**BackupDescribeBackup**](BackupPolicyApi.md#BackupDescribeBackup) | **Get** /api/v1/backup/describe/backup/{projectId}/{name} | 
[**BackupDescribeRestore**](BackupPolicyApi.md#BackupDescribeRestore) | **Get** /api/v1/backup/describe/restore/{projectId}/{name} | 
[**BackupDescribeSchedule**](BackupPolicyApi.md#BackupDescribeSchedule) | **Get** /api/v1/backup/describe/schedule/{projectId}/{name} | 
[**BackupDisableBackup**](BackupPolicyApi.md#BackupDisableBackup) | **Post** /api/v1/backup/disablebackup | 
[**BackupEnableBackup**](BackupPolicyApi.md#BackupEnableBackup) | **Post** /api/v1/backup/enablebackup | 
[**BackupImportBackupStorage**](BackupPolicyApi.md#BackupImportBackupStorage) | **Post** /api/v1/backup/location | 
[**BackupListAllBackupStorages**](BackupPolicyApi.md#BackupListAllBackupStorages) | **Get** /api/v1/backup/location/{projectId} | 
[**BackupListAllBackups**](BackupPolicyApi.md#BackupListAllBackups) | **Get** /api/v1/backup/backups/{projectId} | 
[**BackupListAllDeleteBackupRequests**](BackupPolicyApi.md#BackupListAllDeleteBackupRequests) | **Get** /api/v1/backup/delete-requests/{projectId} | 
[**BackupListAllRestores**](BackupPolicyApi.md#BackupListAllRestores) | **Get** /api/v1/backup/restores/{projectId} | 
[**BackupListAllSchedules**](BackupPolicyApi.md#BackupListAllSchedules) | **Get** /api/v1/backup/schedules/{projectId} | 
[**BackupRestoreBackup**](BackupPolicyApi.md#BackupRestoreBackup) | **Post** /api/v1/backup/restore | 



## BackupByName

> BackupDto BackupByName(ctx, projectId, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyApi.BackupByName(context.Background(), projectId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupByName`: BackupDto
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyApi.BackupByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BackupDto**](BackupDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupClearProject

> BackupClearProject(ctx).ClearProjectBackupCommand(clearProjectBackupCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    clearProjectBackupCommand := *openapiclient.NewClearProjectBackupCommand() // ClearProjectBackupCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupClearProject(context.Background()).ClearProjectBackupCommand(clearProjectBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupClearProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupClearProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clearProjectBackupCommand** | [**ClearProjectBackupCommand**](ClearProjectBackupCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupCreate

> BackupCreate(ctx).CreateBackupPolicyCommand(createBackupPolicyCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    createBackupPolicyCommand := *openapiclient.NewCreateBackupPolicyCommand("Name_example", "CronPeriod_example", "RetentionPeriod_example", int32(123)) // CreateBackupPolicyCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupCreate(context.Background()).CreateBackupPolicyCommand(createBackupPolicyCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBackupPolicyCommand** | [**CreateBackupPolicyCommand**](CreateBackupPolicyCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDeleteBackup

> BackupDeleteBackup(ctx).DeleteBackupCommand(deleteBackupCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    deleteBackupCommand := *openapiclient.NewDeleteBackupCommand() // DeleteBackupCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupDeleteBackup(context.Background()).DeleteBackupCommand(deleteBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupDeleteBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBackupCommand** | [**DeleteBackupCommand**](DeleteBackupCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDeleteBackupLocation

> BackupDeleteBackupLocation(ctx).DeleteBackupStorageLocationCommand(deleteBackupStorageLocationCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    deleteBackupStorageLocationCommand := *openapiclient.NewDeleteBackupStorageLocationCommand() // DeleteBackupStorageLocationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupDeleteBackupLocation(context.Background()).DeleteBackupStorageLocationCommand(deleteBackupStorageLocationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupDeleteBackupLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteBackupLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBackupStorageLocationCommand** | [**DeleteBackupStorageLocationCommand**](DeleteBackupStorageLocationCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDeleteRestore

> BackupDeleteRestore(ctx).DeleteRestoreCommand(deleteRestoreCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    deleteRestoreCommand := *openapiclient.NewDeleteRestoreCommand() // DeleteRestoreCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupDeleteRestore(context.Background()).DeleteRestoreCommand(deleteRestoreCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupDeleteRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRestoreCommand** | [**DeleteRestoreCommand**](DeleteRestoreCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDeleteSchedule

> BackupDeleteSchedule(ctx).DeleteScheduleCommand(deleteScheduleCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    deleteScheduleCommand := *openapiclient.NewDeleteScheduleCommand() // DeleteScheduleCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupDeleteSchedule(context.Background()).DeleteScheduleCommand(deleteScheduleCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupDeleteSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteScheduleCommand** | [**DeleteScheduleCommand**](DeleteScheduleCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDescribeBackup

> string BackupDescribeBackup(ctx, projectId, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyApi.BackupDescribeBackup(context.Background(), projectId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupDescribeBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupDescribeBackup`: string
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyApi.BackupDescribeBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupDescribeBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDescribeRestore

> string BackupDescribeRestore(ctx, projectId, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyApi.BackupDescribeRestore(context.Background(), projectId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupDescribeRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupDescribeRestore`: string
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyApi.BackupDescribeRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupDescribeRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDescribeSchedule

> string BackupDescribeSchedule(ctx, projectId, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyApi.BackupDescribeSchedule(context.Background(), projectId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupDescribeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupDescribeSchedule`: string
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyApi.BackupDescribeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupDescribeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDisableBackup

> BackupDisableBackup(ctx).DisableBackupCommand(disableBackupCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    disableBackupCommand := *openapiclient.NewDisableBackupCommand() // DisableBackupCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupDisableBackup(context.Background()).DisableBackupCommand(disableBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupDisableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDisableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableBackupCommand** | [**DisableBackupCommand**](DisableBackupCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupEnableBackup

> BackupEnableBackup(ctx).EnableBackupCommand(enableBackupCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    enableBackupCommand := *openapiclient.NewEnableBackupCommand() // EnableBackupCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupEnableBackup(context.Background()).EnableBackupCommand(enableBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupEnableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupEnableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableBackupCommand** | [**EnableBackupCommand**](EnableBackupCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupImportBackupStorage

> BackupImportBackupStorage(ctx).ImportBackupStorageLocationCommand(importBackupStorageLocationCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    importBackupStorageLocationCommand := *openapiclient.NewImportBackupStorageLocationCommand() // ImportBackupStorageLocationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupImportBackupStorage(context.Background()).ImportBackupStorageLocationCommand(importBackupStorageLocationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupImportBackupStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupImportBackupStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importBackupStorageLocationCommand** | [**ImportBackupStorageLocationCommand**](ImportBackupStorageLocationCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllBackupStorages

> ListAllBackupStorageLocations BackupListAllBackupStorages(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyApi.BackupListAllBackupStorages(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupListAllBackupStorages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllBackupStorages`: ListAllBackupStorageLocations
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyApi.BackupListAllBackupStorages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllBackupStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**ListAllBackupStorageLocations**](ListAllBackupStorageLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllBackups

> map[string]interface{} BackupListAllBackups(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyApi.BackupListAllBackups(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupListAllBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllBackups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyApi.BackupListAllBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllDeleteBackupRequests

> map[string]interface{} BackupListAllDeleteBackupRequests(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyApi.BackupListAllDeleteBackupRequests(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupListAllDeleteBackupRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllDeleteBackupRequests`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyApi.BackupListAllDeleteBackupRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllDeleteBackupRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllRestores

> map[string]interface{} BackupListAllRestores(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyApi.BackupListAllRestores(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupListAllRestores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllRestores`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyApi.BackupListAllRestores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllRestoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllSchedules

> ListAllSchedules BackupListAllSchedules(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyApi.BackupListAllSchedules(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupListAllSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllSchedules`: ListAllSchedules
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyApi.BackupListAllSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**ListAllSchedules**](ListAllSchedules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreBackup

> BackupRestoreBackup(ctx).RestoreBackupCommand(restoreBackupCommand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/testgoclient"
)

func main() {
    restoreBackupCommand := *openapiclient.NewRestoreBackupCommand() // RestoreBackupCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyApi.BackupRestoreBackup(context.Background()).RestoreBackupCommand(restoreBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyApi.BackupRestoreBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restoreBackupCommand** | [**RestoreBackupCommand**](RestoreBackupCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

