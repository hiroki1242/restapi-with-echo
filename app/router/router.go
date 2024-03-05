package router

import (
	"api-with-echo/app/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() (e *echo.Echo) {
	e = echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/movies", handler.GetAllMovies)
	e.GET("/movies/:id", handler.GetSpecifiedMovie)
	e.POST("/movies", handler.CreateMovie)
	e.PUT("/movies/:id", handler.UpdateMovie)
	e.DELETE("/movies/:id", handler.DeleteMovie)
	return
}
