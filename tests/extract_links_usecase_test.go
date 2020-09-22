package tests

import (
	"reflect"
	"testing"

	"lucasmontano.com/yt-links/models"
	"lucasmontano.com/yt-links/services"
)

func TestExtractLinks(t *testing.T) {
	tables := []struct {
		videoItems    []models.VideoDomainModel
		expectedLinks []models.LinkDomainModel
	}{
		{
			[]models.VideoDomainModel{
				{
					ID:          "Video_1",
					Description: "Meu nome eh Lucas Montano do canal https://youtube.com/LucasMontano",
				},
				{
					ID:          "Video_2",
					Description: "Meu nome eh Lucas Montano do canal https://youtube.com/LucasMontano e Twitch https://twitch.com/Lucas_Montano",
				},
			},
			[]models.LinkDomainModel{
				{
					URL: "youtube.com/LucasMontano",
					Videos: []string{
						"Video_1",
						"Video_2",
					},
				},
				{
					URL: "twitch.com/Lucas_Montano",
					Videos: []string{
						"Video_2",
					},
				},
			},
		},
	}

	for _, table := range tables {
		expected := table.expectedLinks
		got := services.ExtractLinksUseCase(table.videoItems)
		if !reflect.DeepEqual(expected, got) {
			t.Error("Wrong mapping of links extract")
		}
	}
}
