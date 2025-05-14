# \PlannedCoursesAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPlannedCourseById**](PlannedCoursesAPI.md#CancelPlannedCourseById) | **Put** /planned_courses/{id}/cancel | Cancel a planned course.
[**CreatePlannedCourse**](PlannedCoursesAPI.md#CreatePlannedCourse) | **Post** /planned_courses | Create a planned course.
[**GetPlannedCourseById**](PlannedCoursesAPI.md#GetPlannedCourseById) | **Get** /planned_courses/{id} | Get a planned course record
[**GetPlannedCourses**](PlannedCoursesAPI.md#GetPlannedCourses) | **Get** /planned_courses | Get all planned course records
[**GetPlannedCoursesByCourseId**](PlannedCoursesAPI.md#GetPlannedCoursesByCourseId) | **Get** /courses/{course_id}/planned_courses | Get all planned course records of a single course
[**GetPlannedCoursesByIdAndCourseId**](PlannedCoursesAPI.md#GetPlannedCoursesByIdAndCourseId) | **Get** /courses/{course_id}/planned_courses/{id} | Get a planned course record of a single course
[**UpdatePlannedCourseById**](PlannedCoursesAPI.md#UpdatePlannedCourseById) | **Patch** /planned_courses/{id} | Update a planned course.



## CancelPlannedCourseById

> PlannedCourseWithIncludes CancelPlannedCourseById(ctx, id).Execute()

Cancel a planned course.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlannedCoursesAPI.CancelPlannedCourseById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlannedCoursesAPI.CancelPlannedCourseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelPlannedCourseById`: PlannedCourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PlannedCoursesAPI.CancelPlannedCourseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPlannedCourseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlannedCourseWithIncludes**](PlannedCourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlannedCourse

> PlannedCourseWithIncludes CreatePlannedCourse(ctx).PlannedCoursePayload(plannedCoursePayload).Execute()

Create a planned course.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	plannedCoursePayload := *openapiclient.NewPlannedCoursePayload(int32(123), "Type_example", NullableFloat32(123)) // PlannedCoursePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlannedCoursesAPI.CreatePlannedCourse(context.Background()).PlannedCoursePayload(plannedCoursePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlannedCoursesAPI.CreatePlannedCourse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlannedCourse`: PlannedCourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PlannedCoursesAPI.CreatePlannedCourse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlannedCourseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plannedCoursePayload** | [**PlannedCoursePayload**](PlannedCoursePayload.md) |  | 

### Return type

[**PlannedCourseWithIncludes**](PlannedCourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlannedCourseById

> PlannedCourseWithIncludes GetPlannedCourseById(ctx, id).Execute()

Get a planned course record

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlannedCoursesAPI.GetPlannedCourseById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlannedCoursesAPI.GetPlannedCourseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlannedCourseById`: PlannedCourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PlannedCoursesAPI.GetPlannedCourseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlannedCourseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlannedCourseWithIncludes**](PlannedCourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlannedCourses

> []PlannedCourseWithIncludes GetPlannedCourses(ctx).Search(search).Type_(type_).ParentsPublished(parentsPublished).PublishedPublic(publishedPublic).StartDateFrom(startDateFrom).StartDateUntil(startDateUntil).AvailabilityState(availabilityState).CourseId(courseId).Status(status).Sort(sort).Cursor(cursor).PerPage(perPage).Execute()

Get all planned course records

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	search := "search_example" // string | Filter results on search (optional)
	type_ := "type__example" // string | Filter results on type (optional)
	parentsPublished := "parentsPublished_example" // string | Filter results on parents_published (optional)
	publishedPublic := "publishedPublic_example" // string | Only show courses that are published and are either planned or in progress (optional)
	startDateFrom := "startDateFrom_example" // string | Filter results on start_date_from (optional)
	startDateUntil := "startDateUntil_example" // string | Filter results on start_date_until (optional)
	availabilityState := "availabilityState_example" // string | Filter results on availability_state (optional)
	courseId := []int32{int32(123)} // []int32 | Filter results on course_id (optional)
	status := []string{"Status_example"} // []string | Filter results on status (optional)
	sort := []string{"Sort_example"} // []string | Sort the results. Can change order by using `<sort_by>:<direction>` where `<direction>` is either `asc` or `desc` (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlannedCoursesAPI.GetPlannedCourses(context.Background()).Search(search).Type_(type_).ParentsPublished(parentsPublished).PublishedPublic(publishedPublic).StartDateFrom(startDateFrom).StartDateUntil(startDateUntil).AvailabilityState(availabilityState).CourseId(courseId).Status(status).Sort(sort).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlannedCoursesAPI.GetPlannedCourses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlannedCourses`: []PlannedCourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PlannedCoursesAPI.GetPlannedCourses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlannedCoursesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Filter results on search | 
 **type_** | **string** | Filter results on type | 
 **parentsPublished** | **string** | Filter results on parents_published | 
 **publishedPublic** | **string** | Only show courses that are published and are either planned or in progress | 
 **startDateFrom** | **string** | Filter results on start_date_from | 
 **startDateUntil** | **string** | Filter results on start_date_until | 
 **availabilityState** | **string** | Filter results on availability_state | 
 **courseId** | **[]int32** | Filter results on course_id | 
 **status** | **[]string** | Filter results on status | 
 **sort** | **[]string** | Sort the results. Can change order by using &#x60;&lt;sort_by&gt;:&lt;direction&gt;&#x60; where &#x60;&lt;direction&gt;&#x60; is either &#x60;asc&#x60; or &#x60;desc&#x60; | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]PlannedCourseWithIncludes**](PlannedCourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlannedCoursesByCourseId

> []PlannedCourseWithIncludes GetPlannedCoursesByCourseId(ctx, courseId).Search(search).Type_(type_).ParentsPublished(parentsPublished).PublishedPublic(publishedPublic).StartDateFrom(startDateFrom).StartDateUntil(startDateUntil).AvailabilityState(availabilityState).Status(status).Sort(sort).Cursor(cursor).PerPage(perPage).Execute()

Get all planned course records of a single course

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	courseId := int32(56) // int32 | 
	search := "search_example" // string | Filter results on search (optional)
	type_ := "type__example" // string | Filter results on type (optional)
	parentsPublished := "parentsPublished_example" // string | Filter results on parents_published (optional)
	publishedPublic := "publishedPublic_example" // string | Only show courses that are published and are either planned or in progress (optional)
	startDateFrom := "startDateFrom_example" // string | Filter results on start_date_from (optional)
	startDateUntil := "startDateUntil_example" // string | Filter results on start_date_until (optional)
	availabilityState := openapiclient.AvailabilityState("open") // AvailabilityState | Filter results on availability_state (optional)
	status := openapiclient.PlannedCourseStatus("planned") // PlannedCourseStatus | Filter results on status (optional)
	sort := []string{"Sort_example"} // []string | Sort the results. Can change order by using `<sort_by>:<direction>` where `<direction>` is either `asc` or `desc` (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlannedCoursesAPI.GetPlannedCoursesByCourseId(context.Background(), courseId).Search(search).Type_(type_).ParentsPublished(parentsPublished).PublishedPublic(publishedPublic).StartDateFrom(startDateFrom).StartDateUntil(startDateUntil).AvailabilityState(availabilityState).Status(status).Sort(sort).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlannedCoursesAPI.GetPlannedCoursesByCourseId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlannedCoursesByCourseId`: []PlannedCourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PlannedCoursesAPI.GetPlannedCoursesByCourseId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlannedCoursesByCourseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Filter results on search | 
 **type_** | **string** | Filter results on type | 
 **parentsPublished** | **string** | Filter results on parents_published | 
 **publishedPublic** | **string** | Only show courses that are published and are either planned or in progress | 
 **startDateFrom** | **string** | Filter results on start_date_from | 
 **startDateUntil** | **string** | Filter results on start_date_until | 
 **availabilityState** | [**AvailabilityState**](AvailabilityState.md) | Filter results on availability_state | 
 **status** | [**PlannedCourseStatus**](PlannedCourseStatus.md) | Filter results on status | 
 **sort** | **[]string** | Sort the results. Can change order by using &#x60;&lt;sort_by&gt;:&lt;direction&gt;&#x60; where &#x60;&lt;direction&gt;&#x60; is either &#x60;asc&#x60; or &#x60;desc&#x60; | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]PlannedCourseWithIncludes**](PlannedCourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlannedCoursesByIdAndCourseId

> PlannedCourseWithIncludes GetPlannedCoursesByIdAndCourseId(ctx, id, courseId).Execute()

Get a planned course record of a single course

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | 
	courseId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlannedCoursesAPI.GetPlannedCoursesByIdAndCourseId(context.Background(), id, courseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlannedCoursesAPI.GetPlannedCoursesByIdAndCourseId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlannedCoursesByIdAndCourseId`: PlannedCourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PlannedCoursesAPI.GetPlannedCoursesByIdAndCourseId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**courseId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlannedCoursesByIdAndCourseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlannedCourseWithIncludes**](PlannedCourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlannedCourseById

> PlannedCourseWithIncludes UpdatePlannedCourseById(ctx, id).PlannedCoursePatchPayload(plannedCoursePatchPayload).Execute()

Update a planned course.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | 
	plannedCoursePatchPayload := *openapiclient.NewPlannedCoursePatchPayload() // PlannedCoursePatchPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlannedCoursesAPI.UpdatePlannedCourseById(context.Background(), id).PlannedCoursePatchPayload(plannedCoursePatchPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlannedCoursesAPI.UpdatePlannedCourseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlannedCourseById`: PlannedCourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PlannedCoursesAPI.UpdatePlannedCourseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlannedCourseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **plannedCoursePatchPayload** | [**PlannedCoursePatchPayload**](PlannedCoursePatchPayload.md) |  | 

### Return type

[**PlannedCourseWithIncludes**](PlannedCourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

