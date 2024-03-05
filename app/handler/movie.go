package handler

import (
	"api-with-echo/app/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var movies = []model.Movie{
	{1, "Avengers", 2012},
	{2, "Harry Potter", 2001},
	{3, "Star Wars", 1978},
}

func GetAllMovies(c echo.Context) error {
	return c.JSON(http.StatusOK, movies)
}

func GetSpecifiedMovie(c echo.Context) error {
	requestMovieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Id\n")
	}

	for _, movie := range movies {
		if movie.Id == requestMovieId {
			return c.JSON(http.StatusOK, movie)
		}
	}

	return c.String(http.StatusNotFound, "Movie was not found.\n")
}

func CreateMovie(c echo.Context) error {
	var newMovie model.Movie
	if err := c.Bind(&newMovie); err != nil {
		return err
	}

	for _, movie := range movies {
		if movie.Id == newMovie.Id {
			return c.String(http.StatusBadRequest, "Id already exitsts\n")
		}
	}

	movies = append(movies, newMovie)
	return c.JSON(http.StatusCreated, movies)
}

func UpdateMovie(c echo.Context) error {
	requestMovieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Id\n")
	}

	var updateMovieInfo model.Movie
	if err := c.Bind(&updateMovieInfo); err != nil {
		return err
	}

	if requestMovieId == updateMovieInfo.Id {
		return c.String(http.StatusBadRequest, "Invalid Id\n")
	}

	for i, movie := range movies {
		if movie.Id == requestMovieId {
			movies[i] = updateMovieInfo
			return c.JSON(http.StatusOK, movies)
		}
	}

	return c.String(http.StatusNotFound, "Movie was not found.\n")
}

func DeleteMovie(c echo.Context) error {
	requestMovieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Id\n")
	}

	for i, movie := range movies {
		if movie.Id == requestMovieId {
			movies = append(movies[:i], movies[i+1:]...)
			return c.JSON(http.StatusOK, movies)
		}
	}

	return c.String(http.StatusNotFound, "Movie was not found\n")
}
