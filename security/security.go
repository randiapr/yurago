package security

import "golang.org/x/crypto/bcrypt"

// Hashing password with cost (min = 4, max = 31, default = 10)
func Hashing(password string, cost int) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes)
}

// IsStrHashValid Check Hasing
func IsStrHashValid(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
