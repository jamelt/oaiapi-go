// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type GetVectorStoreRequest struct {
	// The ID of the vector store to retrieve.
	VectorStoreID string `pathParam:"style=simple,explode=false,name=vector_store_id"`
}

func (o *GetVectorStoreRequest) GetVectorStoreID() string {
	if o == nil {
		return ""
	}
	return o.VectorStoreID
}

type GetVectorStoreResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	VectorStoreObject *components.VectorStoreObject
}

func (o *GetVectorStoreResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetVectorStoreResponse) GetVectorStoreObject() *components.VectorStoreObject {
	if o == nil {
		return nil
	}
	return o.VectorStoreObject
}
