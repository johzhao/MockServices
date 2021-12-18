package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mock.services/common/config"
	"mock.services/common/router"
	middleware2 "mock.services/common/server/middleware"
	"net/http"
)

func NewServer(logger *zap.Logger) Server {
	return Server{
		logger: logger,
	}
}

type Server struct {
	logger *zap.Logger
	engine *gin.Engine
	srv    *http.Server
	config config.Server
}

func (s *Server) SetupServer(config config.Server) error {
	engine := gin.New()
	engine.Use(
		cors.New(s.corsConfig()),
		middleware2.Log(s.logger),
		middleware2.Recovery(s.logger),
	)
	engine.NoRoute(middleware2.NoRouteHandler(s.logger))

	s.engine = engine
	s.config = config

	return nil
}

func (s *Server) corsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true

	return corsConfig
}

func (s *Server) RunServer() error {
	srv := &http.Server{
		Addr:    s.config.Address,
		Handler: s.engine,
	}
	s.srv = srv
	return s.srv.ListenAndServe()
}

//goland:noinspection GoUnusedParameter
func (s *Server) StopServer(err error) {
	s.logger.Info("stop server")
	_ = s.srv.Shutdown(context.Background())
	s.srv = nil
}

func (s *Server) HandleRouter(router router.Router) {
	s.engine.Handle(router.Method(), router.Path(), router.HandleRequest)
}
