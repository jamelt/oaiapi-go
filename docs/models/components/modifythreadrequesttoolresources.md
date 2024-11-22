# ModifyThreadRequestToolResources

A set of resources that are made available to the assistant's tools in this thread. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs.



## Fields

| Field                                                                                                           | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `CodeInterpreter`                                                                                               | [*components.ModifyThreadRequestCodeInterpreter](../../models/components/modifythreadrequestcodeinterpreter.md) | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |
| `FileSearch`                                                                                                    | [*components.ModifyThreadRequestFileSearch](../../models/components/modifythreadrequestfilesearch.md)           | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |