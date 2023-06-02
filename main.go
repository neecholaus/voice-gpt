package main

import (
	"bufio"
	"fmt"
	text_to_voice "neecholaus/profinabox/text-to-voice"
	"os"
	"strings"
)

func main() {

	t2s := make(chan string)

	text_to_voice.KeepConverting(&t2s)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		t2s <- text
	}
}
