package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Cmd struct {
	Messages []string `json:"messages" xml:"messages" form:"messages" query:"messages"`
}
type Godis struct {
	db map[string]string
}

func NewGodis() *Godis {
	return &Godis{make(map[string]string)}
}

func (g *Godis) HandlePing(c echo.Context) error {
	cmd := new(Cmd)
	if err := c.Bind(cmd); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if len(cmd.Messages) == 0 {
		return c.JSON(http.StatusOK, "PONG")
	}
	return c.JSON(http.StatusOK, cmd.Messages[0])
}

func (g *Godis) HandleEcho(c echo.Context) error {

	cmd := new(Cmd)
	if err := c.Bind(cmd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Problem serializing params with body: %v", err))
	}
	if len(cmd.Messages) == 1 {
		return c.JSON(http.StatusOK, cmd.Messages[0])
	}
	return echo.NewHTTPError(http.StatusBadRequest, "ERR wrong number of arguments for 'echo' command")
}

func (g *Godis) HandleSet(c echo.Context) error {

	cmd := new(Cmd)
	if err := c.Bind(cmd); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if len(cmd.Messages) < 2 {
		return echo.NewHTTPError(http.StatusBadRequest, "ERR wrong number of arguments for 'set' command")
	}
	g.db[cmd.Messages[0]] = cmd.Messages[1]
	return c.JSON(http.StatusOK, "OK")
}

func (g *Godis) HandleGet(c echo.Context) error {

	cmd := new(Cmd)
	if err := c.Bind(cmd); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, g.db[cmd.Messages[0]])
}
