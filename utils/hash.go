package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

// HashToString returns the sha256 hash of a byte slice in string format
func HashToString(data []byte) string {
	hashbytes := sha256.Sum256(data)
	encoding := base64.URLEncoding.WithPadding(base64.NoPadding)
	return encoding.EncodeToString(hashbytes[:])
}
