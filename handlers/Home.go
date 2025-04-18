package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Home(mux chi.Router,dir string){
    mux.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir(dir))))
}