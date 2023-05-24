# \AutoscalingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoscalingDisable**](AutoscalingApi.md#AutoscalingDisable) | **Post** /api/v1/autoscaling/disable | Disable autoscaling
[**AutoscalingEdit**](AutoscalingApi.md#AutoscalingEdit) | **Post** /api/v1/autoscaling/edit | Edit autoscaling
[**AutoscalingEnable**](AutoscalingApi.md#AutoscalingEnable) | **Post** /api/v1/autoscaling/enable | Enable autoscaling
[**AutoscalingSync**](AutoscalingApi.md#AutoscalingSync) | **Post** /api/v1/autoscaling/sync | Sync autoscaling



## AutoscalingDisable

> AutoscalingDisable(ctx).DisableAutoscalingCommand(disableAutoscalingCommand).Execute()

Disable autoscaling

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
    disableAutoscalingCommand := *openapiclient.NewDisableAutoscalingCommand() // DisableAutoscalingCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoscalingApi.AutoscalingDisable(context.Background()).DisableAutoscalingCommand(disableAutoscalingCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingApi.AutoscalingDisable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoscalingDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableAutoscalingCommand** | [**DisableAutoscalingCommand**](DisableAutoscalingCommand.md) |  | 

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


## AutoscalingEdit

> AutoscalingEdit(ctx).EditAutoscalingCommand(editAutoscalingCommand).Execute()

Edit autoscaling

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
    editAutoscalingCommand := *openapiclient.NewEditAutoscalingCommand(int32(123)) // EditAutoscalingCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoscalingApi.AutoscalingEdit(context.Background()).EditAutoscalingCommand(editAutoscalingCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingApi.AutoscalingEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoscalingEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editAutoscalingCommand** | [**EditAutoscalingCommand**](EditAutoscalingCommand.md) |  | 

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


## AutoscalingEnable

> AutoscalingEnable(ctx).EnableAutoscalingCommand(enableAutoscalingCommand).Execute()

Enable autoscaling

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
    enableAutoscalingCommand := *openapiclient.NewEnableAutoscalingCommand(int32(123), "AutoscalingGroupName_example", "Flavor_example") // EnableAutoscalingCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoscalingApi.AutoscalingEnable(context.Background()).EnableAutoscalingCommand(enableAutoscalingCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingApi.AutoscalingEnable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoscalingEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAutoscalingCommand** | [**EnableAutoscalingCommand**](EnableAutoscalingCommand.md) |  | 

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


## AutoscalingSync

> AutoscalingSync(ctx).AutoscalingSyncCommand(autoscalingSyncCommand).Execute()

Sync autoscaling

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
    autoscalingSyncCommand := *openapiclient.NewAutoscalingSyncCommand() // AutoscalingSyncCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoscalingApi.AutoscalingSync(context.Background()).AutoscalingSyncCommand(autoscalingSyncCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingApi.AutoscalingSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoscalingSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoscalingSyncCommand** | [**AutoscalingSyncCommand**](AutoscalingSyncCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

