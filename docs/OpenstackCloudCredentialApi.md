# \OpenstackCloudCredentialApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenstackCreate**](OpenstackCloudCredentialApi.md#OpenstackCreate) | **Post** /api/v1/openstack/create | Add Openstack credentials
[**OpenstackList**](OpenstackCloudCredentialApi.md#OpenstackList) | **Get** /api/v1/openstack/list | Retrieve list of openstack cloud credentials
[**OpenstackNetworks**](OpenstackCloudCredentialApi.md#OpenstackNetworks) | **Post** /api/v1/openstack/networks | Openstack network list
[**OpenstackProjects**](OpenstackCloudCredentialApi.md#OpenstackProjects) | **Post** /api/v1/openstack/projects | Openstack project list
[**OpenstackQuotas**](OpenstackCloudCredentialApi.md#OpenstackQuotas) | **Post** /api/v1/openstack/quotas | Openstack quota list
[**OpenstackRegionList**](OpenstackCloudCredentialApi.md#OpenstackRegionList) | **Post** /api/v1/openstack/regions | Retrieve Openstack regions
[**OpenstackSubnet**](OpenstackCloudCredentialApi.md#OpenstackSubnet) | **Post** /api/v1/openstack/subnets | Retrieve Openstack subnets
[**OpenstackUpdate**](OpenstackCloudCredentialApi.md#OpenstackUpdate) | **Post** /api/v1/openstack/update | Update Openstack credentials
[**OpenstackVolumes**](OpenstackCloudCredentialApi.md#OpenstackVolumes) | **Post** /api/v1/openstack/volumes | Openstack volume list
[**OpenstackZones**](OpenstackCloudCredentialApi.md#OpenstackZones) | **Post** /api/v1/openstack/zones | Fetch Openstack zones



## OpenstackCreate

> ApiResponse OpenstackCreate(ctx).CreateOpenstackCloudCommand(createOpenstackCloudCommand).Execute()

Add Openstack credentials

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
    createOpenstackCloudCommand := *openapiclient.NewCreateOpenstackCloudCommand("Name_example", "OpenStackUser_example", "OpenStackPassword_example", "OpenStackUrl_example", "OpenStackPublicNetwork_example", "OpenStackRegion_example") // CreateOpenstackCloudCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenstackCloudCredentialApi.OpenstackCreate(context.Background()).CreateOpenstackCloudCommand(createOpenstackCloudCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenstackCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialApi.OpenstackCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOpenstackCloudCommand** | [**CreateOpenstackCloudCommand**](CreateOpenstackCloudCommand.md) |  | 

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


## OpenstackList

> OpenstackCredentialList OpenstackList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of openstack cloud credentials

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
    resp, r, err := apiClient.OpenstackCloudCredentialApi.OpenstackList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenstackList`: OpenstackCredentialList
    fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialApi.OpenstackList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackListRequest struct via the builder pattern


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

[**OpenstackCredentialList**](OpenstackCredentialList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenstackNetworks

> []CommonStringBasedDropdownDto OpenstackNetworks(ctx).OpenStackNetworkListQuery(openStackNetworkListQuery).Execute()

Openstack network list

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
    openStackNetworkListQuery := *openapiclient.NewOpenStackNetworkListQuery() // OpenStackNetworkListQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenstackCloudCredentialApi.OpenstackNetworks(context.Background()).OpenStackNetworkListQuery(openStackNetworkListQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenstackNetworks`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialApi.OpenstackNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openStackNetworkListQuery** | [**OpenStackNetworkListQuery**](OpenStackNetworkListQuery.md) |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenstackProjects

> []CommonStringBasedDropdownDto OpenstackProjects(ctx).OpenStackProjectListQuery(openStackProjectListQuery).Execute()

Openstack project list

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
    openStackProjectListQuery := *openapiclient.NewOpenStackProjectListQuery() // OpenStackProjectListQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenstackCloudCredentialApi.OpenstackProjects(context.Background()).OpenStackProjectListQuery(openStackProjectListQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenstackProjects`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialApi.OpenstackProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openStackProjectListQuery** | [**OpenStackProjectListQuery**](OpenStackProjectListQuery.md) |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenstackQuotas

> OpenstackQuotaList OpenstackQuotas(ctx).OpenstackQuotasCommand(openstackQuotasCommand).Execute()

Openstack quota list

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
    openstackQuotasCommand := *openapiclient.NewOpenstackQuotasCommand() // OpenstackQuotasCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenstackCloudCredentialApi.OpenstackQuotas(context.Background()).OpenstackQuotasCommand(openstackQuotasCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenstackQuotas`: OpenstackQuotaList
    fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialApi.OpenstackQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openstackQuotasCommand** | [**OpenstackQuotasCommand**](OpenstackQuotasCommand.md) |  | 

### Return type

[**OpenstackQuotaList**](OpenstackQuotaList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenstackRegionList

> []string OpenstackRegionList(ctx).OpenStackRegionListQuery(openStackRegionListQuery).Execute()

Retrieve Openstack regions

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
    openStackRegionListQuery := *openapiclient.NewOpenStackRegionListQuery() // OpenStackRegionListQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenstackCloudCredentialApi.OpenstackRegionList(context.Background()).OpenStackRegionListQuery(openStackRegionListQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackRegionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenstackRegionList`: []string
    fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialApi.OpenstackRegionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackRegionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openStackRegionListQuery** | [**OpenStackRegionListQuery**](OpenStackRegionListQuery.md) |  | 

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


## OpenstackSubnet

> []Subnet OpenstackSubnet(ctx).OpenstackSubnetListQuery(openstackSubnetListQuery).Execute()

Retrieve Openstack subnets

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
    openstackSubnetListQuery := *openapiclient.NewOpenstackSubnetListQuery() // OpenstackSubnetListQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenstackCloudCredentialApi.OpenstackSubnet(context.Background()).OpenstackSubnetListQuery(openstackSubnetListQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenstackSubnet`: []Subnet
    fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialApi.OpenstackSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openstackSubnetListQuery** | [**OpenstackSubnetListQuery**](OpenstackSubnetListQuery.md) |  | 

### Return type

[**[]Subnet**](Subnet.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenstackUpdate

> OpenstackUpdate(ctx).UpdateOpenStackCommand(updateOpenStackCommand).Execute()

Update Openstack credentials

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
    updateOpenStackCommand := *openapiclient.NewUpdateOpenStackCommand(int32(123), "Name_example", "OpenStackUser_example", "OpenStackPassword_example") // UpdateOpenStackCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpenstackCloudCredentialApi.OpenstackUpdate(context.Background()).UpdateOpenStackCommand(updateOpenStackCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOpenStackCommand** | [**UpdateOpenStackCommand**](UpdateOpenStackCommand.md) |  | 

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


## OpenstackVolumes

> []string OpenstackVolumes(ctx).OpenstackVolumeTypeListQuery(openstackVolumeTypeListQuery).Execute()

Openstack volume list

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
    openstackVolumeTypeListQuery := *openapiclient.NewOpenstackVolumeTypeListQuery() // OpenstackVolumeTypeListQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenstackCloudCredentialApi.OpenstackVolumes(context.Background()).OpenstackVolumeTypeListQuery(openstackVolumeTypeListQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenstackVolumes`: []string
    fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialApi.OpenstackVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openstackVolumeTypeListQuery** | [**OpenstackVolumeTypeListQuery**](OpenstackVolumeTypeListQuery.md) |  | 

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


## OpenstackZones

> []string OpenstackZones(ctx).OpenStackZoneListQuery(openStackZoneListQuery).Execute()

Fetch Openstack zones

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
    openStackZoneListQuery := *openapiclient.NewOpenStackZoneListQuery() // OpenStackZoneListQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenstackCloudCredentialApi.OpenstackZones(context.Background()).OpenStackZoneListQuery(openStackZoneListQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialApi.OpenstackZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenstackZones`: []string
    fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialApi.OpenstackZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openStackZoneListQuery** | [**OpenStackZoneListQuery**](OpenStackZoneListQuery.md) |  | 

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

