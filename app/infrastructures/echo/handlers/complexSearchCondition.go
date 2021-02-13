package handlers

import (
	"encoding/json"
	"net/http"
	domainQuery "regulus/app/domain/query"
	"regulus/app/repositories"
	"regulus/app/usecases/maintenance/master/query"

	"github.com/labstack/echo"
)

/*
UpdateQueryCondition 検索条件修正用ハンドラ
*/
func UpdateQueryCondition(c echo.Context) error {
	repo := repositories.NewQueryConditionRepo()
	condition := &domainQuery.Condition{}
	e := c.Bind(condition)
	if e != nil {
		return e
	}
	err := query.UpdateCondition(repo, condition, getAuthStaffID(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "data update error")
	}
	return c.JSON(http.StatusOK, "data update ok")
}

/*
AddQueryCondition 検索条件追加用ハンドラ
*/
func AddQueryCondition(c echo.Context) error {
	repo := repositories.NewQueryConditionRepo()
	condition := &domainQuery.Condition{}
	e := c.Bind(condition)
	if e != nil {
		return e
	}
	id, err := query.AddCondition(repo, condition, getAuthStaffID(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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
		return c.JSON(http.StatusInternalServerError, e.Error())
	}
	result, err := query.Find(conditionData, repo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

/*
FetchQueryConditionSearchItems は詳細検索条件設定フォームを開く際に必要なデータを取得するハンドラです。
*/
func FetchQueryConditionSearchItems(c echo.Context) error {
	groupRepo := repositories.NewStaffGroupRepo()
	categories, e := query.FetchDataInputFormItems(groupRepo)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, e)
	}
	for _, category := range categories {

		if category.Name == "query-condition" {
			return c.JSON(http.StatusOK, category.SearchItems)
		}
	}
	return c.JSON(http.StatusInternalServerError, "search items is blank")
}

/*
FetchQueryConditionDataInputFormItems は検索条件登録フォームを開く際に必要なデータを取得するハンドラです。
*/
func FetchQueryConditionDataInputFormItems(c echo.Context) error {
	groupRepo := repositories.NewStaffGroupRepo()
	categories, e := query.FetchDataInputFormItems(groupRepo)
	if e != nil {
		return e
	}
	return c.JSON(http.StatusOK, categories)
}
