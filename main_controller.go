package main

/**
 * args[1]: playlistId
 */
func MainController(args []string) {
	playlistItemsResponse := GetVideosService(args[1])
	videoSnippetItems := playlistItemsResponse.Items
	var videoDomainItems []VideoDomainModel
	for _, videoSnippet := range videoSnippetItems {
		videoId := videoSnippet.Snippet.ResourceId.VideoId
		description := videoSnippet.Snippet.Description
		videoDomainModel := VideoDomainModel{
			Id:          videoId,
			Description: description,
		}
		videoDomainItems = append(videoDomainItems, videoDomainModel)
	}
	ExtractLinksUseCase(videoDomainItems)
}
