package utils

import (
    "testing"
    "regexp"
)

func TestCreateSessionID(t *testing.T) {
    t.Run("Hash should be the same for the same inputs", func(t *testing.T) {
        expected := CreateSessionID("user", "pass", "time")
        hash := CreateSessionID("user", "pass", "time")

        if expected != hash {
            t.Errorf("Hash differs for the same for the same inputs")
        }
    })

    t.Run("Hash should contain only hex characters", func(t *testing.T) {
        hash := CreateSessionID("user", "pass", "time")
        match, err := regexp.MatchString("^[a-fA-F0-9]+$", hash)

        if !match || err != nil {
            t.Errorf("Hash should be the same for the same inputs")
        }
    })

    t.Run("Hash should be 20 characters length", func(t *testing.T) {
        hash := CreateSessionID("user", "pass", "time")

        if len([]rune(hash)) != 40 {
            t.Errorf("Hash is not 40 characters long")
        }
    })
}
