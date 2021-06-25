package app

import (
	"github.com/iss14036/music-chart/configs"
	"github.com/iss14036/music-chart/internal/app/external"
	"github.com/iss14036/music-chart/internal/app/http"
	"github.com/iss14036/music-chart/internal/app/http/controller"
	"github.com/iss14036/music-chart/internal/app/repository"
	"github.com/iss14036/music-chart/internal/app/service"
	"github.com/iss14036/music-chart/internal/pkg/gateway"
)

func InitApp(config *configs.Config) http.ServerItf {
	databaseSchema := external.ProviderOrm(config)
	auth := gateway.NewAuth(config)
	userRepo := repository.NewUser(databaseSchema)
	userService := service.NewUser(userRepo, auth)
	UserCtrl := controller.NewUser(userService)
	musicRepo := repository.NewMusic(databaseSchema)
	musicService := service.NewMusic(musicRepo)
	musicCtrl := controller.NewMusic(musicService)
	favRepo := repository.NewFavorite(databaseSchema)
	favService := service.NewFavorite(favRepo, musicRepo)
	favCtrl := controller.NewFavorite(favService)
	controllerSchema := controller.New(UserCtrl, musicCtrl, favCtrl)
	serveItf := http.New(config, controllerSchema)
	return serveItf
}
