# go_llm

A Go package for interacting with AI language models from OpenAI and TogetherAI.

## Features

- OpenAI Chat Completion: Send prompts to OpenAI's GPT-3.5 model and receive responses.
- TogetherAI Inference: Post requests to TogetherAI's endpoint and get language model inferences.

## Installation

```sh
go get github.com/lifsys/onepass
go get github.com/sashabaranov/go-openai
```

## Usage

### OpenAI

```go
response, err := Openai_call(systemMessage, userPrompt, modelName)
```

- `systemMessage`: A string representing the system's message.
- `userPrompt`: The user's prompt to the AI.
- `modelName`: The model to use for the completion (optional).

### TogetherAI

```go
response, err := Togetherai_call(systemMessage, userPrompt)
```

- `systemMessage`: A string representing the system's message.
- `userPrompt`: The user's prompt to the AI.

## Configuration

Set up your API keys for OpenAI and TogetherAI using the `onepass` package.

## License

Specify your license here.

## Contributing

Instructions for how to contribute to the project.

## Support

Contact information for support.

## Authors

List of contributors.

---

Please replace placeholders with actual information about your package, such as the correct installation commands, license, and contact details.
