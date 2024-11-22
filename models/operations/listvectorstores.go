// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
	"github.com/jamelt/openai-api/models/components"
)

// ListVectorStoresQueryParamOrder - Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.
type ListVectorStoresQueryParamOrder string

const (
	ListVectorStoresQueryParamOrderAsc  ListVectorStoresQueryParamOrder = "asc"
	ListVectorStoresQueryParamOrderDesc ListVectorStoresQueryParamOrder = "desc"
)

func (e ListVectorStoresQueryParamOrder) ToPointer() *ListVectorStoresQueryParamOrder {
	return &e
}
func (e *ListVectorStoresQueryParamOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = ListVectorStoresQueryParamOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListVectorStoresQueryParamOrder: %v", v)
	}
}

type ListVectorStoresRequest struct {
	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	//
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.
	//
	Order *ListVectorStoresQueryParamOrder `default:"desc" queryParam:"style=form,explode=true,name=order"`
	// A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.
	//
	After *string `queryParam:"style=form,explode=true,name=after"`
	// A cursor for use in pagination. `before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with obj_foo, your subsequent call can include before=obj_foo in order to fetch the previous page of the list.
	//
	Before *string `queryParam:"style=form,explode=true,name=before"`
}

func (l ListVectorStoresRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListVectorStoresRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListVectorStoresRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListVectorStoresRequest) GetOrder() *ListVectorStoresQueryParamOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListVectorStoresRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListVectorStoresRequest) GetBefore() *string {
	if o == nil {
		return nil
	}
	return o.Before
}

type ListVectorStoresResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListVectorStoresResponse *components.ListVectorStoresResponse
}

func (o *ListVectorStoresResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListVectorStoresResponse) GetListVectorStoresResponse() *components.ListVectorStoresResponse {
	if o == nil {
		return nil
	}
	return o.ListVectorStoresResponse
}
