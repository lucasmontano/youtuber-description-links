package main

import (
	"reflect"
	"testing"
)

func TestExtractLinks(t *testing.T) {
	tables := []struct {
		videoItems      []VideoDomainModel
		expectedLinks []LinkDomainModel
	}{
		{
			[]VideoDomainModel{
				{
					Id:          "Video_1",
					Description: "Meu nome eh Lucas Montano do canal https://youtube.com/LucasMontano",
				},
				{
					Id:          "Video_2",
					Description: "Meu nome eh Lucas Montano do canal https://youtube.com/LucasMontano e Twitch https://twitch.com/Lucas_Montano",
				},
			},
			[]LinkDomainModel{
				{
					url: "youtube.com/LucasMontano",
					videos: []string{
						"Video_1",
						"Video_2",
					},
				},
				{
					url: "twitch.com/Lucas_Montano",
					videos: []string{
						"Video_2",
					},
				},
			},
		},
	}

	for _, table := range tables {
		expected := table.expectedLinks
		got := ExtractLinksUseCase(table.videoItems)
		if !reflect.DeepEqual(expected, got) {
			t.Error("Wrong mapping of links extract")
		}
	}
}
