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

func (c *PiperConfig) getCmdString(text string) string {
	// make filename that is beginning of the text
	desiredFileNameLength := 30
	textLength := len(text)
	if desiredFileNameLength > textLength {
		desiredFileNameLength = textLength
	}
	filename := text[:desiredFileNameLength]
	filename = strings.Replace(filename, " ", "-", -1)

	return strings.Join([]string{
		"echo",
		fmt.Sprintf("'%s'", text),
		"|",
		c.Path,
		fmt.Sprintf("--model %s", c.Model),
		fmt.Sprintf("-f /var/opt/responses/%s.wav", filename),
		"-s",
		c.Speaker,
	}, " ")
}

func defaultPiperConfig() *PiperConfig {
	return &PiperConfig{
		Path:    "/var/opt/piper/piper",
		Model:   "/var/opt/piper-voices/en-us-libritts-high.onnx",
		Speaker: "13",
	}
}

func createWav(text string) (string, error) {
	config := defaultPiperConfig()

	cmdString := config.getCmdString(text)

	fmt.Println(cmdString)

	cmd := exec.Command("bash", "-c", cmdString)

	s, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(s), nil
}
