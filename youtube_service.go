package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"lucasmontano.com/yt-links/models"
)

func getVideosService(playlistID string) models.PlaylistItemsResponse {

	// Open our jsonFile
	jsonFile, err := os.Open("playlist-" + playlistID + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var playlistItemsResponse models.PlaylistItemsResponse
	errGJson := json.Unmarshal(byteValue, &playlistItemsResponse)
	if errGJson != nil {
		fmt.Println(err)
	}

	return playlistItemsResponse
}
