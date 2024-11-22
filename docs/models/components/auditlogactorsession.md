# AuditLogActorSession

The session in which the audit logged action was performed.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `User`                                                                        | [*components.AuditLogActorUser](../../models/components/auditlogactoruser.md) | :heavy_minus_sign:                                                            | The user who performed the audit logged action.                               |
| `IPAddress`                                                                   | **string*                                                                     | :heavy_minus_sign:                                                            | The IP address from which the action was performed.                           |