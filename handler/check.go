package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	s "strings"
)

func HandleCheck(c echo.Context) error {
	url := c.FormValue("url")

	if s.HasPrefix(url, "http://") || s.HasPrefix(url, "https://") {
		// url = url

	} else {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusOK, "Site is "+"<strong style='color:red'>"+"down"+"</strong>")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return c.HTML(http.StatusOK, "Site is "+"<strong style='color:green'>"+"Up"+"</strong>")
	} else {
		return c.HTML(http.StatusOK, "Site is "+"<strong style='color:red'>"+"down"+"</strong>")
	}
	// return c.String(http.StatusOK, url)
}
