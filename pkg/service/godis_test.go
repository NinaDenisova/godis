package service

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	noMessage  string = ""
	anyMessage string = "Any Message"
	anyKey     string = "Key"
	anyValue   string = "value"
)

func TestGodis(t *testing.T) {
	godis := NewGodis()

	Convey("GIVEN any command", t, func() {
		Convey("command SHOULD be case insensitive", func() {
			testCases := [...]string{"PING", "ping", "PinG"}
			for _, testCase := range testCases {
				response, _ := godis.Request(testCase)
				So(response, ShouldEqual, "PONG")
			}
		})
	})

	Convey("GIVEN Ping", t, func() {
		Convey("WHEN no message provided, SHOULD respond with PONG", func() {
			response, _ := godis.Request("PING")
			So(response, ShouldEqual, "PONG")
		})
		Convey("WHEN a message is provided, SHOULD response with message", func() {
			response, _ := godis.Request("PING", anyMessage)
			So(response, ShouldEqual, anyMessage)
		})
	})

	Convey("GIVEN Echo", t, func() {
		Convey("When a message is provided, SHOULD respond with that message", func() {
			response, _ := godis.Request("ECHO", anyMessage)
			So(response, ShouldEqual, anyMessage)
		})
		Convey("When no message provided, SHOULD respond with error", func() {
			_, err := godis.Request("ECHO")
			So(err.Error(), ShouldEqual, "ERR wrong number of arguments for 'echo' command")
		})
		Convey("When too many messages provided, SHOULD respond with error", func() {
			_, err := godis.Request("ECHO", "Hello", "world")
			So(err.Error(), ShouldEqual, "ERR wrong number of arguments for 'echo' command")
		})
	})

	Convey("GIVEN Set", t, func() {
		Convey("When given key value pair, value SHOULD be retrievable by that key", func() {
			response, _ := godis.Request("SET", anyKey, anyValue)
			So(response, ShouldEqual, "OK")
			response, _ = godis.Request("GET", anyKey)
			So(response, ShouldEqual, anyValue)
		})
		Convey("When given too few arguments, SHOULD return error", func() {
			_, err := godis.Request("SET")
			So(err.Error(), ShouldEqual, "ERR wrong number of arguments for 'set' command")
		})
	})
}
