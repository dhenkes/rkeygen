# rkeygen

This package provides an easy way to generate random keys that can be used
as passwords or session tokens.

It uses `crypto/rand` which is a cryptographically secure random number
generator. This unfortunately means that string generation may take longer than
with pseudo-random algorithms.

## Usage

```go
package main

import (
  "log"

  "github.com/dhenkes/rkeygen"
)

func main() {
  // NewRandomString returns a random string with the given length.
  s, err := rkeygen.NewRandomString(32)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("Random string: %v", s)
}
```