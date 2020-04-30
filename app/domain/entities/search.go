package entities

import (
	"regulus/app/domain/vo/query"
)

/*
Categories 検索パターン作成時に使用するカテゴリーリスト
*/
var Categories = []Category{
	{
		Name:      "search",
		ViewValue: "検索",
		SearchItems: ComplexSearchItems{
			SearchConditionList:  []query.FieldAttr{},
			IsShowDisplayItem:    false,
			IsShowOrderCondition: false,
			IsShowSaveCondition:  false,
			Groups:               nil,
		},
	},
}

/*
Category is struct as search category
*/
type Category struct {
	Name        string             `json:"name"`
	ViewValue   string             `json:"viewValue"`
	SearchItems ComplexSearchItems `json:" searchItems"`
}

/*
ComplexSearchItems is struct as search condition display
*/
type ComplexSearchItems struct {
	DisplayItemList      []query.FieldAttr `json:"displayItemList"`
	SearchConditionList  []query.FieldAttr `json:"searchConditionList"`
	OrderConditionList   []query.FieldAttr `json:"orderConditionList"`
	IsShowDisplayItem    bool              `json:"isShowDisplayItem"`
	IsShowOrderCondition bool              `json:"isShowOrderCondition"`
	IsShowSaveCondition  bool              `json:"isShowSaveCondition"`
	Groups               []Group           `json:"groups"`
}

/*
QueryCondition is saved query condition data
*/
type QueryCondition struct {
	ID             string              `json:"id"`
	PatternName    string              `json:"patternName"`
	Category       Category            `json:"category"`
	IsDisclose     bool                `json:"isDisclose"`
	DiscloseGroups []Group             `json:"discloseGroups"`
	Owner          Staff               `json:"owner"`
	ConditionData  query.ConditionData `json:"conditionData"`
}
