package main

import (
	"log"
	"net/http"
)

// internalServerError handles 500 errors. 
// Ye internal errors ko log karega taaki developer debug kar sake,
// lekin user ko ek generic message bhejega security ke liye.
func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// Log the error with method and path for debugging
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	// Send generic message to user
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

// badRequestResponse handles 400 errors.
// Isme hum user ko bata sakte hain ki usne kya galat bheja hai.
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

// notFoundResponse handles 404 errors.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	
	writeJSONError(w, http.StatusNotFound, "not found")
}
