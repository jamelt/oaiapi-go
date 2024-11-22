// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CompletionTokensDetails - Breakdown of tokens used in a completion.
type CompletionTokensDetails struct {
	// When using Predicted Outputs, the number of tokens in the
	// prediction that appeared in the completion.
	//
	AcceptedPredictionTokens *int64 `json:"accepted_prediction_tokens,omitempty"`
	// Audio input tokens generated by the model.
	AudioTokens *int64 `json:"audio_tokens,omitempty"`
	// Tokens generated by the model for reasoning.
	ReasoningTokens *int64 `json:"reasoning_tokens,omitempty"`
	// When using Predicted Outputs, the number of tokens in the
	// prediction that did not appear in the completion. However, like
	// reasoning tokens, these tokens are still counted in the total
	// completion tokens for purposes of billing, output, and context window
	// limits.
	//
	RejectedPredictionTokens *int64 `json:"rejected_prediction_tokens,omitempty"`
}

func (o *CompletionTokensDetails) GetAcceptedPredictionTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.AcceptedPredictionTokens
}

func (o *CompletionTokensDetails) GetAudioTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.AudioTokens
}

func (o *CompletionTokensDetails) GetReasoningTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.ReasoningTokens
}

func (o *CompletionTokensDetails) GetRejectedPredictionTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.RejectedPredictionTokens
}

// PromptTokensDetails - Breakdown of tokens used in the prompt.
type PromptTokensDetails struct {
	// Audio input tokens present in the prompt.
	AudioTokens *int64 `json:"audio_tokens,omitempty"`
	// Cached tokens present in the prompt.
	CachedTokens *int64 `json:"cached_tokens,omitempty"`
}

func (o *PromptTokensDetails) GetAudioTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.AudioTokens
}

func (o *PromptTokensDetails) GetCachedTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.CachedTokens
}

// CompletionUsage - Usage statistics for the completion request.
type CompletionUsage struct {
	// Number of tokens in the generated completion.
	CompletionTokens int64 `json:"completion_tokens"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens"`
	// Total number of tokens used in the request (prompt + completion).
	TotalTokens int64 `json:"total_tokens"`
	// Breakdown of tokens used in a completion.
	CompletionTokensDetails *CompletionTokensDetails `json:"completion_tokens_details,omitempty"`
	// Breakdown of tokens used in the prompt.
	PromptTokensDetails *PromptTokensDetails `json:"prompt_tokens_details,omitempty"`
}

func (o *CompletionUsage) GetCompletionTokens() int64 {
	if o == nil {
		return 0
	}
	return o.CompletionTokens
}

func (o *CompletionUsage) GetPromptTokens() int64 {
	if o == nil {
		return 0
	}
	return o.PromptTokens
}

func (o *CompletionUsage) GetTotalTokens() int64 {
	if o == nil {
		return 0
	}
	return o.TotalTokens
}

func (o *CompletionUsage) GetCompletionTokensDetails() *CompletionTokensDetails {
	if o == nil {
		return nil
	}
	return o.CompletionTokensDetails
}

func (o *CompletionUsage) GetPromptTokensDetails() *PromptTokensDetails {
	if o == nil {
		return nil
	}
	return o.PromptTokensDetails
}
