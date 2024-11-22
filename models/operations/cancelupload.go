// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type CancelUploadRequest struct {
	// The ID of the Upload.
	//
	UploadID string `pathParam:"style=simple,explode=false,name=upload_id"`
}

func (o *CancelUploadRequest) GetUploadID() string {
	if o == nil {
		return ""
	}
	return o.UploadID
}

type CancelUploadResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Upload *components.Upload
}

func (o *CancelUploadResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CancelUploadResponse) GetUpload() *components.Upload {
	if o == nil {
		return nil
	}
	return o.Upload
}
