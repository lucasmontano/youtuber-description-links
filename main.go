package main

import (
	"bufio"
	"fmt"
	"lucasmontano.com/yt-links/models"
	"lucasmontano.com/yt-links/services"
	"os"
	"strings"
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
	// TODO add links to a repository
	linksMap := make(map[string]models.LinkDomainModel)
	for _, link := range links {
		linksMap[link.URL] = link
		fmt.Printf("URL: %v, found in (%v) videos\r\n", link.URL, len(link.Videos))
	}

	scanner(linksMap)
}

func scanner(linksMap map[string]models.LinkDomainModel) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urls := strings.Split(scanner.Text(), " ")
		currentURL := urls[0]
		linkToBeUpdated := linksMap[currentURL]
		services.UpdateVideo(linkToBeUpdated, urls[1])
	}
}
