package gohttp

import (
	"fmt"
	"regexp"
	"strings"
)

// withoutSlash will remove "/" on start of path
func withoutSlash(path string) string {
	if strings.HasPrefix(path, "/") {
		return strings.TrimPrefix(path, "/")
	}

	return path
}

// withoutTrailslash will remove "/" on end of path
func withoutTrailslash(path string) string {
	if strings.HasSuffix(path, "/") {
		return strings.TrimSuffix(path, "/")
	}

	return path
}

// isAcceptable will check if accept `literal` string matches
// with received acceptable string
//
// Example:
//
//	isAcceptable(`application\/json`, "application/json")	// true
//	isAcceptable(`application\/json`, "text/html")		// false
func isAcceptable(accept, acceptable string) bool {
	if acceptable == "" || acceptable == "*/*" {
		return true
	}

	test := regexp.MustCompile(fmt.Sprintf(`^%s`, accept)).
		FindString(acceptable)

	return test != ""
}
