package echo

import (
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, c.Request().URL.Host+"/index.html")
}
