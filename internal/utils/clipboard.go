package utils

import (
	"os/exec"
	"strings"
)

func CopyToClipboard(text string) error {
	command := exec.Command("pbcopy")
	command.Stdin = strings.NewReader(text)
	return command.Run()
}
