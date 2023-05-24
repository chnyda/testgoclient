# \OpaProfilesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpaprofilesCreate**](OpaProfilesApi.md#OpaprofilesCreate) | **Post** /api/v1/opaprofiles/create | Add policy profile
[**OpaprofilesDelete**](OpaProfilesApi.md#OpaprofilesDelete) | **Delete** /api/v1/opaprofiles/{id} | Remove Opa profile by Id
[**OpaprofilesDisableGatekeeper**](OpaProfilesApi.md#OpaprofilesDisableGatekeeper) | **Post** /api/v1/opaprofiles/disablegatekeeper | Disable Gatekeeper by the projectId
[**OpaprofilesEnableGatekeeper**](OpaProfilesApi.md#OpaprofilesEnableGatekeeper) | **Post** /api/v1/opaprofiles/enablegatekeeper | Enable Gatekeeper by the projectId
[**OpaprofilesList**](OpaProfilesApi.md#OpaprofilesList) | **Get** /api/v1/opaprofiles | Retrieve all policy profiles
[**OpaprofilesLockManager**](OpaProfilesApi.md#OpaprofilesLockManager) | **Post** /api/v1/opaprofiles/lockmanager | Lock/Unlock policy profile
[**OpaprofilesMakeDefault**](OpaProfilesApi.md#OpaprofilesMakeDefault) | **Post** /api/v1/opaprofiles/make-default | Choose default policy profile
[**OpaprofilesSync**](OpaProfilesApi.md#OpaprofilesSync) | **Post** /api/v1/opaprofiles/sync | Sync policy profile
[**OpaprofilesUpdate**](OpaProfilesApi.md#OpaprofilesUpdate) | **Put** /api/v1/opaprofiles | Update policy profile



## OpaprofilesCreate

> ApiResponse OpaprofilesCreate(ctx).CreateOpaProfileCommand(createOpaProfileCommand).Execute()

Add policy profile

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
    createOpaProfileCommand := *openapiclient.NewCreateOpaProfileCommand("Name_example") // CreateOpaProfileCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpaProfilesApi.OpaprofilesCreate(context.Background()).CreateOpaProfileCommand(createOpaProfileCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpaProfilesApi.OpaprofilesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpaprofilesCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `OpaProfilesApi.OpaprofilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOpaProfileCommand** | [**CreateOpaProfileCommand**](CreateOpaProfileCommand.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpaprofilesDelete

> OpaprofilesDelete(ctx, id).Execute()

Remove Opa profile by Id

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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpaProfilesApi.OpaprofilesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpaProfilesApi.OpaprofilesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpaprofilesDisableGatekeeper

> OpaprofilesDisableGatekeeper(ctx).DisableGatekeeperCommand(disableGatekeeperCommand).Execute()

Disable Gatekeeper by the projectId

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
    disableGatekeeperCommand := *openapiclient.NewDisableGatekeeperCommand() // DisableGatekeeperCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpaProfilesApi.OpaprofilesDisableGatekeeper(context.Background()).DisableGatekeeperCommand(disableGatekeeperCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpaProfilesApi.OpaprofilesDisableGatekeeper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesDisableGatekeeperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableGatekeeperCommand** | [**DisableGatekeeperCommand**](DisableGatekeeperCommand.md) |  | 

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


## OpaprofilesEnableGatekeeper

> OpaprofilesEnableGatekeeper(ctx).EnableGatekeeperCommand(enableGatekeeperCommand).Execute()

Enable Gatekeeper by the projectId

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
    enableGatekeeperCommand := *openapiclient.NewEnableGatekeeperCommand() // EnableGatekeeperCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpaProfilesApi.OpaprofilesEnableGatekeeper(context.Background()).EnableGatekeeperCommand(enableGatekeeperCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpaProfilesApi.OpaprofilesEnableGatekeeper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesEnableGatekeeperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableGatekeeperCommand** | [**EnableGatekeeperCommand**](EnableGatekeeperCommand.md) |  | 

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


## OpaprofilesList

> OpaProfileList OpaprofilesList(ctx).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).SearchId(searchId).Execute()

Retrieve all policy profiles

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
    organizationId := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    searchId := "searchId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpaProfilesApi.OpaprofilesList(context.Background()).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).SearchId(searchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpaProfilesApi.OpaprofilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpaprofilesList`: OpaProfileList
    fmt.Fprintf(os.Stdout, "Response from `OpaProfilesApi.OpaprofilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 
 **searchId** | **string** |  | 

### Return type

[**OpaProfileList**](OpaProfileList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpaprofilesLockManager

> OpaprofilesLockManager(ctx).OpaProfileLockManagerCommand(opaProfileLockManagerCommand).Execute()

Lock/Unlock policy profile

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
    opaProfileLockManagerCommand := *openapiclient.NewOpaProfileLockManagerCommand() // OpaProfileLockManagerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpaProfilesApi.OpaprofilesLockManager(context.Background()).OpaProfileLockManagerCommand(opaProfileLockManagerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpaProfilesApi.OpaprofilesLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaProfileLockManagerCommand** | [**OpaProfileLockManagerCommand**](OpaProfileLockManagerCommand.md) |  | 

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


## OpaprofilesMakeDefault

> OpaprofilesMakeDefault(ctx).OpaMakeDefaultCommand(opaMakeDefaultCommand).Execute()

Choose default policy profile

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
    opaMakeDefaultCommand := *openapiclient.NewOpaMakeDefaultCommand() // OpaMakeDefaultCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpaProfilesApi.OpaprofilesMakeDefault(context.Background()).OpaMakeDefaultCommand(opaMakeDefaultCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpaProfilesApi.OpaprofilesMakeDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesMakeDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaMakeDefaultCommand** | [**OpaMakeDefaultCommand**](OpaMakeDefaultCommand.md) |  | 

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


## OpaprofilesSync

> OpaprofilesSync(ctx).OpaProfileSyncCommand(opaProfileSyncCommand).Execute()

Sync policy profile

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
    opaProfileSyncCommand := *openapiclient.NewOpaProfileSyncCommand(int32(123)) // OpaProfileSyncCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpaProfilesApi.OpaprofilesSync(context.Background()).OpaProfileSyncCommand(opaProfileSyncCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpaProfilesApi.OpaprofilesSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaProfileSyncCommand** | [**OpaProfileSyncCommand**](OpaProfileSyncCommand.md) |  | 

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


## OpaprofilesUpdate

> OpaprofilesUpdate(ctx).OpaProfileUpdateCommand(opaProfileUpdateCommand).Execute()

Update policy profile

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
    opaProfileUpdateCommand := *openapiclient.NewOpaProfileUpdateCommand(int32(123), "Name_example") // OpaProfileUpdateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpaProfilesApi.OpaprofilesUpdate(context.Background()).OpaProfileUpdateCommand(opaProfileUpdateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpaProfilesApi.OpaprofilesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaProfileUpdateCommand** | [**OpaProfileUpdateCommand**](OpaProfileUpdateCommand.md) |  | 

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

