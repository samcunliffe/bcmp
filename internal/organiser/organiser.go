package organiser

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/samcunliffe/bcmp/internal/parser"
)

// Determine default destination for music files
//
// Typically $HOME/Music. If $HOME cannot be determined, use current directory.
// Note: this does not check that the directory exists.
func DefaultDestination() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return filepath.Join(home, "Music")
}

// Create the directory structure for the album under base
//
// e.g. if base is $HOME/Music, create $HOME/Music/Artist/Album
func CreateDestination(album parser.Album, base string) (string, error) {
	if _, err := os.Stat(base); os.IsNotExist(err) {
		// It's noteworthy if, e.g. the user doesn't have a Music folder
		fmt.Printf("Warning: base destination path %s does not exist. Will create it.\n", base)
	}
	destination := filepath.Join(base, album.Artist, album.Title)
	return destination, os.MkdirAll(destination, os.ModePerm)
}

// Ensure a zip file or music file exists and is not a directory
func CheckFile(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("error accessing zip file: %v", err)
	}
	if fi.IsDir() {
		return fmt.Errorf("the path: %v is a directory, not a file", fi.Name())
	}
	if fi.Size() == 0 {
		return fmt.Errorf("the file: %v is empty", fi.Name())
	}
	return nil
}

// Construct the full destination path for a track file
//
// Assumes the directory structure is correct, i.e. that CreateDestination has
// been called.
func TrackDestinationPath(t parser.Track, destination string) string {
	filename := fmt.Sprintf("%s%s", t.FullTrack, t.FileType)
	return filepath.Join(destination, filename)
}

// Move and rename a single music file
//
// "Tidy" a single music file by moving it to the correct directory structure and
// renaming it appropriately. This is the main function called by `bcmp tidy`.
func MoveAndRenameFile(sourcePath, destination string) error {
	sourceFile := filepath.Base(sourcePath)
	album, track, err := parser.ParseMusicFileName(sourceFile)
	if err != nil {
		return err
	}

	destination, err = CreateDestination(album, destination)
	if err != nil {
		return err
	}

	destinationPath := TrackDestinationPath(track, destination)
	return os.Rename(sourcePath, destinationPath)
}
