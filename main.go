package main

import (
	"log"
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	log.Printf("Server starting on %s...", api.addr)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
