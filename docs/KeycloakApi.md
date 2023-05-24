# \KeycloakApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeycloakCreate**](KeycloakApi.md#KeycloakCreate) | **Post** /api/v1/keycloak/create | Create keycloak configuration for organization
[**KeycloakDelete**](KeycloakApi.md#KeycloakDelete) | **Post** /api/v1/keycloak/delete | Delete keycloak configuration
[**KeycloakEdit**](KeycloakApi.md#KeycloakEdit) | **Post** /api/v1/keycloak/edit | Edit keycloak configuration for organization
[**KeycloakList**](KeycloakApi.md#KeycloakList) | **Get** /api/v1/keycloak | Get keycloak configuration



## KeycloakCreate

> KeycloakCreate(ctx).KeycloakCreateCommand(keycloakCreateCommand).Execute()

Create keycloak configuration for organization

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
    keycloakCreateCommand := *openapiclient.NewKeycloakCreateCommand("Name_example", "Url_example", "RealmsName_example", "ClientId_example", "ClientSecret_example") // KeycloakCreateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeycloakApi.KeycloakCreate(context.Background()).KeycloakCreateCommand(keycloakCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeycloakApi.KeycloakCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeycloakCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keycloakCreateCommand** | [**KeycloakCreateCommand**](KeycloakCreateCommand.md) |  | 

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


## KeycloakDelete

> KeycloakDelete(ctx).KeycloakDeleteCommand(keycloakDeleteCommand).Execute()

Delete keycloak configuration

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
    keycloakDeleteCommand := *openapiclient.NewKeycloakDeleteCommand() // KeycloakDeleteCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeycloakApi.KeycloakDelete(context.Background()).KeycloakDeleteCommand(keycloakDeleteCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeycloakApi.KeycloakDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeycloakDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keycloakDeleteCommand** | [**KeycloakDeleteCommand**](KeycloakDeleteCommand.md) |  | 

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


## KeycloakEdit

> KeycloakEdit(ctx).KeycloakEditCommand(keycloakEditCommand).Execute()

Edit keycloak configuration for organization

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
    keycloakEditCommand := *openapiclient.NewKeycloakEditCommand(int32(123), "Name_example", "Url_example", "RealmsName_example", "ClientId_example", "ClientSecret_example") // KeycloakEditCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeycloakApi.KeycloakEdit(context.Background()).KeycloakEditCommand(keycloakEditCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeycloakApi.KeycloakEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeycloakEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keycloakEditCommand** | [**KeycloakEditCommand**](KeycloakEditCommand.md) |  | 

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


## KeycloakList

> KeycloakListDto KeycloakList(ctx).Execute()

Get keycloak configuration

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
    resp, r, err := apiClient.KeycloakApi.KeycloakList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeycloakApi.KeycloakList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeycloakList`: KeycloakListDto
    fmt.Fprintf(os.Stdout, "Response from `KeycloakApi.KeycloakList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKeycloakListRequest struct via the builder pattern


### Return type

[**KeycloakListDto**](KeycloakListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

