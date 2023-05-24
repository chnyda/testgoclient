# \OparProfilesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpaprofilesDropdown**](OparProfilesApi.md#OpaprofilesDropdown) | **Get** /api/v1/opaprofiles/list | Retrieve policy profiles for organization



## OpaprofilesDropdown

> []CommonExtendedDropdownDto OpaprofilesDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve policy profiles for organization

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
    resp, r, err := apiClient.OparProfilesApi.OpaprofilesDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OparProfilesApi.OpaprofilesDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpaprofilesDropdown`: []CommonExtendedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `OparProfilesApi.OpaprofilesDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpaprofilesDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]CommonExtendedDropdownDto**](CommonExtendedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

