package parser

import "testing"

func TestExtractNumberPrefix(t *testing.T) {
	var testCases = []struct {
		input_track    string
		want_number    int
		want_trackname string
	}{
		{"01 The Aftermath", 1, "The Aftermath"},
		{"02 Dark Clouds", 2, "Dark Clouds"},
		{"06 The Other Side of Anger", 6, "The Other Side of Anger"},
		{"12 Lord of Ruins", 12, "Lord of Ruins"},
	}
	for _, testcase := range testCases {
		n, s, err := extractNumberPrefix(testcase.input_track)
		if err != nil {
			t.Errorf("extractNumberPrefix returned error for %q: %v", testcase.input_track, err)
		}
		if n != testcase.want_number || s != testcase.want_trackname {
			t.Errorf("extractNumberPrefix failed for %q: got (%d, %q), want (%d, %q)", testcase.input_track, n, s, testcase.want_number, testcase.want_trackname)
		}
	}
}

func TestParseMusicFileName(t *testing.T) {
	var testCases = []struct {
		input_filename string
		want_number    int
		want_title     string
		want_track     string
		want_suffix    string
	}{
		{"Crypta - Shades of Sorrow - 01 The Aftermath.flac", 1, "The Aftermath", "01 The Aftermath", ".flac"},
		{"Crypta - Shades of Sorrow - 02 Dark Clouds.mp3", 2, "Dark Clouds", "02 Dark Clouds", ".mp3"},
		{"Crypta - Shades of Sorrow - 06 The Other Side of Anger.ogg", 6, "The Other Side of Anger", "06 The Other Side of Anger", ".ogg"},
	}
	for _, testcase := range testCases {
		got, err := ParseMusicFileName(testcase.input_filename)
		if err != nil {
			t.Errorf("ParseMusicFileName(%q) returned error: %v", testcase.input_filename, err)
		}
		if got.Number != testcase.want_number {
			t.Errorf("ParseMusicFileName(%q) Number = %d; want %d", testcase.input_filename, got.Number, testcase.want_number)
		}
		if got.Title != testcase.want_title {
			t.Errorf("ParseMusicFileName(%q) Title = %q; want %q", testcase.input_filename, got.Title, testcase.want_title)
		}
		if got.FullTrack != testcase.want_track {
			t.Errorf("ParseMusicFileName(%q) FullTrack = %q; want %q", testcase.input_filename, got.FullTrack, testcase.want_track)
		}
		if got.FileType != testcase.want_suffix {
			t.Errorf("ParseMusicFileName(%q) FileType = %q; want %q", testcase.input_filename, got.FileType, testcase.want_suffix)
		}
	}
}

func TestParseMusicFilenameErrors(t *testing.T) {
	var errorCases = []string{
		"Crypta - Shades of Sorrow - The Aftermath.flac",          // Missing track number
		"Crypta - Shades of Sorrow - 01 The Aftermath.txt",        // Invalid file extension
		"Crypta - Shades of Sorrow - 01 The Aftermath",            // No file extension
		"Crypta - Shades of Sorrow - 00 The Aftermath.flac",       // Track zero
		"Crypta - Shades of Sorrow - Track One The Aftermath.mp3", // Non-numeric track number
		"Crypta Shades of Sorrow 01 The Aftermath.flac",           // No hyphens
		" - 01 The Aftermath.flac",                                // Missing artist and album
		"Just the Song Title.flac",                                // Missing artist, album, track number
	}
	for _, filename := range errorCases {
		_, err := ParseMusicFileName(filename)
		if err == nil {
			t.Errorf("ParseMusicFileName(%q) expected to return error, but got nil", filename)
		}
	}
}
