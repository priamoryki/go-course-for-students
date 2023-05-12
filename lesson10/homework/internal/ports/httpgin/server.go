package httpgin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"homework10/internal/app"
)

type Server struct {
	a      app.App
	server *http.Server
}

func NewHTTPServer(port string, a app.App) Server {
	gin.SetMode(gin.ReleaseMode)
	s := Server{a: a}
	s.server = &http.Server{
		Addr:    port,
		Handler: s.Handler(),
	}
	return s
}

func (s *Server) Listen() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Handler() http.Handler {
	a := gin.New()
	api := a.Group("/api/v1")
	api.Use(gin.Logger())
	api.Use(gin.Recovery())
	AppRouter(api, s.a)
	return a
}
