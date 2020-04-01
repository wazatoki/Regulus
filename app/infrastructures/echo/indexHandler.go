package echo

import (
	"net/http"

	"github.com/labstack/echo"
)

func root(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, c.Request().URL.Host+"/index.html")
}

func index(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, c.Request().URL.Host+"/index.html")
}
