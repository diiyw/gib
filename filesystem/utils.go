package filesystem

import (
	"net/url"
	"os"
	"path/filepath"
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

func DirList(dirPath string) ([]string, error) {
	var dir []string
	err := filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir = append(dir, path)
				return nil
			}
			return nil
		})
	return dir, err
}
