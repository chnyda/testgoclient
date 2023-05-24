# \CatalogAppApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CatalogAppCreate**](CatalogAppApi.md#CatalogAppCreate) | **Post** /api/v1/catalog-app/create | Create catalog app
[**CatalogAppDelete**](CatalogAppApi.md#CatalogAppDelete) | **Delete** /api/v1/catalog-app/{id} | Delete catalog app
[**CatalogAppDetails**](CatalogAppApi.md#CatalogAppDetails) | **Get** /api/v1/catalog-app/{catalogAppId} | Catalog App details
[**CatalogAppEditParams**](CatalogAppApi.md#CatalogAppEditParams) | **Put** /api/v1/catalog-app/edit/params | Edit catalog app params
[**CatalogAppEditVersion**](CatalogAppApi.md#CatalogAppEditVersion) | **Put** /api/v1/catalog-app/edit/version | Edit catalog app version
[**CatalogAppLockManager**](CatalogAppApi.md#CatalogAppLockManager) | **Post** /api/v1/catalog-app/lockmanager | Lock catalog app
[**CatalogAppParamDetails**](CatalogAppApi.md#CatalogAppParamDetails) | **Get** /api/v1/catalog-app/params/{catalogAppId} | Catalog App param details



## CatalogAppCreate

> CatalogAppCreate(ctx).CreateCatalogAppCommand(createCatalogAppCommand).Execute()

Create catalog app

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
    createCatalogAppCommand := *openapiclient.NewCreateCatalogAppCommand("RepoName_example", "PackageName_example", int32(123)) // CreateCatalogAppCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogAppApi.CatalogAppCreate(context.Background()).CreateCatalogAppCommand(createCatalogAppCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAppApi.CatalogAppCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogAppCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCatalogAppCommand** | [**CreateCatalogAppCommand**](CreateCatalogAppCommand.md) |  | 

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


## CatalogAppDelete

> CatalogAppDelete(ctx, id).Execute()

Delete catalog app

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
    r, err := apiClient.CatalogAppApi.CatalogAppDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAppApi.CatalogAppDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCatalogAppDeleteRequest struct via the builder pattern


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


## CatalogAppDetails

> CatalogAppDetailsDto CatalogAppDetails(ctx, catalogAppId).Execute()

Catalog App details

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
    catalogAppId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogAppApi.CatalogAppDetails(context.Background(), catalogAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAppApi.CatalogAppDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CatalogAppDetails`: CatalogAppDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `CatalogAppApi.CatalogAppDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogAppId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogAppDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CatalogAppDetailsDto**](CatalogAppDetailsDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatalogAppEditParams

> CatalogAppEditParams(ctx).EditCatalogAppParamCommand(editCatalogAppParamCommand).Execute()

Edit catalog app params

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
    editCatalogAppParamCommand := *openapiclient.NewEditCatalogAppParamCommand(int32(123)) // EditCatalogAppParamCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogAppApi.CatalogAppEditParams(context.Background()).EditCatalogAppParamCommand(editCatalogAppParamCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAppApi.CatalogAppEditParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogAppEditParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editCatalogAppParamCommand** | [**EditCatalogAppParamCommand**](EditCatalogAppParamCommand.md) |  | 

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


## CatalogAppEditVersion

> CatalogAppEditVersion(ctx).EditCatalogAppVersionCommand(editCatalogAppVersionCommand).Execute()

Edit catalog app version

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
    editCatalogAppVersionCommand := *openapiclient.NewEditCatalogAppVersionCommand(int32(123), "Version_example") // EditCatalogAppVersionCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogAppApi.CatalogAppEditVersion(context.Background()).EditCatalogAppVersionCommand(editCatalogAppVersionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAppApi.CatalogAppEditVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogAppEditVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editCatalogAppVersionCommand** | [**EditCatalogAppVersionCommand**](EditCatalogAppVersionCommand.md) |  | 

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


## CatalogAppLockManager

> CatalogAppLockManager(ctx).CatalogAppLockManagement(catalogAppLockManagement).Execute()

Lock catalog app

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
    catalogAppLockManagement := *openapiclient.NewCatalogAppLockManagement() // CatalogAppLockManagement | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogAppApi.CatalogAppLockManager(context.Background()).CatalogAppLockManagement(catalogAppLockManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAppApi.CatalogAppLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogAppLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **catalogAppLockManagement** | [**CatalogAppLockManagement**](CatalogAppLockManagement.md) |  | 

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


## CatalogAppParamDetails

> []CatalogAppParamsDetailsDto CatalogAppParamDetails(ctx, catalogAppId).Execute()

Catalog App param details

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
    catalogAppId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogAppApi.CatalogAppParamDetails(context.Background(), catalogAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAppApi.CatalogAppParamDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CatalogAppParamDetails`: []CatalogAppParamsDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `CatalogAppApi.CatalogAppParamDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogAppId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogAppParamDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CatalogAppParamsDetailsDto**](CatalogAppParamsDetailsDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

