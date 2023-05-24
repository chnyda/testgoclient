# \AwsImagesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesAwsImages**](AwsImagesApi.md#ImagesAwsImages) | **Get** /api/v1/images/aws | Retrieve aws images



## ImagesAwsImages

> AwsImagesPostList ImagesAwsImages(ctx).AwsImagesPostListCommand(awsImagesPostListCommand).Execute()

Retrieve aws images

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
    awsImagesPostListCommand := *openapiclient.NewAwsImagesPostListCommand() // AwsImagesPostListCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AwsImagesApi.ImagesAwsImages(context.Background()).AwsImagesPostListCommand(awsImagesPostListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwsImagesApi.ImagesAwsImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesAwsImages`: AwsImagesPostList
    fmt.Fprintf(os.Stdout, "Response from `AwsImagesApi.ImagesAwsImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesAwsImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsImagesPostListCommand** | [**AwsImagesPostListCommand**](AwsImagesPostListCommand.md) |  | 

### Return type

[**AwsImagesPostList**](AwsImagesPostList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

