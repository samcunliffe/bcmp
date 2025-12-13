package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractNumberPrefix(t *testing.T) {
	var testCases = []struct {
		input      string
		wantNumber int
		wantName   string
	}{
		{"01 The Aftermath", 1, "The Aftermath"},
		{"02 Dark Clouds", 2, "Dark Clouds"},
		{"06 The Other Side of Anger", 6, "The Other Side of Anger"},
		{"12 Lord of Ruins", 12, "Lord of Ruins"},
	}
	for _, tc := range testCases {
		n, s, err := numberPrefix(tc.input)
		assert.NoError(t, err, "extractNumberPrefix(%q) returned error: %v", tc.input, err)
		assert.Equal(t, tc.wantNumber, n)
		assert.Equal(t, tc.wantName, s)
	}
}

func TestParseMusicFileName(t *testing.T) {
	var testCases = []struct {
		inputFilename string
		titleCaseCfg  bool
		wantNumber    int
		wantTitle     string
		wantTrack     string
		wantSuffix    string
	}{
		{"Crypta - Shades of Sorrow - 01 The Aftermath.flac", false, 1, "The Aftermath", "01 The Aftermath", ".flac"},
		{"Crypta - Shades of Sorrow - 02 Dark Clouds.mp3", false, 2, "Dark Clouds", "02 Dark Clouds", ".mp3"},
		{"Crypta - Shades of Sorrow - 06 The Other Side of Anger.ogg", false, 6, "The Other Side of Anger", "06 The Other Side of Anger", ".ogg"},
		{"crypta - shades of sorrow - 06 the other side of anger.flac", true, 6, "The Other Side of Anger", "06 The Other Side of Anger", ".flac"},
		{"CRYPTA - SHADES OF SORROW - 12 LORD OF RUINS.MP3", true, 12, "Lord of Ruins", "12 Lord of Ruins", ".mp3"},
	}
	wantArtist := "Crypta"
	wantAlbum := "Shades of Sorrow"
	for _, tc := range testCases {
		t.Run(tc.inputFilename, func(t *testing.T) {

			// if we're also testing the title case flag
			if tc.titleCaseCfg {
				Config.TitleCase = true
				defer func() { Config.TitleCase = false }()
			}

			// Do the parsing
			gotAlbum, gotTrack, err := ParseMusicFileName(tc.inputFilename)
			assert.NoError(t, err, "ParseMusicFileName(%q) returned error: %v", tc.inputFilename, err)

			assert.Equal(t, wantArtist, gotAlbum.Artist)
			assert.Equal(t, wantAlbum, gotAlbum.Title)
			assert.Equal(t, tc.wantNumber, gotTrack.Number)
			assert.Equal(t, tc.wantTitle, gotTrack.Title)
			assert.Equal(t, tc.wantTrack, gotTrack.FullTrack)
			assert.Equal(t, tc.wantSuffix, gotTrack.FileType)
		})
	}
}

func TestParseMusicFilenameErrors(t *testing.T) {
	var errorCases = []string{
		"Crypta - Shades of Sorrow - The Aftermath.flac",          // Missing track number
		"Crypta - Shades of Sorrow - 00 The Aftermath.flac",       // Track zero
		"Crypta - Shades of Sorrow - Track One The Aftermath.mp3", // Non-numeric track number
		"Crypta Shades of Sorrow 01 The Aftermath.flac",           // No hyphens
		" - 01 The Aftermath.flac",                                // Missing artist and album
		"Just the Song Title.flac",                                // Missing artist, album, track number
	}
	for _, filename := range errorCases {
		_, _, err := ParseMusicFileName(filename)
		assert.Error(t, err, "ParseMusicFileName(%q) expected to return error", filename)
	}
}
