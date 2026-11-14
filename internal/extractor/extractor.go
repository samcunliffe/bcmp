package extractor

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	c "github.com/samcunliffe/bcmp/internal/checker"
	o "github.com/samcunliffe/bcmp/internal/organiser"
	p "github.com/samcunliffe/bcmp/internal/parser"
)

type FileKind uint8

const (
	Unknown FileKind = iota
	Music
	CoverArt
)

// Performs safety checks on the zip file entry.
// If the file is a valid music file it returns Music with no error.
// If it is a valid cover art file it returns CoverArt with no error.
func check(f *zip.File) (FileKind, error) {
	// Check that the file path is not suspicious - in any case we will extract
	// it to the music directory and forget the path. But if we detect this then
	// stop processing as the zip file should be treated with caution.
	//
	// https://security.snyk.io/research/zip-slip-vulnerability
	if !filepath.IsLocal(f.Name) {
		return Unknown, fmt.Errorf("archive contains invalid file path: %s,\ntreat the zip file with caution", f.Name)
	}

	// Any directories - this is likely not a bandcamp file
	if f.FileInfo().IsDir() {
		return Unknown, fmt.Errorf("archive contains a directory, not a valid bandcamp zip")
	}

	switch {
	case c.IsValidMusicFile(filepath.Base(f.Name)):
		return Music, nil
	case c.IsCoverArtFile(filepath.Base(f.Name)):
		return CoverArt, nil
	default:
		// Might make sense to simply ignore unknown files, but for now: treat as an error.
		return Unknown, fmt.Errorf("archive contains an unknown file type: %s", f.Name)
	}
}

// Process a single track file from the zip archive.
// Assumes sanity checks have already been performed.
func processTrack(f *zip.File, destination string) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("impossible to open file in archive: %s", err)
	}
	defer rc.Close()

	_, track, err := p.ParseMusicFileName(f.Name)
	if err != nil {
		return fmt.Errorf("error parsing music file name %s: %v", f.Name, err)
	}

	if o.Config.DryRun {
		fmt.Printf("[Dry run] Extract track → %s\n", o.TrackDestination(track, destination))
		return nil
	}

	// Actually extract and write the file
	fd, err := os.Create(o.TrackDestination(track, destination))
	if err != nil {
		return fmt.Errorf("impossible to create destination file: %s", err)
	}
	defer fd.Close()

	_, err = io.Copy(fd, rc)
	if err != nil {
		return fmt.Errorf("impossible to copy file contents to destination: %s", err)
	}

	fmt.Printf("Extracted track %d: %s\n", track.Number, track.Title)
	return nil
}

// Process the cover art file from the zip archive.
// Also assumes sanity checks have already been performed.
func processCover(f *zip.File, destination string) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("impossible to open file in archive: %s", err)
	}
	defer rc.Close()

	name := filepath.Base(f.Name)
	fullDestination := filepath.Join(destination, name)
	if o.Config.DryRun {
		fmt.Printf("[Dry run] Extract cover art → %s\n", fullDestination)
		return nil
	}

	// Actually extract and write the file
	fd, err := os.Create(fullDestination)
	if err != nil {
		return fmt.Errorf("impossible to create destination file: %s", err)
	}
	defer fd.Close()

	_, err = io.Copy(fd, rc)
	if err != nil {
		return fmt.Errorf("impossible to copy file contents to destination: %s", err)
	}

	fmt.Printf("Extracted cover art: %s\n", name)
	return nil
}

func unzipAndRename(zipPath, destination string) error {
	// Open zip archive for reading.
	rc, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("impossible to open zip reader: %s", err)
	}
	defer rc.Close()

	// Iterate through the files in the archive,
	for _, f := range rc.File {

		kind, err := check(f)
		if err != nil {
			return err
		}

		switch kind {
		case Music:
			err = processTrack(f, destination)
		case CoverArt:
			err = processCover(f, destination)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Extract tracks from a Bandcamp zip file and rename them appropriately
// The function called by `bcmp extract`. Does all checking then calls unzipAndRename.
func Extract(zipFilePath, destination string) error {
	err := c.CheckFile(zipFilePath)
	if err != nil {
		return err
	}
	if !c.IsZipFile(zipFilePath) {
		return fmt.Errorf("file is not a zip archive")
	}

	album, err := p.ParseZipFileName(filepath.Base(zipFilePath))
	if err != nil {
		return err
	}

	destination, err = o.CreateDestination(album, destination)
	if err != nil {
		return err
	}
	if !o.Config.DryRun {
		fmt.Printf("\"%s\"\n", destination)
	}

	return unzipAndRename(zipFilePath, destination)
}
