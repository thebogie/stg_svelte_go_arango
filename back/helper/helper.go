package helper

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	//fish, _ := HashPassword(password)
	//log.Printf(fish)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Authenticate(input, found string) bool {
	return CheckPasswordHash(input, found)
}
