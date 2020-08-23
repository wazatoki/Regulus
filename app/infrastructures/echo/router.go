package echo

import (
	"regulus/app/infrastructures/echo/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func defineRouting(e *echo.Echo) {
	apiPath := "/api/"
	api := e.Group("/api")
	api.Use(middleware.JWT([]byte("secret")))

	e.GET("/", handlers.Root)
	e.GET("/index", handlers.Index)
	e.POST("/login", handlers.Login)
	api.GET(apiPath+"complexSearchCondition", handlers.FindQueryConditionByCondition)
	api.POST(apiPath+"complexSearchCondition", handlers.AddQueryCondition)
	api.GET(apiPath+"complexSearchCondition/categories", handlers.FindAllComplexConditionSearchCategories)
}
