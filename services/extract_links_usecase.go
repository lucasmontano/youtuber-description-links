package services

import "lucasmontano.com/yt-links/models"

//ExtractLinksUseCase : add description
func ExtractLinksUseCase(videos []models.VideoDomainModel) []models.LinkDomainModel {
	linkMap := make(map[string][]string)

	for _, video := range videos {
		links := descriptionToLinksMapper(video.Description)
		for _, link := range links {
			linkMap[link] = append(linkMap[link], video.ID)
		}
	}

	var links []models.LinkDomainModel

	for url, videoItems := range linkMap {
		links = append(links, models.LinkDomainModel{
			URL:    url,
			Videos: videoItems,
		})
	}

	return links
}
