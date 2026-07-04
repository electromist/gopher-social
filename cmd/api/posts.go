package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/electromist/gopher-social.git/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// TODO: Change after auth
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	// 1. URL se parameter nikalna (e.g. /posts/123 me se '123' nikalega)
	idParam := chi.URLParam(r, "postID")

	// 2. String ko int64 me convert karna
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := r.Context()

	// 3. Database method call karna
	post, err := app.store.Posts.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			// Post na milne par 404 Status Not Found bhejo
			writeJSONError(w, http.StatusNotFound, store.ErrNotFound.Error())
		default:
			// Kisi aur problem ke liye 500 error bhejo
			writeJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// 4. Sab sahi hone par HTTP 200 OK aur JSON bhej do
	if err := writeJSON(w, http.StatusOK, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
