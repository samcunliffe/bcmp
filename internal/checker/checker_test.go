package checker

import "testing"

func TestIsZipFile(t *testing.T) {
	var testCases = []struct {
		input string
		want  bool
	}{
		{"archive.zip", true},
		{"ARCHIVE.ZIP", true},
		{"archive.rar", false},
		{"archive", false},
		{"archive.zipx", false},
	}
	for _, testcase := range testCases {
		got := IsZipFile(testcase.input)
		if got != testcase.want {
			t.Errorf("IsZipFile(%q) = %v; want %v", testcase.input, got, testcase.want)
		}
	}
}

func TestIsCoverArtFile(t *testing.T) {
	var testCases = []struct {
		input string
		want  bool
	}{
		{"cover.jpg", true},
		{"COVER.PNG", true},
		{"folder.jpg", true},
		{"FOLDER.PNG", true},
		{"not_cover.jpg", false},
		{"flibble", false},
		{"cover.jpeg", false},
		{"some_other_file.mp3", false},
		{"some_other_file.flac", false},
	}
	for _, testcase := range testCases {
		got := IsCoverArtFile(testcase.input)
		if got != testcase.want {
			t.Errorf("IsCoverArtFile(%q) = %v; want %v", testcase.input, got, testcase.want)
		}
	}
}

func TestIsValidMusicFile(t *testing.T) {
	var testCases = []struct {
		input string
		want  bool
	}{
		{"Crypta - Shades of Sorrow - 01 The Aftermath.txt", false}, // Invalid file extension
		{"Crypta - Shades of Sorrow - 01 The Aftermath", false},     // No file extension
		{"Crypta - Shades of Sorrow - 01 The Aftermath.mp3", true},  // Valid files
		{"Crypta - Shades of Sorrow - 01 The Aftermath.flac", true},
	}
	for _, testcase := range testCases {
		if IsValidMusicFile(testcase.input) != testcase.want {
			t.Errorf("IsValidMusicFile(%q) = %v", testcase.input, !testcase.want)
		}
	}
}

func TestCheckFileDirectory(t *testing.T) {
	err := CheckFile("testdata/directory")
	if err == nil {
		t.Errorf("CheckFile on directory did not return error")
	}
}

func TestCheckFileNonExistent(t *testing.T) {
	err := CheckFile("testdata/nonexistent.zip")
	if err == nil {
		t.Errorf("CheckFile on nonexistent file did not return error")
	}
}

func TestCheckFileEmptyFile(t *testing.T) {
	err := CheckFile("testdata/emptyfile")
	if err == nil {
		t.Errorf("CheckFile on empty file did not return error")
	}
}

func TestCheckFileValidZipFile(t *testing.T) {
	err := CheckFile("testdata/validfile.zip")
	if err != nil {
		t.Errorf("CheckFile on valid file returned error: %v", err)
	}
}

func TestCheckFileValidMusicFile(t *testing.T) {
	err := CheckFile("testdata/ding.flac")
	if err != nil {
		t.Errorf("CheckFile on valid music file returned error: %v", err)
	}
}
