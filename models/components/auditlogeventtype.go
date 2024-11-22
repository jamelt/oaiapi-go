// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AuditLogEventType - The event type.
type AuditLogEventType string

const (
	AuditLogEventTypeAPIKeyCreated         AuditLogEventType = "api_key.created"
	AuditLogEventTypeAPIKeyUpdated         AuditLogEventType = "api_key.updated"
	AuditLogEventTypeAPIKeyDeleted         AuditLogEventType = "api_key.deleted"
	AuditLogEventTypeInviteSent            AuditLogEventType = "invite.sent"
	AuditLogEventTypeInviteAccepted        AuditLogEventType = "invite.accepted"
	AuditLogEventTypeInviteDeleted         AuditLogEventType = "invite.deleted"
	AuditLogEventTypeLoginSucceeded        AuditLogEventType = "login.succeeded"
	AuditLogEventTypeLoginFailed           AuditLogEventType = "login.failed"
	AuditLogEventTypeLogoutSucceeded       AuditLogEventType = "logout.succeeded"
	AuditLogEventTypeLogoutFailed          AuditLogEventType = "logout.failed"
	AuditLogEventTypeOrganizationUpdated   AuditLogEventType = "organization.updated"
	AuditLogEventTypeProjectCreated        AuditLogEventType = "project.created"
	AuditLogEventTypeProjectUpdated        AuditLogEventType = "project.updated"
	AuditLogEventTypeProjectArchived       AuditLogEventType = "project.archived"
	AuditLogEventTypeServiceAccountCreated AuditLogEventType = "service_account.created"
	AuditLogEventTypeServiceAccountUpdated AuditLogEventType = "service_account.updated"
	AuditLogEventTypeServiceAccountDeleted AuditLogEventType = "service_account.deleted"
	AuditLogEventTypeUserAdded             AuditLogEventType = "user.added"
	AuditLogEventTypeUserUpdated           AuditLogEventType = "user.updated"
	AuditLogEventTypeUserDeleted           AuditLogEventType = "user.deleted"
)

func (e AuditLogEventType) ToPointer() *AuditLogEventType {
	return &e
}
func (e *AuditLogEventType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_key.created":
		fallthrough
	case "api_key.updated":
		fallthrough
	case "api_key.deleted":
		fallthrough
	case "invite.sent":
		fallthrough
	case "invite.accepted":
		fallthrough
	case "invite.deleted":
		fallthrough
	case "login.succeeded":
		fallthrough
	case "login.failed":
		fallthrough
	case "logout.succeeded":
		fallthrough
	case "logout.failed":
		fallthrough
	case "organization.updated":
		fallthrough
	case "project.created":
		fallthrough
	case "project.updated":
		fallthrough
	case "project.archived":
		fallthrough
	case "service_account.created":
		fallthrough
	case "service_account.updated":
		fallthrough
	case "service_account.deleted":
		fallthrough
	case "user.added":
		fallthrough
	case "user.updated":
		fallthrough
	case "user.deleted":
		*e = AuditLogEventType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuditLogEventType: %v", v)
	}
}
