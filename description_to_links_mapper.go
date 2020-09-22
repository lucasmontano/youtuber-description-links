package main

import "regexp"

func descriptionToLinksMapper(description string) []string {
	re := regexp.MustCompile(`\b(?:https?:)?(?:(?i:[a-z]+\.)+)[^\s,]+\b`)
	return re.FindAllString(description, -1)
}
