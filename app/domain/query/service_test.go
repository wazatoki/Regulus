package query

import (
	"reflect"
	"regulus/app/domain"
	"testing"
)

func createExpectedStaff1Entity() *domain.Staff {
	return &domain.Staff{
		ID:        "staffid1",
		AccountID: "12345",
		Name:      "name 1",
		Password:  "password 1",
		Groups:    createExpectedStaffGroupEntity1Slice(),
	}
}

func createExpectedStaff2Entity() *domain.Staff {
	return &domain.Staff{
		ID:        "staffid2",
		AccountID: "22345",
		Name:      "name 2",
		Password:  "password 2",
		Groups: []*domain.Group{
			createExpectedStaffGroup1Entity(),
		},
	}
}

func createExpectedStaffGroup1Entity() *domain.Group {
	return &domain.Group{
		ID:   "staffgroupid1",
		Name: "staff group name 1",
	}
}

func createExpectedStaffGroup2Entity() *domain.Group {
	return &domain.Group{
		ID:   "staffgroupid2",
		Name: "staff group name 2",
	}
}

func createExpectedStaffGroup3Entity() *domain.Group {
	return &domain.Group{
		ID:   "staffgroupid3",
		Name: "staff group name 3",
	}
}

func createExpectedStaffGroupEntity1Slice() []*domain.Group {
	return []*domain.Group{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
	}
}

func createExpectedStaffGroupEntity2Slice() []*domain.Group {
	return []*domain.Group{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
		createExpectedStaffGroup3Entity(),
	}
}

var staffSearchConditionList = []FieldAttr{
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

var staffGroupSearchConditionList = []FieldAttr{
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

var queryConditionSearchConditionList = []FieldAttr{
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

var categories = []*Category{
	{
		Name:      "staff",
		ViewValue: "利用者",
		SearchItems: ComplexSearchItems{
			SearchConditionList: staffSearchConditionList,
			DisplayItemList:     []FieldAttr{},
			OrderConditionList:  []FieldAttr{},
			Groups:              createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "staff-group",
		ViewValue: "利用者グループ",
		SearchItems: ComplexSearchItems{
			SearchConditionList: staffGroupSearchConditionList,
			DisplayItemList:     []FieldAttr{},
			OrderConditionList:  []FieldAttr{},
			Groups:              createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "query-condition",
		ViewValue: "検索条件管理",
		SearchItems: ComplexSearchItems{
			SearchConditionList: queryConditionSearchConditionList,
			DisplayItemList:     []FieldAttr{},
			OrderConditionList:  []FieldAttr{},
			Groups:              nil,
		},
	},
}

func createExpectedQueryCondition0Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid0",
		PatternName: "patternName0",
		Category:    categories[0],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "1",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition1Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid1",
		PatternName: "patternName1",
		Category:    categories[0],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "2",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition2Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid2",
		PatternName: "patternName2",
		Category:    categories[0],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "3",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition3Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid3",
		PatternName: "patternName3",
		Category:    categories[0],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "4",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition4Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid4",
		PatternName: "patternName4",
		Category:    categories[0],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "5",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity2Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition5Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid5",
		PatternName: "patternName5",
		Category:    categories[1],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "1",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition6Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid6",
		PatternName: "patternName6",
		Category:    categories[1],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "2",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     true,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition7Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid7",
		PatternName: "patternName7",
		Category:    categories[1],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "3",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition8Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid8",
		PatternName: "patternName8",
		Category:    categories[1],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "4",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition9Entity() *Condition {
	return &Condition{
		ID:          "queryConditionid9",
		PatternName: "patternName9",
		Category:    categories[1],
		ConditionData: ConditionData{
			DisplayItemList: []FieldAttr{},
			SearchConditionList: []SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "5",
					MatchType:      Pertialmatch,
					Operator:       And,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func TestCategoryNameListByMatchType(t *testing.T) {
	type args struct {
		s  string
		mt MatchTypeEnum
	}
	tests := []struct {
		name              string
		args              args
		wantCategoryNames []string
	}{
		{
			name: "match type is match",
			args: args{
				s:  "検索条件管理",
				mt: Match,
			},
			wantCategoryNames: []string{"query-condition"},
		},
		{
			name: "match type is unmatch",
			args: args{
				s:  "検索条件管理",
				mt: Unmatch,
			},
			wantCategoryNames: []string{"staff", "staff-group"},
		},
		{
			name: "match type is pertialmatch",
			args: args{
				s:  "検索条件管",
				mt: Pertialmatch,
			},
			wantCategoryNames: []string{"query-condition"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCategoryNames := CategoryNameListByMatchType(tt.args.s, tt.args.mt); !reflect.DeepEqual(gotCategoryNames, tt.wantCategoryNames) {
				t.Errorf("CategoryNameListByMatchType() = %v, want %v", gotCategoryNames, tt.wantCategoryNames)
			}
		})
	}
}
