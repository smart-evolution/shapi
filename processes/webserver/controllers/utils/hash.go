package utils

import (
	"crypto/sha1"
	"fmt"
)

// HashString - transform string into hash
func HashString(input string) string {
	val := []byte(input)
	h := sha1.New()
	h.Write(val)

	return fmt.Sprintf("%x", h.Sum(nil))
}
