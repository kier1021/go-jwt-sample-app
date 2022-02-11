package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	router     *gin.Engine
	httpServer *http.Server
}

func NewAPIServer() *APIServer {
	router := NewRouter()
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &APIServer{
		router:     router,
		httpServer: httpServer,
	}
}

func (server *APIServer) Run() error {
	if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (server *APIServer) Shutdown(ctx context.Context) error {
	if err := server.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
