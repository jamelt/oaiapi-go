# CompletionTokensDetails

Breakdown of tokens used in a completion.


## Fields

| Field                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AcceptedPredictionTokens`                                                                                                                                                                                                                                             | **int64*                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                     | When using Predicted Outputs, the number of tokens in the<br/>prediction that appeared in the completion.<br/>                                                                                                                                                         |
| `AudioTokens`                                                                                                                                                                                                                                                          | **int64*                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                     | Audio input tokens generated by the model.                                                                                                                                                                                                                             |
| `ReasoningTokens`                                                                                                                                                                                                                                                      | **int64*                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                     | Tokens generated by the model for reasoning.                                                                                                                                                                                                                           |
| `RejectedPredictionTokens`                                                                                                                                                                                                                                             | **int64*                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                     | When using Predicted Outputs, the number of tokens in the<br/>prediction that did not appear in the completion. However, like<br/>reasoning tokens, these tokens are still counted in the total<br/>completion tokens for purposes of billing, output, and context window<br/>limits.<br/> |