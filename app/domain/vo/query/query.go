package query

/*
Item is query condition
*/
type Item struct {
	EntityName string `json:"entityName"`
	FieldName  string `json:"fieldName"`
	Value      string `json:"stringValue"`
	ValueType  string `json:"valueType"` // string number
	MatchType  string `json:"matchType"` // match unmatch pertialmatch gt ge le lt
	Operator   string `json:"operator"`  // and or
}
