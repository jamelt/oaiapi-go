// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
	"github.com/jamelt/openai-api/models/components"
)

// ListVectorStoreFilesQueryParamOrder - Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.
type ListVectorStoreFilesQueryParamOrder string

const (
	ListVectorStoreFilesQueryParamOrderAsc  ListVectorStoreFilesQueryParamOrder = "asc"
	ListVectorStoreFilesQueryParamOrderDesc ListVectorStoreFilesQueryParamOrder = "desc"
)

func (e ListVectorStoreFilesQueryParamOrder) ToPointer() *ListVectorStoreFilesQueryParamOrder {
	return &e
}
func (e *ListVectorStoreFilesQueryParamOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = ListVectorStoreFilesQueryParamOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListVectorStoreFilesQueryParamOrder: %v", v)
	}
}

// QueryParamFilter - Filter by file status. One of `in_progress`, `completed`, `failed`, `cancelled`.
type QueryParamFilter string

const (
	QueryParamFilterInProgress QueryParamFilter = "in_progress"
	QueryParamFilterCompleted  QueryParamFilter = "completed"
	QueryParamFilterFailed     QueryParamFilter = "failed"
	QueryParamFilterCancelled  QueryParamFilter = "cancelled"
)

func (e QueryParamFilter) ToPointer() *QueryParamFilter {
	return &e
}
func (e *QueryParamFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "in_progress":
		fallthrough
	case "completed":
		fallthrough
	case "failed":
		fallthrough
	case "cancelled":
		*e = QueryParamFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamFilter: %v", v)
	}
}

type ListVectorStoreFilesRequest struct {
	// The ID of the vector store that the files belong to.
	VectorStoreID string `pathParam:"style=simple,explode=false,name=vector_store_id"`
	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	//
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.
	//
	Order *ListVectorStoreFilesQueryParamOrder `default:"desc" queryParam:"style=form,explode=true,name=order"`
	// A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.
	//
	After *string `queryParam:"style=form,explode=true,name=after"`
	// A cursor for use in pagination. `before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with obj_foo, your subsequent call can include before=obj_foo in order to fetch the previous page of the list.
	//
	Before *string `queryParam:"style=form,explode=true,name=before"`
	// Filter by file status. One of `in_progress`, `completed`, `failed`, `cancelled`.
	Filter *QueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
}

func (l ListVectorStoreFilesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListVectorStoreFilesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListVectorStoreFilesRequest) GetVectorStoreID() string {
	if o == nil {
		return ""
	}
	return o.VectorStoreID
}

func (o *ListVectorStoreFilesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListVectorStoreFilesRequest) GetOrder() *ListVectorStoreFilesQueryParamOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListVectorStoreFilesRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListVectorStoreFilesRequest) GetBefore() *string {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *ListVectorStoreFilesRequest) GetFilter() *QueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListVectorStoreFilesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListVectorStoreFilesResponse *components.ListVectorStoreFilesResponse
}

func (o *ListVectorStoreFilesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListVectorStoreFilesResponse) GetListVectorStoreFilesResponse() *components.ListVectorStoreFilesResponse {
	if o == nil {
		return nil
	}
	return o.ListVectorStoreFilesResponse
}
