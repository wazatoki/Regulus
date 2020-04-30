# 検索機能

- 表示項目の指定ができる。
- 検索条件の設定ができる。
- 検索結果の並び順の指定ができる。
- 表示項目・検索条件・並び替え条件を保存修正できる。


```puml
@startuml search object diagram

title: search object

left to right direction


enum ValueTypeEnum {
	STRING string = "string"
	NUMBER string = "number"
	BOOLEAN string = "boolean"
	ARRAY string = "array"
}

class FieldAttr {
	ID          string
	ViewValue   string
	FieldType   ValueTypeEnum
}

FieldAttr "1" -- "1" ValueTypeEnum : FieldType >

class Group {
	ID   string
	Name string
}

class ComplexSearchItems {
	DisplayItemList      []FieldAttr
	SearchConditionList  []FieldAttr
	OrderConditionList   []FieldAttr
	IsShowDisplayItem    bool
	IsShowOrderCondition bool
	IsShowSaveCondition  bool
	Groups               []Group
}

ComplexSearchItems "1" -- "1..*" FieldAttr : DisplayItemList >
ComplexSearchItems "1" -- "1..*" FieldAttr : SearchConditionList >
ComplexSearchItems "1" -- "1..*" FieldAttr : OrderConditionList >
ComplexSearchItems "1" -- "1..*" Group : GroupList >

class Category {
	Name        string
	ViewValue   string
	SearchItems ComplexSearchItems
}

Category "1" -- "1" ComplexSearchItems : SearchItems >

enum MatchTypeEnum {
	Match string = "match"
	Unmatch string = "unmatch"
	Pertialmatch string = "pertialmatch"
	Gt string = "gt"
	Ge string = "ge"
	Le string = "le"
	Lt string = "lt"
}

enum OperatorEnum {
    And string = "and"
	Or string = "or"
}

class SearchConditionItem {
	SearchField     FieldAttr
	ConditionValue  string
	MatchType       MatchTypeEnum
	Operator        OperatorEnum
}

SearchConditionItem "1" -- "1" FieldAttr : SearchField >
SearchConditionItem "1" -- "1" MatchTypeEnum : MatchType >
SearchConditionItem "1" -- "1" OperatorEnum : Operator >

enum OrderTypeEnum {
    Asc string = "asc"
	Desc string = "desc"
}

class OrderConditionItem {
	OrderField          FieldAttr
	OrderFieldKeyWord   OrderTypeEnum
}

OrderConditionItem "1" -- "1" FieldAttr : OrderField >
OrderConditionItem "1" -- "1" OrderTypeEnum : OrderFieldKeyWord >

class ConditionData {
    SearchStrings       []string
	DisplayItemList     []FieldAttr
    SearchConditionList []SearchConditionItem
	OrderConditionList  []OrderConditionItem
}

ConditionData "1" *-- "1..n" FieldAttr : DisplayItemList >
ConditionData "1" *-- "1..n" SearchConditionItem : SearchConditionList >
ConditionData "1" *-- "1..n" OrderConditionItem : OrderConditionList >

class Staff {
	ID		string
	Name	string
	Groups	[]Group
}

Staff "1" -- "1..*" Group : Groups >

class QueryCondition {
    ID             string
	PatternName    string
	Category       Category
	IsDisclose     bool
	DiscloseGroups []Group
	Owner          Staff
	ConditionData  ConditionData
}

QueryCondition "1" -- "1" Category : Category >
QueryCondition "1" -- "1" ConditionData : ConditionData >
QueryCondition "1" -- "1" Staff : Owner >

@enduml
```