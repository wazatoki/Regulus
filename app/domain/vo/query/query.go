package query

/*
FieldAttr is query field attribute
*/
type FieldAttr struct {
	ID        string        `json:"id"`
	ViewValue string        `json:"viewValue"`
	FieldType ValueTypeEnum `json:"fieldType"`
}

/*
SearchConditionItem is query condition
*/
type SearchConditionItem struct {
	SearchField    FieldAttr     `json:"searchField"`
	ConditionValue string        `json:"conditionValue"`
	MatchType      MatchTypeEnum `json:"matchType"` // match, unmatch, pertialmatch, gt, ge, le, lt
	Operator       OperatorEnum  `json:"operator"`  // and, or
}

/*
OrderConditionItem is order condition
*/
type OrderConditionItem struct {
	OrderField        FieldAttr     `json:"orderField"`
	OrderFieldKeyWord OrderTypeEnum `json:"orderType"` // asc, desc
}

/*
ConditionData is query condition
*/
type ConditionData struct {
	SearchStrings       []string              `json:"searchStrings"`
	DisplayItemList     []FieldAttr           `json:"displayItemList"`
	SearchConditionList []SearchConditionItem `json:"searchConditionList"`
	OrderConditionList  []OrderConditionItem  `json:"orderConditionList"`
}

/*
ValueTypeEnum is a const type of field value
*/
type ValueTypeEnum string

const (
	// STRING is field value type name
	STRING ValueTypeEnum = "string"

	// NUMBER is field value type name
	NUMBER ValueTypeEnum = "number"

	// BOOLEAN is field value type name
	BOOLEAN ValueTypeEnum = "boolean"

	// ARRAY is field value type name
	ARRAY ValueTypeEnum = "array"
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
	// In is const value of lt
	In MatchTypeEnum = "in"
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
