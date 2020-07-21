package handlers

import (
	"encoding/json"
	"net/http"
	"regulus/app/domain/entities"
	domainQuery "regulus/app/domain/vo/query"
	"regulus/app/repositories"
	"regulus/app/usecases/maintenance/master/query"

	"github.com/labstack/echo"
)

/*
FindQueryConditionByCondition return search result of query condition
*/
func FindQueryConditionByCondition(c echo.Context) error {
	repo := repositories.NewQueryConditionRepo()
	var conditionData *domainQuery.ConditionData
	conditionData = &domainQuery.ConditionData{}
	e := json.Unmarshal([]byte(c.QueryParam("condition")), conditionData)
	if e == nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	result, err := query.Find(repo, conditionData)
	if err == nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, result)
}

/*
FindAllComplexConditionSearchCategories return condition search categories
*/
func FindAllComplexConditionSearchCategories(c echo.Context) error {
	return c.JSON(http.StatusOK, entities.Categories)
}

/*
FindComplexSearchItems return condition search items as query-condition
*/
func FindComplexSearchItems(c echo.Context) error {

	var items entities.ComplexSearchItems

	for _, category := range entities.Categories {

		if category.Name == "query-condition" {

			items = category.SearchItems

			break
		}
	}
	return c.JSON(http.StatusOK, items)
}
