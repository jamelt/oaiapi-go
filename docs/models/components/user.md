# User

Represents an individual `user` within an organization.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Object`                                                       | [components.UserObject](../../models/components/userobject.md) | :heavy_check_mark:                                             | The object type, which is always `organization.user`           |
| `ID`                                                           | *string*                                                       | :heavy_check_mark:                                             | The identifier, which can be referenced in API endpoints       |
| `Name`                                                         | *string*                                                       | :heavy_check_mark:                                             | The name of the user                                           |
| `Email`                                                        | *string*                                                       | :heavy_check_mark:                                             | The email address of the user                                  |
| `Role`                                                         | [components.UserRole](../../models/components/userrole.md)     | :heavy_check_mark:                                             | `owner` or `reader`                                            |
| `AddedAt`                                                      | *int64*                                                        | :heavy_check_mark:                                             | The Unix timestamp (in seconds) of when the user was added.    |