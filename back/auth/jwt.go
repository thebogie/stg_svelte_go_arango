package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("ThisIsASecretKeyThatOnlyIKnowAndWontTellAnybody")
)

var cookieName = "auth-cookie"

type CookieAccess struct {
	HttpReader *http.Request
	Writer     http.ResponseWriter
	Id         string
}

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func (c *CookieAccess) GenerateToken(username string) (string, error) {
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

func (c *CookieAccess) GetAuthCookie() (*http.Cookie, error) {
	return c.HttpReader.Cookie(cookieName)
}

func (c *CookieAccess) CheckAuthCookieForUserid() string {
	var email = "INVALID"
	autocookie, _ := c.HttpReader.Cookie(cookieName)
	log.Printf(autocookie.Name)

	token, _ := jwt.Parse(autocookie.Value, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email = claims["email"].(string)
	}

	return email
}

func (c *CookieAccess) GenerateAuthCookie(tokenString string) {

	//passed save cookie
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	// Set the cookie in the response
	http.SetCookie(c.Writer, cookie)

}
