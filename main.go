package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/heath-check", basicHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server failed to run", err)
	}

	fmt.Printf("Server is running on port %s", server.Addr)
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request received from %s\n", r.RemoteAddr)
	status := http.StatusOK
	res := fmt.Sprintf("Health Check, Status %d\n", status)
	w.Write([]byte(res))
}
