package query

/*
StaffAccountID Staff 表示・検索・並び替え用の項目
*/
var StaffAccountID FieldAttr = FieldAttr{
	ID:        "staff-account-id",
	ViewValue: "職員ID",
	FieldType: STRING,
}

/*
StaffName Staff 表示・検索・並び替え用の項目
*/
var StaffName FieldAttr = FieldAttr{
	ID:        "staff-name",
	ViewValue: "ユーザー名称",
	FieldType: STRING,
}

/*
StaffGroups Staff 検索用の項目
*/
var StaffGroups FieldAttr = FieldAttr{
	ID:        "staff-groups",
	ViewValue: "所属グループ",
	FieldType: ARRAY,
}

/*
StaffGroupNames Staff表示用の項目
*/
var StaffGroupNames FieldAttr = FieldAttr{
	ID:        "staff-group-names",
	ViewValue: "所属グループ名",
	FieldType: STRING,
}
