package query

// StaffGroupSearchConditionList StaffGroup検索用項目
var StaffGroupSearchConditionList []FieldAttr = []FieldAttr{
	{
		ID:        "name",
		ViewValue: "グループ名称",
		FieldType: STRING,
	},
	{
		ID:        "staff-name",
		ViewValue: "利用者名称",
		FieldType: STRING,
	},
	{
		ID:        "staff-account-id",
		ViewValue: "利用者ID",
		FieldType: STRING,
	},
}
