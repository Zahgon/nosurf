package nosurf

import (
	"net/http"
)

// Checks if the given request is exempt from CSRF checks.
// It checks the ExemptFunc first, then the exact paths,
// then the globs and finally the regexps.
func (h *CSRFHandler) IsExempt(r *http.Request) bool { _ = "STUB: not implemented"; return false }

// then the globs

// finally, the regexps

// Exempts an exact path from CSRF checks
// With this (and other Exempt* methods)
// you should take note that Go's paths
// include a leading slash.
func (h *CSRFHandler) ExemptPath(path string) { _ = "STUB: not implemented"; return }

// A variadic argument version of ExemptPath()
func (h *CSRFHandler) ExemptPaths(paths ...string) { _ = "STUB: not implemented"; return }

// Exempts URLs that match the specified glob pattern
// (as used by filepath.Match()) from CSRF checks
//
// Note that ExemptGlob() is unable to detect syntax errors,
// because it doesn't have a path to check it against
// and filepath.Match() doesn't report an error
// if the path is empty.
// If we find a way to check the syntax, ExemptGlob
// MIGHT PANIC on a syntax error in the future.
// ALWAYS check your globs for syntax errors.
func (h *CSRFHandler) ExemptGlob(pattern string) { _ = "STUB: not implemented"; return }

// A variadic argument version of ExemptGlob()
func (h *CSRFHandler) ExemptGlobs(patterns ...string) { _ = "STUB: not implemented"; return }

// Accepts a regular expression string or a compiled *regexp.Regexp
// and exempts URLs that match it from CSRF checks.
//
// If the given argument is neither of the accepted values,
// or the given string fails to compile, ExemptRegexp() panics.
func (h *CSRFHandler) ExemptRegexp(re interface{}) { _ = "STUB: not implemented"; return }

// A variadic argument version of ExemptRegexp()
func (h *CSRFHandler) ExemptRegexps(res ...interface{}) { _ = "STUB: not implemented"; return }

func (h *CSRFHandler) ExemptFunc(fn func(r *http.Request) bool) { _ = "STUB: not implemented"; return }
