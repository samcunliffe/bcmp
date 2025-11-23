package fileparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/samcunliffe/bcmptidy/datamodel"
)

var validMusicFiles = []string{".flac", ".mp3", ".ogg"}

func splitOnHyphen(s string) (string, string) {
	ss := strings.SplitN(s, " - ", 2)
	return ss[0], ss[1]
}

func removeParenthesis(s string) string {
	re := regexp.MustCompile(`\s*\(.*?\)`)
	return strings.TrimSpace(re.ReplaceAllString(s, ""))
}

func extractNumberPrefix(s string) (int, string, error) {
	re := regexp.MustCompile(`^(\d+)\s*(.*)`)

	// Expect "XX Track Name", "XX", "Track Name"
	matches := re.FindStringSubmatch(s)
	if len(matches) != 3 {
		return -1, s, fmt.Errorf("error in regex match for track number extraction")
	}

	// Convert track number from string
	number, err := strconv.Atoi(matches[1])
	if err != nil || number < 1 {
		return -1, s, fmt.Errorf("failed to convert track number '%s' to int: %v", matches[1], err)
	}
	return number, strings.TrimSpace(matches[2]), nil
}

func ParseZipFileName(name string) (datamodel.Album, error) {
	// Trim the .zip suffix
	if !strings.HasSuffix(name, ".zip") {
		return datamodel.Album{}, fmt.Errorf("filename does not have .zip suffix")
	}
	name = strings.TrimSuffix(name, ".zip")

	// Split into artist and album
	if !strings.Contains(name, " - ") {
		return datamodel.Album{}, fmt.Errorf("filename does not contain ' - ' separator")
	}
	artist, album := splitOnHyphen(name)
	return datamodel.Album{Artist: artist, Title: removeParenthesis(album)}, nil
}

func ParseMusicFileName(name string) (datamodel.Track, error) {
	// Trim valid music file suffixes; error if none found
	hadValidSuffix := false
	for _, suffix := range validMusicFiles {
		if strings.HasSuffix(name, suffix) {
			name = strings.TrimSuffix(name, suffix)
			hadValidSuffix = true
			break
		}
	}
	if !hadValidSuffix {
		return datamodel.Track{}, fmt.Errorf("filename does not have a valid music file suffix")
	}

	// Should be two hyphens 'Artist - Album - XX Title'
	if !strings.Contains(name, " - ") {
		return datamodel.Track{}, fmt.Errorf("filename does not contain ' - ' separator")
	}
	if strings.Count(name, " - ") != 2 {
		fmt.Printf("expected two ' - ' separators: '%s'", name)
	}

	// Split into artist, album, number, track title
	_, albumAndTrack := splitOnHyphen(name)
	_, track := splitOnHyphen(albumAndTrack)
	number, title, err := extractNumberPrefix(track)
	if err != nil {
		return datamodel.Track{Title: name}, fmt.Errorf("failed to extract track number and title: %v", err)
	}

	return datamodel.Track{Number: number, Title: title, FullTrack: track}, nil
}
