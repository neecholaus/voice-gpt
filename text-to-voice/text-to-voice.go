package text_to_voice

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func KeepConverting(textChan *chan string) {
	go func() {
		for {
			start := time.Now().UnixMilli()

			text := <-*textChan
			_, _ = createWav(text)

			fmt.Printf("converted text (%d) : %s\n", time.Now().UnixMilli()-start, text)
		}
	}()
}

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

	cmd := exec.Command("bash", "-c", cmdString)

	s, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(s), nil
}
