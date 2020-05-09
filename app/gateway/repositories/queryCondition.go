package repositories

import (
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/sqlboiler"
)

// QueryConditionObjectMap data mapper sqlboiler object to entities object
func QueryConditionObjectMap(sqc *sqlboiler.QueryCondition) (eqc entities.QueryCondition) {
	var category entities.Category
	var staffGroups []entities.StaffGroup
	var displayItemList []query.FieldAttr
	var searchConditionList []query.SearchConditionItem
	var orderConditionList []query.OrderConditionItem

	for _, category = range entities.Categories {
		if category.Name == sqc.CategoryName {
			break
		}
	}

	for _, group := range sqc.R.StaffGroups {
		staffGroups = append(staffGroups, StaffGroupObjectMap(group))
	}

	displayItemList = []query.FieldAttr{}
	for _, item := range sqc.R.QueryDisplayItems {
		var displayItem query.FieldAttr
		for _, displayItem = range category.SearchItems.DisplayItemList {
			if item.DisplayFieldID == displayItem.ID {
				break
			}
		}
		displayItemList = append(displayItemList, displayItem)
	}

	searchConditionList = []query.SearchConditionItem{}
	for _, item := range sqc.R.QuerySearchConditionItems {
		var searchField query.FieldAttr
		for _, searchField = range category.SearchItems.SearchConditionList {
			if item.SearchFieldID == searchField.ID {
				break
			}
		}

		var matchTypeEnum query.MatchTypeEnum
		var operatorEnum query.OperatorEnum

		searchConditionItem := query.SearchConditionItem{
			SearchField:    searchField,
			ConditionValue: item.ConditionValue,
			MatchType:      matchTypeEnum.StrToEnum(item.MatchType),
			Operator:       operatorEnum.StrToEnum(item.Operator),
		}

		searchConditionList = append(searchConditionList, searchConditionItem)
	}

	orderConditionList = []query.OrderConditionItem{}
	for _, item := range sqc.R.QueryOrderConditionItems {
		var orderField query.FieldAttr
		for _, orderField = range category.SearchItems.OrderConditionList {
			if item.OrderFieldID == orderField.ID {
				break
			}
		}
		var orderTypeEnum query.OrderTypeEnum

		orderConditionItem := query.OrderConditionItem{
			OrderField:        orderField,
			OrderFieldKeyWord: orderTypeEnum.StrToEnum(item.OrderFieldKeyWord),
		}
		orderConditionList = append(orderConditionList, orderConditionItem)
	}
	eqc = entities.QueryCondition{
		ID:             sqc.ID,
		PatternName:    sqc.PatternName,
		Category:       category,
		IsDisclose:     sqc.IsDisclose,
		DiscloseGroups: staffGroups,
		ConditionData: query.ConditionData{
			DisplayItemList:     displayItemList,
			SearchConditionList: searchConditionList,
			OrderConditionList:  orderConditionList,
		},
		Owner: StaffObjectMap(sqc.R.Owner),
	}
	return
}

// NewQueryConditionRepo constructor
func NewQueryConditionRepo() *QueryConditionRepo {
	return &QueryConditionRepo{database: createDB()}
}

// QueryConditionRepo repository struct
type QueryConditionRepo struct {
	database db
}
