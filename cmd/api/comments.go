package main

import (
	"net/http"

	"github.com/willystw/golang-simple-social/internal/store"
)

type CreateCommentPayload struct {
	Content string `json:"content" validate:"required,max=250"`
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload CreateCommentPayload
	if err := readJson(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	comment := &store.Comment{
		PostID:  post.ID,
		UserID:  1,
		Content: payload.Content,
	}

	ctx := r.Context()
	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := writeJson(w, http.StatusOK, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
