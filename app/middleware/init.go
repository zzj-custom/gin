package middleware

import (
	"strings"
)

var (
	noCheckApis map[string]int8
)

func init() {
	noCheckApis = map[string]int8{
		"token":      1,
		"index.html": 1,
	}
}

func ignoreTokenCheck(uri string) bool {
	if strings.Contains(uri, "?") {
		index := strings.Index(uri, "?")
		uri = uri[:index]
	}
	if _, ok := noCheckApis[uri]; ok {
		return true
	}
	return false
}
