# \SecurityGroupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecuritygroupCreate**](SecurityGroupApi.md#SecuritygroupCreate) | **Post** /api/v1/securitygroup/create | Create standalonealone profile security group
[**SecuritygroupDelete**](SecurityGroupApi.md#SecuritygroupDelete) | **Delete** /api/v1/securitygroup/{id} | Delete standalone profile security group
[**SecuritygroupEdit**](SecurityGroupApi.md#SecuritygroupEdit) | **Put** /api/v1/securitygroup/edit | Edit standalone profile security group
[**SecuritygroupList**](SecurityGroupApi.md#SecuritygroupList) | **Get** /api/v1/securitygroup/list/{standAloneProfileId} | List stand alone security group by profile id



## SecuritygroupCreate

> ApiResponse SecuritygroupCreate(ctx).CreateSecurityGroupCommand(createSecurityGroupCommand).Execute()

Create standalonealone profile security group

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
    createSecurityGroupCommand := *openapiclient.NewCreateSecurityGroupCommand("Name_example", "RemoteIpPrefix_example", int32(123)) // CreateSecurityGroupCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityGroupApi.SecuritygroupCreate(context.Background()).CreateSecurityGroupCommand(createSecurityGroupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupApi.SecuritygroupCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecuritygroupCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupApi.SecuritygroupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritygroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSecurityGroupCommand** | [**CreateSecurityGroupCommand**](CreateSecurityGroupCommand.md) |  | 

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


## SecuritygroupDelete

> SecuritygroupDelete(ctx, id).Execute()

Delete standalone profile security group

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
    r, err := apiClient.SecurityGroupApi.SecuritygroupDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupApi.SecuritygroupDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSecuritygroupDeleteRequest struct via the builder pattern


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


## SecuritygroupEdit

> SecuritygroupEdit(ctx).EditSecurityGroupCommand(editSecurityGroupCommand).Execute()

Edit standalone profile security group

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
    editSecurityGroupCommand := *openapiclient.NewEditSecurityGroupCommand(int32(123), "Name_example", "RemoteIpPrefix_example") // EditSecurityGroupCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecurityGroupApi.SecuritygroupEdit(context.Background()).EditSecurityGroupCommand(editSecurityGroupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupApi.SecuritygroupEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritygroupEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editSecurityGroupCommand** | [**EditSecurityGroupCommand**](EditSecurityGroupCommand.md) |  | 

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


## SecuritygroupList

> []SecurityGroupListDto SecuritygroupList(ctx, standAloneProfileId).Execute()

List stand alone security group by profile id

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
    standAloneProfileId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityGroupApi.SecuritygroupList(context.Background(), standAloneProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupApi.SecuritygroupList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecuritygroupList`: []SecurityGroupListDto
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupApi.SecuritygroupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**standAloneProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecuritygroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SecurityGroupListDto**](SecurityGroupListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

