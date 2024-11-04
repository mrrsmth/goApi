package api

import (
	"go_api_template/domain"
	"go_api_template/pkg/utils"
	"net/http"
)

// AddRoutes sets up the routes for the server.
func AddRoutes(mux *http.ServeMux, logger utils.ILogger, config utils.IConfig, db utils.PgPool) {
	// Set up domains and handlers

	// User domain
	userHandler := domain.UserDomain(logger, config, db)
	mux.HandleFunc("GET /api/users/{id}", userHandler.GetUserByID(logger))

	// Basic handlers
	mux.HandleFunc("/", NotFound(logger))

	logger.Success("Routes added.")
}

// 404 is a handler for 404 errors.
func NotFound(logger utils.ILogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		logger.Errorf("404 Not Found: %s", path)
		utils.RespondError(w, http.StatusNotFound, "Not Found")
	}
}
