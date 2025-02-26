package main

import (
	"net/http"

	"github.com/willystw/golang-simple-social/internal/store"
)

type CreateCommentPayload struct {
	Content string `json:"content" validate:"required,max=250"`
}

// CreateComment godoc
//
//	@Summary		Creates a comment
//	@Description	Creates a comment by Post ID
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int						true	"Post ID"
//	@Param			payload	body		CreateCommentPayload	true	"Comment payload"
//	@Success		200		{object}	store.Comment			"Comment created"
//	@Success		400		{object}	error					"Post not found"
//	@Success		404		{object}	error					"Post not found"
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/comments/{postID} [post]
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
