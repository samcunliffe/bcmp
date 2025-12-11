package organiser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/samcunliffe/bcmp/internal/parser"
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

	for _, testcase := range testCases {
		got, err := CreateDestination(testcase.album, base)
		if err != nil {
			t.Errorf("CreateDestination(%v, %q) returned error: %v", testcase.album, base, err)
			continue
		}
		if got != testcase.want {
			t.Errorf("CreateDestination(%v, %q) = %q; want %q", testcase.album, base, got, testcase.want)
		}
		if _, err := os.Stat(got); os.IsNotExist(err) {
			t.Errorf("CreateDestination did not create directory %q", got)
		}
	}
}

func TestCreateDestinationNonExistentBase(t *testing.T) {
	base := filepath.Join(t.TempDir(), "some_nonexistent_subdir")

	album := parser.Album{Artist: "Crypta", Title: "Shades of Sorrow"}
	got, err := CreateDestination(album, base)
	if err != nil {
		t.Errorf("CreateDestination(%v, %q) returned error: %v", album, base, err)
	}
	want := filepath.Join(base, "Crypta", "Shades of Sorrow")
	if got != want {
		t.Errorf("CreateDestination(%v, %q) = %q; want %q", album, base, got, want)
	}
	// TODO: Might actually be better to do something else than write a warning to stdout...
	// Investigate logging and a -v/--verbose flag?
}

// Utility test helper to recreate a moved file after testing is done.
func putBackFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic("unable to put back file in testdata: " + err.Error())
	}
	if err := os.WriteFile(path, []byte("Just a non-empty test file."), 0644); err != nil {
		panic("unable to write test file in testdata: " + err.Error())
	}
	f.Close()
}

func TestTidy(t *testing.T) {
	destination := t.TempDir()
	source := "testdata/Artist - Album - 01 Track.flac"
	defer putBackFile(source)

	err := Tidy(source, destination)
	if err != nil {
		t.Fatalf("Tidy(%q, %q) returned error: %v", source, destination, err)
	}

	wantPath := filepath.Join(destination, "Artist", "Album", "01 Track.flac")
	if _, err := os.Stat(wantPath); os.IsNotExist(err) {
		t.Fatalf("Tidy did not create file at %q", wantPath)
	}
}

func TestTidyNonExistentFile(t *testing.T) {
	destination := "./"
	source := "Non Existant Artist - Non Existant Album - 01 Track.flac"

	err := Tidy(source, destination)
	if err == nil {
		t.Fatalf("Tidy(%q, %q) didn't return an error!", source, destination)
	}

	want := "no such file or directory"
	if !strings.Contains(err.Error(), want) {
		t.Fatalf("Tidy(%q, %q) error = %q; want %q", source, destination, err.Error(), want)
	}
}
