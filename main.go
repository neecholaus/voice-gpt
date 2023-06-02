package main

import (
	"bufio"
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	text_to_voice "neecholaus/profinabox/text-to-voice"
	"os"
	"strings"
	"sync"
)

func main() {
	loadDotEnv()
	fmt.Println(os.Getenv("OPENAI_KEY"))

	ai := openai.NewClient(os.Getenv("OPENAI_KEY"))

	tts := make(chan string)
	ttsDone := sync.WaitGroup{}

	text_to_voice.KeepConverting(&tts, &ttsDone)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

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

		ttsDone.Add(1)
		tts <- resp.Choices[0].Message.Content

		ttsDone.Wait()
	}
}

func loadDotEnv() {
	file, _ := os.Open(".env")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "=")

		err := os.Setenv(parts[0], parts[1])
		if err != nil {
			fmt.Println("could not sent env var")
		}
	}
}
