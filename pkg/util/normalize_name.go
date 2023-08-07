package util

import "strings"

func NormalizeName(s string) string {
	var normalized string
	for _, word := range strings.Split(s, " ") {
		if word == "" {
			continue
		}
		if normalized != "" {
			normalized += " "
		}
		normalized += word
	}
	return normalized
}
