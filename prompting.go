package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	text_to_voice "neecholaus/profinabox/text-to-voice"
	"os"
	"strings"
	"sync"
)

type Transport struct {
	Ai      *openai.Client
	Tts     chan string
	TtsDone sync.WaitGroup
}

// Useful for testing without voice.
func waitForText(tp *Transport) {
	text_to_voice.KeepConverting(&tp.Tts, &tp.TtsDone)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		reply := getPromptReply(tp.Ai, text)

		tp.TtsDone.Add(1)
		tp.Tts <- reply

		tp.TtsDone.Wait()
	}
}

func getPromptReply(ai *openai.Client, text string) string {
	resp, err := ai.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Give me the most clear and concise responses you can.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ai error: %s\n", err.Error())
	}

	return resp.Choices[0].Message.Content
}
