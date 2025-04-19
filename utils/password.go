package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"

	"golang.org/x/crypto/argon2"
)

// GenerateSalt creates a random 16-byte salt and encodes it in base64.
func GenerateSalt() (string, error) {
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    if err != nil {
        return "", err
    }
    return base64.RawStdEncoding.EncodeToString(salt), nil
}

// HashPassword hashes a password using Argon2id with the provided salt.
func HashPassword(password string) string {
	var salt string = os.Getenv("PASSWORD_SALT")
    hash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
    return base64.RawStdEncoding.EncodeToString(hash)
}

// CheckPassword verifies if the provided password matches the hashed password using the same salt.
func CheckPassword(password, hashedPassword string) (bool) {
    computedHash := HashPassword(password)
    return computedHash == hashedPassword
}