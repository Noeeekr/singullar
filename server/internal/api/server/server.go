package server

import (
	"log"
	"net/http"
)

type Server struct {
	*http.Server
}

func New(router http.Handler, addr string) *Server {
	return &Server{
		Server: &http.Server{
			Handler: router,
			Addr:    addr,
		},
	}
}

func (s *Server) RegisterErrorLogger(logger *log.Logger) *Server {
	s.ErrorLog = logger
	return s
}
