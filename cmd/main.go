package main

import (
	"github.com/labstack/echo/v4"
	"github.com/NinaDenisova/godis/pkg/service"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "6379"
	CONN_TYPE = "tcp"
)

func main() {
	e := echo.New()

	handlers := godis.NewGodis()

	e.GET("/ping", godis.HandlePing)
	e.GET("/echo", godis.HandleEcho)
	e.GET("/get", godis.HandleGet)
	e.GET("/set", godis.HandleSet)

}
