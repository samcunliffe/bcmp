package organiser

import (
	"path/filepath"
	"testing"

	"github.com/samcunliffe/bcmptidy/internal/parser"
)

func TestDefaultDestination(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	want := filepath.Join(home, "Music")
	got := DefaultDestination()
	if got != want {
		t.Errorf("DefaultDestination() = %q; want %q", got, want)
	}
}

func TestDegermineDefaultDestinationNoHome(t *testing.T) {
	t.Setenv("HOME", "")

	want := filepath.Join(".", "Music")
	got := DefaultDestination()
	if got != want {
		t.Errorf("DefaultDestination() = %q; want %q", got, want)
	}
}

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
