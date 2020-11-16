package internal

import (
	"errors"
	"strings"
)

type Godis struct {
	db map[string]string
}

func NewGodis() *Godis {
	return &Godis{make(map[string]string)}
}

func (g *Godis) Request(command string, message ...string) (string, error) {
	cleanCommand := strings.ToLower(command)
	if cleanCommand == "ping" {
		if len(message) == 0 {
			return "PONG", nil
		}
		return message[0], nil
	}
	if cleanCommand == "echo" {
		if len(message) == 1 {
			return message[0], nil
		}
		return "", errors.New("ERR wrong number of arguments for 'echo' command")
	}
	if cleanCommand == "set" {
		if len(message) < 2 {
			return "", errors.New("ERR wrong number of arguments for 'set' command")
		}
		g.db[message[0]] = message[1]
		return "OK", nil
	}
	if cleanCommand == "get" {
		return g.db[message[0]], nil
	}
	return "", nil
}
