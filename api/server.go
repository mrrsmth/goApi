package api

import (
	"context"
	"fmt"
	"go_api_template/pkg/utils"
	"net/http"
	"time"
)

// LogMiddleware logs the request method, URL, and duration of the request.
// func LogMiddleware(next http.Handler, logger *utils.Logger) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		next.ServeHTTP(w, r)
// 		logger.Infof("%s %s %v", r.Method, r.URL.Path, time.Since(start))
// 	})
// }

// NewServer sets up the server with handlers and middleware.
func NewServer(logger utils.ILogger, config utils.IConfig, db utils.PgPool) http.Handler {
	mux := http.NewServeMux()
	AddRoutes(mux, logger, config, db)
	var handler http.Handler = mux
	// Add middleware as needed

	// handler = LogMiddleware(handler, logger) // Example middleware
	return handler
}

// run contains the main logic of the server.
func Run(ctx context.Context, logger utils.ILogger, config *utils.Config, db utils.PgPool) error {
	if db == nil {
		return fmt.Errorf("database connection is not established")
	}
	defer db.Close()
	srv := NewServer(logger, config, db)

	httpServer := &http.Server{
		Addr:    ":" + config.Port,
		Handler: srv,
	}

	go func() {
		logger.Successf("Server is running on port " + config.Port)
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			logger.Errorf("Server failed: %v\n", err)
		}
	}()

	<-ctx.Done()
	logger.Warn("Server is shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return httpServer.Shutdown(shutdownCtx)
}
