package auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("ThisIsASecretKeyThatOnlyIKnowAndWontTellAnybody")
)

type JwtHeader struct {
	Token string
}

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func (c *JwtHeader) GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["email"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

func (c *JwtHeader) CheckToken(checktoken string) string {
	var email = "INVALID"

	token, _ := jwt.Parse(checktoken, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email = claims["email"].(string)
	}

	return email
}

func (c *JwtHeader) HashPassword(password string) (string, error) {
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

func (c *JwtHeader) Authenticate(input, found string) bool {
	return CheckPasswordHash(input, found)
}
