package services

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func buildOAuthHTTPClient() *http.Client {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/youtube-go-quickstart.json
	config, err := google.ConfigFromJSON(b, youtube.YoutubepartnerScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return getClient(ctx, config)
}
