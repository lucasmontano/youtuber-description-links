package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"lucasmontano.com/yt-links/models"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

const missingClientSecretsMessage = `
Please configure OAuth 2.0
`

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

//GetVideosService : add description
func GetVideosService() models.PlaylistItemsResponse {
	var videoItems []models.VideoSnippet

	service, err := youtube.New(buildOAuthHTTPClient())

	handleError(err, "Error creating YouTube client")

	// Start making YouTube API calls.
	// Call the channels.list method. Set the mine parameter to true to
	// retrieve the playlist ID for uploads to the authenticated user's
	// channel.
	call := service.Channels.List([]string{"contentDetails"}).Id("UCxHdQf7t6ECudqDstXOxRmw")

	response, err := call.Do()
	if err != nil {
		// The channels.list method call returned an error.
		log.Fatalf("Error making API call to list channels: %v", err.Error())
	}

	for _, channel := range response.Items {
		var nextPageToken string
		playlistID := channel.ContentDetails.RelatedPlaylists.Uploads
		// Print the playlist ID for the list of uploaded videos.
		fmt.Printf("Videos in list %s\r\n", playlistID)

		for {
			// Call the playlistItems.list method to retrieve the
			// list of uploaded videos. Each request retrieves 50
			// videos until all videos have been retrieved.
			playlistCall := service.PlaylistItems.List([]string{"snippet"}).
				PlaylistId(playlistID).
				MaxResults(50).
				PageToken(nextPageToken)

			playlistResponse, err := playlistCall.Do()

			if err != nil {
				// The playlistItems.list method call returned an error.
				log.Fatalf("Error fetching playlist items: %v", err.Error())
			}

			for _, playlistItem := range playlistResponse.Items {
				videoItems = append(videoItems, models.VideoSnippet{
					VideoID:     playlistItem.Snippet.ResourceId.VideoId,
					Description: playlistItem.Snippet.Description,
				})
			}

			// Set the token to retrieve the next page of results
			// or exit the loop if all results have been retrieved.
			nextPageToken = playlistResponse.NextPageToken
			if nextPageToken == "" {
				break
			}
			fmt.Println()
		}
	}

	return models.PlaylistItemsResponse{Items: videoItems}
}

func UpdateVideo(link models.LinkDomainModel, newURL string) {
	service, err := youtube.New(buildOAuthHTTPClient())

	handleError(err, "Error creating YouTube client")

	for _, videoID := range link.Videos {
		updateVideo(videoID, link.URL, newURL, service, err)
	}
}

func updateVideo(videoID string, oldURL, newURL string, service *youtube.Service, err error) {
	call := service.Videos.List([]string{"snippet"}).Id(videoID)
	response, err := call.Do()
	if response == nil {
		fmt.Printf("Video not found: " + videoID)
		return
	}
	for _, video := range response.Items {
		originalDescription := video.Snippet.Description
		updatedDescription := strings.Replace(originalDescription, oldURL, newURL, -1)
		video.Snippet.Description = updatedDescription
		updateCall := service.Videos.Update([]string{"snippet"}, video)
		updateResponse, updateErr := updateCall.Do()
		handleError(updateErr, "Error updating YouTube video")
		if updateErr != nil {
			fmt.Printf("New Description: " + updateResponse.Snippet.Description)
		}
	}
}
