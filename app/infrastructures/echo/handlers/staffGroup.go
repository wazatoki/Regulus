package handlers

import (
	"encoding/json"
	"net/http"
	"regulus/app/domain"
	"regulus/app/repositories"
	"regulus/app/usecases"

	"github.com/labstack/echo"
)

/*
DeleteStaffGroup StaffGroup削除用ハンドラ
*/
func DeleteStaffGroup(c echo.Context) error {
	repo := repositories.NewStaffGroupRepo()
	conditionIDs := &[]string{}
	e := c.Bind(conditionIDs)
	if e != nil {
		return e
	}
	result := usecases.GroupDelete(conditionIDs, repo, getAuthStaffID(c))

	return c.JSON(http.StatusOK, result)
}

/*
UpdateStaffGroup StaffGroup修正用ハンドラ
*/
func UpdateStaffGroup(c echo.Context) error {
	repo := repositories.NewStaffGroupRepo()
	staffGroup := &domain.StaffGroup{}
	e := c.Bind(staffGroup)
	if e != nil {
		return e
	}
	err := usecases.GroupUpdate(staffGroup, repo, getAuthStaffID(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorInfo(err.Error(), http.StatusInternalServerError, "Internal Server Error"))
	}
	return c.JSON(http.StatusOK, nil)
}

/*
AddStaffGroup StaffGroup追加用ハンドラ
*/
func AddStaffGroup(c echo.Context) error {
	repo := repositories.NewStaffGroupRepo()
	staffGroup := &domain.StaffGroup{}
	e := c.Bind(staffGroup)
	if e != nil {
		return e
	}
	result, err := usecases.GroupAdd(staffGroup, repo, getAuthStaffID(c))
	if err != nil {

		return c.JSON(http.StatusInternalServerError, newErrorInfo(err.Error(), http.StatusInternalServerError, "Internal Server Error"))
	}

	return c.JSON(http.StatusOK, *result)
}

/*
FindStaffGroupByCondition return search result of StaffGroups
*/
func FindStaffGroupByCondition(c echo.Context) error {
	repo := repositories.NewStaffGroupRepo()
	conditionData := &domain.ConditionData{}
	e := json.Unmarshal([]byte(c.QueryParam("condition")), conditionData)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, e.Error())
	}
	result, err := usecases.GroupFind(conditionData, repo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorInfo(err.Error(), http.StatusInternalServerError, "Internal Server Error"))
	}
	return c.JSON(http.StatusOK, result)
}

/*
FetchStaffGroupSearchItems は詳細検索条件設定フォームを開く際に必要なデータを取得するハンドラです。
*/
func FetchStaffGroupSearchItems(c echo.Context) error {
	groupRepo := repositories.NewStaffGroupRepo()
	categories, e := usecases.QueryFetchDataInputFormItems(groupRepo)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, e)
	}
	for _, category := range categories {

		if category.Name == "staff-group" {
			return c.JSON(http.StatusOK, category.SearchItems)
		}
	}
	return c.JSON(http.StatusInternalServerError, newErrorInfo("search items is blank", http.StatusInternalServerError, "Internal Server Error"))
}
