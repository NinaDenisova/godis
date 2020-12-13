package service

import (
	"errors"

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

func (g *Godis) HandlePing(c echo.Context) (string, error) {
	cmd := new(Cmd)
	if err := c.Bind(cmd); err != nil {
		return "", err
	}

	if len(cmd.Messages) == 0 {
		return "PONG", nil
	}
	return cmd.Messages[0], nil
}

func (g *Godis) HandleEcho(c echo.Context) (string, error) {

	cmd := new(Cmd)
	if err := c.Bind(cmd); err != nil {
		return "", err
	}
	if len(cmd.Messages) == 1 {
		return cmd.Messages[0], nil
	}
	return "", errors.New("ERR wrong number of arguments for 'echo' command")
}

func (g *Godis) HandleSet(c echo.Context) (string, error) {

	cmd := new(Cmd)
	if err := c.Bind(cmd); err != nil {
		return "", err
	}
	if len(cmd.Messages) < 2 {
		return "", errors.New("ERR wrong number of arguments for 'set' command")
	}
	g.db[cmd.Messages[0]] = cmd.Messages[1]
	return "OK", nil
}

func (g *Godis) HandleGet(c echo.Context) (string, error) {

	cmd := new(Cmd)
	if err := c.Bind(cmd); err != nil {
		return "", err
	}
	return g.db[cmd.Messages[0]], nil
}
