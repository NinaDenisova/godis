package main

import (
	"fmt"
	"github.com/jayShepard/godis/pkg"
	"net"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "6379"
	CONN_TYPE = "tcp"
)

func main() {
	e := echo.New()

	handlers := godis.NewGodis()

	e.GET(/ping, godis.HandlePing)
	e.GET(/echo, godis.HandleEcho)
	e.GET(/get, godis.HandleGet)
	e.GET(/set, godis.HandleSet)

}

