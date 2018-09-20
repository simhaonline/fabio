package route

import (
	"strings"
)

// matcher determines whether a host/path matches a route
type matcher func(uri string, r *Route) bool

// Matcher contains the available matcher functions.
// Update config/load.go#load after updating.
var Matcher = map[string]matcher{
	"prefix": prefixMatcher,
	"glob":   globMatcher,
	"nocase": noCaseMatcher,
}

// prefixMatcher matches path to the routes' path.
func prefixMatcher(uri string, r *Route) bool {
	return strings.HasPrefix(uri, r.Path)
}

// globMatcher matches path to the routes' path using gobwas/glob.
func globMatcher(uri string, r *Route) bool {
	return r.Glob.Match(uri)
}

// noCase matches path to the routes' path ignoring case
func noCaseMatcher(uri string, r *Route) bool {
	lowerURI := strings.ToLower(uri)
	lowerPath := strings.ToLower(r.Path)
	return strings.HasPrefix(lowerURI, lowerPath)
}
