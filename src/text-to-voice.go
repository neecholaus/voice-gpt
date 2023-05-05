package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type PiperConfig struct {
	Path    string
	Model   string
	Speaker string
}

func defaultPiperConfig() *PiperConfig {
	return &PiperConfig{
		Path:    "/var/opt/piper/piper",
		Model:   "/var/opt/piper-voices/en-us-libritts-high.onnx",
		Speaker: "13",
	}
}

func createWav(text string) {
	config := defaultPiperConfig()

	cmdPieces := []string{
		"echo",
		fmt.Sprintf("'%s'", text),
		"|",
		config.Path,
		fmt.Sprintf("--model %s", config.Model),
		"--f test.wav",
		"--d /var/opt/responses",
		"-s",
		config.Speaker,
	}

	cmd := exec.Command("bash", "-c", strings.Join(cmdPieces, " "))

	_, err := cmd.Output()
	if err != nil {
		fmt.Println(fmt.Errorf("createWav failed: %s", err))
	}
}
