package utils

import "strings"

func SplitStringToArray(text string) []string {
	arr := make([]string, 0)
	if len(text) > 0 {
		arr = strings.Split(text, ",")
	}
	return arr
}

func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
