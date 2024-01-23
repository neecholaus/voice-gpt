package main

import (
	"bufio"
	"context"
	"fmt"
	text_to_voice "neecholaus/voicegpt/text-to-voice"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/sashabaranov/go-openai"
)

type Transport struct {
	Ai             *openai.Client
	Tts            chan string
	TtsDone        sync.WaitGroup
	MessageHistory *[]openai.ChatCompletionMessage
}

// Useful for testing without voice.
func waitForText(tp *Transport) {
	text_to_voice.KeepConverting(&tp.Tts, &tp.TtsDone)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		*tp.MessageHistory = append(*tp.MessageHistory, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		if len(*tp.MessageHistory) > 10 {
			amount := len(*tp.MessageHistory) - 10
			fmt.Printf("removing %d messages from history", amount)
			*tp.MessageHistory = (*tp.MessageHistory)[amount:]
		}

		fmt.Printf("messages: %d\n", len(*tp.MessageHistory))
		addMessageToLog(&(*tp.MessageHistory)[len(*tp.MessageHistory)-1])

		reply := getPromptReply(tp.Ai, tp.MessageHistory)

		// createWav

		tp.TtsDone.Add(1)
		tp.Tts <- reply

		tp.TtsDone.Wait()
	}
}

func waitForAudioIn(tp *Transport) {
	for {
		files, err := os.ReadDir("./audio-in")
		if err != nil {
			fmt.Printf("could not read audio-in dir: %s\n", err.Error())
		}

		if len(files) < 1 {
			fmt.Println("no files")
			time.Sleep(time.Second)
			continue
		}

		sortedFilenames := []string{}
		for _, file := range files {
			sortedFilenames = append(sortedFilenames, file.Name())
		}
		sort.Strings(sortedFilenames)

		for _, file := range files {
			if file.IsDir() == true {
				continue
			}

			fmt.Printf("next for stt: %s\n", file.Name())
			break
		}

		time.Sleep(time.Second)
	}
}

func getPromptReply(ai *openai.Client, history *[]openai.ChatCompletionMessage) string {
	// add role always to beginning
	messages := append([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: os.Getenv("GPT_ROLE"),
		},
	}, *history...)

	resp, err := ai.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: append([]openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: os.Getenv("GPT_ROLE"),
				},
			},
				messages...,
			),
		},
	)

	if err != nil {
		fmt.Printf("ai error: %s\n", err.Error())
		return ""
	}

	*history = append(*history, resp.Choices[0].Message)
	addMessageToLog(&resp.Choices[0].Message)

	return resp.Choices[0].Message.Content
}

func addMessageToLog(message *openai.ChatCompletionMessage) {
	f, _ := os.OpenFile("./.conversation.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	defer f.Close()

	_, err := f.WriteString(fmt.Sprintf("\n (%s)\n%s\n", message.Role, message.Content))
	if err != nil {
		fmt.Println(err)
	}
}
