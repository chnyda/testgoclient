# \OrganizationSubscriptionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsubcriptionsList**](OrganizationSubscriptionsApi.md#OrganizationsubcriptionsList) | **Get** /api/v1/organizationsubcriptions | Retrieve all organization subscriptions
[**OrganizationsubcriptionsUpdate**](OrganizationSubscriptionsApi.md#OrganizationsubcriptionsUpdate) | **Post** /api/v1/organizationsubcriptions/update | Update subscription



## OrganizationsubcriptionsList

> OrganizationSubscriptionList OrganizationsubcriptionsList(ctx).Offset(offset).Limit(limit).OrganizationId(organizationId).Search(search).Execute()

Retrieve all organization subscriptions

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
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationSubscriptionsApi.OrganizationsubcriptionsList(context.Background()).Offset(offset).Limit(limit).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSubscriptionsApi.OrganizationsubcriptionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsubcriptionsList`: OrganizationSubscriptionList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationSubscriptionsApi.OrganizationsubcriptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsubcriptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**OrganizationSubscriptionList**](OrganizationSubscriptionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsubcriptionsUpdate

> OrganizationsubcriptionsUpdate(ctx).UpdateOrganizationSubscriptionCommand(updateOrganizationSubscriptionCommand).Execute()

Update subscription

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
    updateOrganizationSubscriptionCommand := *openapiclient.NewUpdateOrganizationSubscriptionCommand() // UpdateOrganizationSubscriptionCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationSubscriptionsApi.OrganizationsubcriptionsUpdate(context.Background()).UpdateOrganizationSubscriptionCommand(updateOrganizationSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSubscriptionsApi.OrganizationsubcriptionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsubcriptionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrganizationSubscriptionCommand** | [**UpdateOrganizationSubscriptionCommand**](UpdateOrganizationSubscriptionCommand.md) |  | 

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

