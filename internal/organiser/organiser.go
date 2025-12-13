package organiser

import (
	"fmt"
	"os"
	"path/filepath"

	c "github.com/samcunliffe/bcmp/internal/checker"
	p "github.com/samcunliffe/bcmp/internal/parser"
)

// Config holds package-level configuration for the organiser.
var Config = struct {
	DryRun bool
}{}

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
func CreateDestination(album p.Album, base string) (string, error) {
	if _, err := os.Stat(base); os.IsNotExist(err) {
		// It's noteworthy if, e.g. the user doesn't have a Music folder
		fmt.Printf("Warning: base destination path %s does not exist. Will create it.\n", base)
	}
	destination := filepath.Join(base, album.Artist, album.Title)

	if Config.DryRun {
		fmt.Printf("[Dry run] Create directory: %s\n", destination)
		return destination, nil
	}
	return destination, os.MkdirAll(destination, os.ModePerm)
}

// Construct the full destination path for a track file
//
// Assumes the directory structure is correct, i.e. that CreateDestination has
// been called.
func TrackDestination(t p.Track, destination string) string {
	filename := fmt.Sprintf("%s%s", t.FullTrack, t.FileType)
	return filepath.Join(destination, filename)
}

// Move and rename a single music file
//
// "Tidy" a single music file by moving it to the correct directory structure and
// renaming it appropriately. This is a private helper function used by Tidy.
func moveAndRenameFile(sourcePath, destination string) error {
	sourceFile := filepath.Base(sourcePath)

	if !c.IsValidMusicFile(sourceFile) {
		return fmt.Errorf("file %s is not a valid music file", sourceFile)
	}

	album, track, err := p.ParseMusicFileName(sourceFile)
	if err != nil {
		return err
	}

	destination, err = CreateDestination(album, destination)
	if err != nil {
		return err
	}

	if Config.DryRun {
		fmt.Printf("[Dry run] %s → %s\n", sourcePath, TrackDestination(track, destination))
		return nil
	}

	return os.Rename(sourcePath, TrackDestination(track, destination))
}

// Tidy checks a music file and moves it to the correct directory structure.
func Tidy(musicFile, destination string) error {
	if err := c.CheckFile(musicFile); err != nil {
		return err
	}
	return moveAndRenameFile(musicFile, destination)
}
