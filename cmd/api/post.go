package main

import (
	"errors"
	"net/http"
	"social-post/internal/store"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
    Title   string   `json:"title"`
    Content string   `json:"content"`
    Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	// Decode JSON
	if err := readJSON(w, r, &payload); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	
	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  1, // temporary
		Tags: payload.Tags,
	}

	ctx := r.Context()

	// Store
	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create post")
		return
	}

	//  Return created resource
	if err := writeJSON(w, http.StatusCreated, envelope{"post": post}); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to encode response")
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
	ctx := r.Context()

	post, err := app.store.Posts.GetByID(ctx, id)

	if err != nil{
		switch {
		case errors.Is(err, store.ErrNotFound):
			writeError(w, http.StatusNotFound, err.Error())
			default :
			writeError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err = writeJSON(w,http.StatusOK, post); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}
