package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// HashToString returns the sha256 hash of a byte slice in string format
func HashToString(data []byte) string {
	hashbytes := sha256.Sum256(data)
	return hex.EncodeToString(hashbytes[:])
}

// ScriptHash is a custom hash for scripts that cannot be confused with hashes for modules
func ScriptHash(data []byte) string {
	return fmt.Sprintf("script%s", HashToString(data))
}

func ModuleHash(data []byte) string {
	return fmt.Sprintf("module%s", HashToString(data))
}
