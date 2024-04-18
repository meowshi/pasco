package utils

import (
	"strings"
	"unicode"
)

func CutEventName(eventName string) string {
	firstNumberIdx := 0
	for i, ch := range eventName {
		if unicode.IsNumber(ch) {
			firstNumberIdx = i
			break
		}
	}
	return strings.Trim(eventName[:firstNumberIdx], " \n\t")
}

func HaveEmpty(slice []interface{}) bool {
	for _, el := range slice {
		if len(el.(string)) == 0 {
			return true
		}
	}
	return false
}
