package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/stenstromen/go-snapnote-backend/controller"
)

func main() {
	router := mux.NewRouter()

	router.Use(authMiddleware)

	router.HandleFunc("/post", controller.CreateFormData).Methods("POST")
	router.HandleFunc("/get/{noteid}", controller.GetFormData).Methods("GET")

	corsOptions := cors.Options{
		AllowedOrigins: []string{"http://example.com", "http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	corsHandler := cors.New(corsOptions).Handler(router)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != os.Getenv("AUTHORIZATION") {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
