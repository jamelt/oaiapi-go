// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type CreateThreadResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ThreadObject *components.ThreadObject
}

func (o *CreateThreadResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateThreadResponse) GetThreadObject() *components.ThreadObject {
	if o == nil {
		return nil
	}
	return o.ThreadObject
}