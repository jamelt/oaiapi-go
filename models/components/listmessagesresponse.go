// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListMessagesResponse struct {
	Object  string          `json:"object"`
	Data    []MessageObject `json:"data"`
	FirstID string          `json:"first_id"`
	LastID  string          `json:"last_id"`
	HasMore bool            `json:"has_more"`
}

func (o *ListMessagesResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}

func (o *ListMessagesResponse) GetData() []MessageObject {
	if o == nil {
		return []MessageObject{}
	}
	return o.Data
}

func (o *ListMessagesResponse) GetFirstID() string {
	if o == nil {
		return ""
	}
	return o.FirstID
}

func (o *ListMessagesResponse) GetLastID() string {
	if o == nil {
		return ""
	}
	return o.LastID
}

func (o *ListMessagesResponse) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}
