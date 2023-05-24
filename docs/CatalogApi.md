# \CatalogApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CatalogBindProject**](CatalogApi.md#CatalogBindProject) | **Post** /api/v1/catalog/bind-project | Bind projects to catalog
[**CatalogCreate**](CatalogApi.md#CatalogCreate) | **Post** /api/v1/catalog/create | Create catalog
[**CatalogDelete**](CatalogApi.md#CatalogDelete) | **Delete** /api/v1/catalog/{id} | Delete catalog
[**CatalogDropdown**](CatalogApi.md#CatalogDropdown) | **Get** /api/v1/catalog/list | Catalog dropdown list for organization
[**CatalogEdit**](CatalogApi.md#CatalogEdit) | **Put** /api/v1/catalog/edit | Edit catalog
[**CatalogList**](CatalogApi.md#CatalogList) | **Get** /api/v1/catalog | Catalog list for organization
[**CatalogLock**](CatalogApi.md#CatalogLock) | **Post** /api/v1/catalog/lockmanager | Lock catalog



## CatalogBindProject

> CatalogBindProject(ctx).BindProjectsToCatalogCommand(bindProjectsToCatalogCommand).Execute()

Bind projects to catalog

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
    bindProjectsToCatalogCommand := *openapiclient.NewBindProjectsToCatalogCommand([]openapiclient.UpdateCatalogDto{*openapiclient.NewUpdateCatalogDto()}, int32(123)) // BindProjectsToCatalogCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogApi.CatalogBindProject(context.Background()).BindProjectsToCatalogCommand(bindProjectsToCatalogCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.CatalogBindProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogBindProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindProjectsToCatalogCommand** | [**BindProjectsToCatalogCommand**](BindProjectsToCatalogCommand.md) |  | 

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


## CatalogCreate

> CatalogCreate(ctx).CreateCatalogCommand(createCatalogCommand).Execute()

Create catalog

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
    createCatalogCommand := *openapiclient.NewCreateCatalogCommand("Name_example") // CreateCatalogCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogApi.CatalogCreate(context.Background()).CreateCatalogCommand(createCatalogCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.CatalogCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCatalogCommand** | [**CreateCatalogCommand**](CreateCatalogCommand.md) |  | 

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


## CatalogDelete

> CatalogDelete(ctx, id).Execute()

Delete catalog

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
    r, err := apiClient.CatalogApi.CatalogDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.CatalogDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCatalogDeleteRequest struct via the builder pattern


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


## CatalogDropdown

> []CatalogDropdownDto CatalogDropdown(ctx).Search(search).Execute()

Catalog dropdown list for organization

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogApi.CatalogDropdown(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.CatalogDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CatalogDropdown`: []CatalogDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.CatalogDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**[]CatalogDropdownDto**](CatalogDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatalogEdit

> CatalogEdit(ctx).EditCatalogCommand(editCatalogCommand).Execute()

Edit catalog

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
    editCatalogCommand := *openapiclient.NewEditCatalogCommand(int32(123), "Name_example") // EditCatalogCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogApi.CatalogEdit(context.Background()).EditCatalogCommand(editCatalogCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.CatalogEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editCatalogCommand** | [**EditCatalogCommand**](EditCatalogCommand.md) |  | 

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


## CatalogList

> CatalogList CatalogList(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Catalog list for organization

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogApi.CatalogList(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.CatalogList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CatalogList`: CatalogList
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.CatalogList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**CatalogList**](CatalogList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatalogLock

> CatalogLock(ctx).CatalogLockManagementCommand(catalogLockManagementCommand).Execute()

Lock catalog

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
    catalogLockManagementCommand := *openapiclient.NewCatalogLockManagementCommand() // CatalogLockManagementCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogApi.CatalogLock(context.Background()).CatalogLockManagementCommand(catalogLockManagementCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.CatalogLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **catalogLockManagementCommand** | [**CatalogLockManagementCommand**](CatalogLockManagementCommand.md) |  | 

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

