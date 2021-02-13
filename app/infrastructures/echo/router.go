package echo

import (
	"regulus/app/infrastructures/echo/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func defineRouting(e *echo.Echo) {

	e.GET("/", handlers.Root)
	e.GET("/index", handlers.Index)
	e.POST("/login", handlers.Login)

	apiPath := "/api"
	api := e.Group(apiPath)
	api.Use(middleware.JWT([]byte("secret")))
	api.GET("/complexSearchCondition", handlers.FindQueryConditionByCondition)
	api.POST("/complexSearchCondition", handlers.AddQueryCondition)
	api.PUT("/complexSearchCondition", handlers.UpdateQueryCondition)
	api.GET("/complexSearchCondition/DataInputFormItems", handlers.FetchQueryConditionDataInputFormItems)
	api.GET("/complexSearchCondition/complexSearchItems", handlers.FetchQueryConditionSearchItems)
}
