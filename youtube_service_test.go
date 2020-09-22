package main

import (
	"testing"
)

func TestYoutubeServiceSerialization(t *testing.T) {
	playListItems := []PlaylistItem{
		{
			Snippet: VideoSnippet{
				Description: "Falando sobre trade-off na programação.\n\nConfira o nosso canal oficial: https://www.youtube.com/lucasmontano\n\nEsse clip foi recortado de uma live na Twitch\nhttps://www.twitch.tv/lucas_montano",
			},
		},
		{
			Snippet: VideoSnippet{
				Description: "Esse clip foi recortado de uma streaming onde eu estava comentando um Tweet que tinha feito no mesmo dia.\n\nCanal oficial\nhttps://www.youtube.com/lucasmontano\n\nConfira nosso canal na Twitch: twitch.tv/lucas_montano\nTwitter: https://twitter.com/lucas_montano",
			},
		},
	}

	tables := []struct {
		playlistID string
		expected   PlaylistItemsResponse
	}{
		{
			"PL986prvoaFkZiVvqtGvb0rdGUemYjUg2t",
			PlaylistItemsResponse(struct{ Items []PlaylistItem }{Items: playListItems}),
		},
	}

	for _, table := range tables {
		expected := table.expected
		got := GetVideosService(table.playlistID)

		if expected.Items[0].Snippet.Description != got.Items[0].Snippet.Description {
			t.Errorf("Description doenst match!")
		}
	}
}
