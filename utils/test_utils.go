package utils

import "strings"

// TrimWhitespaces removes all whitespace, \n and \t from a string
func TrimWhitespaces(old string) string {
	old = strings.ReplaceAll(old, "\n", "")
	old = strings.ReplaceAll(old, "\t", "")
	old = strings.ReplaceAll(old, " ", "")
	old = strings.TrimSpace(old)

	return old
}
