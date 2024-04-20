package main

import (
	"github.com/0ussamabernou/go-checksite/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handler.HandleIndex)
	e.POST("/check", handler.HandleCheck)
	e.Logger.Fatal(e.Start(":1323"))

}
