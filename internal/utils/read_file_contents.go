package utils

import (
	"os"
	"strings"
)

func ReadFileContents(inputFilePath *string) (*string, error) {
	content, err := os.ReadFile(*inputFilePath)
	if err != nil {
		return nil, err
	}

	result := strings.TrimRight(string(content), "\n")
	return &result, nil
}
