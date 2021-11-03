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
	defineComplexSearchConditionRouting(api)
	defineStaffGroupRouting(api)
}

func defineComplexSearchConditionRouting(api *echo.Group) {
	api.GET("/complexSearchCondition", handlers.FindQueryConditionByCondition)
	api.POST("/complexSearchCondition", handlers.AddQueryCondition)
	api.PUT("/complexSearchCondition", handlers.UpdateQueryCondition)
	api.DELETE("/complexSearchCondition", handlers.DeleteQueryCondition)
	api.GET("/complexSearchCondition/dataInputFormItems", handlers.FetchQueryConditionDataInputFormItems)
	api.GET("/complexSearchCondition/complexSearchItems", handlers.FetchQueryConditionSearchItems)
	api.GET("/complexSearchCondition/updateFavoriteConditions", handlers.UpdateFavoriteConditions)
}

func defineStaffGroupRouting(api *echo.Group) {
	api.GET("/staffGroup", handlers.FindStaffGroupByCondition)
	api.POST("/staffGroup", handlers.AddStaffGroup)
	api.PUT("/staffGroup", handlers.UpdateStaffGroup)
	api.DELETE("/staffGroup", handlers.DeleteStaffGroup)
	api.GET("/staffGroup/complexSearchItems", handlers.FetchStaffGroupSearchItems)
}
