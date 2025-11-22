package fileparser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/samcunliffe/bcmptidy/datamodel"
)

func splitOnHyphen(s string) (string, string) {
	ss := strings.SplitN(s, " - ", 2)
	return ss[0], ss[1]
}

func removeParenthesis(s string) string {
	re := regexp.MustCompile(`\s*\(.*?\)`)
	return strings.TrimSpace(re.ReplaceAllString(s, ""))
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

func ParseMusicFileName(name string) (map[string]string, error) {
	return nil, nil
}
