# BatchData


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Code`                                                                     | **string*                                                                  | :heavy_minus_sign:                                                         | An error code identifying the error type.                                  |
| `Message`                                                                  | **string*                                                                  | :heavy_minus_sign:                                                         | A human-readable message providing more details about the error.           |
| `Param`                                                                    | **string*                                                                  | :heavy_minus_sign:                                                         | The name of the parameter that caused the error, if applicable.            |
| `Line`                                                                     | **int64*                                                                   | :heavy_minus_sign:                                                         | The line number of the input file where the error occurred, if applicable. |