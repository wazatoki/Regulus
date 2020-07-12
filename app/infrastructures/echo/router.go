package echo

import (
	"regulus/app/infrastructures/echo/handlers"

	"github.com/labstack/echo"
)

func defineRouting(e *echo.Echo) {
	//apiPath := "/api/"
	e.GET("/", handlers.Root)
	e.GET("/index", handlers.Index)
	//e.GET(apiPath+"maker/ComplexSearchItems", handle.MakerComplexSearchItems)
}
