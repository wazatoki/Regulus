package echo

import (
	"github.com/labstack/echo"
)

func defineRouting(e *echo.Echo) {
	e.GET("/", index)
}
