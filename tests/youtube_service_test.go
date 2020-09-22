package tests

import (
	"testing"

	"lucasmontano.com/yt-links/models"
)

func TestYoutubeServiceSerialization(t *testing.T) {
	playListItems := []models.PlaylistItem{
		{
			Snippet: models.VideoSnippet{
				Description: "Falando sobre trade-off na programação.\n\nConfira o nosso canal oficial: https://www.youtube.com/lucasmontano\n\nEsse clip foi recortado de uma live na Twitch\nhttps://www.twitch.tv/lucas_montano",
			},
		},
		{
			Snippet: models.VideoSnippet{
				Description: "Esse clip foi recortado de uma streaming onde eu estava comentando um Tweet que tinha feito no mesmo dia.\n\nCanal oficial\nhttps://www.youtube.com/lucasmontano\n\nConfira nosso canal na Twitch: twitch.tv/lucas_montano\nTwitter: https://twitter.com/lucas_montano",
			},
		},
	}

	tables := []struct {
		playlistID string
		expected   models.PlaylistItemsResponse
	}{
		{
			"PL986prvoaFkZiVvqtGvb0rdGUemYjUg2t",
			models.PlaylistItemsResponse(struct{ Items []models.PlaylistItem }{Items: playListItems}),
		},
	}

	for _, table := range tables {
		expected := table.expected
		got := getVideosService(table.playlistID)

		if expected.Items[0].Snippet.Description != got.Items[0].Snippet.Description {
			t.Errorf("Description doenst match!")
		}
	}
}
