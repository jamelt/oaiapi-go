# ProjectAPIKey

Represents an individual API key in a project.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Object`                                                                         | [components.ProjectAPIKeyObject](../../models/components/projectapikeyobject.md) | :heavy_check_mark:                                                               | The object type, which is always `organization.project.api_key`                  |
| `RedactedValue`                                                                  | *string*                                                                         | :heavy_check_mark:                                                               | The redacted value of the API key                                                |
| `Name`                                                                           | *string*                                                                         | :heavy_check_mark:                                                               | The name of the API key                                                          |
| `CreatedAt`                                                                      | *int64*                                                                          | :heavy_check_mark:                                                               | The Unix timestamp (in seconds) of when the API key was created                  |
| `ID`                                                                             | *string*                                                                         | :heavy_check_mark:                                                               | The identifier, which can be referenced in API endpoints                         |
| `Owner`                                                                          | [components.Owner](../../models/components/owner.md)                             | :heavy_check_mark:                                                               | N/A                                                                              |