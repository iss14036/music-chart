package http

import (
	"fmt"
	"github.com/iss14036/music-chart/configs"
	"github.com/iss14036/music-chart/internal/app/http/controller"
	"github.com/iss14036/music-chart/internal/app/http/routes"

	"github.com/labstack/echo/v4"
	"github.com/tylerb/graceful"
)

type Server struct {
	server *graceful.Server
	echo   *echo.Echo
	cfg    *configs.Config
	ctrl   *controller.Schema
}

type ServerItf interface {
	ListenAndServe() error
}

func (s *Server) ListenAndServe() error {
	port := fmt.Sprintf(`:%v`, s.cfg.AppPort)
	s.echo.Logger.Fatal(s.echo.Start(port))
	return nil
}

func (s *Server) connectCoreWithEcho() {
	e := s.echo
	c := s.ctrl

	routes.Register(e, c, s.cfg)
}

func New(
	cfg *configs.Config,
	ctrl *controller.Schema,
) ServerItf {
	server := &Server{
		echo: echo.New(),
		cfg:  cfg,
		ctrl: ctrl,
	}

	server.connectCoreWithEcho()
	return server
}
