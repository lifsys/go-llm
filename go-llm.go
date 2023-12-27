package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func general_response(system string, prompt string, model string) (string, error) {
	client := openai.NewClient(get_key("OPEN-AI","Mamba")) // Replace "YOUR_API_KEY" with your actual API key
	// Select Model
	if model == "" {
			// Set a default model if no model is provided
			model = "gpt-3.5-turbo-1106"
		}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: system,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v", err)
	}

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", nil
}
