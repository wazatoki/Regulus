package echo

import (
	"regulus/app/gateway/handle"

	"github.com/labstack/echo"
)

func defineRouting(e *echo.Echo) {
	e.GET("/", root)
	e.GET("/index", index)
	e.GET("/maker/ComplexSearchItems", handle.MakerComplexSearchItems)
}
