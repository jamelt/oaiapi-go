# CodeInterpreter

The Code Interpreter tool call definition.


## Fields

| Field                                                                                                                                                                                                  | Type                                                                                                                                                                                                   | Required                                                                                                                                                                                               | Description                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Input`                                                                                                                                                                                                | *string*                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                     | The input to the Code Interpreter tool call.                                                                                                                                                           |
| `Outputs`                                                                                                                                                                                              | [][components.Outputs](../../models/components/outputs.md)                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                     | The outputs from the Code Interpreter tool call. Code Interpreter can output one or more items, including text (`logs`) or images (`image`). Each of these are represented by a different object type. |