package main

import (
	"net/http"
)

// HealthCheck godoc
//
//	@Summary		Check application health
//	@Description	Check application health
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		500	{object}	error
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := writeJson(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
