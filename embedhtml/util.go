package embedhtml

import (
	"net/url"
	"path/filepath"
	"strings"
	"unicode"
)

// from https://github.com/golang/tools/blob/master/internal/span/uri.go
func pathToURI(path string) string {
	if path == "" {
		return ""
	}
	if !isWindowsDrivePath(path) {
		if abs, err := filepath.Abs(path); err == nil {
			path = abs
		}
	}
	if isWindowsDrivePath(path) {
		path = "/" + strings.ToUpper(string(path[0])) + path[1:]
	}
	path = filepath.ToSlash(path)
	u := url.URL{
		Scheme: "file",
		Path:   path,
	}
	return u.String()
}
func isWindowsDrivePath(path string) bool {
	if len(path) < 3 {
		return false
	}
	return unicode.IsLetter(rune(path[0])) && path[1] == ':'
}
