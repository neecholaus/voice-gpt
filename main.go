package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	t2s := make(chan string)

	go func() {
		for {
			text := <-t2s
			_, _ = createWav(text)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		t2s <- text
	}
}
