package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gogogomoku/gomoku/internal/brain"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type SimpleResponse struct {
	ResCode int    `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func StartServer() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://192.168.1.40:8080/"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", home)
	e.GET("/start", StartGame)
	e.GET("/restart", RestartGame)
	e.GET("/move/:pos/id/:id", MakeMove)
	e.Start(":4242")
}

func home(c echo.Context) error {
	if brain.GameRound.Status == brain.NotStarted {
		return c.JSON(http.StatusOK, &SimpleResponse{1, "Game hasn't started yet"})
	} else if brain.GameRound.Status == brain.Running {
		return c.JSON(http.StatusOK, brain.GameRound)
	} else {
		return c.JSON(http.StatusOK, &SimpleResponse{3, "Game finished"})
	}
}

func StartGame(c echo.Context) error {
	brain.StartRound()
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func RestartGame(c echo.Context) error {
	brain.InitializeValues()
	brain.StartRound()
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func MakeMove(c echo.Context) error {
	id, err1 := strconv.Atoi(c.Param("id"))
	position, err2 := strconv.Atoi(c.Param("pos"))
	if err1 != nil || err2 != nil {
		txtErr := fmt.Sprint("Error in params", err1, err2)
		return c.JSON(http.StatusOK, &SimpleResponse{-1, txtErr})
	}
	code, msg := brain.HandleMove(id, position)
	if code == 0 {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		return c.JSON(http.StatusBadRequest, &SimpleResponse{-1, msg})
	}
}
