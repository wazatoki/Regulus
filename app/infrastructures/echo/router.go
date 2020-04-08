package echo

import (
	"regulus/app/gateway/handle"

	"github.com/labstack/echo"
)

func defineRouting(e *echo.Echo) {
	apiPath := "/api"
	e.GET("/", root)
	e.GET("/index", index)
	e.GET(apiPath+"/maker/ComplexSearchItems", handle.MakerComplexSearchItems)
}
