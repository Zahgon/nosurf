// Package nosurf implements an HTTP handler that
// mitigates Cross-Site Request Forgery Attacks.
package nosurf

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
)

const (
	// the name of CSRF cookie
	CookieName = "csrf_token"
	// the name of the form field
	FormFieldName = "csrf_token"
	// the name of CSRF header
	HeaderName = "X-CSRF-Token"
	// the HTTP status code for the default failure handler
	FailureCode = 400

	// Max-Age in seconds for the default base cookie. 365 days.
	MaxAge = 365 * 24 * 60 * 60
)

var safeMethods = []string{"GET", "HEAD", "OPTIONS", "TRACE"}

// reasons for CSRF check failures
var (
	ErrNoReferer  = errors.New("A secure request contained no Referer or its value was malformed")
	ErrBadReferer = errors.New("A secure request's Referer comes from a different origin" +
		" from the request's URL")
	ErrBadOrigin = errors.New("Request was made with a disallowed origin specified in the Origin header")
	ErrBadToken  = errors.New("The CSRF token in the cookie doesn't match the one" +
		" received in a form/header.")

	// Internal error. When this is raised, and the request is secure, we additionally check for Referer.
	errNoOrigin = errors.New("Origin header was not present")
)

type CSRFHandler struct {
	// Handlers that CSRFHandler wraps.
	successHandler http.Handler
	failureHandler http.Handler

	// The base cookie that CSRF cookies will be built upon.
	// This should be a better solution of customizing the options
	// than a bunch of methods SetCookieExpiration(), etc.
	baseCookie http.Cookie

	// Slices of paths that are exempt from CSRF checks.
	// All of those will be matched against Request.URL.Path,
	// So they should take the leading slash into account
	// Paths can be specified by...
	// ...an exact path,
	exemptPaths []string
	// ...a regexp,
	exemptRegexps []*regexp.Regexp
	// ...or a glob (as used by path.Match()).
	exemptGlobs []string
	// ...or a custom matcher function
	exemptFunc func(r *http.Request) bool

	isTLS           func(r *http.Request) bool
	isAllowedOrigin func(r *url.URL) bool
}

func defaultFailureHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Extracts the "sent" token from the request
// and returns an unmasked version of it
func extractToken(r *http.Request) []byte {
	_ = "STUB: not implemented"
	// Prefer the header over form value
	return nil
}

// Then POST values

// If all else fails, try a multipart value.
// PostFormValue() will already have called ParseMultipartForm()

// Constructs a new CSRFHandler that calls
// the specified handler if the CSRF check succeeds.
func New(handler http.Handler) *CSRFHandler { _ = "STUB: not implemented"; return nil }

// The same as New(), but has an interface return type.
func NewPure(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (h CSRFHandler) getCookieName() string { _ = "STUB: not implemented"; return "" }

func (h *CSRFHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// If the length of the real token isn't what it should be,
// it has either been tampered with,
// or we're migrating onto a new algorithm for generating tokens,
// or it hasn't ever been set so far.
// In any case of those, we should regenerate it.
//
// As a consequence, CSRF check will fail when comparing the tokens later on,
// so we don't have to fail it just yet.

// short-circuit with a success for safe methods

// Finally, we check the token itself.

// Everything else passed, handle the success.

// handleSuccess simply calls the successHandler.
// Everything else, like setting a token in the context
// is taken care of by h.ServeHTTP()
func (h *CSRFHandler) handleSuccess(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Same applies here: h.ServeHTTP() sets the failure reason, the token,
// and only then calls handleFailure()
func (h *CSRFHandler) handleFailure(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *CSRFHandler) ensureSameOrigin(r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// If no `Sec-Fetch-Site: same-origin` is present, fallback to Origin or Referer,
// including considering custom allowed origins.

// If Origin header was not present, fall back on Referer check for both secure and insecure requests.
// This is opposite of Django's behavior, but should be fine, as neither of the three headers existing is an edge case.
// https://github.com/django/django/blob/8be0c0d6901669661fca578f474cd51cd284d35a/django/middleware/csrf.py#L460

func (h *CSRFHandler) checkReferer(selfOrigin *url.URL, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *CSRFHandler) checkOrigin(selfOrigin *url.URL, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Generates a new token, sets it on the given request and returns it
func (h *CSRFHandler) RegenerateToken(w http.ResponseWriter, r *http.Request) string {
	_ = "STUB: not implemented"
	return ""
}

func (h *CSRFHandler) setTokenCookie(w http.ResponseWriter, r *http.Request, token []byte) {
	_ = "STUB: not implemented"
	// ctxSetToken() does the masking for us
	return
}

// Sets the handler to call in case the CSRF check
// fails. By default it's defaultFailureHandler.
func (h *CSRFHandler) SetFailureHandler(handler http.Handler) { _ = "STUB: not implemented"; return }

// Sets the base cookie to use when building a CSRF token cookie
// This way you can specify the Domain, Path, HttpOnly, Secure, etc.
func (h *CSRFHandler) SetBaseCookie(cookie http.Cookie) { _ = "STUB: not implemented"; return }

// SetIsTLSFunc sets a delegate function which determines, on a per-request basis, whether the request is made over a secure connection.
// This should return `true` iff the URL that the user uses to access the application begins with https://.
// For example, if the Go web application is served via plain-text HTTP,
// but the user is accessing it through HTTPS via a TLS-terminating reverse-proxy, this should return `true`.
//
// Examples:
//
// 1. If you're using the Go TLS stack (no TLS-terminating proxies in between the user and the app), you may use:
//
//	h.SetIsTLSFunc(func(r *http.Request) bool { return r.TLS != nil })
//
// 2. If your application is behind a reverse proxy that terminates TLS, you should configure the reverse proxy
// to report the protocol that the request was made over via an HTTP header,
// e.g. `X-Forwarded-Proto`.
// You should also validate that the request is coming in from an IP of a trusted reverse proxy
// to ensure that this header has not been spoofed by an attacker. For example:
//
//	var trustedProxies = []string{"198.51.100.1", "198.51.100.2"}
//	h.SetIsTLSFunc(func(r *http.Request) bool {
//		ip, _, _ := strings.Cut(r.RemoteAddr, ":")
//		proto := r.Header.Get("X-Forwarded-Proto")
//		return slices.Contains(trustedProxies, ip) && proto == "https"
//	})
func (h *CSRFHandler) SetIsTLSFunc(f func(*http.Request) bool) {
	_ = "STUB: not implemented"

	// SetAllowedOrigins defines a function that checks whether the request comes from an allowed origin.
	// This function will be invoked when the request is not considered a same-origin request.
	// If this function returns `false`, request will be disallowed.
	//
	// In most cases, this will be used with [StaticOrigins].
	return
}

func (h *CSRFHandler) SetIsAllowedOriginFunc(f func(*url.URL) bool) {
	_ = "STUB: not implemented"
	return

	// StaticOrigins returns a delegate, suitable for passing to [CSRFHandler.SetIsAllowedOriginFunc],
	// that validates the request origin against a static list of allowed origins.
	// This function expects each element to be of form `scheme://host`, e.g.: `https://example.com`, `http://example.org`.
	// If any element of the slice is an invalid URL, this function will return an error.
	// If an element includes additional URL parts (e.g. a path), these parts will be ignored,
	// as origin checks only take the scheme and host into account.
	//
	// Example:
	//
	//	h := nosurf.New()
	//	origins, err := nosurf.StaticOrigins("https://api.example.com", "http://insecure.example.com")
	//	if err != nil {
	//		panic(err)
	//	}
	//	h.SetIsAllowedOriginFunc(origins)
}

func StaticOrigins(origins ...string) (func(r *url.URL) bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
