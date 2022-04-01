package server

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
	config "payment-service/configs"
)

type Server struct {
	server *http.Server
}

type ServerGRPC struct {
	server *grpc.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
