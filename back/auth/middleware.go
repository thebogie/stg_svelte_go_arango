package auth

import (
	"back/graph/model"
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"strings"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	ID string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			jwtheader := JwtHeader{}
			ctx := context.WithValue(r.Context(), "jwtheader", &jwtheader)
			r = r.WithContext(ctx)

			// Check if the authorization header is present and valid.
			auth := r.Header.Get("Authorization")
			if auth == "" {

				// Peek at the request body without consuming it entirely
				peekBody, err := io.ReadAll(io.LimitReader(r.Body, 1024))
				if err != nil {
					return
				}

				// Restore the original request body after peeking
				r.Body = io.NopCloser(io.MultiReader(bytes.NewReader(peekBody), r.Body))

				// Check if this a login request. send through
				if strings.Contains(string(peekBody), "LoginUser") {

					next.ServeHTTP(w, r)
					return
				} else {
					//PASSTHROUGH
					//next.ServeHTTP(w, r)
					//return
					http.Error(w, "Login or Authorization header is required", http.StatusUnauthorized)
					return
				}

			}

			//check auth header to ensure allowed to make calls
			email := jwtheader.CheckToken(auth)

			//TODO: better check?
			//TODO: check for expired too
			if email == "INVALID" {
				log.Printf("Incorrect JWT token.")
				http.Error(w, "Wrong user or not valid", http.StatusBadRequest)
				return
			}
			ctx = context.WithValue(r.Context(), "authuser", email)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
