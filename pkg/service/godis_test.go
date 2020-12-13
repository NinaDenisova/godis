package service

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
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
	var c echo.Context
	var rec *httptest.ResponseRecorder
	var anyMessage string = "Any message"
	Convey("Scenario: Ping with different # of messages", t, func() {
		Convey("WHEN no message provided", func() {
			c, rec = createContext(http.MethodGet, "ping")
			err := godis.HandlePing(c)
			Convey("SHOULD respond with PONG and return OK", func() {
				So(err, ShouldEqual, nil)
				So(rec.Code, ShouldEqual, http.StatusOK)
				So(extractResponse(rec.Body), ShouldEqual, "PONG")
			})
		})
		Convey("WHEN a message is provided", func() {
			c, rec = createContext(http.MethodGet, "ping", anyMessage)
			err := godis.HandlePing(c)
			Convey("SHOULD respond with initial message", func() {
				So(err, ShouldEqual, nil)
				So(rec.Code, ShouldEqual, http.StatusOK)
				So(extractResponse(rec.Body), ShouldEqual, anyMessage)
			})
		})
		//TODO add use case of multiple messages
	})

	Convey("Scenario: ECHO with different # of messages", t, func() {
		Convey("WHEN a message is provided", func() {
			c, rec = createContext(http.MethodGet, "ECHO", anyMessage)
			err := godis.HandleEcho(c)
			Convey("SHOULD respond with the initial message", func() {
				So(err, ShouldEqual, nil)
				So(rec.Code, ShouldEqual, http.StatusOK)
				So(extractResponse(rec.Body), ShouldEqual, anyMessage)
			})
		})
		Convey("WHEN no message provided, SHOULD respond with error", func() {
			c, rec = createContext(http.MethodGet, "ECHO")
			err := godis.HandleEcho(c)
			Convey("SHOULD respond with error", func() {
				he, _ := err.(*echo.HTTPError)
				So(he.Code, ShouldEqual, http.StatusBadRequest)
				// TODO deserialize to compare text
				//So(err.Error(), ShouldEqual, "ERR wrong number of arguments for 'echo' command")
			})
		})
		Convey("WHEN too many messages provided, SHOULD respond with error", func() {
			c, rec = createContext(http.MethodGet, "ECHO", "Hello", "world")
			err := godis.HandleEcho(c)
			So(err, ShouldNotEqual, nil)
			he, _ := err.(*echo.HTTPError)
			So(he.Code, ShouldEqual, http.StatusBadRequest)
			// TODO deserialize to compare text
			//So(err.Error(), ShouldEqual, "ERR wrong number of arguments for 'echo' command")
		})
	})

	// Convey("GIVEN Set", t, func() {
	// 	Convey("When given key value pair, value SHOULD be retrievable by that key", func() {
	// 		c, rec = createContext(http.MethodPost, "Hello", "world")
	// 		response, _ := godis.HandleSet(http.MethodPost, "SET", anyKey, anyValue)
	// 		So(response, ShouldEqual, "OK")
	// 		response, _ = godis.Request(anyKey)
	// 		So(response, ShouldEqual, anyValue)
	// 	})
	// 	Convey("When given too few arguments, SHOULD return error", func() {
	// 		_, err := godis.Request("SET")
	// 		So(err.Error(), ShouldEqual, "ERR wrong number of arguments for 'set' command")
	// 	})
	// })
}

func NewInputBody(messages ...string) string {
	if len(messages) > 0 {
		if len(messages) > 1 {
			return `{"messages": ["` + messages[0] + `","` + messages[1] + `"]}`
		} else {
			return `{"messages": ["` + messages[0] + `"]}`
		}
	}

	return "{\"messages\": []}"
}

func createContext(httpMethod string, command string, messages ...string) (echo.Context, *httptest.ResponseRecorder) {

	e := echo.New()
	req := httptest.NewRequest(httpMethod, "/", strings.NewReader(NewInputBody(messages...)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(fmt.Sprintf("/%s", command))

	return c, rec
}
func extractResponse(response *bytes.Buffer) string {
	trimmedResponse := strings.TrimPrefix(response.String(), "\"")
	return strings.TrimSuffix(trimmedResponse, "\"\n")
}
