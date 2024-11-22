// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type DeleteAssistantRequest struct {
	// The ID of the assistant to delete.
	AssistantID string `pathParam:"style=simple,explode=false,name=assistant_id"`
}

func (o *DeleteAssistantRequest) GetAssistantID() string {
	if o == nil {
		return ""
	}
	return o.AssistantID
}

type DeleteAssistantResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	DeleteAssistantResponse *components.DeleteAssistantResponse
}

func (o *DeleteAssistantResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteAssistantResponse) GetDeleteAssistantResponse() *components.DeleteAssistantResponse {
	if o == nil {
		return nil
	}
	return o.DeleteAssistantResponse
}
