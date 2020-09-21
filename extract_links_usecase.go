package main

import (
	"regexp"
)

func ExtractLinksUseCase(playlistItemsJson string) []string {
	return extractLinksFromDescription(playlistItemsJson)
}

func extractLinksFromDescription(description string) []string {
	re := regexp.MustCompile(`\b(?:https?:)?(?:(?i:[a-z]+\.)+)[^\s,]+\b`)
	return re.FindAllString(description, -1)
}
