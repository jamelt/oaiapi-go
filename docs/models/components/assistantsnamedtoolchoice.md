# AssistantsNamedToolChoice

Specifies a tool the model should use. Use to force the model to call a specific tool.


## Fields

| Field                                                                                                         | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `Type`                                                                                                        | [components.AssistantsNamedToolChoiceType](../../models/components/assistantsnamedtoolchoicetype.md)          | :heavy_check_mark:                                                                                            | The type of the tool. If type is `function`, the function name must be set                                    |
| `Function`                                                                                                    | [*components.AssistantsNamedToolChoiceFunction](../../models/components/assistantsnamedtoolchoicefunction.md) | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |