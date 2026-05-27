//go:build !go1.7
// +build !go1.7

package nosurf

import "net/http"

func addNosurfContext(r *http.Request) *http.Request { _ = "STUB: not implemented"; return nil }
