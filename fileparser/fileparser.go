package fileparser

import (
	"fmt"
	"strings"
)

func ParseZipFileName(name string) (map[string]string, error) {
	// Trim the .zip suffix
	if !strings.HasSuffix(name, ".zip") {
		return nil, fmt.Errorf("filename does not have .zip suffix")
	}
	name = strings.TrimSuffix(name, ".zip")

	// Split into artist and album
	if !strings.Contains(name, " - ") {
		return nil, fmt.Errorf("filename does not contain ' - ' separator")
	}
	nameAndAlbum := strings.SplitN(name, " - ", 2)
	return map[string]string{"artist": nameAndAlbum[0], "album": nameAndAlbum[1]}, nil
}

func ParseMusicFileName(name string) (map[string]string, error) {
	return nil, nil
}
