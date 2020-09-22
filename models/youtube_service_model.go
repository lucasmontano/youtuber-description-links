package models

//PlaylistItemsResponse : add description
type PlaylistItemsResponse struct {
	Items []PlaylistItem `json:"items"`
}

//PlaylistItem : add description
type PlaylistItem struct {
	Snippet VideoSnippet `json:"snippet"`
}

//VideoSnippet : add description
type VideoSnippet struct {
	ResourceID  ResourceID `json:"resourceId"`
	Description string     `json:"description"`
}

//ResourceID : add description
type ResourceID struct {
	VideoID string `json:"videoId"`
}
