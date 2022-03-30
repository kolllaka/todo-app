package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

const (
	timeout = 10 * time.Second
)

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    timeout,
		WriteTimeout:   timeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
