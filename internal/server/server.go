package server

import (
	"fmt"
	"net/http"
	"strconv"

	"gomoku/internal/brain"

	"github.com/labstack/echo"
)

type SimpleResponse struct {
	Message string `json:"message" xml:"message"`
}

func StartServer() {
	e := echo.New()
	e.GET("/", home)
	e.GET("/start", StartGame)
	e.GET("/move/:pos/id/:id", MakeMove)
	e.Start(":4242")
}

func home(c echo.Context) error {
	if brain.GameRound.Status == brain.NotStarted {
		return c.JSON(http.StatusOK, &SimpleResponse{"Game hasn't started yet"})
	} else if brain.GameRound.Status == brain.Running {
		return c.JSON(http.StatusOK, brain.GameRound)
	} else {
		return c.JSON(http.StatusOK, &SimpleResponse{"Game finished"})
	}
}

func StartGame(c echo.Context) error {
	brain.StartRound()
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func MakeMove(c echo.Context) error {
	id, err1 := strconv.Atoi(c.Param("id"))
	position, err2 := strconv.Atoi(c.Param("pos"))
	if err1 != nil || err2 != nil {
		txtErr := fmt.Sprint("Error in params", err1, err2)
		return c.JSON(http.StatusOK, &SimpleResponse{txtErr})
	}
	code, msg := brain.HandleMove(id, position)
	if code == 0 {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		return c.JSON(http.StatusBadRequest, &SimpleResponse{msg})
	}
}
