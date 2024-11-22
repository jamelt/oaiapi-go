# VectorStoreFileBatchObjectFileCounts


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `InProgress`                                            | *int64*                                                 | :heavy_check_mark:                                      | The number of files that are currently being processed. |
| `Completed`                                             | *int64*                                                 | :heavy_check_mark:                                      | The number of files that have been processed.           |
| `Failed`                                                | *int64*                                                 | :heavy_check_mark:                                      | The number of files that have failed to process.        |
| `Cancelled`                                             | *int64*                                                 | :heavy_check_mark:                                      | The number of files that where cancelled.               |
| `Total`                                                 | *int64*                                                 | :heavy_check_mark:                                      | The total number of files.                              |