# \DnsServersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsserversCreate**](DnsServersApi.md#DnsserversCreate) | **Post** /api/v1/dnsservers/create | Create dns servers for access profile
[**DnsserversEdit**](DnsServersApi.md#DnsserversEdit) | **Put** /api/v1/dnsservers/edit/{id} | Edit dns server
[**DnsserversList**](DnsServersApi.md#DnsserversList) | **Get** /api/v1/dnsservers/{accessProfileId} | 



## DnsserversCreate

> ApiResponse DnsserversCreate(ctx).CreateDnsServerCommand(createDnsServerCommand).Execute()

Create dns servers for access profile

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
    createDnsServerCommand := *openapiclient.NewCreateDnsServerCommand("Address_example", int32(123)) // CreateDnsServerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsServersApi.DnsserversCreate(context.Background()).CreateDnsServerCommand(createDnsServerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersApi.DnsserversCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsserversCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `DnsServersApi.DnsserversCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsserversCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDnsServerCommand** | [**CreateDnsServerCommand**](CreateDnsServerCommand.md) |  | 

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


## DnsserversEdit

> DnsserversEdit(ctx, id).DnsNtpAddressEditDto(dnsNtpAddressEditDto).Execute()

Edit dns server

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
    dnsNtpAddressEditDto := *openapiclient.NewDnsNtpAddressEditDto() // DnsNtpAddressEditDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DnsServersApi.DnsserversEdit(context.Background(), id).DnsNtpAddressEditDto(dnsNtpAddressEditDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersApi.DnsserversEdit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDnsserversEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsNtpAddressEditDto** | [**DnsNtpAddressEditDto**](DnsNtpAddressEditDto.md) |  | 

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


## DnsserversList

> []DnsServersListDto DnsserversList(ctx, accessProfileId).Search(search).Execute()



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
    accessProfileId := int32(56) // int32 | 
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsServersApi.DnsserversList(context.Background(), accessProfileId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersApi.DnsserversList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsserversList`: []DnsServersListDto
    fmt.Fprintf(os.Stdout, "Response from `DnsServersApi.DnsserversList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsserversListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**[]DnsServersListDto**](DnsServersListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

