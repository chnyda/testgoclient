# \OperationCredentialApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpscredentialsList**](OperationCredentialApi.md#OpscredentialsList) | **Get** /api/v1/opscredentials/list | Retrieve all operation credentials
[**OpscredentialsListByOrganizationId**](OperationCredentialApi.md#OpscredentialsListByOrganizationId) | **Get** /api/v1/opscredentials | Retrieve operation credentials by organization Id



## OpscredentialsList

> OperationCredentials OpscredentialsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve all operation credentials

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
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationCredentialApi.OpscredentialsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialApi.OpscredentialsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpscredentialsList`: OperationCredentials
    fmt.Fprintf(os.Stdout, "Response from `OperationCredentialApi.OpscredentialsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**OperationCredentials**](OperationCredentials.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpscredentialsListByOrganizationId

> []OperationCredentialsForOrganizationEntity OpscredentialsListByOrganizationId(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve operation credentials by organization Id

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
    resp, r, err := apiClient.OperationCredentialApi.OpscredentialsListByOrganizationId(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialApi.OpscredentialsListByOrganizationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpscredentialsListByOrganizationId`: []OperationCredentialsForOrganizationEntity
    fmt.Fprintf(os.Stdout, "Response from `OperationCredentialApi.OpscredentialsListByOrganizationId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsListByOrganizationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]OperationCredentialsForOrganizationEntity**](OperationCredentialsForOrganizationEntity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

