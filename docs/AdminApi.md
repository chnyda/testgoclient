# \AdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAddBalance**](AdminApi.md#AdminAddBalance) | **Post** /api/v1/admin/organizations/add/balance | Add balance for organization
[**AdminBillingOperations**](AdminApi.md#AdminBillingOperations) | **Post** /api/v1/admin/cloudcredentials/billing | Billing operations: enable/disable billing 
[**AdminCreateUser**](AdminApi.md#AdminCreateUser) | **Post** /api/v1/admin/users/create | User creation for admin
[**AdminDeleteOrg**](AdminApi.md#AdminDeleteOrg) | **Post** /api/v1/admin/organizations/delete | Delete organization
[**AdminKeycloakList**](AdminApi.md#AdminKeycloakList) | **Get** /api/v1/admin/keycloak/list | Keycloak list for admin
[**AdminMakeCsm**](AdminApi.md#AdminMakeCsm) | **Post** /api/v1/admin/users/make/csm | User csm update for admin
[**AdminMakeOwner**](AdminApi.md#AdminMakeOwner) | **Post** /api/v1/admin/users/make/owner | User choose owner for admin
[**AdminOrganizations**](AdminApi.md#AdminOrganizations) | **Get** /api/v1/admin/organizations/list |  Organizations for admin
[**AdminProjectList**](AdminApi.md#AdminProjectList) | **Get** /api/v1/admin/projects/list | Projects for admin
[**AdminUpdateProject**](AdminApi.md#AdminUpdateProject) | **Post** /api/v1/admin/projects/update/version | Projects update for admin
[**AdminUpdateProjectKube**](AdminApi.md#AdminUpdateProjectKube) | **Post** /api/v1/admin/projects/update/kubeconfig | Projects update kube for admin
[**AdminUpdateUser**](AdminApi.md#AdminUpdateUser) | **Post** /api/v1/admin/users/update/password | User password update for admin
[**AdminUpdateUserKube**](AdminApi.md#AdminUpdateUserKube) | **Post** /api/v1/admin/projects/update/userkube | Projects update kube for admin
[**AdminUpdateUsers**](AdminApi.md#AdminUpdateUsers) | **Post** /api/v1/admin/users/update/email | User email update for admin
[**AdminUsersList**](AdminApi.md#AdminUsersList) | **Get** /api/v1/admin/users/list | Users for admin



## AdminAddBalance

> AdminAddBalance(ctx).AdminAddBalanceCommand(adminAddBalanceCommand).Execute()

Add balance for organization

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
    adminAddBalanceCommand := *openapiclient.NewAdminAddBalanceCommand("CustomerId_example") // AdminAddBalanceCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminAddBalance(context.Background()).AdminAddBalanceCommand(adminAddBalanceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminAddBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAddBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminAddBalanceCommand** | [**AdminAddBalanceCommand**](AdminAddBalanceCommand.md) |  | 

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


## AdminBillingOperations

> AdminBillingOperations(ctx).AdminBillingOperationCommand(adminBillingOperationCommand).Execute()

Billing operations: enable/disable billing 

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
    adminBillingOperationCommand := *openapiclient.NewAdminBillingOperationCommand(int32(123)) // AdminBillingOperationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminBillingOperations(context.Background()).AdminBillingOperationCommand(adminBillingOperationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminBillingOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminBillingOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminBillingOperationCommand** | [**AdminBillingOperationCommand**](AdminBillingOperationCommand.md) |  | 

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


## AdminCreateUser

> AdminCreateUser(ctx).AdminUserCreateCommand(adminUserCreateCommand).Execute()

User creation for admin

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
    adminUserCreateCommand := *openapiclient.NewAdminUserCreateCommand("Email_example", "Username_example", "Password_example") // AdminUserCreateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminCreateUser(context.Background()).AdminUserCreateCommand(adminUserCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminCreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUserCreateCommand** | [**AdminUserCreateCommand**](AdminUserCreateCommand.md) |  | 

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


## AdminDeleteOrg

> AdminDeleteOrg(ctx).AdminOrganizationsDeleteCommand(adminOrganizationsDeleteCommand).Execute()

Delete organization

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
    adminOrganizationsDeleteCommand := *openapiclient.NewAdminOrganizationsDeleteCommand() // AdminOrganizationsDeleteCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminDeleteOrg(context.Background()).AdminOrganizationsDeleteCommand(adminOrganizationsDeleteCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminDeleteOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminOrganizationsDeleteCommand** | [**AdminOrganizationsDeleteCommand**](AdminOrganizationsDeleteCommand.md) |  | 

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


## AdminKeycloakList

> AdminProjectsList AdminKeycloakList(ctx).Limit(limit).Offset(offset).Execute()

Keycloak list for admin

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminKeycloakList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminKeycloakList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminKeycloakList`: AdminProjectsList
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminKeycloakList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminKeycloakListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**AdminProjectsList**](AdminProjectsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminMakeCsm

> AdminMakeCsm(ctx).MakeCsmCommand(makeCsmCommand).Execute()

User csm update for admin

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
    makeCsmCommand := *openapiclient.NewMakeCsmCommand("UserId_example", "Mode_example") // MakeCsmCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminMakeCsm(context.Background()).MakeCsmCommand(makeCsmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminMakeCsm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminMakeCsmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **makeCsmCommand** | [**MakeCsmCommand**](MakeCsmCommand.md) |  | 

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


## AdminMakeOwner

> AdminMakeOwner(ctx).MakeOwnerCommand(makeOwnerCommand).Execute()

User choose owner for admin

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
    makeOwnerCommand := *openapiclient.NewMakeOwnerCommand("UserId_example") // MakeOwnerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminMakeOwner(context.Background()).MakeOwnerCommand(makeOwnerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminMakeOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminMakeOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **makeOwnerCommand** | [**MakeOwnerCommand**](MakeOwnerCommand.md) |  | 

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


## AdminOrganizations

> AdminOrganizationsList AdminOrganizations(ctx).Limit(limit).Offset(offset).PartnerId(partnerId).Search(search).Execute()

 Organizations for admin

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
    partnerId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminOrganizations(context.Background()).Limit(limit).Offset(offset).PartnerId(partnerId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminOrganizations`: AdminOrganizationsList
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **partnerId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**AdminOrganizationsList**](AdminOrganizationsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminProjectList

> AdminProjectsList AdminProjectList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).Execute()

Projects for admin

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminProjectList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminProjectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminProjectList`: AdminProjectsList
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminProjectList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminProjectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**AdminProjectsList**](AdminProjectsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUpdateProject

> AdminUpdateProject(ctx).AdminProjectUpdateCommand(adminProjectUpdateCommand).Execute()

Projects update for admin

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
    adminProjectUpdateCommand := *openapiclient.NewAdminProjectUpdateCommand(int32(123), "KubernetesCurrentVersion_example", "KubesprayCurrentVersion_example") // AdminProjectUpdateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminUpdateProject(context.Background()).AdminProjectUpdateCommand(adminProjectUpdateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminUpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminProjectUpdateCommand** | [**AdminProjectUpdateCommand**](AdminProjectUpdateCommand.md) |  | 

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


## AdminUpdateProjectKube

> AdminUpdateProjectKube(ctx).AdminUpdateProjectKubeConfigCommand(adminUpdateProjectKubeConfigCommand).Execute()

Projects update kube for admin

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
    adminUpdateProjectKubeConfigCommand := *openapiclient.NewAdminUpdateProjectKubeConfigCommand(int32(123)) // AdminUpdateProjectKubeConfigCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminUpdateProjectKube(context.Background()).AdminUpdateProjectKubeConfigCommand(adminUpdateProjectKubeConfigCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminUpdateProjectKube``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateProjectKubeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUpdateProjectKubeConfigCommand** | [**AdminUpdateProjectKubeConfigCommand**](AdminUpdateProjectKubeConfigCommand.md) |  | 

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


## AdminUpdateUser

> AdminUpdateUser(ctx).AdminUsersUpdatePasswordCommand(adminUsersUpdatePasswordCommand).Execute()

User password update for admin

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
    adminUsersUpdatePasswordCommand := *openapiclient.NewAdminUsersUpdatePasswordCommand("Id_example", "Password_example") // AdminUsersUpdatePasswordCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminUpdateUser(context.Background()).AdminUsersUpdatePasswordCommand(adminUsersUpdatePasswordCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminUpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUsersUpdatePasswordCommand** | [**AdminUsersUpdatePasswordCommand**](AdminUsersUpdatePasswordCommand.md) |  | 

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


## AdminUpdateUserKube

> AdminUpdateUserKube(ctx).AdminUpdateUserKubeConfigCommand(adminUpdateUserKubeConfigCommand).Execute()

Projects update kube for admin

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
    adminUpdateUserKubeConfigCommand := *openapiclient.NewAdminUpdateUserKubeConfigCommand(int32(123)) // AdminUpdateUserKubeConfigCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminUpdateUserKube(context.Background()).AdminUpdateUserKubeConfigCommand(adminUpdateUserKubeConfigCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminUpdateUserKube``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateUserKubeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUpdateUserKubeConfigCommand** | [**AdminUpdateUserKubeConfigCommand**](AdminUpdateUserKubeConfigCommand.md) |  | 

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


## AdminUpdateUsers

> AdminUpdateUsers(ctx).AdminUsersUpdateEmailCommand(adminUsersUpdateEmailCommand).Execute()

User email update for admin

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
    adminUsersUpdateEmailCommand := *openapiclient.NewAdminUsersUpdateEmailCommand("Id_example") // AdminUsersUpdateEmailCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminUpdateUsers(context.Background()).AdminUsersUpdateEmailCommand(adminUsersUpdateEmailCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminUpdateUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUsersUpdateEmailCommand** | [**AdminUsersUpdateEmailCommand**](AdminUsersUpdateEmailCommand.md) |  | 

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


## AdminUsersList

> AdminUsersList AdminUsersList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).Execute()

Users for admin

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminUsersList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminUsersList`: AdminUsersList
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**AdminUsersList**](AdminUsersList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

