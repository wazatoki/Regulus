package services

import (
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/utils"
	"sort"
)

/*
CreateCategories 検索パターン作成時に使用するカテゴリーリストを返す
*/
func CreateCategories(groups []entities.StaffGroup) (categories []entities.Category) {

	categories = []entities.Category{}

	categories = append(categories, createQueryConditionCategory(groups))

	categories = append(categories, createStaffCategory(groups))

	categories = append(categories, createStaffGroupCategory(groups))

	return
}

func createQueryConditionCategory(groups []entities.StaffGroup) (category entities.Category) {

	optionItems := []query.OptionItem{}

	for _, g := range groups {
		optionItems = append(optionItems, query.OptionItem{ID: g.ID, ViewValue: g.Name})
	}

	category = entities.Category{
		Name:      "query-condition",
		ViewValue: "検索条件管理",
		SearchItems: entities.ComplexSearchItems{
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
			DisplayItemList:    []query.FieldAttr{},
			OrderConditionList: []query.FieldAttr{},
			Groups:             groups,
		},
	}

	return
}

func createStaffCategory(groups []entities.StaffGroup) (category entities.Category) {
	category = entities.Category{
		Name:      "staff",
		ViewValue: "利用者",
		SearchItems: entities.ComplexSearchItems{
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

func createStaffGroupCategory(groups []entities.StaffGroup) (category entities.Category) {
	category = entities.Category{
		Name:      "staff-group",
		ViewValue: "利用者グループ",
		SearchItems: entities.ComplexSearchItems{
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

/*
Sort is sort maker slice by orderItems
*/
func Sort(queryConditions []entities.QueryCondition, orderItems ...query.OrderConditionItem) []entities.QueryCondition {
	sort.Slice(queryConditions, func(i int, j int) bool {
		return compare(queryConditions[i], queryConditions[j], orderItems, 0)
	})
	return queryConditions
}

func compare(queryCondition1 entities.QueryCondition, queryCondition2 entities.QueryCondition, orderItems []query.OrderConditionItem, orderIndex int) bool {

	if len(orderItems) <= orderIndex {
		return false
	}

	switch orderItems[orderIndex].OrderField.ID {
	case "pattern-name":
		if queryCondition1.PatternName == queryCondition2.PatternName {
			orderIndex++
			return compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return queryCondition1.PatternName > queryCondition2.PatternName
		}
		return queryCondition1.PatternName < queryCondition2.PatternName
	case "category-view-value":
		if queryCondition1.Category.ViewValue == queryCondition2.Category.ViewValue {
			orderIndex++
			return compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return queryCondition1.Category.ViewValue > queryCondition2.Category.ViewValue
		}
		return queryCondition1.Category.ViewValue < queryCondition2.Category.ViewValue
	case "is-disclose":
		if queryCondition1.IsDisclose == queryCondition2.IsDisclose {
			orderIndex++
			return compare(queryCondition1, queryCondition2, orderItems, orderIndex)
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
			return compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return queryCondition1.Owner.Name > queryCondition2.Owner.Name
		}
		return queryCondition1.Owner.Name < queryCondition2.Owner.Name

	default:
		if queryCondition1.PatternName == queryCondition2.PatternName {
			orderIndex++
			return compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return queryCondition1.PatternName > queryCondition2.PatternName
		}
		return queryCondition1.PatternName < queryCondition2.PatternName

	}
}
