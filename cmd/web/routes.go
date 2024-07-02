package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// handles and recovers from any Panic and responds with the error code
	mux.Use(middleware.Recoverer)

	// prevent any command or request taking more than 60 seconds to time out
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Get("/", app.ShowHome)

	return mux
}
