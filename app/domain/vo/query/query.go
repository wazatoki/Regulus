package query

/*
ConditionItem is query condition
*/
type ConditionItem struct {
	EntityName EntityEnum    `json:"entityName"`
	FieldName  FieldEnum     `json:"fieldName"`
	Value      string        `json:"stringValue"`
	ValueType  ValueTypeEnum `json:"valueType"` // string, number
	MatchType  MatchTypeEnum `json:"matchType"` // match, unmatch, pertialmatch, gt, ge, le, lt
	Operator   OperatorEnum  `json:"operator"`  // and, or
}

/*
OrderItem is order condition
*/
type OrderItem struct {
	EntityName EntityEnum    `json:"entityName"`
	FieldName  FieldEnum     `json:"fieldName"`
	OrderType  OrderTypeEnum `json:"orderType"` // asc, desc
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
