// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type DeleteMessageRequest struct {
	// The ID of the thread to which this message belongs.
	ThreadID string `pathParam:"style=simple,explode=false,name=thread_id"`
	// The ID of the message to delete.
	MessageID string `pathParam:"style=simple,explode=false,name=message_id"`
}

func (o *DeleteMessageRequest) GetThreadID() string {
	if o == nil {
		return ""
	}
	return o.ThreadID
}

func (o *DeleteMessageRequest) GetMessageID() string {
	if o == nil {
		return ""
	}
	return o.MessageID
}

type DeleteMessageResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	DeleteMessageResponse *components.DeleteMessageResponse
}

func (o *DeleteMessageResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteMessageResponse) GetDeleteMessageResponse() *components.DeleteMessageResponse {
	if o == nil {
		return nil
	}
	return o.DeleteMessageResponse
}
