# Project

Represents an individual project.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ID`                                                                        | *string*                                                                    | :heavy_check_mark:                                                          | The identifier, which can be referenced in API endpoints                    |
| `Object`                                                                    | [components.ProjectObject](../../models/components/projectobject.md)        | :heavy_check_mark:                                                          | The object type, which is always `organization.project`                     |
| `Name`                                                                      | *string*                                                                    | :heavy_check_mark:                                                          | The name of the project. This appears in reporting.                         |
| `CreatedAt`                                                                 | *int64*                                                                     | :heavy_check_mark:                                                          | The Unix timestamp (in seconds) of when the project was created.            |
| `ArchivedAt`                                                                | **int64*                                                                    | :heavy_minus_sign:                                                          | The Unix timestamp (in seconds) of when the project was archived or `null`. |
| `Status`                                                                    | [components.ProjectStatus](../../models/components/projectstatus.md)        | :heavy_check_mark:                                                          | `active` or `archived`                                                      |