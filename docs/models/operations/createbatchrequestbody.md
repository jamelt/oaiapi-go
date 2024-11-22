# CreateBatchRequestBody


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                           | Type                                                                                                                                                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                                                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `InputFileID`                                                                                                                                                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                              | The ID of an uploaded file that contains requests for the new batch.<br/><br/>See [upload file](/docs/api-reference/files/create) for how to upload a file.<br/><br/>Your input file must be formatted as a [JSONL file](/docs/api-reference/batch/request-input), and must be uploaded with the purpose `batch`. The file can contain up to 50,000 requests, and can be up to 100 MB in size.<br/> |
| `Endpoint`                                                                                                                                                                                                                                                                                                                                                                      | [operations.Endpoint](../../models/operations/endpoint.md)                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                              | The endpoint to be used for all requests in the batch. Currently `/v1/chat/completions`, `/v1/embeddings`, and `/v1/completions` are supported. Note that `/v1/embeddings` batches are also restricted to a maximum of 50,000 embedding inputs across all requests in the batch.                                                                                                |
| `CompletionWindow`                                                                                                                                                                                                                                                                                                                                                              | [operations.CompletionWindow](../../models/operations/completionwindow.md)                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                              | The time frame within which the batch should be processed. Currently only `24h` is supported.                                                                                                                                                                                                                                                                                   |
| `Metadata`                                                                                                                                                                                                                                                                                                                                                                      | map[string]*string*                                                                                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                              | Optional custom metadata for the batch.                                                                                                                                                                                                                                                                                                                                         |