# ToolResources

A set of resources that are used by the assistant's tools. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs.



## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `CodeInterpreter`                                                                                       | [*components.AssistantObjectCodeInterpreter](../../models/components/assistantobjectcodeinterpreter.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `FileSearch`                                                                                            | [*components.AssistantObjectFileSearch](../../models/components/assistantobjectfilesearch.md)           | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |