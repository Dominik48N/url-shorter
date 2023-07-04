package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func CheckPassword(password, hashedPassword string) bool {
	hash := sha256.Sum256([]byte(password))
	hashed := hex.EncodeToString(hash[:])
	return hashed == hashedPassword
}
