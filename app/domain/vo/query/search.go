package query

// QueryConditionPatternName 検索パターン管理用表示・検索・並び替え用の項目
var QueryConditionPatternName = FieldAttr{
	ID:        "query-condition-pattern-name",
	ViewValue: "検索パターン名称",
	FieldType: STRING,
}

// QueryConditionCategoryViewValue 検索パターン管理用表示・検索・並び替え用の項目
var QueryConditionCategoryViewValue = FieldAttr{
	ID:        "query-condition-category-view-value",
	ViewValue: "カテゴリー名称",
	FieldType: STRING,
}

// QueryConditionIsDisclose 検索パターン管理用表示・検索・並び替え用の項目
var QueryConditionIsDisclose = FieldAttr{
	ID:        "query-condition-is-disclose",
	ViewValue: "公開",
	FieldType: BOOLEAN,
}

// QueryConditionDiscloseGroups 検索パターン管理用表示・検索・並び替え用の項目
var QueryConditionDiscloseGroups = FieldAttr{
	ID:        "query-condition-disclose-groups",
	ViewValue: "公開先グループ",
	FieldType: ARRAY,
}

// QueryConditionOwner 検索パターン管理用表示・検索・並び替え用の項目
var QueryConditionOwner = FieldAttr{
	ID:        "query-condition-owner",
	ViewValue: "所有者",
	FieldType: STRING,
}
