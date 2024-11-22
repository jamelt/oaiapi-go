// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
	"github.com/jamelt/openai-api/models/components"
)

// ListRunStepsQueryParamOrder - Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.
type ListRunStepsQueryParamOrder string

const (
	ListRunStepsQueryParamOrderAsc  ListRunStepsQueryParamOrder = "asc"
	ListRunStepsQueryParamOrderDesc ListRunStepsQueryParamOrder = "desc"
)

func (e ListRunStepsQueryParamOrder) ToPointer() *ListRunStepsQueryParamOrder {
	return &e
}
func (e *ListRunStepsQueryParamOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = ListRunStepsQueryParamOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListRunStepsQueryParamOrder: %v", v)
	}
}

type QueryParamInclude string

const (
	QueryParamIncludeStepDetailsToolCallsWildcardFileSearchResultsWildcardContent QueryParamInclude = "step_details.tool_calls[*].file_search.results[*].content"
)

func (e QueryParamInclude) ToPointer() *QueryParamInclude {
	return &e
}
func (e *QueryParamInclude) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "step_details.tool_calls[*].file_search.results[*].content":
		*e = QueryParamInclude(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamInclude: %v", v)
	}
}

type ListRunStepsRequest struct {
	// The ID of the thread the run and run steps belong to.
	ThreadID string `pathParam:"style=simple,explode=false,name=thread_id"`
	// The ID of the run the run steps belong to.
	RunID string `pathParam:"style=simple,explode=false,name=run_id"`
	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	//
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.
	//
	Order *ListRunStepsQueryParamOrder `default:"desc" queryParam:"style=form,explode=true,name=order"`
	// A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.
	//
	After *string `queryParam:"style=form,explode=true,name=after"`
	// A cursor for use in pagination. `before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with obj_foo, your subsequent call can include before=obj_foo in order to fetch the previous page of the list.
	//
	Before *string `queryParam:"style=form,explode=true,name=before"`
	// A list of additional fields to include in the response. Currently the only supported value is `step_details.tool_calls[*].file_search.results[*].content` to fetch the file search result content.
	//
	// See the [file search tool documentation](/docs/assistants/tools/file-search#customizing-file-search-settings) for more information.
	//
	Include []QueryParamInclude `queryParam:"style=form,explode=true,name=include[]"`
}

func (l ListRunStepsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListRunStepsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListRunStepsRequest) GetThreadID() string {
	if o == nil {
		return ""
	}
	return o.ThreadID
}

func (o *ListRunStepsRequest) GetRunID() string {
	if o == nil {
		return ""
	}
	return o.RunID
}

func (o *ListRunStepsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListRunStepsRequest) GetOrder() *ListRunStepsQueryParamOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListRunStepsRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListRunStepsRequest) GetBefore() *string {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *ListRunStepsRequest) GetInclude() []QueryParamInclude {
	if o == nil {
		return nil
	}
	return o.Include
}

type ListRunStepsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListRunStepsResponse *components.ListRunStepsResponse
}

func (o *ListRunStepsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListRunStepsResponse) GetListRunStepsResponse() *components.ListRunStepsResponse {
	if o == nil {
		return nil
	}
	return o.ListRunStepsResponse
}
