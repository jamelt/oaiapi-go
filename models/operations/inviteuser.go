// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type InviteUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// User invited successfully.
	Invite *components.Invite
}

func (o *InviteUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *InviteUserResponse) GetInvite() *components.Invite {
	if o == nil {
		return nil
	}
	return o.Invite
}