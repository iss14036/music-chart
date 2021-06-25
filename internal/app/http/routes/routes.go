package routes

import (
	"github.com/iss14036/music-chart/configs"
	"github.com/iss14036/music-chart/internal/app/http/controller"
	"github.com/iss14036/music-chart/internal/pkg/gateway"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, ctrl *controller.Schema, cfg *configs.Config) {
	e.GET("/ping", ctrl.HealthCheck)
	auth := gateway.NewAuth(cfg)
	middleware := e.Group("", auth.Middleware)

	e.POST("/login", ctrl.Login)
	e.POST("/user", ctrl.InsertUser)
	middleware.GET("/user/:id", ctrl.GetUser)
	middleware.GET("/music", ctrl.FetchMusic)
	middleware.GET("/music/:id", ctrl.DetailMusic)
	middleware.POST("/favorite-music", ctrl.SetFavorite)
	middleware.GET("/favorite-music", ctrl.FetchFavorite)
}
