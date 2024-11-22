# UploadPart

The upload Part represents a chunk of bytes we can add to an Upload object.



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ID`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | The upload Part unique identifier, which can be referenced in API endpoints. |
| `CreatedAt`                                                                  | *int64*                                                                      | :heavy_check_mark:                                                           | The Unix timestamp (in seconds) for when the Part was created.               |
| `UploadID`                                                                   | *string*                                                                     | :heavy_check_mark:                                                           | The ID of the Upload object that this Part was added to.                     |
| `Object`                                                                     | [components.UploadPartObject](../../models/components/uploadpartobject.md)   | :heavy_check_mark:                                                           | The object type, which is always `upload.part`.                              |