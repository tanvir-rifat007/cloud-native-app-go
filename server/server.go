// Package server contains everything for setting up and running the HTTP server.
package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)


type Server struct{
	address string
	mux     chi.Router
	server  *http.Server
	log     *zap.Logger
}

type Options struct{
	Host  string
	Port  int
	Log  *zap.Logger
}

func New(opts Options) *Server{

	if opts.Log == nil {
		opts.Log = zap.NewNop()
	}


	addr:= net.JoinHostPort(opts.Host, strconv.Itoa(opts.Port))
	mux:= chi.NewMux()

	return &Server{
		address: addr,
		mux    : mux,
		log    : opts.Log,
		
		server : &http.Server{
			Addr: addr,
			Handler: mux,
			ReadTimeout: 5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout: 5 * time.Second,


		},
	}
	
}

func (s *Server) Start()error{

// start setting up routes	
	s.setupRoutes()

	s.log.Info("Starting server on", zap.String("address", s.address))


	if err:= s.server.ListenAndServe();err!=nil && !errors.Is(err,http.ErrServerClosed){
		s.log.Error("could not start server", zap.Error(err))
		return fmt.Errorf("could not start server: %w", err)
	}
	return nil
}

func (s *Server) Stop()error{
	s.log.Info("Stopping server on", zap.String("address", s.address))

	ctx,cancel:= context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	if err:= s.server.Shutdown(ctx);err!=nil{
		s.log.Error("could not stop server", zap.Error(err))
		return fmt.Errorf("could not stop server: %w", err)
	}
	return nil
}
