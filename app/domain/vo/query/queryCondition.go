package query

// QueryConditionSearchConditionList QueryCondition検索用項目
var QueryConditionSearchConditionList []FieldAttr = []FieldAttr{
	{
		ID:        "pattern-name",
		ViewValue: "検索パターン名称",
		FieldType: STRING,
	},
	{
		ID:        "category-view-value",
		ViewValue: "カテゴリー名称",
		FieldType: STRING,
	},
	{
		ID:        "is-disclose",
		ViewValue: "公開",
		FieldType: BOOLEAN,
	},
	{
		ID:        "disclose-groups",
		ViewValue: "公開先グループ",
		FieldType: ARRAY,
	},
	{
		ID:        "owner",
		ViewValue: "所有者",
		FieldType: STRING,
	},
}
