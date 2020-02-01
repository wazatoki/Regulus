package query

/*
Item is query condition
*/
type Item struct {
	EntityName   string  `json:"entityName"`
	FieldName    string  `json:"fieldName"`
	StringValue  string  `json:"stringValue"`
	Int32Value   int32   `json:"int32Value"`
	Int64Value   int64   `json:"int64Value"`
	Float32Value float32 `json:"float32Value"`
	Float64Value float64 `json:"float64Value"`
	ValueType    string  `json:"valueType"` // string int32 int64 float32 float64
	MatchType    string  `json:"matchType"` // match unmatch pertialmatch gt ge le lt
	Operator     string  `json:"operator"`  // not and or
}
