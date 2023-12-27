package go-llm

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/lifsys/onepass"
	openai "github.com/sashabaranov/go-openai"
)

type InferenceRequest struct {
	Model             string   `json:"model"`
	MaxTokens         int      `json:"max_tokens"`
	Prompt            string   `json:"prompt"`
	RequestType       string   `json:"request_type"`
	Temperature       float64  `json:"temperature"`
	TopP              float64  `json:"top_p"`
	TopK              int      `json:"top_k"`
	RepetitionPenalty float64  `json:"repetition_penalty"`
	Stop              []string `json:"stop"`
	NegativePrompt    string   `json:"negative_prompt"`
	SessionKey        string   `json:"sessionKey"`
	RepetitivePenalty float64  `json:"repetitive_penalty"`
}

// InferenceResponse represents the JSON payload for the response.
type InferenceResponse struct {
	Output struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	} `json:"output"`
}

func Openai_call(system string, prompt string, model string) (string, error) {
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

// togetherai sends a POST request to the specified endpoint with the given payload and returns the response body as a string.
func Togetherai_call(system string, prompt string) (string, error) {
	endpoint := "https://api.together.xyz/inference"
	reqBody := InferenceRequest{
		Model:             "mistralai/Mixtral-8x7B-Instruct-v0.1",
		MaxTokens:         4096,
		Prompt:            "[System] " + system + "[/System]" + "[INST] " + prompt + " [/INST]",
		RequestType:       "language-model-inference",
		Temperature:       0.7,
		TopP:              0.7,
		TopK:              50,
		RepetitionPenalty: 1,
		Stop:              []string{"</s>", "[INST]"},
		NegativePrompt:    "",
		SessionKey:        "feeebaa6-4e1b-4d3a-912e-c5abc1a44f52",
		RepetitivePenalty: 1,
	}

	// Marshal the request body to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// Set the headers
	authToken := onepass.Get_key("TogetherAI", "API") // Load the token from an environment variable
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the request using a shared client to enable connection reuse
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Decode the response data into the InferenceResponse struct
	var response InferenceResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	// Extract the "choices" from the response and concatenate them
	var choices []string
	for _, choice := range response.Output.Choices {
		choices = append(choices, choice.Text)
	}

	// Join the choices into a single string
	result := strings.Join(choices, "\n")

	return result, nil
}