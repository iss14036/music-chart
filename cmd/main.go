package main

import (
	"fmt"
	"github.com/iss14036/music-chart/configs"
	"github.com/iss14036/music-chart/internal/app"
)

func main() {
	cfg := configs.Reader()
	app := app.InitApp(cfg)
	err := app.ListenAndServe()
	if err != nil {
		fmt.Println('x')
	}
}
