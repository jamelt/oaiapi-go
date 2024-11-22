# Invite

Represents an individual `invite` to the organization.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Object`                                                           | [components.InviteObject](../../models/components/inviteobject.md) | :heavy_check_mark:                                                 | The object type, which is always `organization.invite`             |
| `ID`                                                               | *string*                                                           | :heavy_check_mark:                                                 | The identifier, which can be referenced in API endpoints           |
| `Email`                                                            | *string*                                                           | :heavy_check_mark:                                                 | The email address of the individual to whom the invite was sent    |
| `Role`                                                             | [components.InviteRole](../../models/components/inviterole.md)     | :heavy_check_mark:                                                 | `owner` or `reader`                                                |
| `Status`                                                           | [components.InviteStatus](../../models/components/invitestatus.md) | :heavy_check_mark:                                                 | `accepted`,`expired`, or `pending`                                 |
| `InvitedAt`                                                        | *int64*                                                            | :heavy_check_mark:                                                 | The Unix timestamp (in seconds) of when the invite was sent.       |
| `ExpiresAt`                                                        | *int64*                                                            | :heavy_check_mark:                                                 | The Unix timestamp (in seconds) of when the invite expires.        |
| `AcceptedAt`                                                       | **int64*                                                           | :heavy_minus_sign:                                                 | The Unix timestamp (in seconds) of when the invite was accepted.   |