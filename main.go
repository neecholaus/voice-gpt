package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/sashabaranov/go-openai"
)

type Config struct {
	AiKey string
	Mode  string
}

func main() {
	config := initConfig()

	tp := Transport{
		Ai:             openai.NewClient(os.Getenv("OPENAI_KEY")),
		Tts:            make(chan string),
		TtsDone:        sync.WaitGroup{},
		MessageHistory: &[]openai.ChatCompletionMessage{},
	}

	if config.Mode == "text" {
		waitForText(&tp)
	} else if config.Mode == "voice" {
		waitForAudioIn(&tp)
	} else {
		fmt.Printf("the mode (%s) is not supported", config.Mode)
	}
}

func initConfig() *Config {
	file, _ := os.Open(".env")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "=")

		if len(parts) < 2 {
			continue
		}

		err := os.Setenv(parts[0], parts[1])
		if err != nil {
			fmt.Println("could not set env var")
		}
	}

	config := Config{}

	if os.Getenv("OPENAI_KEY") != "" {
		config.AiKey = os.Getenv("OPENAI_KEY")
	} else {
		panic("missing keys")
	}

	if os.Getenv("MODE") != "" {
		config.Mode = os.Getenv("MODE")
	} else {
		config.Mode = "text"
	}

	return &config
}
