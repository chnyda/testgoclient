# \AWSCloudCredentialApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AwsCreate**](AWSCloudCredentialApi.md#AwsCreate) | **Post** /api/v1/aws/create | Add Aws credentials
[**AwsDeviceNames**](AWSCloudCredentialApi.md#AwsDeviceNames) | **Post** /api/v1/aws/device-names | Aws device name list
[**AwsList**](AWSCloudCredentialApi.md#AwsList) | **Get** /api/v1/aws/list | Retrieve list of aws cloud credentials
[**AwsOwners**](AWSCloudCredentialApi.md#AwsOwners) | **Get** /api/v1/aws/owners | Retrieve aws verified owner list
[**AwsRegionlist**](AWSCloudCredentialApi.md#AwsRegionlist) | **Post** /api/v1/aws/regions | Retrieve aws regions list
[**AwsUpdate**](AWSCloudCredentialApi.md#AwsUpdate) | **Post** /api/v1/aws/update | Update AWS credentials
[**AwsZones**](AWSCloudCredentialApi.md#AwsZones) | **Post** /api/v1/aws/zones | Fetch Aws zones



## AwsCreate

> ApiResponse AwsCreate(ctx).CreateAwsCloudCommand(createAwsCloudCommand).Execute()

Add Aws credentials

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
    createAwsCloudCommand := *openapiclient.NewCreateAwsCloudCommand("Name_example", "AwsSecretAccessKey_example", "AwsAccessKeyId_example", "AwsRegion_example") // CreateAwsCloudCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialApi.AwsCreate(context.Background()).CreateAwsCloudCommand(createAwsCloudCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialApi.AwsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialApi.AwsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAwsCloudCommand** | [**CreateAwsCloudCommand**](CreateAwsCloudCommand.md) |  | 

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


## AwsDeviceNames

> []string AwsDeviceNames(ctx).AwsBlockDeviceMappingsCommand(awsBlockDeviceMappingsCommand).Execute()

Aws device name list

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
    awsBlockDeviceMappingsCommand := *openapiclient.NewAwsBlockDeviceMappingsCommand() // AwsBlockDeviceMappingsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialApi.AwsDeviceNames(context.Background()).AwsBlockDeviceMappingsCommand(awsBlockDeviceMappingsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialApi.AwsDeviceNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsDeviceNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialApi.AwsDeviceNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsDeviceNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsBlockDeviceMappingsCommand** | [**AwsBlockDeviceMappingsCommand**](AwsBlockDeviceMappingsCommand.md) |  | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsList

> AwsCredentialList AwsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of aws cloud credentials

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
    resp, r, err := apiClient.AWSCloudCredentialApi.AwsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialApi.AwsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsList`: AwsCredentialList
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialApi.AwsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsListRequest struct via the builder pattern


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

[**AwsCredentialList**](AwsCredentialList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsOwners

> []CommonStringBasedDropdownDto AwsOwners(ctx).Execute()

Retrieve aws verified owner list

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialApi.AwsOwners(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialApi.AwsOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsOwners`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialApi.AwsOwners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAwsOwnersRequest struct via the builder pattern


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


## AwsRegionlist

> []AwsRegionDto AwsRegionlist(ctx).RegionListCommand(regionListCommand).Execute()

Retrieve aws regions list

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
    regionListCommand := *openapiclient.NewRegionListCommand() // RegionListCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialApi.AwsRegionlist(context.Background()).RegionListCommand(regionListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialApi.AwsRegionlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsRegionlist`: []AwsRegionDto
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialApi.AwsRegionlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsRegionlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionListCommand** | [**RegionListCommand**](RegionListCommand.md) |  | 

### Return type

[**[]AwsRegionDto**](AwsRegionDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsUpdate

> AwsUpdate(ctx).UpdateAwsCommand(updateAwsCommand).Execute()

Update AWS credentials

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
    updateAwsCommand := *openapiclient.NewUpdateAwsCommand(int32(123), "Name_example", "AwsSecretAccessKey_example", "AwsAccessKeyId_example") // UpdateAwsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AWSCloudCredentialApi.AwsUpdate(context.Background()).UpdateAwsCommand(updateAwsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialApi.AwsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAwsCommand** | [**UpdateAwsCommand**](UpdateAwsCommand.md) |  | 

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


## AwsZones

> AzResult AwsZones(ctx).AmazonAvailabilityZonesCommand(amazonAvailabilityZonesCommand).Execute()

Fetch Aws zones

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
    amazonAvailabilityZonesCommand := *openapiclient.NewAmazonAvailabilityZonesCommand() // AmazonAvailabilityZonesCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialApi.AwsZones(context.Background()).AmazonAvailabilityZonesCommand(amazonAvailabilityZonesCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialApi.AwsZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsZones`: AzResult
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialApi.AwsZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amazonAvailabilityZonesCommand** | [**AmazonAvailabilityZonesCommand**](AmazonAvailabilityZonesCommand.md) |  | 

### Return type

[**AzResult**](AzResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

