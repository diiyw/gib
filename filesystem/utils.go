package filesystem

import (
	"net/url"
	"strings"
)

// GetFilename get filename from a url string
func GetFilename(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	u.Path = strings.Replace(u.Path, "/", "_", -1)

	return strings.Trim(u.Path, "_")
}
