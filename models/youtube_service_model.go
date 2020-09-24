package models

type PlaylistItemsResponse struct {
	Items []VideoSnippet
}

type VideoSnippet struct {
	VideoID  string
	Description string
}
