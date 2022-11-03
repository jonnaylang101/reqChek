package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const reqIdHeaderKey string = "X-Request-ID"

func main() {
	e := echo.New()
	e.GET("/requestid", func(c echo.Context) error {
		return c.String(
			http.StatusOK,
			fmt.Sprintf(
				"this is the %s header: %s",
				reqIdHeaderKey,
				c.Request().Header.Get(reqIdHeaderKey),
			),
		)
	})
	log.Fatal(e.Start(":5000"))
}
