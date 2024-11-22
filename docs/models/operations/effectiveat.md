# EffectiveAt

Return only events whose `effective_at` (Unix seconds) is in this range.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `Gt`                                                                                           | **int64*                                                                                       | :heavy_minus_sign:                                                                             | Return only events whose `effective_at` (Unix seconds) is greater than this value.             |
| `Gte`                                                                                          | **int64*                                                                                       | :heavy_minus_sign:                                                                             | Return only events whose `effective_at` (Unix seconds) is greater than or equal to this value. |
| `Lt`                                                                                           | **int64*                                                                                       | :heavy_minus_sign:                                                                             | Return only events whose `effective_at` (Unix seconds) is less than this value.                |
| `Lte`                                                                                          | **int64*                                                                                       | :heavy_minus_sign:                                                                             | Return only events whose `effective_at` (Unix seconds) is less than or equal to this value.    |