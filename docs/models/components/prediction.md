# Prediction

Configuration for a [Predicted Output](/docs/guides/latency-optimization#use-predicted-outputs),
which can greatly improve response times when large parts of the model
response are known ahead of time. This is most common when you are
regenerating a file with only minor changes to most of the content.



## Supported Types

### StaticContent

```go
prediction := components.CreatePredictionStaticContent(components.StaticContent{/* values here */})
```

