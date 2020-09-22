package main

func ExtractLinksUseCase(videos []VideoDomainModel) []LinkDomainModel {
	linkMap := make(map[string][]string)

	for _, video := range videos {
		links := DescriptionToLinksMapper(video.Description)
		for _, link := range links {
			linkMap[link] = append(linkMap[link], video.Id)
		}
	}

	var links []LinkDomainModel

	for url, videoItems := range linkMap {
		links = append(links, LinkDomainModel{
			url:    url,
			videos: videoItems,
		})
	}

	return links
}
