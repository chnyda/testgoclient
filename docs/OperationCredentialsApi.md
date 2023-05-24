# \OperationCredentialsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpscredentialsCreate**](OperationCredentialsApi.md#OpscredentialsCreate) | **Post** /api/v1/opscredentials | Add Operation credentials
[**OpscredentialsDelete**](OperationCredentialsApi.md#OpscredentialsDelete) | **Delete** /api/v1/opscredentials/{id} | Remove Operation credential by Id
[**OpscredentialsLockManager**](OperationCredentialsApi.md#OpscredentialsLockManager) | **Post** /api/v1/opscredentials/lockmanager | Lock/Unlock operation credential
[**OpscredentialsMakeDefault**](OperationCredentialsApi.md#OpscredentialsMakeDefault) | **Post** /api/v1/opscredentials/makedefault | Choose default operation credential



## OpscredentialsCreate

> ApiResponse OpscredentialsCreate(ctx).OperationCredentialsCreateCommand(operationCredentialsCreateCommand).Execute()

Add Operation credentials

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
    operationCredentialsCreateCommand := *openapiclient.NewOperationCredentialsCreateCommand("Name_example", "PrometheusUsername_example", "PrometheusPassword_example", "PrometheusUrl_example") // OperationCredentialsCreateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationCredentialsApi.OpscredentialsCreate(context.Background()).OperationCredentialsCreateCommand(operationCredentialsCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsApi.OpscredentialsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpscredentialsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `OperationCredentialsApi.OpscredentialsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationCredentialsCreateCommand** | [**OperationCredentialsCreateCommand**](OperationCredentialsCreateCommand.md) |  | 

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


## OpscredentialsDelete

> OpscredentialsDelete(ctx, id).Execute()

Remove Operation credential by Id

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
    r, err := apiClient.OperationCredentialsApi.OpscredentialsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsApi.OpscredentialsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOpscredentialsDeleteRequest struct via the builder pattern


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


## OpscredentialsLockManager

> OpscredentialsLockManager(ctx).OperationCredentialLockManagerCommand(operationCredentialLockManagerCommand).Execute()

Lock/Unlock operation credential

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
    operationCredentialLockManagerCommand := *openapiclient.NewOperationCredentialLockManagerCommand() // OperationCredentialLockManagerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OperationCredentialsApi.OpscredentialsLockManager(context.Background()).OperationCredentialLockManagerCommand(operationCredentialLockManagerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsApi.OpscredentialsLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationCredentialLockManagerCommand** | [**OperationCredentialLockManagerCommand**](OperationCredentialLockManagerCommand.md) |  | 

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


## OpscredentialsMakeDefault

> OpscredentialsMakeDefault(ctx).OperationCredentialsMakeDefaultCommand(operationCredentialsMakeDefaultCommand).Execute()

Choose default operation credential

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
    operationCredentialsMakeDefaultCommand := *openapiclient.NewOperationCredentialsMakeDefaultCommand() // OperationCredentialsMakeDefaultCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OperationCredentialsApi.OpscredentialsMakeDefault(context.Background()).OperationCredentialsMakeDefaultCommand(operationCredentialsMakeDefaultCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsApi.OpscredentialsMakeDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsMakeDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationCredentialsMakeDefaultCommand** | [**OperationCredentialsMakeDefaultCommand**](OperationCredentialsMakeDefaultCommand.md) |  | 

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

