# Prompt

The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.

Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document.



## Supported Types

### 

```go
prompt := components.CreatePromptStr(string{/* values here */})
```

### 

```go
prompt := components.CreatePromptArrayOfStr([]string{/* values here */})
```

### 

```go
prompt := components.CreatePromptArrayOfInteger([]int64{/* values here */})
```

### 

```go
prompt := components.CreatePromptArrayOfArrayOfInteger([][]int64{/* values here */})
```

