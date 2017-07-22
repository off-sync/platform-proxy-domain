// Package frontends contains the domain models for the Frontend aggregate root.
package frontends

import "errors"

// Errors
var (
	ErrFrontendURLMissing = errors.New("frontend URL missing")
	ErrServiceNameMissing = errors.New("service name missing")
)
