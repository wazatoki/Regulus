package entities

import (
	"regulus/app/domain/authentication"
	"regulus/app/domain/vo/query"
)

/*
Category is struct as search category
*/
type Category struct {
	Name        string             `json:"name"`
	ViewValue   string             `json:"viewValue"`
	SearchItems ComplexSearchItems `json:"searchItems"`
}

/*
ComplexSearchItems is struct as search condition display
*/
type ComplexSearchItems struct {
	DisplayItemList     []query.FieldAttr       `json:"displayItemList"`
	SearchConditionList []query.FieldAttr       `json:"searchConditionList"`
	OrderConditionList  []query.FieldAttr       `json:"orderConditionList"`
	Groups              []*authentication.Group `json:"groups"`
}

/*
QueryCondition is saved query condition data
*/
type QueryCondition struct {
	ID             string                  `json:"id"`
	PatternName    string                  `json:"patternName"`
	Category       *Category               `json:"category"`
	IsDisclose     bool                    `json:"isDisclose"`
	DiscloseGroups []*authentication.Group `json:"discloseGroups"`
	Owner          *authentication.Staff   `json:"owner"`
	ConditionData  query.ConditionData     `json:"conditionData"`
}
