package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	for _, tc := range testCases {
		got := IsZipFile(tc.input)
		assert.Equal(t, tc.want, got, "IsZipFile(%q) = %v; want %v", tc.input, got, tc.want)
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
	for _, tc := range testCases {
		got := IsCoverArtFile(tc.input)
		assert.Equal(t, tc.want, got, "IsCoverArtFile(%q) = %v; want %v", tc.input, got, tc.want)
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
	for _, tc := range testCases {
		got := IsValidMusicFile(tc.input)
		assert.Equal(t, tc.want, got, "IsValidMusicFile(%q) = %v; want %v", tc.input, got, tc.want)
	}
}

func TestCheckFileDirectory(t *testing.T) {
	err := CheckFile("testdata/directory")
	assert.Error(t, err, "CheckFile on directory did not return error")
	assert.Contains(t, err.Error(), "is a directory", "CheckFile on directory returned wrong error: %v", err)
}

func TestCheckFileNonExistent(t *testing.T) {
	err := CheckFile("testdata/nonexistent.zip")
	assert.Error(t, err, "CheckFile on nonexistent file did not return error")
	assert.Contains(t, err.Error(), "no such file or directory", "CheckFile on nonexistent file returned wrong error: %v", err)
}

func TestCheckFileEmptyFile(t *testing.T) {
	err := CheckFile("testdata/emptyfile")
	assert.Error(t, err, "CheckFile on empty file did not return error")
	assert.Contains(t, err.Error(), "is empty", "CheckFile on empty file returned wrong error: %v", err)
}

func TestCheckFileValidZipFile(t *testing.T) {
	err := CheckFile("testdata/validfile.zip")
	assert.NoError(t, err, "CheckFile on valid zip file returned error: %v", err)
}

func TestCheckFileValidMusicFile(t *testing.T) {
	err := CheckFile("testdata/ding.flac")
	assert.NoError(t, err, "CheckFile on valid music file returned error: %v", err)
}
