package utils

import (
	"regexp"
	"testing"
)

func TestHashString(t *testing.T) {
	t.Run("Hash should be the same for the same inputs", func(t *testing.T) {
		expected := HashString("somestring")
		hash := HashString("somestring")

		if expected != hash {
			t.Errorf("Hash differs for the same for the same inputs")
		}
	})

	t.Run("Hash should contain only hex characters", func(t *testing.T) {
		hash := HashString("somestring")
		matched, err := regexp.MatchString("^[a-f0-9]+$", hash)

		if !matched || err != nil {
			t.Errorf("Hash doesn't contain only hex characters")
		}
	})

	t.Run("Hash should be of proper characters length", func(t *testing.T) {
		hash := HashString("somestring")

		if len(hash) != 40 {
			t.Errorf("Hash is not 40 characters long")
		}
	})
}
