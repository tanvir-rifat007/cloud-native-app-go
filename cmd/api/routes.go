package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() chi.Router{
	mux:= chi.NewRouter()


	 mux.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))

	// Health check route

	mux.Get("/api/v1/health", app.healthCheck)

	// Newsletter signup route
	mux.Post("/api/v1/newsletter", app.NewsletterSignup)

	return mux
}