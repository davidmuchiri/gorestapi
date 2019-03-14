package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//Port is the default port for now
const Port = ":8080"

// New Returns a new serverwith
func New(mux *mux.Router) *http.Server {
	srv := &http.Server{
		Addr:         Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	return srv

}
