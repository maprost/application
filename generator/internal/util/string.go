package util

import "strings"

func ReplaceNewLine(str string) string {
	return strings.Replace(str, "\n", "\n\n", -1)
}
