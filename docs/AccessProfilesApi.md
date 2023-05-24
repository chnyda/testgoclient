# \AccessProfilesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessprofilesCreate**](AccessProfilesApi.md#AccessprofilesCreate) | **Post** /api/v1/accessprofiles/create | Create access profile
[**AccessprofilesDelete**](AccessProfilesApi.md#AccessprofilesDelete) | **Delete** /api/v1/accessprofiles/{id} | Delete access profile by Id
[**AccessprofilesDropdown**](AccessProfilesApi.md#AccessprofilesDropdown) | **Get** /api/v1/accessprofiles/list | Retrieve access profiles by organization Id
[**AccessprofilesList**](AccessProfilesApi.md#AccessprofilesList) | **Get** /api/v1/accessprofiles | Retrieve all access profiles
[**AccessprofilesLockManager**](AccessProfilesApi.md#AccessprofilesLockManager) | **Post** /api/v1/accessprofiles/lockmanager | Lock/unlock access profiles
[**AccessprofilesUpdate**](AccessProfilesApi.md#AccessprofilesUpdate) | **Put** /api/v1/accessprofiles/update/{id} | Update access profile



## AccessprofilesCreate

> ApiResponse AccessprofilesCreate(ctx).CreateAccessProfileCommand(createAccessProfileCommand).Execute()

Create access profile

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
    createAccessProfileCommand := *openapiclient.NewCreateAccessProfileCommand("Name_example") // CreateAccessProfileCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.AccessprofilesCreate(context.Background()).CreateAccessProfileCommand(createAccessProfileCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.AccessprofilesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessprofilesCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.AccessprofilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessprofilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessProfileCommand** | [**CreateAccessProfileCommand**](CreateAccessProfileCommand.md) |  | 

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


## AccessprofilesDelete

> AccessprofilesDelete(ctx, id).Execute()

Delete access profile by Id

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
    r, err := apiClient.AccessProfilesApi.AccessprofilesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.AccessprofilesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccessprofilesDeleteRequest struct via the builder pattern


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


## AccessprofilesDropdown

> []CommonDropdownDto AccessprofilesDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve access profiles by organization Id

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.AccessprofilesDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.AccessprofilesDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessprofilesDropdown`: []CommonDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.AccessprofilesDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessprofilesDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]CommonDropdownDto**](CommonDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessprofilesList

> AccessProfilesList AccessprofilesList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve all access profiles

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.AccessprofilesList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.AccessprofilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessprofilesList`: AccessProfilesList
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.AccessprofilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessprofilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**AccessProfilesList**](AccessProfilesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessprofilesLockManager

> AccessprofilesLockManager(ctx).AccessProfilesLockManagementCommand(accessProfilesLockManagementCommand).Execute()

Lock/unlock access profiles

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
    accessProfilesLockManagementCommand := *openapiclient.NewAccessProfilesLockManagementCommand(int32(123), "Mode_example") // AccessProfilesLockManagementCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessProfilesApi.AccessprofilesLockManager(context.Background()).AccessProfilesLockManagementCommand(accessProfilesLockManagementCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.AccessprofilesLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessprofilesLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessProfilesLockManagementCommand** | [**AccessProfilesLockManagementCommand**](AccessProfilesLockManagementCommand.md) |  | 

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


## AccessprofilesUpdate

> AccessprofilesUpdate(ctx, id).UpdateAccessProfileDto(updateAccessProfileDto).Execute()

Update access profile

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
    updateAccessProfileDto := *openapiclient.NewUpdateAccessProfileDto() // UpdateAccessProfileDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessProfilesApi.AccessprofilesUpdate(context.Background(), id).UpdateAccessProfileDto(updateAccessProfileDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.AccessprofilesUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccessprofilesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccessProfileDto** | [**UpdateAccessProfileDto**](UpdateAccessProfileDto.md) |  | 

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

