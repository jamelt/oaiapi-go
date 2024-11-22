# ModifyUserRequest


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `UserID`                                                                             | *string*                                                                             | :heavy_check_mark:                                                                   | The ID of the user.                                                                  |
| `UserRoleUpdateRequest`                                                              | [components.UserRoleUpdateRequest](../../models/components/userroleupdaterequest.md) | :heavy_check_mark:                                                                   | The new user role to modify. This must be one of `owner` or `member`.                |