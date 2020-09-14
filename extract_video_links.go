package main

import (
	"regexp"
)

func ExtractVideoLinks(description string) []string {
	re := regexp.MustCompile(`\b(?:https?:)?(?:(?i:[a-z]+\.)+)[^\s,]+\b`)
	return re.FindAllString(description, -1)
}
