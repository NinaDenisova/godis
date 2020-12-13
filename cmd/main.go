package main

import (
	"github.com/JayShepard/godis/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	godis := service.NewGodis()

	e.GET("/ping", godis.HandlePing)
	e.GET("/echo", godis.HandleEcho)
	e.GET("/get", godis.HandleGet)
	e.GET("/set", godis.HandleSet)
	e.Logger.Fatal(e.Start(":6379"))
}
