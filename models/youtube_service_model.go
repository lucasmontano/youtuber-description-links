package models

type PlaylistItemsResponse struct {
	Items []PlaylistItem `json:"items"`
}

type PlaylistItem struct {
	Snippet VideoSnippet `json:"snippet"`
}

type VideoSnippet struct {
	ResourceId  ResourceId `json:"resourceId"`
	Description string     `json:"description"`
}

type ResourceId struct {
	VideoId string `json:"videoId"`
}
