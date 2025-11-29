package utils

import (
	"os"
	"strings"
)

func ReadFile(inputFilePath *string) (string, error) {
	content, err := os.ReadFile(*inputFilePath)
	if err != nil {
		panic(err)
	}

	return strings.TrimRight(string(content), "\n"), nil
}
