package main

import "fmt"

func main() {
	path, _ := createWav("a test from the main go file")

	fmt.Printf("new wave is at: %s\n", path)
}
