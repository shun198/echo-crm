package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
	_ "github.com/shun198/go-crm/application/docs"
)

// @title Swagger Example API
// @version 1.0
func main() {
	e := echo.New()
	// https://github.com/swaggo/echo-swagger?tab=readme-ov-file
	// https://medium.com/@chaewonkong/a-five-step-guide-to-integrating-swagger-with-echo-in-go-79be49cfedbe
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hot Reload")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
