package handlers

import (
	"encoding/json"
	"net/http"
	"regulus/app/domain/entities"
	"regulus/app/domain/services"
	domainQuery "regulus/app/domain/vo/query"
	"regulus/app/repositories"
	"regulus/app/usecases/maintenance/master/query"

	"github.com/labstack/echo"
)

func AddQueryCondition(c echo.Context) error {
	repo := repositories.NewQueryConditionRepo()
	condition := &entities.QueryCondition{}
	e := c.Bind(condition)
	if e != nil {
		return e
	}
	id, err := query.AddCondition(repo, condition)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "data insert error")
	}
	return c.JSON(http.StatusOK, id)
}

/*
FindQueryConditionByCondition return search result of query condition
*/
func FindQueryConditionByCondition(c echo.Context) error {
	repo := repositories.NewQueryConditionRepo()
	var conditionData *domainQuery.ConditionData
	conditionData = &domainQuery.ConditionData{}
	e := json.Unmarshal([]byte(c.QueryParam("condition")), conditionData)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, "request unmarshal error")
	}
	result, err := query.Find(repo, conditionData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "data fetch error")
	}
	return c.JSON(http.StatusOK, result)
}

/*
FindAllComplexConditionSearchCategories return condition search categories
*/
func FindAllComplexConditionSearchCategories(c echo.Context) error {
	r := repositories.NewStaffGroupRepo()
	groups, _ := r.SelectAll()
	return c.JSON(http.StatusOK, services.CreateCategories(groups))
}

/*
FindComplexSearchItems return condition search items as query-condition
*/
func FindComplexSearchItems(c echo.Context) error {

	var items entities.ComplexSearchItems
	r := repositories.NewStaffGroupRepo()
	groups, _ := r.SelectAll()

	for _, category := range services.CreateCategories(groups) {

		if category.Name == "query-condition" {

			items = category.SearchItems

			break
		}
	}
	return c.JSON(http.StatusOK, items)
}
