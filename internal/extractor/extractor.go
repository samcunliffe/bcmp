package extractor

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	o "github.com/samcunliffe/bcmp/internal/organiser"
	p "github.com/samcunliffe/bcmp/internal/parser"
)

// Performs safety checks on the zip file entry.
func check(f *zip.File) error {
	// Check that the file path is not suspicious - in any case we will extract
	// it to the music directory and forget the path. But if we detect this then
	// stop processing as the zip file should be treated with caution.
	//
	// https://security.snyk.io/research/zip-slip-vulnerability
	if !filepath.IsLocal(f.Name) {
		return fmt.Errorf("archive contains invalid file path: %s,\ntreat the zip file with caution", f.Name)
	}

	// Any directories - this is likely not a bandcamp file
	if f.FileInfo().IsDir() {
		return fmt.Errorf("archive contains a directory, not a valid bandcamp zip")
	}

	if !p.IsValidMusicFile(filepath.Base(f.Name)) {
		return fmt.Errorf("filename does not have a valid music file suffix: %s", f.Name)
	}

	return nil
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

func unzipAndRename(zipPath, destination string) error {
	// Open zip archive for reading.
	rc, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("impossible to open zip reader: %s", err)
	}
	defer rc.Close()

	// Iterate through the files in the archive,
	for _, f := range rc.File {

		// Ignore cover art for now
		if p.IsCoverArtFile(filepath.Base(f.Name)) {
			fmt.Printf("Skipping: %s\n", f.Name)
			continue
		}

		// Preflight checks
		if err := check(f); err != nil {
			return err
		}

		err := processTrack(f, destination)
		if err != nil {
			return err
		}
	}
	return nil
}

// Extract tracks from a Bandcamp zip file and rename them appropriately
// The functon called by `bcmp extract`. Does all checking then calls extractAndRename.
func Extract(zipFilePath, destination string) error {
	err := o.CheckFile(zipFilePath)
	if err != nil {
		return err
	}

	album, err := p.ParseZipFileName(filepath.Base(zipFilePath))
	if err != nil {
		return err
	}

	destination, err = o.CreateDestination(album, destination)
	if err != nil {
		return err
	}

	return unzipAndRename(zipFilePath, destination)
}
