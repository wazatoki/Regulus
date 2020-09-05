package services

import (
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/utils"
	"sort"
	"strings"
)

/*
QueryConditions *entities.QueryConditionのスライス
Findで条件抽出
Sortで並び替え
*/
type QueryConditions []*entities.QueryCondition

/*
Sort is sort QueryConditions slice by orderItems
*/
func (q QueryConditions) Sort(orderItems ...query.OrderConditionItem) (result QueryConditions) {

	result = make(QueryConditions, 0)

	for _, condition := range q {

		result = append(result, condition)
	}

	sort.Slice(result, func(i int, j int) bool {
		return q.compare(result[i], result[j], orderItems, 0)
	})
	return
}

func (q QueryConditions) compare(
	queryCondition1 *entities.QueryCondition,
	queryCondition2 *entities.QueryCondition,
	orderItems []query.OrderConditionItem,
	orderIndex int) bool {

	if len(orderItems) <= orderIndex {
		return false
	}

	switch orderItems[orderIndex].OrderField.ID {
	case "pattern-name":
		if queryCondition1.PatternName == queryCondition2.PatternName {
			orderIndex++
			return q.compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return queryCondition1.PatternName > queryCondition2.PatternName
		}
		return queryCondition1.PatternName < queryCondition2.PatternName
	case "category-view-value":
		if queryCondition1.Category.ViewValue == queryCondition2.Category.ViewValue {
			orderIndex++
			return q.compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return queryCondition1.Category.ViewValue > queryCondition2.Category.ViewValue
		}
		return queryCondition1.Category.ViewValue < queryCondition2.Category.ViewValue
	case "is-disclose":
		if queryCondition1.IsDisclose == queryCondition2.IsDisclose {
			orderIndex++
			return q.compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		qc1 := utils.BoolToInt(queryCondition1.IsDisclose)
		qc2 := utils.BoolToInt(queryCondition2.IsDisclose)
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return qc1 > qc2
		}
		return qc1 < qc2

	case "owner":
		if queryCondition1.Owner.ID == queryCondition2.Owner.ID {
			orderIndex++
			return q.compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return queryCondition1.Owner.Name > queryCondition2.Owner.Name
		}
		return queryCondition1.Owner.Name < queryCondition2.Owner.Name

	default:
		if queryCondition1.PatternName == queryCondition2.PatternName {
			orderIndex++
			return q.compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return queryCondition1.PatternName > queryCondition2.PatternName
		}
		return queryCondition1.PatternName < queryCondition2.PatternName

	}
}

/*
FindBySearchStrings キーワード検索
*/
func (q QueryConditions) FindBySearchStrings(serchStrings ...string) (result QueryConditions) {

	if len(serchStrings) == 0 {
		return q
	}

	result = make(QueryConditions, 0)
	for _, c := range q {
		result = append(result, c)
	}

	for _, str := range serchStrings {

		temp := make(QueryConditions, 0)

		for _, c := range result {
			if strings.Index(c.PatternName, str) != -1 || strings.Index(c.Category.ViewValue, str) != -1 || strings.Index(c.Owner.Name, str) != -1 {
				temp = append(temp, c)
			}
		}

		result = temp
	}

	return
}

/*
FindBySearchConditionItem 条件抽出
*/
func (q QueryConditions) FindBySearchConditionItem(queryItems ...query.SearchConditionItem) (result QueryConditions) {

	if len(queryItems) == 0 {
		return q
	}

	var temp QueryConditions
	var temp2 QueryConditions
	result = make(QueryConditions, 0)

	for _, condition := range q {

		result = append(result, condition)
	}

	for _, queryItem := range queryItems {

		temp = make(QueryConditions, 0)

		if queryItem.Operator == query.And {

			for _, queryCondition := range result {

				if q.isMatchCondition(queryItem, queryCondition) {

					temp = append(temp, queryCondition)
				}

			}

			result = temp
		} else {

			for _, queryCondition := range q {

				if q.isMatchCondition(queryItem, queryCondition) {

					temp = append(temp, queryCondition)
				}
			}

			temp2 = make(QueryConditions, 0)

			for _, tempCondition := range temp {

				f := false

				for _, resultCondition := range result {

					if tempCondition == resultCondition {

						f = true
						break
					}
				}

				if !f {

					temp2 = append(temp2, tempCondition)
				}
			}

			result = append(result, temp2...)
		}
	}

	return
}

func (q QueryConditions) isMatchCondition(sc query.SearchConditionItem, qc *entities.QueryCondition) bool {

	switch sc.SearchField.FieldType {
	case query.STRING:

		switch sc.MatchType {
		case query.Match:

			switch sc.SearchField.ID {
			case "pattern-name":
				return qc.PatternName == sc.ConditionValue

			case "category-view-value":
				return qc.Category.ViewValue == sc.ConditionValue

			case "owner":
				return qc.Owner.Name == sc.SearchField.ViewValue
			}

		case query.Unmatch:

			switch sc.SearchField.ID {
			case "pattern-name":
				return qc.PatternName != sc.ConditionValue

			case "category-view-value":
				return qc.Category.ViewValue != sc.ConditionValue

			case "owner":
				return qc.Owner.Name != sc.SearchField.ViewValue
			}

		case query.Pertialmatch:

			switch sc.SearchField.ID {
			case "pattern-name":
				if strings.Index(qc.PatternName, sc.ConditionValue) == -1 {
					return false
				}
				return true

			case "category-view-value":
				if strings.Index(qc.Category.ViewValue, sc.ConditionValue) == -1 {
					return false
				}
				return true

			case "owner":
				if strings.Index(qc.Owner.Name, sc.ConditionValue) == -1 {
					return false
				}
				return true
			}
		}

	case query.BOOLEAN:

		switch sc.SearchField.ID {
		case "is-disclose":
			if qc.IsDisclose {
				return sc.ConditionValue == "true"
			}
			return sc.ConditionValue == "false"
		}

	case query.ARRAY:

		switch sc.SearchField.ID {
		case "disclose-groups":
			if qc.DiscloseGroups != nil {

				for _, group := range qc.DiscloseGroups {

					if group.ID == sc.ConditionValue {
						return true
					}
				}
				return false
			}
		}

	}

	return false
}

/*
CreateCategories 検索パターン作成時に使用するカテゴリーリストを返す
optionのグループにはすべてのstaffGroupを渡す。
*/
func CreateCategories(groups []*entities.StaffGroup) (categories []*entities.Category) {

	categories = []*entities.Category{}

	categories = append(categories, CreateQueryConditionCategory(groups))

	categories = append(categories, createStaffCategory(groups))

	categories = append(categories, createStaffGroupCategory(groups))

	return
}

/*
CreateQueryConditionCategory 詳細検索条件設定時に使用する条件項目を返す
optionのグループにはすべてのstaffGroupを渡す。
*/
func CreateQueryConditionCategory(groups []*entities.StaffGroup) (category *entities.Category) {

	optionItems := []query.OptionItem{}

	for _, g := range groups {
		optionItems = append(optionItems, query.OptionItem{ID: g.ID, ViewValue: g.Name})
	}

	category = &entities.Category{
		Name:      "query-condition",
		ViewValue: "検索条件管理",
		SearchItems: entities.ComplexSearchItems{
			IsShowDisplayItem:    false,
			IsShowOrderCondition: true,
			IsShowSaveCondition:  true,
			SearchConditionList: []query.FieldAttr{
				{
					ID:        "pattern-name",
					ViewValue: "検索パターン名称",
					FieldType: query.STRING,
				},
				{
					ID:        "category-view-value",
					ViewValue: "カテゴリー名称",
					FieldType: query.STRING,
				},
				{
					ID:        "is-disclose",
					ViewValue: "公開",
					FieldType: query.BOOLEAN,
				},
				{
					ID:          "disclose-groups",
					ViewValue:   "公開先グループ",
					FieldType:   query.ARRAY,
					OptionItems: optionItems,
				},
				{
					ID:        "owner",
					ViewValue: "所有者",
					FieldType: query.STRING,
				},
			},
			DisplayItemList: []query.FieldAttr{},
			OrderConditionList: []query.FieldAttr{
				{
					ID:        "pattern-name",
					ViewValue: "検索パターン名称",
					FieldType: query.STRING,
				},
				{
					ID:        "category-view-value",
					ViewValue: "カテゴリー名称",
					FieldType: query.STRING,
				},
				{
					ID:        "is-disclose",
					ViewValue: "公開",
					FieldType: query.BOOLEAN,
				},
				{
					ID:        "owner",
					ViewValue: "所有者",
					FieldType: query.STRING,
				},
			},
			Groups: groups,
		},
	}

	return
}

func createStaffCategory(groups []*entities.StaffGroup) (category *entities.Category) {
	category = &entities.Category{
		Name:      "staff",
		ViewValue: "利用者",
		SearchItems: entities.ComplexSearchItems{
			IsShowDisplayItem:    false,
			IsShowOrderCondition: false,
			IsShowSaveCondition:  true,
			SearchConditionList: []query.FieldAttr{
				{
					ID:        "account-id",
					ViewValue: "利用者ID",
					FieldType: query.STRING,
				},
				{
					ID:        "name",
					ViewValue: "利用者名称",
					FieldType: query.STRING,
				},
				{
					ID:        "groups",
					ViewValue: "所属グループ",
					FieldType: query.ARRAY,
				},
				{
					ID:        "group-name",
					ViewValue: "所属グループ名",
					FieldType: query.STRING,
				},
			},
			DisplayItemList:    []query.FieldAttr{},
			OrderConditionList: []query.FieldAttr{},
			Groups:             groups,
		},
	}
	return
}

func createStaffGroupCategory(groups []*entities.StaffGroup) (category *entities.Category) {
	category = &entities.Category{
		Name:      "staff-group",
		ViewValue: "利用者グループ",
		SearchItems: entities.ComplexSearchItems{
			IsShowDisplayItem:    false,
			IsShowOrderCondition: false,
			IsShowSaveCondition:  false,
			SearchConditionList: []query.FieldAttr{
				{
					ID:        "name",
					ViewValue: "グループ名称",
					FieldType: query.STRING,
				},
				{
					ID:        "staff-name",
					ViewValue: "利用者名称",
					FieldType: query.STRING,
				},
				{
					ID:        "staff-account-id",
					ViewValue: "利用者ID",
					FieldType: query.STRING,
				},
			},
			DisplayItemList:    []query.FieldAttr{},
			OrderConditionList: []query.FieldAttr{},
			Groups:             groups,
		},
	}
	return
}
