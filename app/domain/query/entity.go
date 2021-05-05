package query

import (
	"regulus/app/domain"
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
	DisplayItemList     []FieldAttr     `json:"displayItemList"`
	SearchConditionList []FieldAttr     `json:"searchConditionList"`
	OrderConditionList  []FieldAttr     `json:"orderConditionList"`
	Groups              []*domain.Group `json:"groups"`
}

/*
Condition is saved query condition data
*/
type Condition struct {
	ID             string          `json:"id"`
	PatternName    string          `json:"patternName"`
	Category       *Category       `json:"category"`
	IsDisclose     bool            `json:"isDisclose"`
	DiscloseGroups []*domain.Group `json:"discloseGroups"`
	Owner          *domain.Staff   `json:"owner"`
	ConditionData  ConditionData   `json:"conditionData"`
}
