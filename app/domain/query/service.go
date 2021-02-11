package query

import (
	"regulus/app/domain/authentication"
	"regulus/app/utils"
	"sort"
	"strings"
)

/*
Conditions *entities.Conditionのスライス
Findで条件抽出
Sortで並び替え
*/
type Conditions []*Condition

/*
Find 条件抽出
*/
func (q Conditions) Find(queryItems ...SearchConditionItem) (result Conditions) {

	var temp Conditions
	var temp2 Conditions
	result = make(Conditions, 0)

	for _, condition := range q {

		result = append(result, condition)
	}

	for _, queryItem := range queryItems {

		temp = make(Conditions, 0)

		if queryItem.Operator == And {

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

			temp2 = make(Conditions, 0)

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

func (q Conditions) isMatchCondition(sc SearchConditionItem, qc *Condition) bool {

	switch sc.SearchField.FieldType {
	case STRING:

		switch sc.MatchType {
		case Match:

			switch sc.SearchField.ID {
			case "pattern-name":
				return qc.PatternName == sc.ConditionValue

			case "category-view-value":
				return qc.Category.ViewValue == sc.ConditionValue

			case "owner":
				return qc.Owner.Name == sc.SearchField.ViewValue
			}

		case Unmatch:

			switch sc.SearchField.ID {
			case "pattern-name":
				return qc.PatternName != sc.ConditionValue

			case "category-view-value":
				return qc.Category.ViewValue != sc.ConditionValue

			case "owner":
				return qc.Owner.Name != sc.SearchField.ViewValue
			}

		case Pertialmatch:

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

	case BOOLEAN:

		switch sc.SearchField.ID {
		case "is-disclose":
			if qc.IsDisclose {
				return sc.ConditionValue == "true"
			}
			return sc.ConditionValue == "false"
		}

	case ARRAY:

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
	case "pattern-name":
		if queryCondition1.PatternName == queryCondition2.PatternName {
			orderIndex++
			return compare(queryCondition1, queryCondition2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == Desc {
			return queryCondition1.PatternName > queryCondition2.PatternName
		}
		return queryCondition1.PatternName < queryCondition2.PatternName
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

	default:
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
