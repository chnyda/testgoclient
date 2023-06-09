/*
Taikun - WebApi

Testing OperationCredentialsApiService

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

func Test_openapi_OperationCredentialsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OperationCredentialsApiService OpscredentialsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OperationCredentialsApi.OpscredentialsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OperationCredentialsApiService OpscredentialsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.OperationCredentialsApi.OpscredentialsDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OperationCredentialsApiService OpscredentialsLockManager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OperationCredentialsApi.OpscredentialsLockManager(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OperationCredentialsApiService OpscredentialsMakeDefault", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OperationCredentialsApi.OpscredentialsMakeDefault(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
