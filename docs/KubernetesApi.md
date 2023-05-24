# \KubernetesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubernetesAddK8sAlert**](KubernetesApi.md#KubernetesAddK8sAlert) | **Post** /api/v1/kubernetes/alert/{projectId} | Add k8s alert
[**KubernetesAddK8sEvents**](KubernetesApi.md#KubernetesAddK8sEvents) | **Post** /api/v1/kubernetes/event/{projectId} | Add k8s event
[**KubernetesAlertList**](KubernetesApi.md#KubernetesAlertList) | **Get** /api/v1/kubernetes/{projectId}/alerts | Retrieve a list of alerts for project
[**KubernetesCli**](KubernetesApi.md#KubernetesCli) | **Post** /api/v1/kubernetes/cli | Execute k8s web socket namespaced pod
[**KubernetesConfigMapList**](KubernetesApi.md#KubernetesConfigMapList) | **Get** /api/v1/kubernetes/{projectId}/configmap | Retrieve a list of k8s config map for all namespaces
[**KubernetesCrdList**](KubernetesApi.md#KubernetesCrdList) | **Get** /api/v1/kubernetes/{projectId}/crd | Retrieve a list of crd
[**KubernetesCronJobList**](KubernetesApi.md#KubernetesCronJobList) | **Get** /api/v1/kubernetes/{projectId}/cronjobs | Retrieve a list of k8s cron jobs for all namespaces
[**KubernetesDaemonSetList**](KubernetesApi.md#KubernetesDaemonSetList) | **Get** /api/v1/kubernetes/{projectId}/daemonset | Retrieve list of k8s daemonset
[**KubernetesDashboardList**](KubernetesApi.md#KubernetesDashboardList) | **Get** /api/v1/kubernetes/{projectId}/dashboard | Retrieve a list of crd
[**KubernetesDeploymentList**](KubernetesApi.md#KubernetesDeploymentList) | **Get** /api/v1/kubernetes/{projectId}/deployment | Retrieve a list of k8s deployment for all namespaces
[**KubernetesDescribeConfigMap**](KubernetesApi.md#KubernetesDescribeConfigMap) | **Post** /api/v1/kubernetes/describe/configmap | Describe configmap
[**KubernetesDescribeCrd**](KubernetesApi.md#KubernetesDescribeCrd) | **Post** /api/v1/kubernetes/describe/crd | Describe crd
[**KubernetesDescribeCronjob**](KubernetesApi.md#KubernetesDescribeCronjob) | **Post** /api/v1/kubernetes/describe/cronjob | Describe cronjob
[**KubernetesDescribeDaemonSet**](KubernetesApi.md#KubernetesDescribeDaemonSet) | **Post** /api/v1/kubernetes/describe/daemonset | Describe daemonset
[**KubernetesDescribeDeployment**](KubernetesApi.md#KubernetesDescribeDeployment) | **Post** /api/v1/kubernetes/describe/deployment | Describe deployment
[**KubernetesDescribeIngress**](KubernetesApi.md#KubernetesDescribeIngress) | **Post** /api/v1/kubernetes/describe/ingress | Describe ingress
[**KubernetesDescribeJob**](KubernetesApi.md#KubernetesDescribeJob) | **Post** /api/v1/kubernetes/describe/job | Describe job
[**KubernetesDescribeNetworkPolicy**](KubernetesApi.md#KubernetesDescribeNetworkPolicy) | **Post** /api/v1/kubernetes/describe/network-policy | Describe network policy
[**KubernetesDescribeNode**](KubernetesApi.md#KubernetesDescribeNode) | **Post** /api/v1/kubernetes/describe/node | Describe node
[**KubernetesDescribePdb**](KubernetesApi.md#KubernetesDescribePdb) | **Post** /api/v1/kubernetes/describe/pdb | Describe pdb
[**KubernetesDescribePod**](KubernetesApi.md#KubernetesDescribePod) | **Post** /api/v1/kubernetes/describe/pod | Describe pod
[**KubernetesDescribePvc**](KubernetesApi.md#KubernetesDescribePvc) | **Post** /api/v1/kubernetes/describe/pvc | Describe pvc
[**KubernetesDescribeSecret**](KubernetesApi.md#KubernetesDescribeSecret) | **Post** /api/v1/kubernetes/describe/secret | Describe secret
[**KubernetesDescribeService**](KubernetesApi.md#KubernetesDescribeService) | **Post** /api/v1/kubernetes/describe/service | Describe service
[**KubernetesDescribeStorageClass**](KubernetesApi.md#KubernetesDescribeStorageClass) | **Post** /api/v1/kubernetes/describe/storageclass | Describe storage class
[**KubernetesDescribeSts**](KubernetesApi.md#KubernetesDescribeSts) | **Post** /api/v1/kubernetes/describe/sts | Describe stateful set
[**KubernetesDownload**](KubernetesApi.md#KubernetesDownload) | **Get** /api/v1/kubernetes/{projectId}/download | Download kube config file
[**KubernetesGetSupportedList**](KubernetesApi.md#KubernetesGetSupportedList) | **Get** /api/v1/kubernetes/supported/list | Retrieve Taikun supported kubernetes versions
[**KubernetesHelmReleaseList**](KubernetesApi.md#KubernetesHelmReleaseList) | **Get** /api/v1/kubernetes/{projectId}/helmreleases | Retrieve a list of k8s helm releases for all namespaces
[**KubernetesIngressList**](KubernetesApi.md#KubernetesIngressList) | **Get** /api/v1/kubernetes/{projectId}/ingress | Retrieve a list of k8s ingress for all namespaces
[**KubernetesInteractiveShell**](KubernetesApi.md#KubernetesInteractiveShell) | **Post** /api/v1/kubernetes/interactive-shell | Produce interactive shell command
[**KubernetesJobsList**](KubernetesApi.md#KubernetesJobsList) | **Get** /api/v1/kubernetes/{projectId}/jobs | Retrieve a list of k8s jobs for all namespaces
[**KubernetesKillPod**](KubernetesApi.md#KubernetesKillPod) | **Post** /api/v1/kubernetes/{projectId}/deletepod/{metadataName}/{namespace} | 
[**KubernetesKubeConfig**](KubernetesApi.md#KubernetesKubeConfig) | **Get** /api/v1/kubernetes/{projectId}/kubeconfig | Retrieve kube config file
[**KubernetesNamespaceList**](KubernetesApi.md#KubernetesNamespaceList) | **Get** /api/v1/kubernetes/{projectId}/namespaces | Retrieve kube config file
[**KubernetesNetworkPolicyList**](KubernetesApi.md#KubernetesNetworkPolicyList) | **Get** /api/v1/kubernetes/{projectId}/network-policies | Retrieve a list of k8s network-policies for all namespaces
[**KubernetesNodeList**](KubernetesApi.md#KubernetesNodeList) | **Get** /api/v1/kubernetes/{projectId}/node | Retrieve a list of k8s node
[**KubernetesOverview**](KubernetesApi.md#KubernetesOverview) | **Get** /api/v1/kubernetes/overview | Overview kubernetes nodes and pods by organization id
[**KubernetesPatchCrd**](KubernetesApi.md#KubernetesPatchCrd) | **Post** /api/v1/kubernetes/patch/crd | Patch crd
[**KubernetesPatchCronJob**](KubernetesApi.md#KubernetesPatchCronJob) | **Post** /api/v1/kubernetes/patch/cronjob | Patch cron-job
[**KubernetesPatchIngress**](KubernetesApi.md#KubernetesPatchIngress) | **Post** /api/v1/kubernetes/patch/ingress | Patch ingress
[**KubernetesPatchJob**](KubernetesApi.md#KubernetesPatchJob) | **Post** /api/v1/kubernetes/patch/job | Patch job
[**KubernetesPatchNode**](KubernetesApi.md#KubernetesPatchNode) | **Post** /api/v1/kubernetes/patch/node | Patch node
[**KubernetesPatchPdb**](KubernetesApi.md#KubernetesPatchPdb) | **Post** /api/v1/kubernetes/patch/pdb | Patch pdb
[**KubernetesPatchPod**](KubernetesApi.md#KubernetesPatchPod) | **Post** /api/v1/kubernetes/patch/pod | Patch pod
[**KubernetesPatchPvc**](KubernetesApi.md#KubernetesPatchPvc) | **Post** /api/v1/kubernetes/patch/pvc | Patch pvc
[**KubernetesPatchSecret**](KubernetesApi.md#KubernetesPatchSecret) | **Post** /api/v1/kubernetes/patch/secret | Patch secret
[**KubernetesPatchSts**](KubernetesApi.md#KubernetesPatchSts) | **Post** /api/v1/kubernetes/patch/sts | Patch sts
[**KubernetesPdbList**](KubernetesApi.md#KubernetesPdbList) | **Get** /api/v1/kubernetes/{projectId}/pdb | Retrieve a list of k8s pdb for all namespaces
[**KubernetesPodList**](KubernetesApi.md#KubernetesPodList) | **Get** /api/v1/kubernetes/{projectId}/pod | Retrieve a list of k8s pod for all namespaces
[**KubernetesPodLogs**](KubernetesApi.md#KubernetesPodLogs) | **Post** /api/v1/kubernetes/podLogs | Retrieve k8s pod logs
[**KubernetesPvcList**](KubernetesApi.md#KubernetesPvcList) | **Get** /api/v1/kubernetes/{projectId}/pvc | Retrieve a list of k8s pvc for all namespaces
[**KubernetesRemovealerts**](KubernetesApi.md#KubernetesRemovealerts) | **Post** /api/v1/kubernetes/removealerts | Remove k8s alerts
[**KubernetesRestartDaemonSet**](KubernetesApi.md#KubernetesRestartDaemonSet) | **Post** /api/v1/kubernetes/restart/daemonset | Restart daemon set
[**KubernetesRestartDeployment**](KubernetesApi.md#KubernetesRestartDeployment) | **Post** /api/v1/kubernetes/restart/deployment | Restart deployment
[**KubernetesRestartSts**](KubernetesApi.md#KubernetesRestartSts) | **Post** /api/v1/kubernetes/restart/sts | Restart stateful set
[**KubernetesSecretList**](KubernetesApi.md#KubernetesSecretList) | **Get** /api/v1/kubernetes/{projectId}/secret | Retrieve a list of k8s secret for all namespaces
[**KubernetesServiceList**](KubernetesApi.md#KubernetesServiceList) | **Get** /api/v1/kubernetes/{projectId}/service | Retrieve a list of k8s service for all namespaces
[**KubernetesSilenceManager**](KubernetesApi.md#KubernetesSilenceManager) | **Post** /api/v1/kubernetes/silencemanager | Silence management for k8s alerts
[**KubernetesStorageClassList**](KubernetesApi.md#KubernetesStorageClassList) | **Get** /api/v1/kubernetes/{projectId}/storageclass | Retrieve a list of k8s storageclass for all namespaces
[**KubernetesStsList**](KubernetesApi.md#KubernetesStsList) | **Get** /api/v1/kubernetes/{projectId}/sts | Retrieve a list of k8s sts for all namespaces
[**KubernetesUpdateAlert**](KubernetesApi.md#KubernetesUpdateAlert) | **Put** /api/v1/kubernetes/updatealert/{alertId} | Update k8s alert



## KubernetesAddK8sAlert

> KubernetesAddK8sAlert(ctx, projectId).CreateAlertDto(createAlertDto).Execute()

Add k8s alert

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
    projectId := int32(56) // int32 | 
    createAlertDto := *openapiclient.NewCreateAlertDto() // CreateAlertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesAddK8sAlert(context.Background(), projectId).CreateAlertDto(createAlertDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesAddK8sAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesAddK8sAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAlertDto** | [**CreateAlertDto**](CreateAlertDto.md) |  | 

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


## KubernetesAddK8sEvents

> KubernetesAddK8sEvents(ctx, projectId).KubernetesEventCreateDto(kubernetesEventCreateDto).Execute()

Add k8s event

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
    projectId := int32(56) // int32 | 
    kubernetesEventCreateDto := *openapiclient.NewKubernetesEventCreateDto() // KubernetesEventCreateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesAddK8sEvents(context.Background(), projectId).KubernetesEventCreateDto(kubernetesEventCreateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesAddK8sEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesAddK8sEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesEventCreateDto** | [**KubernetesEventCreateDto**](KubernetesEventCreateDto.md) |  | 

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


## KubernetesAlertList

> KubernetesAlertList KubernetesAlertList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Type_(type_).StartDate(startDate).EndDate(endDate).Execute()

Retrieve a list of alerts for project

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesAlertList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Type_(type_).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesAlertList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesAlertList`: KubernetesAlertList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesAlertList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesAlertListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **type_** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 

### Return type

[**KubernetesAlertList**](KubernetesAlertList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCli

> string KubernetesCli(ctx).KubernetesCliCommand(kubernetesCliCommand).Execute()

Execute k8s web socket namespaced pod

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
    kubernetesCliCommand := *openapiclient.NewKubernetesCliCommand() // KubernetesCliCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesCli(context.Background()).KubernetesCliCommand(kubernetesCliCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesCli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesCli`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesCli`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCliCommand** | [**KubernetesCliCommand**](KubernetesCliCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesConfigMapList

> ConfigMaps KubernetesConfigMapList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s config map for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesConfigMapList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesConfigMapList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesConfigMapList`: ConfigMaps
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesConfigMapList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesConfigMapListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**ConfigMaps**](ConfigMaps.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCrdList

> interface{} KubernetesCrdList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of crd

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesCrdList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesCrdList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesCrdList`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesCrdList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCrdListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCronJobList

> KubernetesCronJobsList KubernetesCronJobList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s cron jobs for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesCronJobList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesCronJobList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesCronJobList`: KubernetesCronJobsList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesCronJobList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCronJobListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**KubernetesCronJobsList**](KubernetesCronJobsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDaemonSetList

> DaemonSets KubernetesDaemonSetList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve list of k8s daemonset

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDaemonSetList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDaemonSetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDaemonSetList`: DaemonSets
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDaemonSetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDaemonSetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**DaemonSets**](DaemonSets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDashboardList

> string KubernetesDashboardList(ctx, projectId).Execute()

Retrieve a list of crd

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
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDashboardList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDashboardList`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDashboardList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDeploymentList

> Deployments KubernetesDeploymentList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s deployment for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDeploymentList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDeploymentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDeploymentList`: Deployments
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDeploymentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDeploymentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Deployments**](Deployments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeConfigMap

> string KubernetesDescribeConfigMap(ctx).DescribeConfigMapCommand(describeConfigMapCommand).Execute()

Describe configmap

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
    describeConfigMapCommand := *openapiclient.NewDescribeConfigMapCommand(int32(123), "Name_example", "Namespace_example") // DescribeConfigMapCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeConfigMap(context.Background()).DescribeConfigMapCommand(describeConfigMapCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeConfigMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeConfigMap`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeConfigMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeConfigMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeConfigMapCommand** | [**DescribeConfigMapCommand**](DescribeConfigMapCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeCrd

> string KubernetesDescribeCrd(ctx).DescribeCrdCommand(describeCrdCommand).Execute()

Describe crd

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
    describeCrdCommand := *openapiclient.NewDescribeCrdCommand(int32(123), "Name_example") // DescribeCrdCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeCrd(context.Background()).DescribeCrdCommand(describeCrdCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeCrd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeCrd`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeCrd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeCrdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeCrdCommand** | [**DescribeCrdCommand**](DescribeCrdCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeCronjob

> string KubernetesDescribeCronjob(ctx).DescribeCronJobCommand(describeCronJobCommand).Execute()

Describe cronjob

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
    describeCronJobCommand := *openapiclient.NewDescribeCronJobCommand(int32(123), "Name_example", "Namespace_example") // DescribeCronJobCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeCronjob(context.Background()).DescribeCronJobCommand(describeCronJobCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeCronjob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeCronjob`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeCronjob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeCronjobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeCronJobCommand** | [**DescribeCronJobCommand**](DescribeCronJobCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeDaemonSet

> string KubernetesDescribeDaemonSet(ctx).DescribeDaemonSetCommand(describeDaemonSetCommand).Execute()

Describe daemonset

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
    describeDaemonSetCommand := *openapiclient.NewDescribeDaemonSetCommand(int32(123), "Name_example", "Namespace_example") // DescribeDaemonSetCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeDaemonSet(context.Background()).DescribeDaemonSetCommand(describeDaemonSetCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeDaemonSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeDaemonSet`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeDaemonSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeDaemonSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeDaemonSetCommand** | [**DescribeDaemonSetCommand**](DescribeDaemonSetCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeDeployment

> string KubernetesDescribeDeployment(ctx).DescribeDeploymentCommand(describeDeploymentCommand).Execute()

Describe deployment

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
    describeDeploymentCommand := *openapiclient.NewDescribeDeploymentCommand(int32(123), "Name_example", "Namespace_example") // DescribeDeploymentCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeDeployment(context.Background()).DescribeDeploymentCommand(describeDeploymentCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeDeployment`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeDeploymentCommand** | [**DescribeDeploymentCommand**](DescribeDeploymentCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeIngress

> string KubernetesDescribeIngress(ctx).DescribeIngressCommand(describeIngressCommand).Execute()

Describe ingress

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
    describeIngressCommand := *openapiclient.NewDescribeIngressCommand(int32(123), "Name_example", "Namespace_example") // DescribeIngressCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeIngress(context.Background()).DescribeIngressCommand(describeIngressCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeIngress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeIngress`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeIngress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeIngressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeIngressCommand** | [**DescribeIngressCommand**](DescribeIngressCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeJob

> string KubernetesDescribeJob(ctx).DescribeJobCommand(describeJobCommand).Execute()

Describe job

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
    describeJobCommand := *openapiclient.NewDescribeJobCommand(int32(123), "Name_example", "Namespace_example") // DescribeJobCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeJob(context.Background()).DescribeJobCommand(describeJobCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeJob`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeJobCommand** | [**DescribeJobCommand**](DescribeJobCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeNetworkPolicy

> string KubernetesDescribeNetworkPolicy(ctx).DescribeNetworkPolicyCommand(describeNetworkPolicyCommand).Execute()

Describe network policy

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
    describeNetworkPolicyCommand := *openapiclient.NewDescribeNetworkPolicyCommand(int32(123), "Name_example", "Namespace_example") // DescribeNetworkPolicyCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeNetworkPolicy(context.Background()).DescribeNetworkPolicyCommand(describeNetworkPolicyCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeNetworkPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeNetworkPolicy`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeNetworkPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeNetworkPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeNetworkPolicyCommand** | [**DescribeNetworkPolicyCommand**](DescribeNetworkPolicyCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeNode

> string KubernetesDescribeNode(ctx).DescribeNodeCommand(describeNodeCommand).Execute()

Describe node

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
    describeNodeCommand := *openapiclient.NewDescribeNodeCommand(int32(123), "Name_example") // DescribeNodeCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeNode(context.Background()).DescribeNodeCommand(describeNodeCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeNode`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeNodeCommand** | [**DescribeNodeCommand**](DescribeNodeCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribePdb

> string KubernetesDescribePdb(ctx).DescribePodDisruptionCommand(describePodDisruptionCommand).Execute()

Describe pdb

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
    describePodDisruptionCommand := *openapiclient.NewDescribePodDisruptionCommand(int32(123), "Name_example", "Namespace_example") // DescribePodDisruptionCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribePdb(context.Background()).DescribePodDisruptionCommand(describePodDisruptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribePdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribePdb`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribePdb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribePdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describePodDisruptionCommand** | [**DescribePodDisruptionCommand**](DescribePodDisruptionCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribePod

> string KubernetesDescribePod(ctx).DescribePodCommand(describePodCommand).Execute()

Describe pod

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
    describePodCommand := *openapiclient.NewDescribePodCommand(int32(123), "Name_example", "Namespace_example") // DescribePodCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribePod(context.Background()).DescribePodCommand(describePodCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribePod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribePod`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribePod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describePodCommand** | [**DescribePodCommand**](DescribePodCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribePvc

> string KubernetesDescribePvc(ctx).DescribePvcCommand(describePvcCommand).Execute()

Describe pvc

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
    describePvcCommand := *openapiclient.NewDescribePvcCommand(int32(123), "Name_example", "Namespace_example") // DescribePvcCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribePvc(context.Background()).DescribePvcCommand(describePvcCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribePvc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribePvc`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribePvc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribePvcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describePvcCommand** | [**DescribePvcCommand**](DescribePvcCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeSecret

> string KubernetesDescribeSecret(ctx).DescribeSecretCommand(describeSecretCommand).Execute()

Describe secret

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
    describeSecretCommand := *openapiclient.NewDescribeSecretCommand(int32(123), "Name_example", "Namespace_example") // DescribeSecretCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeSecret(context.Background()).DescribeSecretCommand(describeSecretCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeSecret`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeSecretCommand** | [**DescribeSecretCommand**](DescribeSecretCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeService

> string KubernetesDescribeService(ctx).DescribeServiceCommand(describeServiceCommand).Execute()

Describe service

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
    describeServiceCommand := *openapiclient.NewDescribeServiceCommand(int32(123), "Name_example", "Namespace_example") // DescribeServiceCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeService(context.Background()).DescribeServiceCommand(describeServiceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeService`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeServiceCommand** | [**DescribeServiceCommand**](DescribeServiceCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeStorageClass

> string KubernetesDescribeStorageClass(ctx).DescribeStorageClassCommand(describeStorageClassCommand).Execute()

Describe storage class

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
    describeStorageClassCommand := *openapiclient.NewDescribeStorageClassCommand(int32(123), "Name_example") // DescribeStorageClassCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeStorageClass(context.Background()).DescribeStorageClassCommand(describeStorageClassCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeStorageClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeStorageClass`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeStorageClass`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeStorageClassCommand** | [**DescribeStorageClassCommand**](DescribeStorageClassCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeSts

> string KubernetesDescribeSts(ctx).DescribeStsCommand(describeStsCommand).Execute()

Describe stateful set

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
    describeStsCommand := *openapiclient.NewDescribeStsCommand(int32(123), "Name_example", "Namespace_example") // DescribeStsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDescribeSts(context.Background()).DescribeStsCommand(describeStsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDescribeSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeSts`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDescribeSts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeStsCommand** | [**DescribeStsCommand**](DescribeStsCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDownload

> interface{} KubernetesDownload(ctx, projectId).Execute()

Download kube config file

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
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesDownload(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDownload`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetSupportedList

> []KubernetesVersionListDto KubernetesGetSupportedList(ctx).Execute()

Retrieve Taikun supported kubernetes versions

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
    resp, r, err := apiClient.KubernetesApi.KubernetesGetSupportedList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesGetSupportedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetSupportedList`: []KubernetesVersionListDto
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesGetSupportedList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetSupportedListRequest struct via the builder pattern


### Return type

[**[]KubernetesVersionListDto**](KubernetesVersionListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesHelmReleaseList

> HelmReleasesList KubernetesHelmReleaseList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s helm releases for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesHelmReleaseList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesHelmReleaseList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesHelmReleaseList`: HelmReleasesList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesHelmReleaseList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesHelmReleaseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**HelmReleasesList**](HelmReleasesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesIngressList

> Ingresses KubernetesIngressList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s ingress for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesIngressList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesIngressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesIngressList`: Ingresses
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesIngressList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesIngressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Ingresses**](Ingresses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesInteractiveShell

> string KubernetesInteractiveShell(ctx).InteractiveShellSendCommand(interactiveShellSendCommand).Execute()

Produce interactive shell command

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
    interactiveShellSendCommand := *openapiclient.NewInteractiveShellSendCommand(int32(123), "Token_example") // InteractiveShellSendCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesInteractiveShell(context.Background()).InteractiveShellSendCommand(interactiveShellSendCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesInteractiveShell``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesInteractiveShell`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesInteractiveShell`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesInteractiveShellRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interactiveShellSendCommand** | [**InteractiveShellSendCommand**](InteractiveShellSendCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesJobsList

> KubernetesJobList KubernetesJobsList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s jobs for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesJobsList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesJobsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesJobsList`: KubernetesJobList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesJobsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesJobsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**KubernetesJobList**](KubernetesJobList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesKillPod

> KubernetesKillPod(ctx, projectId, metadataName, namespace).Execute()



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
    projectId := int32(56) // int32 | 
    metadataName := "metadataName_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesKillPod(context.Background(), projectId, metadataName, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesKillPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**metadataName** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesKillPodRequest struct via the builder pattern


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


## KubernetesKubeConfig

> KubeConfigResponse KubernetesKubeConfig(ctx, projectId).Execute()

Retrieve kube config file

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
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesKubeConfig(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesKubeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesKubeConfig`: KubeConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesKubeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesKubeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubeConfigResponse**](KubeConfigResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesNamespaceList

> []string KubernetesNamespaceList(ctx, projectId).Execute()

Retrieve kube config file

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
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesNamespaceList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesNamespaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesNamespaceList`: []string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesNamespaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNamespaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesNetworkPolicyList

> NetworkPolicies KubernetesNetworkPolicyList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s network-policies for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesNetworkPolicyList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesNetworkPolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesNetworkPolicyList`: NetworkPolicies
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesNetworkPolicyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNetworkPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**NetworkPolicies**](NetworkPolicies.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesNodeList

> interface{} KubernetesNodeList(ctx, projectId).SearchId(searchId).Execute()

Retrieve a list of k8s node

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
    projectId := int32(56) // int32 | 
    searchId := "searchId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesNodeList(context.Background(), projectId).SearchId(searchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesNodeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesNodeList`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesNodeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNodeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchId** | **string** |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesOverview

> string KubernetesOverview(ctx).OrganizationId(organizationId).Execute()

Overview kubernetes nodes and pods by organization id

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesOverview(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesOverview`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchCrd

> KubernetesPatchCrd(ctx).PatchCrdCommand(patchCrdCommand).Execute()

Patch crd

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
    patchCrdCommand := *openapiclient.NewPatchCrdCommand(int32(123), "Yaml_example", "Name_example") // PatchCrdCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchCrd(context.Background()).PatchCrdCommand(patchCrdCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchCrd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchCrdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchCrdCommand** | [**PatchCrdCommand**](PatchCrdCommand.md) |  | 

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


## KubernetesPatchCronJob

> KubernetesPatchCronJob(ctx).PatchCronJobCommand(patchCronJobCommand).Execute()

Patch cron-job

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
    patchCronJobCommand := *openapiclient.NewPatchCronJobCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchCronJobCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchCronJob(context.Background()).PatchCronJobCommand(patchCronJobCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchCronJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchCronJobCommand** | [**PatchCronJobCommand**](PatchCronJobCommand.md) |  | 

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


## KubernetesPatchIngress

> KubernetesPatchIngress(ctx).PatchIngressCommand(patchIngressCommand).Execute()

Patch ingress

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
    patchIngressCommand := *openapiclient.NewPatchIngressCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchIngressCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchIngress(context.Background()).PatchIngressCommand(patchIngressCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchIngress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchIngressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchIngressCommand** | [**PatchIngressCommand**](PatchIngressCommand.md) |  | 

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


## KubernetesPatchJob

> KubernetesPatchJob(ctx).PatchJobCommand(patchJobCommand).Execute()

Patch job

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
    patchJobCommand := *openapiclient.NewPatchJobCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchJobCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchJob(context.Background()).PatchJobCommand(patchJobCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchJobCommand** | [**PatchJobCommand**](PatchJobCommand.md) |  | 

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


## KubernetesPatchNode

> KubernetesPatchNode(ctx).PatchNodeCommand(patchNodeCommand).Execute()

Patch node

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
    patchNodeCommand := *openapiclient.NewPatchNodeCommand(int32(123), "Name_example") // PatchNodeCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchNode(context.Background()).PatchNodeCommand(patchNodeCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchNodeCommand** | [**PatchNodeCommand**](PatchNodeCommand.md) |  | 

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


## KubernetesPatchPdb

> KubernetesPatchPdb(ctx).PatchPdbCommand(patchPdbCommand).Execute()

Patch pdb

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
    patchPdbCommand := *openapiclient.NewPatchPdbCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchPdbCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchPdb(context.Background()).PatchPdbCommand(patchPdbCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchPdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchPdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchPdbCommand** | [**PatchPdbCommand**](PatchPdbCommand.md) |  | 

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


## KubernetesPatchPod

> KubernetesPatchPod(ctx).PatchPodCommand(patchPodCommand).Execute()

Patch pod

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
    patchPodCommand := *openapiclient.NewPatchPodCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchPodCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchPod(context.Background()).PatchPodCommand(patchPodCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchPodCommand** | [**PatchPodCommand**](PatchPodCommand.md) |  | 

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


## KubernetesPatchPvc

> KubernetesPatchPvc(ctx).PatchPvcCommand(patchPvcCommand).Execute()

Patch pvc

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
    patchPvcCommand := *openapiclient.NewPatchPvcCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchPvcCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchPvc(context.Background()).PatchPvcCommand(patchPvcCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchPvc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchPvcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchPvcCommand** | [**PatchPvcCommand**](PatchPvcCommand.md) |  | 

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


## KubernetesPatchSecret

> KubernetesPatchSecret(ctx).PatchSecretCommand(patchSecretCommand).Execute()

Patch secret

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
    patchSecretCommand := *openapiclient.NewPatchSecretCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchSecretCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchSecret(context.Background()).PatchSecretCommand(patchSecretCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchSecretCommand** | [**PatchSecretCommand**](PatchSecretCommand.md) |  | 

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


## KubernetesPatchSts

> KubernetesPatchSts(ctx).PatchStsCommand(patchStsCommand).Execute()

Patch sts

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
    patchStsCommand := *openapiclient.NewPatchStsCommand(int32(123), "Yaml_example", "Name_example", "Namespace_example") // PatchStsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesPatchSts(context.Background()).PatchStsCommand(patchStsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPatchSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchStsCommand** | [**PatchStsCommand**](PatchStsCommand.md) |  | 

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


## KubernetesPdbList

> PodDisruptions KubernetesPdbList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s pdb for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesPdbList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPdbList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesPdbList`: PodDisruptions
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesPdbList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPdbListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**PodDisruptions**](PodDisruptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPodList

> Pods KubernetesPodList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s pod for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesPodList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPodList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesPodList`: Pods
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesPodList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPodListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Pods**](Pods.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPodLogs

> interface{} KubernetesPodLogs(ctx).KubernetesPodLogsCommand(kubernetesPodLogsCommand).Execute()

Retrieve k8s pod logs

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
    kubernetesPodLogsCommand := *openapiclient.NewKubernetesPodLogsCommand() // KubernetesPodLogsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesPodLogs(context.Background()).KubernetesPodLogsCommand(kubernetesPodLogsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPodLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesPodLogs`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesPodLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPodLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesPodLogsCommand** | [**KubernetesPodLogsCommand**](KubernetesPodLogsCommand.md) |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPvcList

> Pvcs KubernetesPvcList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s pvc for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesPvcList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesPvcList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesPvcList`: Pvcs
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesPvcList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPvcListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Pvcs**](Pvcs.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesRemovealerts

> KubernetesRemovealerts(ctx).DeleteAlertCommand(deleteAlertCommand).Execute()

Remove k8s alerts

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
    deleteAlertCommand := *openapiclient.NewDeleteAlertCommand(int32(123)) // DeleteAlertCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesRemovealerts(context.Background()).DeleteAlertCommand(deleteAlertCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesRemovealerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRemovealertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAlertCommand** | [**DeleteAlertCommand**](DeleteAlertCommand.md) |  | 

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


## KubernetesRestartDaemonSet

> KubernetesRestartDaemonSet(ctx).RestartDaemonSetCommand(restartDaemonSetCommand).Execute()

Restart daemon set

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
    restartDaemonSetCommand := *openapiclient.NewRestartDaemonSetCommand(int32(123), "Name_example", "Namespace_example") // RestartDaemonSetCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesRestartDaemonSet(context.Background()).RestartDaemonSetCommand(restartDaemonSetCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesRestartDaemonSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRestartDaemonSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restartDaemonSetCommand** | [**RestartDaemonSetCommand**](RestartDaemonSetCommand.md) |  | 

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


## KubernetesRestartDeployment

> KubernetesRestartDeployment(ctx).RestartDeploymentCommand(restartDeploymentCommand).Execute()

Restart deployment

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
    restartDeploymentCommand := *openapiclient.NewRestartDeploymentCommand(int32(123), "Name_example", "Namespace_example") // RestartDeploymentCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesRestartDeployment(context.Background()).RestartDeploymentCommand(restartDeploymentCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesRestartDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRestartDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restartDeploymentCommand** | [**RestartDeploymentCommand**](RestartDeploymentCommand.md) |  | 

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


## KubernetesRestartSts

> KubernetesRestartSts(ctx).RestartStsCommand(restartStsCommand).Execute()

Restart stateful set

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
    restartStsCommand := *openapiclient.NewRestartStsCommand(int32(123), "Name_example", "Namespace_example") // RestartStsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesRestartSts(context.Background()).RestartStsCommand(restartStsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesRestartSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRestartStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restartStsCommand** | [**RestartStsCommand**](RestartStsCommand.md) |  | 

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


## KubernetesSecretList

> Secrets KubernetesSecretList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s secret for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesSecretList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesSecretList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesSecretList`: Secrets
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesSecretList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesSecretListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Secrets**](Secrets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesServiceList

> Services KubernetesServiceList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s service for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesServiceList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesServiceList`: Services
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Services**](Services.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesSilenceManager

> KubernetesSilenceManager(ctx).SilenceOperationsCommand(silenceOperationsCommand).Execute()

Silence management for k8s alerts

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
    silenceOperationsCommand := *openapiclient.NewSilenceOperationsCommand(int32(123), "Mode_example") // SilenceOperationsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesSilenceManager(context.Background()).SilenceOperationsCommand(silenceOperationsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesSilenceManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesSilenceManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **silenceOperationsCommand** | [**SilenceOperationsCommand**](SilenceOperationsCommand.md) |  | 

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


## KubernetesStorageClassList

> StorageClasses KubernetesStorageClassList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s storageclass for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesStorageClassList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesStorageClassList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesStorageClassList`: StorageClasses
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesStorageClassList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesStorageClassListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**StorageClasses**](StorageClasses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesStsList

> StsList KubernetesStsList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s sts for all namespaces

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.KubernetesStsList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesStsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesStsList`: StsList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.KubernetesStsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesStsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**StsList**](StsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesUpdateAlert

> KubernetesUpdateAlert(ctx, alertId).UpdateKubernetesAlertDto(updateKubernetesAlertDto).Execute()

Update k8s alert

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
    alertId := int32(56) // int32 | 
    updateKubernetesAlertDto := *openapiclient.NewUpdateKubernetesAlertDto() // UpdateKubernetesAlertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesApi.KubernetesUpdateAlert(context.Background(), alertId).UpdateKubernetesAlertDto(updateKubernetesAlertDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.KubernetesUpdateAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateKubernetesAlertDto** | [**UpdateKubernetesAlertDto**](UpdateKubernetesAlertDto.md) |  | 

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

