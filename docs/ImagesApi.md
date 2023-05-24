# \ImagesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesAwsCommonImages**](ImagesApi.md#ImagesAwsCommonImages) | **Get** /api/v1/images/aws/common/{cloudId} | Commonly used aws images
[**ImagesAwsPersonalImages**](ImagesApi.md#ImagesAwsPersonalImages) | **Get** /api/v1/images/aws/personal/{cloudId} | Aws personal images
[**ImagesAzureImages**](ImagesApi.md#ImagesAzureImages) | **Get** /api/v1/images/azure/{cloudId}/{publisherName}/{offer}/{sku} | 
[**ImagesAzurePersonalImages**](ImagesApi.md#ImagesAzurePersonalImages) | **Get** /api/v1/images/azure/personal/{cloudId} | Azure personal images
[**ImagesBindImagesToProject**](ImagesApi.md#ImagesBindImagesToProject) | **Get** /api/v1/images/bind | Bind images to project
[**ImagesCommonGoogleImages**](ImagesApi.md#ImagesCommonGoogleImages) | **Get** /api/v1/images/google/common/{cloudId} | Commonly used google images
[**ImagesGoogleImages**](ImagesApi.md#ImagesGoogleImages) | **Get** /api/v1/images/google/{cloudId}/{type} | 
[**ImagesImageDetails**](ImagesApi.md#ImagesImageDetails) | **Get** /api/v1/images/details | Get image details
[**ImagesOpenstackImages**](ImagesApi.md#ImagesOpenstackImages) | **Get** /api/v1/images/openstack/{cloudId} | Retrieve openstack images
[**ImagesSelectedImagesForProject**](ImagesApi.md#ImagesSelectedImagesForProject) | **Get** /api/v1/images/projects/lis | Retrieve selected images for projects
[**ImagesTanzuImages**](ImagesApi.md#ImagesTanzuImages) | **Get** /api/v1/images/tanzu/{cloudId} | Retrieve tanzu images
[**ImagesUnbindImagesFromProject**](ImagesApi.md#ImagesUnbindImagesFromProject) | **Get** /api/v1/images/unbind | Unbind images from project



## ImagesAwsCommonImages

> []AwsOwnerDetails ImagesAwsCommonImages(ctx, cloudId).Execute()

Commonly used aws images

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ImagesAwsCommonImages(context.Background(), cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesAwsCommonImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAwsCommonImages`: []AwsOwnerDetails
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesAwsCommonImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesAwsCommonImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AwsOwnerDetails**](AwsOwnerDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesAwsPersonalImages

> []CommonStringBasedDropdownDto ImagesAwsPersonalImages(ctx, cloudId).Search(search).Execute()

Aws personal images

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ImagesAwsPersonalImages(context.Background(), cloudId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesAwsPersonalImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAwsPersonalImages`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesAwsPersonalImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesAwsPersonalImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesAzureImages

> AzureImageList ImagesAzureImages(ctx, cloudId, publisherName, offer, sku).Latest(latest).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



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
    publisherName := "publisherName_example" // string | 
    offer := "offer_example" // string | 
    sku := "sku_example" // string | 
    latest := true // bool | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ImagesAzureImages(context.Background(), cloudId, publisherName, offer, sku).Latest(latest).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesAzureImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAzureImages`: AzureImageList
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesAzureImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**publisherName** | **string** |  | 
**offer** | **string** |  | 
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesAzureImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **latest** | **bool** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**AzureImageList**](AzureImageList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesAzurePersonalImages

> []CommonStringBasedDropdownDto ImagesAzurePersonalImages(ctx, cloudId).Execute()

Azure personal images

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ImagesAzurePersonalImages(context.Background(), cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesAzurePersonalImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAzurePersonalImages`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesAzurePersonalImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesAzurePersonalImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesBindImagesToProject

> ImagesBindImagesToProject(ctx).BindImageToProjectCommand(bindImageToProjectCommand).Execute()

Bind images to project

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
    bindImageToProjectCommand := *openapiclient.NewBindImageToProjectCommand() // BindImageToProjectCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesApi.ImagesBindImagesToProject(context.Background()).BindImageToProjectCommand(bindImageToProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesBindImagesToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesBindImagesToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindImageToProjectCommand** | [**BindImageToProjectCommand**](BindImageToProjectCommand.md) |  | 

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


## ImagesCommonGoogleImages

> []GoogleOwnerDetails ImagesCommonGoogleImages(ctx, cloudId).Execute()

Commonly used google images

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ImagesCommonGoogleImages(context.Background(), cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesCommonGoogleImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesCommonGoogleImages`: []GoogleOwnerDetails
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesCommonGoogleImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesCommonGoogleImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GoogleOwnerDetails**](GoogleOwnerDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesGoogleImages

> GoogleImageList ImagesGoogleImages(ctx, cloudId, type_).Latest(latest).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



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
    type_ := "type__example" // string | 
    latest := true // bool | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ImagesGoogleImages(context.Background(), cloudId, type_).Latest(latest).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesGoogleImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesGoogleImages`: GoogleImageList
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesGoogleImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesGoogleImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **latest** | **bool** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**GoogleImageList**](GoogleImageList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesImageDetails

> ImagesImageDetails(ctx).ImageByIdCommand(imageByIdCommand).Execute()

Get image details

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
    imageByIdCommand := *openapiclient.NewImageByIdCommand() // ImageByIdCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesApi.ImagesImageDetails(context.Background()).ImageByIdCommand(imageByIdCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesImageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesImageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageByIdCommand** | [**ImageByIdCommand**](ImageByIdCommand.md) |  | 

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


## ImagesOpenstackImages

> OpenstackImageList ImagesOpenstackImages(ctx, cloudId).Personal(personal).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

Retrieve openstack images

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
    personal := true // bool | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ImagesOpenstackImages(context.Background(), cloudId).Personal(personal).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesOpenstackImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesOpenstackImages`: OpenstackImageList
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesOpenstackImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesOpenstackImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **personal** | **bool** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**OpenstackImageList**](OpenstackImageList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesSelectedImagesForProject

> BoundImagesForProjectsList ImagesSelectedImagesForProject(ctx).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).Execute()

Retrieve selected images for projects

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
    projectId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ImagesSelectedImagesForProject(context.Background()).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesSelectedImagesForProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesSelectedImagesForProject`: BoundImagesForProjectsList
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesSelectedImagesForProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesSelectedImagesForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **projectId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **filterBy** | **string** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**BoundImagesForProjectsList**](BoundImagesForProjectsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTanzuImages

> TanzuImageList ImagesTanzuImages(ctx, cloudId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

Retrieve tanzu images

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
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ImagesTanzuImages(context.Background(), cloudId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesTanzuImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTanzuImages`: TanzuImageList
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesTanzuImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTanzuImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**TanzuImageList**](TanzuImageList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesUnbindImagesFromProject

> ImagesUnbindImagesFromProject(ctx).DeleteImageFromProjectCommand(deleteImageFromProjectCommand).Execute()

Unbind images from project

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
    deleteImageFromProjectCommand := *openapiclient.NewDeleteImageFromProjectCommand() // DeleteImageFromProjectCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesApi.ImagesUnbindImagesFromProject(context.Background()).DeleteImageFromProjectCommand(deleteImageFromProjectCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesUnbindImagesFromProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesUnbindImagesFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteImageFromProjectCommand** | [**DeleteImageFromProjectCommand**](DeleteImageFromProjectCommand.md) |  | 

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

