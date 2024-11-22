# ProjectUser

Represents an individual user in a project.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Object`                                                                     | [components.ProjectUserObject](../../models/components/projectuserobject.md) | :heavy_check_mark:                                                           | The object type, which is always `organization.project.user`                 |
| `ID`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | The identifier, which can be referenced in API endpoints                     |
| `Name`                                                                       | *string*                                                                     | :heavy_check_mark:                                                           | The name of the user                                                         |
| `Email`                                                                      | *string*                                                                     | :heavy_check_mark:                                                           | The email address of the user                                                |
| `Role`                                                                       | [components.ProjectUserRole](../../models/components/projectuserrole.md)     | :heavy_check_mark:                                                           | `owner` or `member`                                                          |
| `AddedAt`                                                                    | *int64*                                                                      | :heavy_check_mark:                                                           | The Unix timestamp (in seconds) of when the project was added.               |