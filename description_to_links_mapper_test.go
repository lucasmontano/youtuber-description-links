package main

import "testing"

func TestExtractLinksQt(t *testing.T) {

	tables := []struct {
		description     string
		expectedQtLinks int
	}{
		{
			"Meu nome eh Lucas Montano do canal https://youtube.com/LucasMontano",
			1,
		},
		{
			"Meu nome eh Lucas Montano do canal https://youtube.com/LucasMontano e https://twitch.com/Lucas_Montano",
			2,
		},
		{
			"Meu nome eh Lucas Montano do canal https://instagram.com/LucasMontano https://youtube.com/LucasMontano e https://twitch.com/Lucas_Montano",
			3,
		},
		{
			"Meu nome eh Lucas Montano do canal Lucas Montano",
			0,
		},
	}

	for _, table := range tables {
		expected := table.expectedQtLinks
		got := DescriptionToLinksMapper(table.description)
		if expected != len(got) {
			t.Errorf("Wrong qt. of links extract, got: %d, want: %d.", len(got), expected)
		}
	}
}
