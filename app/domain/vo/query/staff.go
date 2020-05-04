package query

// StaffSearchConditionList Staff検索用項目
var StaffSearchConditionList []FieldAttr = []FieldAttr{
	{
		ID:        "account-id",
		ViewValue: "利用者ID",
		FieldType: STRING,
	},
	{
		ID:        "name",
		ViewValue: "利用者名称",
		FieldType: STRING,
	},
	{
		ID:        "groups",
		ViewValue: "所属グループ",
		FieldType: ARRAY,
	},
	{
		ID:        "group-name",
		ViewValue: "所属グループ名",
		FieldType: STRING,
	},
}
