package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gogogomoku/gomoku/internal/brain"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type SimpleResponse struct {
	ResCode int    `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

type StartGameParams struct {
	AiStatus1 int16 `json:"AiStatus1" xml:"AiStatus1"`
	AiStatus2 int16 `json:"AiStatus2" xml:"AiStatus2"`
}

func StartServer() {
	e := echo.New()
	url := fmt.Sprintf("http://%s:%s", os.Getenv("APP_SERVER"), os.Getenv("VUE_APP_PORT"))
	fmt.Println(url)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://127.0.0.1:8080",
			"http://192.168.1.19:8080",
			url,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
		},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.GET("/", home)
	// e.POST("/end", ClearGame) // keep for now
	e.POST("/start", StartGame)
	e.POST("/restart", RestartGame)
	e.GET("/move/:pos/id/:id", MakeMove)
	e.Debug = true
	e.Start(":4243")
}

func returnGameInfo(c echo.Context) error {
	if brain.Game.Status == brain.NotStarted {
		return c.JSON(http.StatusOK, &SimpleResponse{1, "Game hasn't started yet"})
	} else if brain.Game.Status == brain.Running {
		return c.JSON(http.StatusOK, brain.Game)
	} else if brain.Game.Status == brain.Concluded {
		return c.JSON(http.StatusOK, brain.Game)
	} else {
		return c.JSON(http.StatusOK, &SimpleResponse{1, "Game status not recognized"})
	}
}

func home(c echo.Context) error {
	brain.Game.Winner = 0 // todo: reset whole state here ?
	return returnGameInfo(c)
}

func StartGame(c echo.Context) error {
	m := StartGameParams{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	brain.StartRound(m.AiStatus1, m.AiStatus2)
	return returnGameInfo(c)
}

func RestartGame(c echo.Context) error { // todo: rmv if same as StartGame
	m := StartGameParams{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	brain.StartRound(m.AiStatus1, m.AiStatus2)
	return returnGameInfo(c)
}

// keep for now
// func ClearGame(c echo.Context) error {
// 	m := StartGameParams{}
// 	if err := c.Bind(&m); err != nil {
// 		return err
// 	}
// 	brain.ClearRound(m.AiStatus1, m.AiStatus2)
// 	return returnGameInfo(c)
// }

func MakeMove(c echo.Context) error {
	id, err1 := (strconv.Atoi(c.Param("id")))
	position, err2 := strconv.Atoi(c.Param("pos"))
	if err1 != nil || err2 != nil {
		txtErr := fmt.Sprint("Error in params", err1, err2)
		return c.JSON(http.StatusOK, &SimpleResponse{-1, txtErr})
	}
	code, msg := brain.HandleMove(int16(id), int16(position))
	// After turn change
	brain.GetInvalidMovesForCurrentPlayer()
	fmt.Printf("(debug) brain.HandleMove() msg :: %#v\n", msg)
	if code == 0 {
		return c.JSON(http.StatusOK, brain.Game)
	} else {
		return c.JSON(http.StatusBadRequest, &SimpleResponse{-1, msg})
	}
}
