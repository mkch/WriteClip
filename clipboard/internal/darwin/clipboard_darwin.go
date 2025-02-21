package darwin

import (
	"errors"
	"os/exec"
	"strings"
)

// ClipboardText gets the text in clipboard.
func ClipboardText() (str string, err error) {
	var stdout strings.Builder
	var stderr strings.Builder
	cmd := exec.Command("pbpaste")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err = cmd.Run(); err != nil {
		return
	}
	if stderr.Len() != 0 {
		err = errors.New("stderr: " + stderr.String())
		return
	}
	str = stdout.String()
	return
}

// SetClipboardText sets str into clipboard.
func SetClipboardText(str string) (err error) {
	var stdout strings.Builder
	var stderr strings.Builder
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(str)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err = cmd.Run(); err != nil {
		return
	}
	if stderr.Len() != 0 {
		err = errors.New("stderr: " + stderr.String())
		return
	}
	return
}
