/*
Taikun - WebApi

Testing ImagesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/chnyda/testgoclient"
)

func Test_openapi_ImagesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImagesApiService ImagesAwsCommonImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesApi.ImagesAwsCommonImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesAwsPersonalImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesApi.ImagesAwsPersonalImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesAzureImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32
		var publisherName string
		var offer string
		var sku string

		resp, httpRes, err := apiClient.ImagesApi.ImagesAzureImages(context.Background(), cloudId, publisherName, offer, sku).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesAzurePersonalImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesApi.ImagesAzurePersonalImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesBindImagesToProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ImagesApi.ImagesBindImagesToProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesCommonGoogleImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesApi.ImagesCommonGoogleImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesGoogleImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32
		var type_ string

		resp, httpRes, err := apiClient.ImagesApi.ImagesGoogleImages(context.Background(), cloudId, type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesImageDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ImagesApi.ImagesImageDetails(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesOpenstackImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesApi.ImagesOpenstackImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesSelectedImagesForProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ImagesApi.ImagesSelectedImagesForProject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesTanzuImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.ImagesApi.ImagesTanzuImages(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ImagesUnbindImagesFromProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ImagesApi.ImagesUnbindImagesFromProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
