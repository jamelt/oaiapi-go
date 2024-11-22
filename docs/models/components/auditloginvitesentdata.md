# AuditLogInviteSentData

The payload used to create the invite.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Email`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | The email invited to the organization.                               |
| `Role`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | The role the email was invited to be. Is either `owner` or `member`. |