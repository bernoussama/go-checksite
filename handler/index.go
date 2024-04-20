package handler

import (
	"github.com/0ussamabernou/go-checksite/view/index"
	"github.com/labstack/echo/v4"
)

// type IndexHandler struct{}

func HandleIndex(c echo.Context) error {
	return render(c, index.Home())

}
