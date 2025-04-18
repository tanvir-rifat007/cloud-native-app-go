package server

import "canvas/handlers"

func (s *Server) setupRoutes(){
  handlers.Health(s.mux)

  dir := s.opts.PublicDir
	if dir == "" {
		dir = "./public"
	}
  handlers.Home(s.mux,dir)
}