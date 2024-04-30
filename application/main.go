package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
	"github.com/shun198/go-crm/application/docs"
)

func main() {
	e := echo.New()
	// https://github.com/swaggo/echo-swagger?tab=readme-ov-file
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hot Reload")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
