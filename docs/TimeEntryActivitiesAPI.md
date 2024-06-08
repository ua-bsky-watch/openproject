# \TimeEntryActivitiesAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTimeEntriesActivity**](TimeEntryActivitiesAPI.md#GetTimeEntriesActivity) | **Get** /api/v3/time_entries/activity/{id} | View time entries activity



## GetTimeEntriesActivity

> TimeEntryActivityModel GetTimeEntriesActivity(ctx, id).Execute()

View time entries activity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func main() {
	id := int64(1) // int64 | Time entries activity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeEntryActivitiesAPI.GetTimeEntriesActivity(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntryActivitiesAPI.GetTimeEntriesActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeEntriesActivity`: TimeEntryActivityModel
	fmt.Fprintf(os.Stdout, "Response from `TimeEntryActivitiesAPI.GetTimeEntriesActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Time entries activity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeEntriesActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimeEntryActivityModel**](TimeEntryActivityModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
