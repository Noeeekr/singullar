package server

import (
	"log"
	"net/http"
)

type Server struct {
	s http.Server
}

func New() *Server {
	return &Server{
		s: http.Server{},
	}
}
func (s *Server) WithRouter(router http.Handler) *Server {
	s.s.Handler = router
	return s
}
func (s *Server) WithAddr(addr string) *Server {
	s.s.Addr = addr
	return s
}
func (s *Server) WithErrLogger(logger *log.Logger) *Server {
	s.s.ErrorLog = logger
	return s
}

func (s *Server) ListenAndServe() error {
	return s.s.ListenAndServe()
}
