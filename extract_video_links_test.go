package main

import (
	"testing"

	"./models"
)

func TestExtractVideoLinksQt(t *testing.T) {

	tables := []models.Video{
		{
			Description:     "Meu nome eh Lucas Montano do canal https://youtube.com/LucasMontano",
			ExpectedQtLinks: 1,
		},
		{
			Description:     "Meu nome eh Lucas Montano do canal https://youtube.com/LucasMontano e https://twitch.com/Lucas_Montano",
			ExpectedQtLinks: 2,
		},
		{
			Description:     "Meu nome eh Lucas Montano do canal https://instagram.com/LucasMontano https://youtube.com/LucasMontano e https://twitch.com/Lucas_Montano",
			ExpectedQtLinks: 3,
		},
		{
			Description:     "Meu nome eh Lucas Montano do canal Lucas Montano",
			ExpectedQtLinks: 0,
		},
	}

	for _, table := range tables {
		expected := table.ExpectedQtLinks
		got := ExtractVideoLinks(table.Description)
		if expected != len(got) {
			t.Errorf("Wrong qt. of links extract, got: %d, want: %d.", len(got), expected)
		}
	}
}
