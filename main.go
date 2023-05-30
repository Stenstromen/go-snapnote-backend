package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/stenstromen/go-snapnote-backend/controller"
)

func main() {
	router := mux.NewRouter()

	// Add auth middleware after CORS middleware
	router.Use(authMiddleware)

	router.HandleFunc("/post", controller.CreateFormData).Methods("POST")
	router.HandleFunc("/get/{noteid}", controller.GetFormData).Methods("GET")

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	corsOptions := cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	corsHandler := cors.New(corsOptions).Handler(router)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimSpace(r.Header.Get("Authorization"))
		envToken := strings.TrimSpace(os.Getenv("AUTHORIZATION"))
		log.Printf("Authorization header: '%s'", token)
		log.Printf("Authorization Env Variable: '%s'", envToken)

		if token != os.Getenv("AUTHORIZATION") {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
