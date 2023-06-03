package ai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func AskChatGPT(content, token string) string {
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err.Error()
	}

	return resp.Choices[0].Message.Content
}
