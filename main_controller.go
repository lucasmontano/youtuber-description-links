package main

import "lucasmontano.com/yt-links/models"

/**
 * args[1]: playlistId
 */

func mainController(args []string) {
	playlistItemsResponse := getVideosService(args[1])
	videoSnippetItems := playlistItemsResponse.Items
	var videoDomainItems []models.VideoDomainModel
	for _, videoSnippet := range videoSnippetItems {
		videoID := videoSnippet.Snippet.ResourceID.VideoID
		description := videoSnippet.Snippet.Description
		videoDomainModel := models.VideoDomainModel{
			ID:          videoID,
			Description: description,
		}
		videoDomainItems = append(videoDomainItems, videoDomainModel)
	}
	extractLinksUseCase(videoDomainItems)
}
