package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Hash encrypt password using bcrypt.
// Remember that each call will generate different hash
// even if the password is the same.
func Hash(password string) (string, error) {
	if password == "" {
		return "", errors.New("password must not empty")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// IsMatch returns true if password equals to hash, false otherwise.
// Input password is commonly from user's input,
// while the hash should be from database record.
func IsMatch(password, hash string) bool {
	if password == "" || hash == "" {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
