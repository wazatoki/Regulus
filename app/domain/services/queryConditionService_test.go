package services

import (
	"reflect"
	"regulus/app/domain/authentication"
	"regulus/app/domain/entities"
	"regulus/app/domain/query"
	"testing"
)

func createExpectedStaff1Entity() *authentication.Staff {
	return &authentication.Staff{
		ID:        "staffid1",
		AccountID: "12345",
		Name:      "name 1",
		Password:  "password 1",
		Groups:    createExpectedStaffGroupEntity1Slice(),
	}
}

func createExpectedStaff2Entity() *authentication.Staff {
	return &authentication.Staff{
		ID:        "staffid2",
		AccountID: "22345",
		Name:      "name 2",
		Password:  "password 2",
		Groups: []*authentication.Group{
			createExpectedStaffGroup1Entity(),
		},
	}
}

func createExpectedStaffGroup1Entity() *authentication.Group {
	return &authentication.Group{
		ID:   "staffgroupid1",
		Name: "staff group name 1",
	}
}

func createExpectedStaffGroup2Entity() *authentication.Group {
	return &authentication.Group{
		ID:   "staffgroupid2",
		Name: "staff group name 2",
	}
}

func createExpectedStaffGroup3Entity() *authentication.Group {
	return &authentication.Group{
		ID:   "staffgroupid3",
		Name: "staff group name 3",
	}
}

func createExpectedStaffGroupEntity1Slice() []*authentication.Group {
	return []*authentication.Group{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
	}
}

func createExpectedStaffGroupEntity2Slice() []*authentication.Group {
	return []*authentication.Group{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
		createExpectedStaffGroup3Entity(),
	}
}

var staffSearchConditionList = []query.FieldAttr{
	{
		ID:        "account-id",
		ViewValue: "利用者ID",
		FieldType: query.STRING,
	},
	{
		ID:        "name",
		ViewValue: "利用者名称",
		FieldType: query.STRING,
	},
	{
		ID:        "groups",
		ViewValue: "所属グループ",
		FieldType: query.ARRAY,
	},
	{
		ID:        "group-name",
		ViewValue: "所属グループ名",
		FieldType: query.STRING,
	},
}

var staffGroupSearchConditionList = []query.FieldAttr{
	{
		ID:        "name",
		ViewValue: "グループ名称",
		FieldType: query.STRING,
	},
	{
		ID:        "staff-name",
		ViewValue: "利用者名称",
		FieldType: query.STRING,
	},
	{
		ID:        "staff-account-id",
		ViewValue: "利用者ID",
		FieldType: query.STRING,
	},
}

var queryConditionSearchConditionList = []query.FieldAttr{
	{
		ID:        "pattern-name",
		ViewValue: "検索パターン名称",
		FieldType: query.STRING,
	},
	{
		ID:        "category-view-value",
		ViewValue: "カテゴリー名称",
		FieldType: query.STRING,
	},
	{
		ID:        "is-disclose",
		ViewValue: "公開",
		FieldType: query.BOOLEAN,
	},
	{
		ID:        "disclose-groups",
		ViewValue: "公開先グループ",
		FieldType: query.ARRAY,
	},
	{
		ID:        "owner",
		ViewValue: "所有者",
		FieldType: query.STRING,
	},
}

var categories = []*entities.Category{
	{
		Name:      "staff",
		ViewValue: "利用者",
		SearchItems: entities.ComplexSearchItems{
			SearchConditionList: staffSearchConditionList,
			DisplayItemList:     []query.FieldAttr{},
			OrderConditionList:  []query.FieldAttr{},
			Groups:              createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "staff-group",
		ViewValue: "利用者グループ",
		SearchItems: entities.ComplexSearchItems{
			SearchConditionList: staffGroupSearchConditionList,
			DisplayItemList:     []query.FieldAttr{},
			OrderConditionList:  []query.FieldAttr{},
			Groups:              createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "query-condition",
		ViewValue: "検索条件管理",
		SearchItems: entities.ComplexSearchItems{
			SearchConditionList: queryConditionSearchConditionList,
			DisplayItemList:     []query.FieldAttr{},
			OrderConditionList:  []query.FieldAttr{},
			Groups:              nil,
		},
	},
}

func createExpectedQueryCondition0Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid0",
		PatternName: "patternName0",
		Category:    categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "1",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition1Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid1",
		PatternName: "patternName1",
		Category:    categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "2",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition2Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid2",
		PatternName: "patternName2",
		Category:    categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "3",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition3Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid3",
		PatternName: "patternName3",
		Category:    categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "4",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition4Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid4",
		PatternName: "patternName4",
		Category:    categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "5",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity2Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition5Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid5",
		PatternName: "patternName5",
		Category:    categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "1",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition6Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid6",
		PatternName: "patternName6",
		Category:    categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "2",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     true,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition7Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid7",
		PatternName: "patternName7",
		Category:    categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "3",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition8Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid8",
		PatternName: "patternName8",
		Category:    categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "4",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition9Entity() *entities.QueryCondition {
	return &entities.QueryCondition{
		ID:          "queryConditionid9",
		PatternName: "patternName9",
		Category:    categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "5",
					MatchType:      query.Pertialmatch,
					Operator:       query.And,
				},
			},
			OrderConditionList: []query.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func TestQueryConditions_Find(t *testing.T) {
	type args struct {
		queryItems []query.SearchConditionItem
	}
	tests := []struct {
		name       string
		q          QueryConditions
		args       args
		wantResult QueryConditions
	}{
		{
			name: "Find with condition",
			q: QueryConditions{
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
				queryItems: []query.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[0],
						ConditionValue: "5",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
					{
						SearchField:    queryConditionSearchConditionList[1],
						ConditionValue: "利用者グ",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
				},
			},
			wantResult: QueryConditions{
				createExpectedQueryCondition5Entity(),
			},
		},
		{
			name: "Find with boole condition",
			q: QueryConditions{
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
				queryItems: []query.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[2],
						ConditionValue: "true",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
				},
			},
			wantResult: QueryConditions{
				createExpectedQueryCondition6Entity(),
			},
		},
		{
			name: "Find with array, or condition",
			q: QueryConditions{
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
				queryItems: []query.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[3],
						ConditionValue: "staffgroupid3",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
					{
						SearchField:    queryConditionSearchConditionList[0],
						ConditionValue: "5",
						MatchType:      query.Pertialmatch,
						Operator:       query.Or,
					},
				},
			},
			wantResult: QueryConditions{
				createExpectedQueryCondition4Entity(),
				createExpectedQueryCondition5Entity(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.q.Find(tt.args.queryItems...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("QueryConditions.Find() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
