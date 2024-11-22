# PredictionContent

The content that should be matched when generating a model response.
If generated tokens would match this content, the entire model response
can be returned much more quickly.



## Supported Types

### 

```go
predictionContent := components.CreatePredictionContentStr(string{/* values here */})
```

### 

```go
predictionContent := components.CreatePredictionContentArrayOfChatCompletionRequestMessageContentPartText([]components.ChatCompletionRequestMessageContentPartText{/* values here */})
```

