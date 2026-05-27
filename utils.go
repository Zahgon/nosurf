package nosurf

import (
	"net/url"
)

func sContains(slice []string, s string) bool {
	_ = "STUB: not implemented"
	// checks if the given slice contains the given string
	return false
}

// Checks if the given URLs have the same origin
// (that is, they share the host, the port and the scheme)
func sameOrigin(u1, u2 *url.URL) bool {
	_ = "STUB: not implemented"
	// we take pointers, as url.Parse() returns a pointer
	// and http.Request.URL is a pointer as well
	return false
}

// Host is either host or host:port
