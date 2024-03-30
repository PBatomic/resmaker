package server

import (
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	responseDelay := r.URL.Query().Get("delay")
	responseStatusCodes := r.URL.Query().Get("statusCode")
}
