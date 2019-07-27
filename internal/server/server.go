package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gogogomoku/gomoku/internal/brain"
	"github.com/gorilla/mux"
)

type SimpleResponse struct {
	ResCode int    `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func home(w http.ResponseWriter, r *http.Request) {
	brain.StartRound()
	if brain.GameRound.Status == brain.NotStarted {
		json.NewEncoder(w).Encode(SimpleResponse{ResCode: 1, Message: "Game hasn't started yet"})
	} else if brain.GameRound.Status == brain.Running {
		json.NewEncoder(w).Encode(brain.GameRound)
	} else {
		json.NewEncoder(w).Encode(SimpleResponse{3, "Game finished"})
	}
}

func startGame(w http.ResponseWriter, r *http.Request) {
	brain.StartRound()
	json.NewEncoder(w).Encode(brain.GameRound)
}

func restartGame(w http.ResponseWriter, r *http.Request) {
	brain.InitializeValues()
	brain.StartRound()
	json.NewEncoder(w).Encode(brain.GameRound)
}

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/", home)
	router.HandleFunc("/start", startGame)
	router.HandleFunc("/restart", restartGame)
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:4242",
		WriteTimeout: 45 * time.Second,
		ReadTimeout:  45 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
//
// 	"github.com/gogogomoku/gomoku/internal/brain"
//
// 	"github.com/labstack/echo"
// 	"github.com/labstack/echo/middleware"
// )
//
// type SimpleResponse struct {
// 	ResCode int    `json:"code" xml:"code"`
// 	Message string `json:"message" xml:"message"`
// }
//
// func StartServer() {
// 	e := echo.New()
// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{
// 			"http://localhost:8080",
// 			"http://192.168.1.40:8080/",
// 			"http://127.0.0.1:8080",
// 		},
// 		AllowHeaders: []string{
// 			echo.HeaderOrigin,
// 			echo.HeaderContentType,
// 			echo.HeaderAccept,
// 			http.MethodGet,
// 			http.MethodPut,
// 			http.MethodPost,
// 		},
// 	}))
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
// 		StackSize: 1 << 10, // 1 KB
// 	}))
// 	e.GET("/", home)
// 	e.GET("/start", StartGame)
// 	e.GET("/restart", RestartGame)
// 	e.GET("/move/:pos/id/:id", MakeMove)
// 	e.Debug = true
// 	e.Start(":4242")
//
// }
//
// func home(c echo.Context) error {
// 	if brain.GameRound.Status == brain.NotStarted {
// 		return c.JSON(http.StatusOK, &SimpleResponse{1, "Game hasn't started yet"})
// 	} else if brain.GameRound.Status == brain.Running {
// 		return c.JSON(http.StatusOK, brain.GameRound)
// 	} else {
// 		return c.JSON(http.StatusOK, &SimpleResponse{3, "Game finished"})
// 	}
// }
//
// func StartGame(c echo.Context) error {
// 	brain.StartRound()
// 	return c.Redirect(http.StatusTemporaryRedirect, "/")
// }
//
// func RestartGame(c echo.Context) error {
// 	brain.InitializeValues()
// 	brain.StartRound()
// 	return c.Redirect(http.StatusTemporaryRedirect, "/")
// }
//
// func MakeMove(c echo.Context) error {
// 	id, err1 := strconv.Atoi(c.Param("id"))
// 	position, err2 := strconv.Atoi(c.Param("pos"))
// 	if err1 != nil || err2 != nil {
// 		txtErr := fmt.Sprint("Error in params", err1, err2)
// 		return c.JSON(http.StatusOK, &SimpleResponse{-1, txtErr})
// 	}
// 	code, msg := brain.HandleMove(id, position)
// 	if code == 0 {
// 		return c.JSON(http.StatusOK, brain.GameRound)
// 	} else {
// 		return c.JSON(http.StatusBadRequest, &SimpleResponse{-1, msg})
// 	}
// }
