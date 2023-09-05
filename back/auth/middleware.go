package auth

import (
	"context"
	"log"
	"net/http"

	"back/graph/model"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	ID string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			cookieaccess := CookieAccess{
				HttpReader: r,
				Writer:     w,
			}

			ctx := context.WithValue(r.Context(), "cookiemaker", &cookieaccess)
			ctx = context.WithValue(ctx, "ResponseWriter", w)
			ctx = context.WithValue(ctx, "Request", r)

			// and call the next with our new context
			r = r.WithContext(ctx)

			c, err := cookieaccess.GetAuthCookie()

			//cookie doesnt exist send through
			if err != nil || c == nil {

				next.ServeHTTP(w, r)
				return
			}

			//cookie exists. is it the correct cookie...
			email := cookieaccess.CheckAuthCookieForUserid()

			//TODO: better check?
			//TODO: check for expired too
			if email == "INVALID" {
				log.Printf("Incorrect JWT token.")
				http.Error(w, "Wrong user", http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
