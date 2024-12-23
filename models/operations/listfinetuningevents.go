// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/internal/utils"
	"github.com/jamelt/openai-api/models/components"
)

type ListFineTuningEventsRequest struct {
	// The ID of the fine-tuning job to get events for.
	//
	FineTuningJobID string `pathParam:"style=simple,explode=false,name=fine_tuning_job_id"`
	// Identifier for the last event from the previous pagination request.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Number of events to retrieve.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
}

func (l ListFineTuningEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListFineTuningEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListFineTuningEventsRequest) GetFineTuningJobID() string {
	if o == nil {
		return ""
	}
	return o.FineTuningJobID
}

func (o *ListFineTuningEventsRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListFineTuningEventsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type ListFineTuningEventsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListFineTuningJobEventsResponse *components.ListFineTuningJobEventsResponse
}

func (o *ListFineTuningEventsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListFineTuningEventsResponse) GetListFineTuningJobEventsResponse() *components.ListFineTuningJobEventsResponse {
	if o == nil {
		return nil
	}
	return o.ListFineTuningJobEventsResponse
}
