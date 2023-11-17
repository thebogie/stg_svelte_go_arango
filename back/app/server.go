package main

import (
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"back/app/db"
	"back/auth"
	"back/graph/resolver"
	"back/pkg/adapter/controller"
	"back/pkg/registry"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

const defaultPort = "50002"

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	initdb := db.InitDB()

	router := chi.NewRouter()

	router.Use(Logger)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://smacktalkgaming.com", "http://localhost:50003", "http://192.168.86.45:50003", "http://192.168.86.101:50003"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link", "Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(cors.Handler)
	router.Use(auth.Middleware())
	router.Use(middleware.Recoverer)

	ctrl := newController(initdb)
	srv := handler.NewDefaultServer(resolver.NewSchema(ctrl))

	router.Handle("/graphql/playground", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/graphql/playground for GraphQL playground", port)
	//log.Fatal(http.ListenAndServe("192.168.86.45:"+port, router))
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func newController(db driver.Database) controller.Controller {
	r := registry.New(db)
	return r.NewController()
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("**Request:")
		fmt.Println(formatRequest(r))

		wrapped := wrapResponseWriter(w)
		fmt.Println("**END OF REQUEST**")
		next.ServeHTTP(wrapped, r)

		fmt.Println("**Response:")
		fmt.Println("Status:", wrapped.Status())
		fmt.Println("Header:", wrapped.Header())
		fmt.Println("Body:", wrapped.Body.String())
		fmt.Println("**END OF RESPONSE**")
	})
}

type responseWriterWrapper struct {
	http.ResponseWriter
	StatusNum int
	Body      strings.Builder
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriterWrapper {
	return &responseWriterWrapper{
		ResponseWriter: w,
		StatusNum:      http.StatusOK,
	}
}

func (rw *responseWriterWrapper) Write(b []byte) (int, error) {
	rw.Body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriterWrapper) WriteHeader(statusCode int) {
	rw.StatusNum = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriterWrapper) Status() int {
	return rw.StatusNum
}

func formatRequest(r *http.Request) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Method: %s\n", r.Method))
	sb.WriteString(fmt.Sprintf("URL: %s\n", r.URL.String()))
	sb.WriteString("Headers:\n")
	r.Header.Write(&sb)
	sb.WriteString("\n")

	// Copy the original body to a new buffer
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		sb.WriteString(fmt.Sprintf("Failed to read request body: %s\n", err.Error()))
	} else {
		sb.WriteString(fmt.Sprintf("Body: %s\n", buf.String()))
	}

	// Restore the original body
	r.Body = io.NopCloser(strings.NewReader(buf.String()))

	return sb.String()
}
