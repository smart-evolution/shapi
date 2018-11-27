package utils

import (
    "fmt"
    "crypto/sha1"
)

// Hash - transform string into hash
func HashString(input string) string {
    val := []byte(input)
    h := sha1.New()
    h.Write(val)

    return fmt.Sprintf("%x", h.Sum(nil))
}
