package query

import (
	"regulus/app/domain/entities/group"
)

/*
ComplexSearchItems is struct as search condition display
*/
type ComplexSearchItems struct {
	DisplayItemList      []FieldAttr   `json:"displayItemList"`
	SearchConditionList  []FieldAttr   `json:"searchConditionList"`
	OrderConditionList   []FieldAttr   `json:"orderConditionList"`
	IsShowDisplayItem    bool          `json:"isShowDisplayItem"`
	IsShowOrderCondition bool          `json:"isShowOrderCondition"`
	IsShowSaveCondition  bool          `json:"isShowSaveCondition"`
	GroupList            []group.Group `json:"groupList"`
}

/*
FieldAttr is query field attribute
*/
type FieldAttr struct {
	ID         string        `json:"id"`
	EntityName EntityEnum    `json:"entityName"`
	FieldName  FieldEnum     `json:"fieldName"`
	ViewValue  string        `json:"viewValue"`
	FieldType  ValueTypeEnum `json:"fieldType"`
}

/*
ConditionItem is query condition
*/
type ConditionItem struct {
	Field          FieldAttr     `json:"field"`
	ConditionValue string        `json:"conditionValue"`
	MatchType      MatchTypeEnum `json:"matchType"` // match, unmatch, pertialmatch, gt, ge, le, lt
	Operator       OperatorEnum  `json:"operator"`  // and, or
}

/*
OrderItem is order condition
*/
type OrderItem struct {
	OrderField FieldAttr     `json:"orderField"`
	OrderType  OrderTypeEnum `json:"orderType"` // asc, desc
}

/*
ConditionData is query condition
*/
type ConditionData struct {
	SearchStrings       []string        `json:"searchStrings"`
	DisplayItemList     []FieldAttr     `json:"displayItemList"`
	SearchConditionList []ConditionItem `json:"searchConditionList"`
	OrderConditionList  []OrderItem     `json:"orderConditionList"`
}

/*
SaveData is save query save data
*/
type SaveData struct {
	ID             string        `json:"id"`
	PatternName    string        `json:"patternName"`
	Category       string        `json:"category"`
	IsDisclose     bool          `json:"isDisclose"`
	DiscloseGroups []string      `json:"discloseGroups"`
	OwnerID        string        `json:"ownerID"`
	ConditionData  ConditionData `json:"conditionData"`
}

/*
EntityEnum is entity name by snake case
*/
type EntityEnum string

/*
FieldEnum is field name by snake case
*/
type FieldEnum string

/*
ValueTypeEnum is a const type of field value
*/
type ValueTypeEnum string

const (
	// String is field value type name
	String ValueTypeEnum = "string"
	// Number is field value type name
	Number ValueTypeEnum = "number"
)

/*
MatchTypeEnum is a const type of match type
*/
type MatchTypeEnum string

const (
	// Match is const value of match
	Match MatchTypeEnum = "match"
	// Unmatch is const value of unmatch
	Unmatch MatchTypeEnum = "unmatch"
	// Pertialmatch is const value of pertialmatch
	Pertialmatch MatchTypeEnum = "pertialmatch"
	// Gt is const value of gt
	Gt MatchTypeEnum = "gt"
	// Ge is const value of ge
	Ge MatchTypeEnum = "ge"
	// Le is const value of le
	Le MatchTypeEnum = "le"
	// Lt is const value of lt
	Lt MatchTypeEnum = "lt"
)

/*
OperatorEnum is a const type of Operator
*/
type OperatorEnum string

const (
	// And is const value of  and
	And OperatorEnum = "and"
	// Or is const value of or
	Or OperatorEnum = "or"
)

/*
OrderTypeEnum is a const type of OrderType
*/
type OrderTypeEnum string

const (
	// Asc is const value of  asc
	Asc OrderTypeEnum = "asc"
	// Desc is const value of desc
	Desc OrderTypeEnum = "desc"
)
