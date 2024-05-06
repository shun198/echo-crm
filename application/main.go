package main

import (
	"github.com/shun198/go-crm/config"
	_ "github.com/shun198/go-crm/docs"
	"github.com/shun198/go-crm/route"

	// https://github.com/labstack/echo-contrib/issues/8
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title	Swagger Example API
// @version	1.0
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db, err := config.StartDatabase()

	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	env := &config.Env{Echo: e, DB: db}
	route.SetUserRoutes(env)
	// https://github.com/swaggo/echo-swagger?tab=readme-ov-file
	// https://medium.com/@chaewonkong/a-five-step-guide-to-integrating-swagger-with-echo-in-go-79be49cfedbe
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
