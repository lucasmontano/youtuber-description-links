package main

import (
	"fmt"
	"lucasmontano.com/yt-links/models"
	"lucasmontano.com/yt-links/services"
)

func main() {
	playlistItemsResponse := services.GetVideosService()
	videoSnippetItems := playlistItemsResponse.Items
	var videoDomainItems []models.VideoDomainModel
	for _, videoSnippet := range videoSnippetItems {
		videoID := videoSnippet.VideoID
		description := videoSnippet.Description
		videoDomainModel := models.VideoDomainModel{
			ID:          videoID,
			Description: description,
		}
		videoDomainItems = append(videoDomainItems, videoDomainModel)
	}
	links := services.ExtractLinksUseCase(videoDomainItems)
	for _, link := range links {
		fmt.Printf("URL: %v, in (%v) videos\r\n", link.URL, len(link.Videos))
	}
}
