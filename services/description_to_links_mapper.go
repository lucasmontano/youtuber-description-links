package services

import "regexp"

//DescriptionToLinksMapper : add description
func DescriptionToLinksMapper(description string) []string {
	re := regexp.MustCompile(`\b(?:https?:)?(?:(?i:[a-z]+\.)+)[^\s,]+\b`)
	return re.FindAllString(description, -1)
}
