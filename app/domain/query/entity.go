package query

import (
	"regulus/app/domain/authentication"
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
	DisplayItemList     []FieldAttr             `json:"displayItemList"`
	SearchConditionList []FieldAttr             `json:"searchConditionList"`
	OrderConditionList  []FieldAttr             `json:"orderConditionList"`
	Groups              []*authentication.Group `json:"groups"`
}

/*
Condition is saved query condition data
*/
type Condition struct {
	ID             string                  `json:"id"`
	PatternName    string                  `json:"patternName"`
	Category       *Category               `json:"category"`
	IsDisclose     bool                    `json:"isDisclose"`
	DiscloseGroups []*authentication.Group `json:"discloseGroups"`
	Owner          *authentication.Staff   `json:"owner"`
	ConditionData  ConditionData           `json:"conditionData"`
}
