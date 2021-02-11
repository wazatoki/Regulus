package query

import (
	"regulus/app/domain/authentication"
	"regulus/app/utils"
	"sort"
	"strings"
)

/*
Conditions *query.Conditionのスライス

*/
type Conditions []*Condition

/*
Sort is sort maker slice by orderItems
*/
func Sort(queryConditions []*Condition, orderItems ...OrderConditionItem) []*Condition {
	sort.Slice(queryConditions, func(i int, j int) bool {
		return compare(queryConditions[i], queryConditions[j], orderItems, 0)
	})
	return queryConditions
}

func compare(queryCondition1 *Condition, queryCondition2 *Condition, orderItems []OrderConditionItem, orderIndex int) bool {

	if len(orderItems) <= orderIndex {
		return false
	}

	switch orderItems[orderIndex].OrderField.ID {
	case "category-view-value":
		if queryCondition1.Category.ViewValue == queryCondition2.Category.ViewValue {
			orderIndex++
			return compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == Desc {
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
		if orderItems[orderIndex].OrderFieldKeyWord == Desc {
			return qc1 > qc2
		}
		return qc1 < qc2

	case "owner":
		if queryCondition1.Owner.ID == queryCondition2.Owner.ID {
			orderIndex++
			return compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == Desc {
			return queryCondition1.Owner.Name > queryCondition2.Owner.Name
		}
		return queryCondition1.Owner.Name < queryCondition2.Owner.Name

	default: // PatternNameで並び替え
		if queryCondition1.PatternName == queryCondition2.PatternName {
			orderIndex++
			return compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == Desc {
			return queryCondition1.PatternName > queryCondition2.PatternName
		}
		return queryCondition1.PatternName < queryCondition2.PatternName

	}
}

/*
CategoryNameListByMatchType ViewValueの値が引数の文字列sとMatchType mtに合致する
NameのSliceを返す。

*/
func CategoryNameListByMatchType(s string, mt MatchTypeEnum) (categoryNames []string) {

	categories := CreateCategories([]*authentication.Group{})

	for _, category := range categories {

		switch mt {
		case Match:
			if category.ViewValue == s {
				categoryNames = append(categoryNames, category.Name)
			}
		case Unmatch:
			if category.ViewValue != s {
				categoryNames = append(categoryNames, category.Name)
			}
		default: // Pertialmatch
			if strings.Contains(category.ViewValue, s) {
				categoryNames = append(categoryNames, category.Name)
			}
		}
	}

	return
}

/*
CreateCategories 検索パターン作成時に使用するカテゴリーリストを返す
optionのグループにはすべてのauthentication.Groupを渡す。
*/
func CreateCategories(groups []*authentication.Group) (categories []*Category) {

	categories = []*Category{}

	categories = append(categories, createQueryConditionCategory(groups))

	categories = append(categories, createStaffCategory(groups))

	categories = append(categories, createStaffGroupCategory(groups))

	return
}

func createQueryConditionCategory(groups []*authentication.Group) (category *Category) {

	optionItems := []OptionItem{}

	for _, g := range groups {
		optionItems = append(optionItems, OptionItem{ID: g.ID, ViewValue: g.Name})
	}

	category = &Category{
		Name:      "query-condition",
		ViewValue: "検索条件管理",
		SearchItems: ComplexSearchItems{
			SearchConditionList: []FieldAttr{
				{
					ID:        "pattern-name",
					ViewValue: "検索パターン名称",
					FieldType: STRING,
				},
				{
					ID:        "category-view-value",
					ViewValue: "カテゴリー名称",
					FieldType: STRING,
				},
				{
					ID:        "is-disclose",
					ViewValue: "公開",
					FieldType: BOOLEAN,
				},
				{
					ID:          "disclose-groups",
					ViewValue:   "公開先グループ",
					FieldType:   ARRAY,
					OptionItems: optionItems,
				},
				{
					ID:        "owner",
					ViewValue: "所有者",
					FieldType: STRING,
				},
			},
			DisplayItemList:    []FieldAttr{},
			OrderConditionList: []FieldAttr{},
			Groups:             groups,
		},
	}

	return
}

func createStaffCategory(groups []*authentication.Group) (category *Category) {
	category = &Category{
		Name:      "staff",
		ViewValue: "利用者",
		SearchItems: ComplexSearchItems{
			SearchConditionList: []FieldAttr{
				{
					ID:        "account-id",
					ViewValue: "利用者ID",
					FieldType: STRING,
				},
				{
					ID:        "name",
					ViewValue: "利用者名称",
					FieldType: STRING,
				},
				{
					ID:        "groups",
					ViewValue: "所属グループ",
					FieldType: ARRAY,
				},
				{
					ID:        "group-name",
					ViewValue: "所属グループ名",
					FieldType: STRING,
				},
			},
			DisplayItemList:    []FieldAttr{},
			OrderConditionList: []FieldAttr{},
			Groups:             groups,
		},
	}
	return
}

func createStaffGroupCategory(groups []*authentication.Group) (category *Category) {
	category = &Category{
		Name:      "staff-group",
		ViewValue: "利用者グループ",
		SearchItems: ComplexSearchItems{
			SearchConditionList: []FieldAttr{
				{
					ID:        "name",
					ViewValue: "グループ名称",
					FieldType: STRING,
				},
				{
					ID:        "staff-name",
					ViewValue: "利用者名称",
					FieldType: STRING,
				},
				{
					ID:        "staff-account-id",
					ViewValue: "利用者ID",
					FieldType: STRING,
				},
			},
			DisplayItemList:    []FieldAttr{},
			OrderConditionList: []FieldAttr{},
			Groups:             groups,
		},
	}
	return
}
