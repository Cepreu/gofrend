package utils

import (
	"crypto/sha256"
	"math/rand"
)

// HashToString returns the sha256 hash of a byte slice in string format
func HashToString(data []byte) string {
	hashbytes := sha256.Sum256(data)
	return encoding.EncodeToString(hashbytes[:])
}

// HashPassword returns the sha256 hash of a password and salt
func HashPassword(password string, salt string) [32]byte {
	return sha256.Sum256([]byte(password + salt))
}

// VerifyPassword returns true iff the sha256 hash of a password and salt matches the given hash
func VerifyPassword(password string, salt string, hash [32]byte) bool {
	return HashPassword(password, salt) == hash
}

// GenerateSalt generates a random salt to be used for hashing
func GenerateSalt() string {
	return randomString(32)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
