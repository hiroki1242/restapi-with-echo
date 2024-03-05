package main

import (
	"api-with-echo/app/router"

	"github.com/labstack/echo/middleware"
)

func main() {
	e := router.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
