package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
	"github.com/stenstromen/go-snapnote-backend/controller"
)

func main() {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("POST /post", controller.CreateFormData)
	mux.HandleFunc("GET /get/{noteid}", controller.GetFormData)

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	corsOptions := cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	// Use the auth middleware
	handler := authMiddleware(mux)

	// Wrap the handler in the CORS handler
	corsHandler := cors.New(corsOptions).Handler(handler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimSpace(r.Header.Get("Authorization"))

		if token != os.Getenv("AUTHORIZATION") {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
