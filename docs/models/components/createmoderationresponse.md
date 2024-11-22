# CreateModerationResponse

Represents if a given text input is potentially harmful.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ID`                                                       | *string*                                                   | :heavy_check_mark:                                         | The unique identifier for the moderation request.          |
| `Model`                                                    | *string*                                                   | :heavy_check_mark:                                         | The model used to generate the moderation results.         |
| `Results`                                                  | [][components.Results](../../models/components/results.md) | :heavy_check_mark:                                         | A list of moderation objects.                              |