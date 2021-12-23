package rkeygen_test

import (
	"testing"

	"github.com/dhenkes/rkeygen"
)

func TestNewRandomString(t *testing.T) {
	t.Run("InvalidInteger", func(t *testing.T) {
		if _, err := rkeygen.NewRandomString(-1); err == nil {
			t.Fatal(err)
		} else if err != rkeygen.ErrInvalidLength {
			t.Fatal("Error should be ErrInvalidLength.")
		}
	})

	t.Run("ValidString", func(t *testing.T) {
		if s, err := rkeygen.NewRandomString(32); err != nil {
			t.Fatal(err)
		} else if len(s) != 32 {
			t.Fatal("String should be 32 characters.")
		}
	})
}
