//go:build !go1.7
// +build !go1.7

package nosurf

import (
	"net/http"
	"sync"
)

// This file implements a context similar to one found
// in gorilla/context, but tailored specifically for our use case
// and not using gorilla's package just because.

type csrfContext struct {
	// The masked, base64 encoded token
	// That's suitable for use in form fields, etc.
	token string
	// reason for the failure of CSRF check
	reason error
}

var (
	contextMap = make(map[*http.Request]*csrfContext)
	cmMutex    = new(sync.RWMutex)
)

// Token() takes an HTTP request and returns
// the CSRF token for that request
// or an empty string if the token does not exist.
//
// Note that the token won't be available after
// CSRFHandler finishes
// (that is, in another handler that wraps it,
// or after the request has been served)
func Token(req *http.Request) string { _ = "STUB: not implemented"; return "" }

// Reason() takes an HTTP request and returns
// the reason of failure of the CSRF check for that request
//
// Note that the same availability restrictions apply for Reason() as for Token().
func Reason(req *http.Request) error { _ = "STUB: not implemented"; return nil }

// Takes a raw token, masks it with a per-request key,
// encodes in base64 and makes it available to the wrapped handler
func ctxSetToken(req *http.Request, token []byte) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

func ctxSetReason(req *http.Request, reason error) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

func ctxClear(req *http.Request) { _ = "STUB: not implemented"; return }
