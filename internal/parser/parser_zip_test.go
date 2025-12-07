package parser

import "testing"

func TestParseZipFileName(t *testing.T) {
	var testCases = []struct {
		inputFilename string
		titleCaseCfg  bool
		wantArtist    string
		wantAlbum     string
	}{
		{"Architects - For Those That Wish To Exist.zip", false, "Architects", "For Those That Wish To Exist"},
		{"Bloodywood - Nu Delhi.zip", false, "Bloodywood", "Nu Delhi"},
		{"Crypta - Shades of Sorrow (pre-order).zip", false, "Crypta", "Shades of Sorrow"},
		{"Crypta - Shades of Sorrow.zip", false, "Crypta", "Shades of Sorrow"},
		{"Enslaved - Heimdal.zip", false, "Enslaved", "Heimdal"},
		{"Immovable Stone - Sylosis.zip", false, "Immovable Stone", "Sylosis"},
		{"Lokust - Infidel.zip", false, "Lokust", "Infidel"},
		{"Malevolence - Malicious Intent.zip", false, "Malevolence", "Malicious Intent"},
		{"Nervosa - Jailbreak (pre-order).zip", false, "Nervosa", "Jailbreak"},
		{"Nervosa - Jailbreak.zip", false, "Nervosa", "Jailbreak"},
		{"Orbit Culture - Death Above Life (24-bit HD audio) (pre-order).zip", false, "Orbit Culture", "Death Above Life"},
		{"Orbit Culture - Rasen.zip", false, "Orbit Culture", "Rasen"},
		{"Pallbearer - Mind Burns Alive (pre-order).zip", false, "Pallbearer", "Mind Burns Alive"},
		{"NERVOSA - JAILBREAK.ZIP", true, "Nervosa", "Jailbreak"},
		{"BLEED FROM WITHIN - SHRINE.ZIP", true, "Bleed from Within", "Shrine"},
	}
	for _, testcase := range testCases {
		// if we're also testing the title case flag
		if testcase.titleCaseCfg {
			Config.TitleCase = true
		}
		got, err := ParseZipFileName(testcase.inputFilename)
		Config.TitleCase = false // reset after test execution

		if err != nil {
			t.Errorf("ParseZipFileName(%q) returned error: %v", testcase.inputFilename, err)
		}
		if got.Artist != testcase.wantArtist || got.Title != testcase.wantAlbum {
			t.Errorf("ParseZipFileName(%q) = artist: %q, album: %q; want artist: %q, album: %q",
				testcase.inputFilename, got.Artist, got.Title, testcase.wantArtist, testcase.wantAlbum)
		}

	}
}

func TestParseZipFileNameErrors(t *testing.T) {
	var errorCases = []string{
		"NoHyphenHere.zip",
		"Too - Many - Hyphens.zip",
		"MissingSuffix",
		"Also Missing Hyphen.zip",
		"Not a zip file.tar.gz",
	}
	for _, filename := range errorCases {
		_, err := ParseZipFileName(filename)
		if err == nil {
			t.Errorf("ParseZipFileName(%q) expected error but got nil", filename)
		}
	}
}
