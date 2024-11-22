# ProjectServiceAccount

Represents an individual service account in a project.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `Object`                                                                                         | [components.ProjectServiceAccountObject](../../models/components/projectserviceaccountobject.md) | :heavy_check_mark:                                                                               | The object type, which is always `organization.project.service_account`                          |
| `ID`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | The identifier, which can be referenced in API endpoints                                         |
| `Name`                                                                                           | *string*                                                                                         | :heavy_check_mark:                                                                               | The name of the service account                                                                  |
| `Role`                                                                                           | [components.ProjectServiceAccountRole](../../models/components/projectserviceaccountrole.md)     | :heavy_check_mark:                                                                               | `owner` or `member`                                                                              |
| `CreatedAt`                                                                                      | *int64*                                                                                          | :heavy_check_mark:                                                                               | The Unix timestamp (in seconds) of when the service account was created                          |