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
		// Yahan badRequestResponse use karein
		app.badRequestResponse(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		// Yahan internalServerError use karein
		app.internalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		// Yahan internalServerError use karein
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		// Yahan internalServerError (ya badRequest bhi kar sakte ho)
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	post, err := app.store.Posts.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			// Yahan notFoundResponse use karein
			app.notFoundResponse(w, r, err)
		default:
			// Yahan internalServerError use karein
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		// Yahan internalServerError use karein
		app.internalServerError(w, r, err)
		return
	}
}

