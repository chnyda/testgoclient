# \KubernetesProfilesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubernetesprofilesCreate**](KubernetesProfilesApi.md#KubernetesprofilesCreate) | **Post** /api/v1/kubernetesprofiles | Add kubernetes profile
[**KubernetesprofilesDelete**](KubernetesProfilesApi.md#KubernetesprofilesDelete) | **Delete** /api/v1/kubernetesprofiles/{id} | Delete kubernetes profile
[**KubernetesprofilesDropdown**](KubernetesProfilesApi.md#KubernetesprofilesDropdown) | **Get** /api/v1/kubernetesprofiles | Retrieve all kubernetes profiles for organization
[**KubernetesprofilesList**](KubernetesProfilesApi.md#KubernetesprofilesList) | **Get** /api/v1/kubernetesprofiles/list | Retrieve all kubernetes profiles
[**KubernetesprofilesLockManager**](KubernetesProfilesApi.md#KubernetesprofilesLockManager) | **Post** /api/v1/kubernetesprofiles/lockmanager | Kubernetes profile lock/unlock
[**NotificationsList**](KubernetesProfilesApi.md#NotificationsList) | **Get** /api/v1/notifications/list | Retrieve all notifications



## KubernetesprofilesCreate

> ApiResponse KubernetesprofilesCreate(ctx).CreateKubernetesProfileCommand(createKubernetesProfileCommand).Execute()

Add kubernetes profile

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
    createKubernetesProfileCommand := *openapiclient.NewCreateKubernetesProfileCommand("Name_example") // CreateKubernetesProfileCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesProfilesApi.KubernetesprofilesCreate(context.Background()).CreateKubernetesProfileCommand(createKubernetesProfileCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesApi.KubernetesprofilesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesprofilesCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesProfilesApi.KubernetesprofilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesprofilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKubernetesProfileCommand** | [**CreateKubernetesProfileCommand**](CreateKubernetesProfileCommand.md) |  | 

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


## KubernetesprofilesDelete

> KubernetesprofilesDelete(ctx, id).Execute()

Delete kubernetes profile

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
    r, err := apiClient.KubernetesProfilesApi.KubernetesprofilesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesApi.KubernetesprofilesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKubernetesprofilesDeleteRequest struct via the builder pattern


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


## KubernetesprofilesDropdown

> []KubernetesProfilesEntity KubernetesprofilesDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve all kubernetes profiles for organization

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
    resp, r, err := apiClient.KubernetesProfilesApi.KubernetesprofilesDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesApi.KubernetesprofilesDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesprofilesDropdown`: []KubernetesProfilesEntity
    fmt.Fprintf(os.Stdout, "Response from `KubernetesProfilesApi.KubernetesprofilesDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesprofilesDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]KubernetesProfilesEntity**](KubernetesProfilesEntity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesprofilesList

> KubernetesProfilesList KubernetesprofilesList(ctx).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve all kubernetes profiles

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
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesProfilesApi.KubernetesprofilesList(context.Background()).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesApi.KubernetesprofilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesprofilesList`: KubernetesProfilesList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesProfilesApi.KubernetesprofilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesprofilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**KubernetesProfilesList**](KubernetesProfilesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesprofilesLockManager

> KubernetesprofilesLockManager(ctx).KubernetesProfilesLockManagerCommand(kubernetesProfilesLockManagerCommand).Execute()

Kubernetes profile lock/unlock

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
    kubernetesProfilesLockManagerCommand := *openapiclient.NewKubernetesProfilesLockManagerCommand() // KubernetesProfilesLockManagerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesProfilesApi.KubernetesprofilesLockManager(context.Background()).KubernetesProfilesLockManagerCommand(kubernetesProfilesLockManagerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesApi.KubernetesprofilesLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesprofilesLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesProfilesLockManagerCommand** | [**KubernetesProfilesLockManagerCommand**](KubernetesProfilesLockManagerCommand.md) |  | 

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


## NotificationsList

> NotificationHistory NotificationsList(ctx).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).FilterBy(filterBy).ProjectId(projectId).UserId(userId).IsDeleted(isDeleted).Search(search).Execute()

Retrieve all notifications

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
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    projectId := int32(56) // int32 |  (optional)
    userId := "userId_example" // string |  (optional)
    isDeleted := true // bool |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesProfilesApi.NotificationsList(context.Background()).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).FilterBy(filterBy).ProjectId(projectId).UserId(userId).IsDeleted(isDeleted).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesApi.NotificationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsList`: NotificationHistory
    fmt.Fprintf(os.Stdout, "Response from `KubernetesProfilesApi.NotificationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **organizationId** | **int32** |  | 
 **filterBy** | **string** |  | 
 **projectId** | **int32** |  | 
 **userId** | **string** |  | 
 **isDeleted** | **bool** |  | 
 **search** | **string** |  | 

### Return type

[**NotificationHistory**](NotificationHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

