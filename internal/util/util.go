package util

import (
	"errors"
	"strings"
)

func NormalizeContent(content string) string {
	var normalized []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			normalized = append(normalized, line)
		}
	}
	return strings.Join(normalized, "\n")
}

// ErrorPrepend returns a new error with a message prepended to the given error.
func ErrorPrepend(m string, err error) error {
	return errors.Join(errors.New(m), err)
}
