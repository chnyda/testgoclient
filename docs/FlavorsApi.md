# \FlavorsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlavorsAwsInstanceTypes**](FlavorsApi.md#FlavorsAwsInstanceTypes) | **Get** /api/v1/flavors/aws/{cloudId} | 
[**FlavorsAzureVmSizes**](FlavorsApi.md#FlavorsAzureVmSizes) | **Get** /api/v1/flavors/azure/{cloudId} | 
[**FlavorsBindToProject**](FlavorsApi.md#FlavorsBindToProject) | **Get** /api/v1/flavors/bind | 
[**FlavorsDropdownFlavors**](FlavorsApi.md#FlavorsDropdownFlavors) | **Get** /api/v1/flavors/credentials/dropdown/list | 
[**FlavorsGoogleMachineTypes**](FlavorsApi.md#FlavorsGoogleMachineTypes) | **Get** /api/v1/flavors/google/{cloudId} | 
[**FlavorsOpenstackFlavors**](FlavorsApi.md#FlavorsOpenstackFlavors) | **Get** /api/v1/flavors/openstack/{cloudId} | 
[**FlavorsProxmoxFlavors**](FlavorsApi.md#FlavorsProxmoxFlavors) | **Get** /api/v1/flavors/proxmox/{cloudId} | 
[**FlavorsSelectedFlavorsForProject**](FlavorsApi.md#FlavorsSelectedFlavorsForProject) | **Get** /api/v1/flavors/projects/list | 
[**FlavorsTanzuFlavors**](FlavorsApi.md#FlavorsTanzuFlavors) | **Get** /api/v1/flavors/tanzu/{cloudId} | 
[**FlavorsUnbindFromProject**](FlavorsApi.md#FlavorsUnbindFromProject) | **Get** /api/v1/flavors/unbind | 



## FlavorsAwsInstanceTypes

> AwsFlavorList FlavorsAwsInstanceTypes(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()



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
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    startRam := "startRam_example" // string |  (optional)
    endRam := "endRam_example" // string |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsAwsInstanceTypes(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsAwsInstanceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsAwsInstanceTypes`: AwsFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsAwsInstanceTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsAwsInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **string** |  | 
 **endRam** | **string** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**AwsFlavorList**](AwsFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsAzureVmSizes

> AzureFlavorList FlavorsAzureVmSizes(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()



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
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    startRam := int32(56) // int32 |  (optional)
    endRam := int32(56) // int32 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsAzureVmSizes(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsAzureVmSizes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsAzureVmSizes`: AzureFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsAzureVmSizes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsAzureVmSizesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **int32** |  | 
 **endRam** | **int32** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**AzureFlavorList**](AzureFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsBindToProject

> FlavorsBindToProject(ctx).BindFlavorToProjectCommand(bindFlavorToProjectCommand).Execute()



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
    bindFlavorToProjectCommand := *openapiclient.NewBindFlavorToProjectCommand() // BindFlavorToProjectCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FlavorsApi.FlavorsBindToProject(context.Background()).BindFlavorToProjectCommand(bindFlavorToProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsBindToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsBindToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindFlavorToProjectCommand** | [**BindFlavorToProjectCommand**](BindFlavorToProjectCommand.md) |  | 

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


## FlavorsDropdownFlavors

> []CloudCredentialsDropdownRecordDto FlavorsDropdownFlavors(ctx).OrganizationId(organizationId).FilterBy(filterBy).Search(search).Execute()



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
    filterBy := "filterBy_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsDropdownFlavors(context.Background()).OrganizationId(organizationId).FilterBy(filterBy).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsDropdownFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsDropdownFlavors`: []CloudCredentialsDropdownRecordDto
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsDropdownFlavors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsDropdownFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **filterBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]CloudCredentialsDropdownRecordDto**](CloudCredentialsDropdownRecordDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsGoogleMachineTypes

> GoogleFlavorList FlavorsGoogleMachineTypes(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()



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
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    startRam := "startRam_example" // string |  (optional)
    endRam := "endRam_example" // string |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsGoogleMachineTypes(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsGoogleMachineTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsGoogleMachineTypes`: GoogleFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsGoogleMachineTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsGoogleMachineTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **string** |  | 
 **endRam** | **string** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**GoogleFlavorList**](GoogleFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsOpenstackFlavors

> OpenstackFlavorList FlavorsOpenstackFlavors(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()



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
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    startRam := "startRam_example" // string |  (optional)
    endRam := "endRam_example" // string |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsOpenstackFlavors(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsOpenstackFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsOpenstackFlavors`: OpenstackFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsOpenstackFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsOpenstackFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **string** |  | 
 **endRam** | **string** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**OpenstackFlavorList**](OpenstackFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsProxmoxFlavors

> ProxmoxFlavorList FlavorsProxmoxFlavors(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()



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
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    startRam := int32(56) // int32 |  (optional)
    endRam := int32(56) // int32 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsProxmoxFlavors(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsProxmoxFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsProxmoxFlavors`: ProxmoxFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsProxmoxFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsProxmoxFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **int32** |  | 
 **endRam** | **int32** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**ProxmoxFlavorList**](ProxmoxFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsSelectedFlavorsForProject

> BoundFlavorsForProjectsList FlavorsSelectedFlavorsForProject(ctx).WithPrice(withPrice).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).FlavorName(flavorName).Execute()



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
    withPrice := true // bool | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    projectId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    flavorName := "flavorName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsSelectedFlavorsForProject(context.Background()).WithPrice(withPrice).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).FlavorName(flavorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsSelectedFlavorsForProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsSelectedFlavorsForProject`: BoundFlavorsForProjectsList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsSelectedFlavorsForProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsSelectedFlavorsForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withPrice** | **bool** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **projectId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **filterBy** | **string** |  | 
 **organizationId** | **int32** |  | 
 **flavorName** | **string** |  | 

### Return type

[**BoundFlavorsForProjectsList**](BoundFlavorsForProjectsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsTanzuFlavors

> TanzuFlavorList FlavorsTanzuFlavors(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()



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
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    startRam := int64(789) // int64 |  (optional)
    endRam := int64(789) // int64 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlavorsApi.FlavorsTanzuFlavors(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsTanzuFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlavorsTanzuFlavors`: TanzuFlavorList
    fmt.Fprintf(os.Stdout, "Response from `FlavorsApi.FlavorsTanzuFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsTanzuFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **int64** |  | 
 **endRam** | **int64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**TanzuFlavorList**](TanzuFlavorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsUnbindFromProject

> FlavorsUnbindFromProject(ctx).UnbindFlavorFromProjectCommand(unbindFlavorFromProjectCommand).Execute()



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
    unbindFlavorFromProjectCommand := *openapiclient.NewUnbindFlavorFromProjectCommand() // UnbindFlavorFromProjectCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FlavorsApi.FlavorsUnbindFromProject(context.Background()).UnbindFlavorFromProjectCommand(unbindFlavorFromProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlavorsApi.FlavorsUnbindFromProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsUnbindFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unbindFlavorFromProjectCommand** | [**UnbindFlavorFromProjectCommand**](UnbindFlavorFromProjectCommand.md) |  | 

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

