// Package rkeygen provides an easy way to generate random keys that can be
// used as passwords or session tokens.
package rkeygen

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

var (
	// ErrInvalidLength is returned by NewRandomString if the provided length
	// is invalid.
	ErrInvalidLength = errors.New("rkeygen: length must be a positive integer.")
)

// NewRandomString returns a new random string based on a given length
func NewRandomString(n int) (string, error) {
	if n < 0 {
		return "", ErrInvalidLength
	}

	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b)[:n], nil
}
