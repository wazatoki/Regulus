package domain

import (
	"reflect"
	"testing"
)

func createExpectedStaff1Entity() *Staff {
	return &Staff{
		ID:          "staffid1",
		AccountID:   "12345",
		Name:        "name 1",
		Password:    "password 1",
		StaffGroups: createExpectedStaffGroupEntity1Slice(),
	}
}

func createExpectedStaff2Entity() *Staff {
	return &Staff{
		ID:        "staffid2",
		AccountID: "22345",
		Name:      "name 2",
		Password:  "password 2",
		StaffGroups: StaffGroups{
			createExpectedStaffGroup1Entity(),
		},
	}
}

func createExpectedStaffGroup1Entity() *StaffGroup {
	return &StaffGroup{
		ID:   "staffgroupid1",
		Name: "staff group name 1",
	}
}

func createExpectedStaffGroup2Entity() *StaffGroup {
	return &StaffGroup{
		ID:   "staffgroupid2",
		Name: "staff group name 2",
	}
}

func createExpectedStaffGroup3Entity() *StaffGroup {
	return &StaffGroup{
		ID:   "staffgroupid3",
		Name: "staff group name 3",
	}
}

func createExpectedStaffGroupEntity1Slice() StaffGroups {
	return StaffGroups{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
	}
}

func createExpectedStaffGroupEntity2Slice() StaffGroups {
	return StaffGroups{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
		createExpectedStaffGroup3Entity(),
	}
}

var staffSearchConditionList = []FieldAttr{
	{
		ID:        "account-id",
		ViewValue: "利用者ID",
		FieldType: QueryValueTypeEnum.STRING,
	},
	{
		ID:        "name",
		ViewValue: "利用者名称",
		FieldType: QueryValueTypeEnum.STRING,
	},
	{
		ID:        "groups",
		ViewValue: "所属グループ",
		FieldType: QueryValueTypeEnum.ARRAY,
	},
	{
		ID:        "group-name",
		ViewValue: "所属グループ名",
		FieldType: QueryValueTypeEnum.STRING,
	},
}

var staffGroupSearchConditionList = []FieldAttr{
	{
		ID:        "name",
		ViewValue: "グループ名称",
		FieldType: QueryValueTypeEnum.STRING,
	},
	{
		ID:        "staff-name",
		ViewValue: "利用者名称",
		FieldType: QueryValueTypeEnum.STRING,
	},
	{
		ID:        "staff-account-id",
		ViewValue: "利用者ID",
		FieldType: QueryValueTypeEnum.STRING,
	},
}

var queryConditionOrderConditionList = []FieldAttr{
	{
		ID:        "pattern-name",
		ViewValue: "検索パターン名称",
		FieldType: QueryValueTypeEnum.STRING,
	},
	{
		ID:        "category-view-value",
		ViewValue: "カテゴリー名称",
		FieldType: QueryValueTypeEnum.STRING,
	},
	{
		ID:        "is-disclose",
		ViewValue: "公開",
		FieldType: QueryValueTypeEnum.BOOLEAN,
	},
	{
		ID:        "owner",
		ViewValue: "所有者",
		FieldType: QueryValueTypeEnum.STRING,
	},
}

var queryConditionSearchConditionList = []FieldAttr{
	{
		ID:        "pattern-name",
		ViewValue: "検索パターン名称",
		FieldType: QueryValueTypeEnum.STRING,
	},
	{
		ID:        "category-view-value",
		ViewValue: "カテゴリー名称",
		FieldType: QueryValueTypeEnum.STRING,
	},
	{
		ID:        "is-disclose",
		ViewValue: "公開",
		FieldType: QueryValueTypeEnum.BOOLEAN,
	},
	{
		ID:        "disclose-groups",
		ViewValue: "公開先グループ",
		FieldType: QueryValueTypeEnum.ARRAY,
	},
	{
		ID:        "owner",
		ViewValue: "所有者",
		FieldType: QueryValueTypeEnum.STRING,
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
			StaffGroups:         createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "staff-group",
		ViewValue: "利用者グループ",
		SearchItems: ComplexSearchItems{
			SearchConditionList: staffGroupSearchConditionList,
			DisplayItemList:     []FieldAttr{},
			OrderConditionList:  []FieldAttr{},
			StaffGroups:         createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "query-condition",
		ViewValue: "検索条件管理",
		SearchItems: ComplexSearchItems{
			SearchConditionList: queryConditionSearchConditionList,
			DisplayItemList:     []FieldAttr{},
			OrderConditionList:  []FieldAttr{},
			StaffGroups:         nil,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
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
					MatchType:      QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func TestConditions_Sort(t *testing.T) {
	type args struct {
		orderItems []OrderConditionItem
	}
	tests := []struct {
		name       string
		c          *Conditions
		args       args
		wantResult Conditions
	}{
		{
			name: "Sort with order condition",
			c: &Conditions{
				createExpectedQueryCondition0Entity(),
				createExpectedQueryCondition1Entity(),
				createExpectedQueryCondition2Entity(),
				createExpectedQueryCondition3Entity(),
				createExpectedQueryCondition4Entity(),
				createExpectedQueryCondition5Entity(),
				createExpectedQueryCondition6Entity(),
				createExpectedQueryCondition7Entity(),
				createExpectedQueryCondition8Entity(),
				createExpectedQueryCondition9Entity(),
			},
			args: args{
				orderItems: []OrderConditionItem{
					{
						OrderField:        queryConditionOrderConditionList[1],
						OrderFieldKeyWord: QueryOrderTypeEnum.ASC,
					},
					{
						OrderField:        queryConditionOrderConditionList[0],
						OrderFieldKeyWord: QueryOrderTypeEnum.DESC,
					},
				},
			},
			wantResult: Conditions{
				createExpectedQueryCondition4Entity(),
				createExpectedQueryCondition3Entity(),
				createExpectedQueryCondition2Entity(),
				createExpectedQueryCondition1Entity(),
				createExpectedQueryCondition0Entity(),
				createExpectedQueryCondition9Entity(),
				createExpectedQueryCondition8Entity(),
				createExpectedQueryCondition7Entity(),
				createExpectedQueryCondition6Entity(),
				createExpectedQueryCondition5Entity(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Sort(tt.args.orderItems...); !reflect.DeepEqual(got, tt.wantResult) {
				t.Errorf("Conditions.Sort() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}
