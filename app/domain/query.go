package domain

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
	DisplayItemList     []FieldAttr `json:"displayItemList"`
	SearchConditionList []FieldAttr `json:"searchConditionList"`
	OrderConditionList  []FieldAttr `json:"orderConditionList"`
	Groups              []*Group    `json:"groups"`
}

/*
Condition is saved query condition data
*/
type Condition struct {
	ID             string        `json:"id"`
	PatternName    string        `json:"patternName"`
	Category       *Category     `json:"category"`
	IsDisclose     bool          `json:"isDisclose"`
	DiscloseGroups []*Group      `json:"discloseGroups"`
	Owner          *Staff        `json:"owner"`
	ConditionData  ConditionData `json:"conditionData"`
}

/*
FieldAttr is query field attribute
*/
type FieldAttr struct {
	ID          string         `json:"id"`
	ViewValue   string         `json:"viewValue"`
	FieldType   QueryValueType `json:"fieldType"`
	OptionItems []OptionItem   `json:"optionItems"`
}

/*
OptionItem is option for select-ish tag select or radio, checkbox
*/
type OptionItem struct {
	ID        string `json:"id"`
	ViewValue string `json:"viewValue"`
}

/*
SearchConditionItem is query condition
*/
type SearchConditionItem struct {
	SearchField    FieldAttr      `json:"searchField"`
	ConditionValue string         `json:"conditionValue"`
	MatchType      QueryMatchType `json:"matchType"` // match, unmatch, pertialmatch, gt, ge, le, lt, in
	Operator       QueryOperator  `json:"operator"`  // and, or
}

/*
OrderConditionItem is order condition
*/
type OrderConditionItem struct {
	OrderField        FieldAttr      `json:"orderField"`
	OrderFieldKeyWord QueryOrderType `json:"orderType"` // asc, desc
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
QueryValueType is a const type of field value
*/
type QueryValueType struct{ Value string }

func (s QueryValueType) String() string {
	return s.Value
}

var QueryValueTypeEnum = struct {
	STRING  QueryValueType
	NUMBER  QueryValueType
	BOOLEAN QueryValueType
	ARRAY   QueryValueType
}{
	STRING:  QueryValueType{"string"},
	NUMBER:  QueryValueType{"number"},
	BOOLEAN: QueryValueType{"boolean"},
	ARRAY:   QueryValueType{"array"},
}

/*
QueryMatchType is a const type of match type
*/
type QueryMatchType struct{ Value string }

func (s QueryMatchType) String() string {
	return s.Value
}

func (m QueryMatchType) StrToEnum(s string) QueryMatchType {
	switch s {
	case string(QueryMatchTypeEnum.MATCH.Value):
		return QueryMatchTypeEnum.MATCH
	case string(QueryMatchTypeEnum.UNMATCH.Value):
		return QueryMatchTypeEnum.UNMATCH
	case string(QueryMatchTypeEnum.PERTIALMATCH.Value):
		return QueryMatchTypeEnum.PERTIALMATCH
	case string(QueryMatchTypeEnum.GT.Value):
		return QueryMatchTypeEnum.GT
	case string(QueryMatchTypeEnum.GE.Value):
		return QueryMatchTypeEnum.GE
	case string(QueryMatchTypeEnum.LE.Value):
		return QueryMatchTypeEnum.LE
	case string(QueryMatchTypeEnum.LT.Value):
		return QueryMatchTypeEnum.LT
	case string(QueryMatchTypeEnum.IN.Value):
		return QueryMatchTypeEnum.IN
	default:
		return QueryMatchTypeEnum.MATCH
	}
}

var QueryMatchTypeEnum = struct {
	MATCH        QueryMatchType
	UNMATCH      QueryMatchType
	PERTIALMATCH QueryMatchType
	GT           QueryMatchType
	GE           QueryMatchType
	LE           QueryMatchType
	LT           QueryMatchType
	IN           QueryMatchType
}{
	MATCH:        QueryMatchType{"match"},
	UNMATCH:      QueryMatchType{"unmatch"},
	PERTIALMATCH: QueryMatchType{"pertialmatch"},
	GT:           QueryMatchType{"gt"},
	GE:           QueryMatchType{"ge"},
	LE:           QueryMatchType{"le"},
	LT:           QueryMatchType{"lt"},
	IN:           QueryMatchType{"in"},
}

/*
QueryOperator is a const type of Operator
*/
type QueryOperator struct{ Value string }

func (s QueryOperator) String() string {
	return s.Value
}

func (o QueryOperator) StrToEnum(s string) QueryOperator {
	switch s {
	case string(QueryOperatorEnum.AND.Value):
		return QueryOperatorEnum.AND
	case string(QueryOperatorEnum.OR.Value):
		return QueryOperatorEnum.OR
	default:
		return QueryOperatorEnum.AND
	}
}

var QueryOperatorEnum = struct {
	AND QueryOperator
	OR  QueryOperator
}{
	AND: QueryOperator{"and"},
	OR:  QueryOperator{"or"},
}

/*
QueryOrderType is a const type of Operator
*/
type QueryOrderType struct{ Value string }

func (s QueryOrderType) String() string {
	return s.Value
}

func (o QueryOrderType) StrToEnum(s string) QueryOrderType {
	switch s {
	case string(QueryOrderTypeEnum.ASC.Value):
		return QueryOrderTypeEnum.ASC
	case string(QueryOrderTypeEnum.DESC.Value):
		return QueryOrderTypeEnum.DESC
	default:
		return QueryOrderTypeEnum.ASC
	}
}

var QueryOrderTypeEnum = struct {
	ASC  QueryOrderType
	DESC QueryOrderType
}{
	ASC:  QueryOrderType{"asc"},
	DESC: QueryOrderType{"desc"},
}
