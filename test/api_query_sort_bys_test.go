/*
OpenProject API V3 (Stable)

Testing QuerySortBysAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openproject

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func Test_openproject_QuerySortBysAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QuerySortBysAPIService ViewQuerySortBy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.QuerySortBysAPI.ViewQuerySortBy(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}