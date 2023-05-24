# \AlertingProfilesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertingprofilesAssignEmail**](AlertingProfilesApi.md#AlertingprofilesAssignEmail) | **Put** /api/v1/alertingprofiles/assignemails/{id} | Assign Alerting emails
[**AlertingprofilesAssignWebhooks**](AlertingProfilesApi.md#AlertingprofilesAssignWebhooks) | **Put** /api/v1/alertingprofiles/assignwebhooks/{id} | Assign Alerting webhooks
[**AlertingprofilesAttach**](AlertingProfilesApi.md#AlertingprofilesAttach) | **Post** /api/v1/alertingprofiles/attach | Attach alerting profile to project
[**AlertingprofilesCreate**](AlertingProfilesApi.md#AlertingprofilesCreate) | **Post** /api/v1/alertingprofiles/create | Add Alerting profile
[**AlertingprofilesDelete**](AlertingProfilesApi.md#AlertingprofilesDelete) | **Delete** /api/v1/alertingprofiles/{id} | Remove Alerting profiles by Id
[**AlertingprofilesDetach**](AlertingProfilesApi.md#AlertingprofilesDetach) | **Post** /api/v1/alertingprofiles/detach | Detach alerting profile from project
[**AlertingprofilesDropdown**](AlertingProfilesApi.md#AlertingprofilesDropdown) | **Get** /api/v1/alertingprofiles/list | Retrieve all Alerting profiles for organization
[**AlertingprofilesEdit**](AlertingProfilesApi.md#AlertingprofilesEdit) | **Post** /api/v1/alertingprofiles/edit | Update Alerting profile
[**AlertingprofilesList**](AlertingProfilesApi.md#AlertingprofilesList) | **Get** /api/v1/alertingprofiles | Retrieve all Alerting profiles
[**AlertingprofilesLockManager**](AlertingProfilesApi.md#AlertingprofilesLockManager) | **Post** /api/v1/alertingprofiles/lockmanager | Lock/Unlock Alerting profiles
[**AlertingprofilesVerify**](AlertingProfilesApi.md#AlertingprofilesVerify) | **Post** /api/v1/alertingprofiles/verifywebhook | Verify webhook endpoint



## AlertingprofilesAssignEmail

> AlertingprofilesAssignEmail(ctx, id).AlertingEmailDto(alertingEmailDto).Execute()

Assign Alerting emails

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
    alertingEmailDto := []openapiclient.AlertingEmailDto{*openapiclient.NewAlertingEmailDto("Email_example")} // []AlertingEmailDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingProfilesApi.AlertingprofilesAssignEmail(context.Background(), id).AlertingEmailDto(alertingEmailDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesAssignEmail``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertingprofilesAssignEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertingEmailDto** | [**[]AlertingEmailDto**](AlertingEmailDto.md) |  | 

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


## AlertingprofilesAssignWebhooks

> AlertingprofilesAssignWebhooks(ctx, id).AlertingWebhookDto(alertingWebhookDto).Execute()

Assign Alerting webhooks

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
    alertingWebhookDto := []openapiclient.AlertingWebhookDto{*openapiclient.NewAlertingWebhookDto()} // []AlertingWebhookDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingProfilesApi.AlertingprofilesAssignWebhooks(context.Background(), id).AlertingWebhookDto(alertingWebhookDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesAssignWebhooks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertingprofilesAssignWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertingWebhookDto** | [**[]AlertingWebhookDto**](AlertingWebhookDto.md) |  | 

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


## AlertingprofilesAttach

> AlertingprofilesAttach(ctx).AttachDetachAlertingProfileCommand(attachDetachAlertingProfileCommand).Execute()

Attach alerting profile to project

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
    attachDetachAlertingProfileCommand := *openapiclient.NewAttachDetachAlertingProfileCommand() // AttachDetachAlertingProfileCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingProfilesApi.AlertingprofilesAttach(context.Background()).AttachDetachAlertingProfileCommand(attachDetachAlertingProfileCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesAttach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesAttachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachDetachAlertingProfileCommand** | [**AttachDetachAlertingProfileCommand**](AttachDetachAlertingProfileCommand.md) |  | 

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


## AlertingprofilesCreate

> ApiResponse AlertingprofilesCreate(ctx).CreateAlertingProfileCommand(createAlertingProfileCommand).Execute()

Add Alerting profile

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
    createAlertingProfileCommand := *openapiclient.NewCreateAlertingProfileCommand("Name_example") // CreateAlertingProfileCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingProfilesApi.AlertingprofilesCreate(context.Background()).CreateAlertingProfileCommand(createAlertingProfileCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertingprofilesCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesApi.AlertingprofilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAlertingProfileCommand** | [**CreateAlertingProfileCommand**](CreateAlertingProfileCommand.md) |  | 

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


## AlertingprofilesDelete

> AlertingprofilesDelete(ctx, id).Execute()

Remove Alerting profiles by Id

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
    r, err := apiClient.AlertingProfilesApi.AlertingprofilesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertingprofilesDeleteRequest struct via the builder pattern


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


## AlertingprofilesDetach

> AlertingprofilesDetach(ctx).AttachDetachAlertingProfileCommand(attachDetachAlertingProfileCommand).Execute()

Detach alerting profile from project

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
    attachDetachAlertingProfileCommand := *openapiclient.NewAttachDetachAlertingProfileCommand() // AttachDetachAlertingProfileCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingProfilesApi.AlertingprofilesDetach(context.Background()).AttachDetachAlertingProfileCommand(attachDetachAlertingProfileCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesDetach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesDetachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachDetachAlertingProfileCommand** | [**AttachDetachAlertingProfileCommand**](AttachDetachAlertingProfileCommand.md) |  | 

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


## AlertingprofilesDropdown

> []CommonDropdownDto AlertingprofilesDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve all Alerting profiles for organization

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
    resp, r, err := apiClient.AlertingProfilesApi.AlertingprofilesDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertingprofilesDropdown`: []CommonDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesApi.AlertingprofilesDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]CommonDropdownDto**](CommonDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingprofilesEdit

> ApiResponse AlertingprofilesEdit(ctx).UpdateAlertingProfileCommand(updateAlertingProfileCommand).Execute()

Update Alerting profile

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
    updateAlertingProfileCommand := *openapiclient.NewUpdateAlertingProfileCommand(int32(123)) // UpdateAlertingProfileCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingProfilesApi.AlertingprofilesEdit(context.Background()).UpdateAlertingProfileCommand(updateAlertingProfileCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertingprofilesEdit`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesApi.AlertingprofilesEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAlertingProfileCommand** | [**UpdateAlertingProfileCommand**](UpdateAlertingProfileCommand.md) |  | 

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


## AlertingprofilesList

> AlertingProfilesList AlertingprofilesList(ctx).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).SearchId(searchId).Execute()

Retrieve all Alerting profiles

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    searchId := "searchId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingProfilesApi.AlertingprofilesList(context.Background()).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).SearchId(searchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertingprofilesList`: AlertingProfilesList
    fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesApi.AlertingprofilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 
 **searchId** | **string** |  | 

### Return type

[**AlertingProfilesList**](AlertingProfilesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingprofilesLockManager

> AlertingprofilesLockManager(ctx).AlertingProfilesLockManagerCommand(alertingProfilesLockManagerCommand).Execute()

Lock/Unlock Alerting profiles

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
    alertingProfilesLockManagerCommand := *openapiclient.NewAlertingProfilesLockManagerCommand() // AlertingProfilesLockManagerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingProfilesApi.AlertingprofilesLockManager(context.Background()).AlertingProfilesLockManagerCommand(alertingProfilesLockManagerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertingProfilesLockManagerCommand** | [**AlertingProfilesLockManagerCommand**](AlertingProfilesLockManagerCommand.md) |  | 

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


## AlertingprofilesVerify

> AlertingprofilesVerify(ctx).VerifyWebhookCommand(verifyWebhookCommand).Execute()

Verify webhook endpoint

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
    verifyWebhookCommand := *openapiclient.NewVerifyWebhookCommand("Url_example") // VerifyWebhookCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingProfilesApi.AlertingprofilesVerify(context.Background()).VerifyWebhookCommand(verifyWebhookCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesApi.AlertingprofilesVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyWebhookCommand** | [**VerifyWebhookCommand**](VerifyWebhookCommand.md) |  | 

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

