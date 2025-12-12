package organiser

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/samcunliffe/bcmp/internal/bcmptest"
	"github.com/samcunliffe/bcmp/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestCreateDestination(t *testing.T) {
	base := t.TempDir()

	testCases := []struct {
		album parser.Album
		want  string
	}{
		{
			album: parser.Album{Artist: "Crypta", Title: "Shades of Sorrow"},
			want:  filepath.Join(base, "Crypta", "Shades of Sorrow"),
		},
		{
			album: parser.Album{Artist: "Orbit Culture", Title: "Death Above Life"},
			want:  filepath.Join(base, "Orbit Culture", "Death Above Life"),
		},
	}

	for _, tc := range testCases {
		got, err := CreateDestination(tc.album, base)
		assert.NoError(t, err, "CreateDestination(%v, %q) returned error: %v", tc.album, base, err)
		assert.Equal(t, tc.want, got, "CreateDestination(%v, %q) = %q; want %q", tc.album, base, got, tc.want)
		assert.DirExists(t, got, "CreateDestination did not create directory %q", got)
	}
}

func TestCreateDestinationNonExistentBase(t *testing.T) {
	base := filepath.Join(t.TempDir(), "some_nonexistent_subdir")

	album := parser.Album{Artist: "Crypta", Title: "Shades of Sorrow"}
	got, err := CreateDestination(album, base)
	assert.NoError(t, err, "CreateDestination(%v, %q) returned error: %v", album, base, err)

	want := filepath.Join(base, "Crypta", "Shades of Sorrow")
	assert.Equal(t, want, got, "CreateDestination(%v, %q) = %q; want %q", album, base, got, want)
	// TODO: Might actually be better to do something else than write a warning to stdout...
	// Investigate logging and a -v/--verbose flag?
}

func TestTidy(t *testing.T) {
	destination := t.TempDir()
	source := "testdata/Artist - Album - 01 Track.flac"
	defer bcmptest.PutFileBack(t, source)

	err := Tidy(source, destination)
	assert.NoError(t, err, "Tidy(%q, %q) returned error: %v", source, destination, err)

	wantPath := filepath.Join(destination, "Artist", "Album", "01 Track.flac")
	assert.FileExists(t, wantPath, "Tidy did not move file to %q", wantPath)
}

func TestTidyNonExistentFile(t *testing.T) {
	destination := "./"
	source := "Non Existent Artist - Non Existent Album - 01 Track.flac"

	err := Tidy(source, destination)
	assert.Error(t, err, "Tidy(%q, %q) didn't return an error!", source, destination)

	want := "no such file or directory"
	if !strings.Contains(err.Error(), want) {
		t.Fatalf("Tidy(%q, %q) error = %q; want %q", source, destination, err.Error(), want)
	}
}

func TestTidyInvalidFilename(t *testing.T) {
	destination := "./"
	source := "testdata/Un-Parsable Filename.flac"
	defer bcmptest.PutFileBack(t, source)

	err := Tidy(source, destination)
	if err == nil {
		t.Fatalf("Tidy(%q, %q) didn't return an error!", source, destination)
	}
	want := "does not contain ' - '"
	if !strings.Contains(err.Error(), want) {
		t.Fatalf("Tidy(%q, %q) error = %q; want %q", source, destination, err.Error(), want)
	}
}

func TestTidyNonMusicFile(t *testing.T) {
	destination := "./"
	source := "testdata/Not a Music File.txt"
	defer bcmptest.PutFileBack(t, source)

	err := Tidy(source, destination)
	assert.ErrorContains(t, err, "not a valid music file", "Tidy(%q, %q) didn't return a/the expected error", source, destination)
}
