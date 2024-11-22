# AuditLogActor

The actor who performed the audit logged action.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Type`                                                                              | [*components.AuditLogActorType](../../models/components/auditlogactortype.md)       | :heavy_minus_sign:                                                                  | The type of actor. Is either `session` or `api_key`.                                |
| `Session`                                                                           | [*components.AuditLogActorSession](../../models/components/auditlogactorsession.md) | :heavy_minus_sign:                                                                  | The session in which the audit logged action was performed.                         |
| `APIKey`                                                                            | [*components.AuditLogActorAPIKey](../../models/components/auditlogactorapikey.md)   | :heavy_minus_sign:                                                                  | The API Key used to perform the audit logged action.                                |