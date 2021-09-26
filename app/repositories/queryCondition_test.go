package repositories

import (
	"context"
	"reflect"
	"regulus/app/domain"
	"regulus/app/infrastructures/sqlboiler"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func setUpQueryConditionTest() *sqlx.DB {
	db := createDB()
	con, _ := db.Open()

	return con
}

func tearDownQueryConditionTest(con *sqlx.DB) {
	con.Close()
}

var staffSearchConditionList = []domain.FieldAttr{
	{
		ID:        "account-id",
		ViewValue: "利用者ID",
		FieldType: domain.QueryValueTypeEnum.STRING,
	},
	{
		ID:        "name",
		ViewValue: "利用者名称",
		FieldType: domain.QueryValueTypeEnum.STRING,
	},
	{
		ID:        "groups",
		ViewValue: "所属グループ",
		FieldType: domain.QueryValueTypeEnum.ARRAY,
	},
	{
		ID:        "group-name",
		ViewValue: "所属グループ名",
		FieldType: domain.QueryValueTypeEnum.STRING,
	},
}

var staffGroupSearchConditionList = []domain.FieldAttr{
	{
		ID:        "name",
		ViewValue: "グループ名称",
		FieldType: domain.QueryValueTypeEnum.STRING,
	},
	{
		ID:        "staff-name",
		ViewValue: "利用者名称",
		FieldType: domain.QueryValueTypeEnum.STRING,
	},
	{
		ID:        "staff-account-id",
		ViewValue: "利用者ID",
		FieldType: domain.QueryValueTypeEnum.STRING,
	},
}

var queryConditionSearchConditionList = []domain.FieldAttr{
	{
		ID:        "pattern-name",
		ViewValue: "検索パターン名称",
		FieldType: domain.QueryValueTypeEnum.STRING,
	},
	{
		ID:        "category-view-value",
		ViewValue: "カテゴリー名称",
		FieldType: domain.QueryValueTypeEnum.STRING,
	},
	{
		ID:        "is-disclose",
		ViewValue: "公開",
		FieldType: domain.QueryValueTypeEnum.BOOLEAN,
	},
	{
		ID:        "disclose-groups",
		ViewValue: "公開先グループ",
		FieldType: domain.QueryValueTypeEnum.ARRAY,
		OptionItems: []domain.OptionItem{
			{ID: "staffgroupid1", ViewValue: "staff group name 1"},
			{ID: "staffgroupid2", ViewValue: "staff group name 2"},
			{ID: "staffgroupid3", ViewValue: "staff group name 3"},
		},
	},
	{
		ID:        "owner",
		ViewValue: "所有者",
		FieldType: domain.QueryValueTypeEnum.STRING,
	},
}

var categories = []*domain.Category{
	{
		Name:      "staff",
		ViewValue: "利用者",
		SearchItems: domain.ComplexSearchItems{
			SearchConditionList: staffSearchConditionList,
			DisplayItemList:     []domain.FieldAttr{},
			OrderConditionList:  []domain.FieldAttr{},
			StaffGroups:         createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "staff-group",
		ViewValue: "利用者グループ",
		SearchItems: domain.ComplexSearchItems{
			SearchConditionList: staffGroupSearchConditionList,
			DisplayItemList:     []domain.FieldAttr{},
			OrderConditionList:  []domain.FieldAttr{},
			StaffGroups:         createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "query-condition",
		ViewValue: "検索条件管理",
		SearchItems: domain.ComplexSearchItems{
			SearchConditionList: queryConditionSearchConditionList,
			DisplayItemList:     []domain.FieldAttr{},
			OrderConditionList: []domain.FieldAttr{{
				ID:        "pattern-name",
				ViewValue: "検索パターン名称",
				FieldType: domain.QueryValueTypeEnum.STRING,
			},
				{
					ID:        "category-view-value",
					ViewValue: "カテゴリー名称",
					FieldType: domain.QueryValueTypeEnum.STRING,
				},
				{
					ID:        "is-disclose",
					ViewValue: "公開",
					FieldType: domain.QueryValueTypeEnum.BOOLEAN,
				},
				{
					ID:        "owner",
					ViewValue: "所有者",
					FieldType: domain.QueryValueTypeEnum.STRING,
				}},
			StaffGroups: createExpectedStaffGroupEntity2Slice(),
		},
	},
}

func createExpectedQueryCondition0Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid0",
		PatternName: "patternName0",
		Category:    categories[0],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "1",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition1Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid1",
		PatternName: "patternName1",
		Category:    categories[0],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "2",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition2Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid2",
		PatternName: "patternName2",
		Category:    categories[0],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "3",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition3Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid3",
		PatternName: "patternName3",
		Category:    categories[0],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "4",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition4Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid4",
		PatternName: "patternName4",
		Category:    categories[0],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffSearchConditionList[0],
					ConditionValue: "5",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition5Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid5",
		PatternName: "patternName5",
		Category:    categories[1],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "1",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition6Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid6",
		PatternName: "patternName6",
		Category:    categories[1],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "2",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition7Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid7",
		PatternName: "patternName7",
		Category:    categories[1],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "3",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition8Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid8",
		PatternName: "patternName8",
		Category:    categories[1],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "4",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     true,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition9Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid9",
		PatternName: "patternName9",
		Category:    categories[1],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    staffGroupSearchConditionList[0],
					ConditionValue: "5",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: createExpectedStaffGroupEntity2Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition10Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid10",
		PatternName: "patternName10",
		Category:    categories[2],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    queryConditionSearchConditionList[0],
					ConditionValue: "10",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: nil,
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition11Entity() *domain.Condition {
	return &domain.Condition{
		ID:          "queryConditionid11",
		PatternName: "patternName11",
		Category:    categories[2],
		ConditionData: domain.ConditionData{
			DisplayItemList: []domain.FieldAttr{},
			SearchConditionList: []domain.SearchConditionItem{
				{
					SearchField:    queryConditionSearchConditionList[1],
					ConditionValue: "検索",
					MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
					Operator:       domain.QueryOperatorEnum.AND,
				},
			},
			OrderConditionList: []domain.OrderConditionItem{},
		},
		DiscloseGroups: nil,
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func TestQueryConditionObjectMap(t *testing.T) {
	type args struct {
		sqc *sqlboiler.QueryCondition
	}
	tests := []struct {
		name    string
		args    args
		wantEqc *domain.Condition
	}{
		{
			name:    "convert sqlboiler.QueryCondition to entities.QueryConditon",
			wantEqc: createExpectedQueryCondition1Entity(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			con := setUpQueryConditionTest()
			defer tearDownQueryConditionTest(con)
			setupTestData()

			tt.args.sqc, _ = sqlboiler.QueryConditions(
				qm.Where("id=?", "queryConditionid1"),
				qm.Load(qm.Rels(sqlboiler.QueryConditionRels.Owner, sqlboiler.StaffRels.StaffGroups), qm.Where("del != true")),
				qm.Load(sqlboiler.QueryConditionRels.QueryDisplayItems),
				qm.Load(sqlboiler.QueryConditionRels.QueryOrderConditionItems),
				qm.Load(sqlboiler.QueryConditionRels.QuerySearchConditionItems),
				qm.Load(sqlboiler.QueryConditionRels.StaffGroups),
			).One(context.Background(), con)

			gotEqc := QueryConditionObjectMap(tt.args.sqc)

			if diff := cmp.Diff(gotEqc, tt.wantEqc); diff != "" {
				t.Errorf("differs = %s", diff)
			}
		})
	}
}

func TestQueryConditionRepo_SelectAll(t *testing.T) {
	type fields struct {
		database db
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "it should get 10entities as select all ",
			fields: fields{
				database: createDB(),
			},
			want:    12,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			con := setUpQueryConditionTest()
			defer tearDownQueryConditionTest(con)
			setupTestData()

			q := &QueryConditionRepo{
				database: tt.fields.database,
			}
			gotQueryConditions, err := q.SelectAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryConditionRepo.SelectAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotQueryConditions) != tt.want {
				t.Errorf("want = %v, got length = %v", tt.want, len(gotQueryConditions))
			}
		})
	}
}

func TestQueryConditionRepo_Select(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		queryItems []domain.SearchConditionItem
	}
	tests := []struct {
		name                      string
		fields                    fields
		args                      args
		wantResultQueryConditions domain.Conditions
		wantErr                   bool
	}{
		{
			name: "should get sprcified entity as select by condition",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []domain.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[0],
						ConditionValue: "5",
						MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
						Operator:       domain.QueryOperatorEnum.AND,
					},
					{
						SearchField:    queryConditionSearchConditionList[1],
						ConditionValue: "利用者",
						MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
						Operator:       domain.QueryOperatorEnum.AND,
					},
				},
			},
			wantResultQueryConditions: domain.Conditions{
				createExpectedQueryCondition5Entity(),
			},
			wantErr: false,
		},
		{
			name: "should get sprcified entity as select by owner",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []domain.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[2],
						ConditionValue: "true",
						MatchType:      domain.QueryMatchTypeEnum.MATCH,
						Operator:       domain.QueryOperatorEnum.AND,
					},
				},
			},
			wantResultQueryConditions: domain.Conditions{
				createExpectedQueryCondition8Entity(),
			},
			wantErr: false,
		},
		{
			name: "should get sprcified entity as select by disclose groups",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []domain.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[3],
						ConditionValue: "[\"staffgroupid3\"]",
						MatchType:      domain.QueryMatchTypeEnum.MATCH,
						Operator:       domain.QueryOperatorEnum.AND,
					},
				},
			},
			wantResultQueryConditions: domain.Conditions{
				createExpectedQueryCondition9Entity(),
			},
			wantErr: false,
		},
		{
			name: "should get sprcified entity as select by owner",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []domain.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[4],
						ConditionValue: "2",
						MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
						Operator:       domain.QueryOperatorEnum.AND,
					},
					{
						SearchField:    queryConditionSearchConditionList[0],
						ConditionValue: "5",
						MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
						Operator:       domain.QueryOperatorEnum.AND,
					},
				},
			},
			wantResultQueryConditions: domain.Conditions{
				createExpectedQueryCondition5Entity(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			con := setUpQueryConditionTest()
			defer tearDownQueryConditionTest(con)
			setupTestData()

			q := &QueryConditionRepo{
				database: tt.fields.database,
			}

			gotResultQueryConditions, err := q.Select(tt.args.queryItems...)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryConditionRepo.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(gotResultQueryConditions, tt.wantResultQueryConditions); diff != "" {
				t.Errorf("differs = %s", diff)
			}
		})
	}
}

func TestQueryConditionRepo_SelectByID(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		id string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantQueryCondition *domain.Condition
		wantErr            bool
	}{
		{
			name: "should get sprcified entity as select by id",
			fields: fields{
				database: createDB(),
			},
			args: args{
				id: "queryConditionid0",
			},
			wantQueryCondition: createExpectedQueryCondition0Entity(),
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			con := setUpQueryConditionTest()
			defer tearDownQueryConditionTest(con)
			setupTestData()

			q := &QueryConditionRepo{
				database: tt.fields.database,
			}
			gotQueryCondition, err := q.SelectByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryConditionRepo.SelectByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotQueryCondition, tt.wantQueryCondition) {
				t.Errorf("QueryConditionRepo.SelectByID() = %v, want %v", gotQueryCondition, tt.wantQueryCondition)
			}
		})
	}
}

func TestQueryConditionRepo_SelectByIDs(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		ids []string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		wantQueryConditions domain.Conditions
		wantErr             bool
	}{
		{
			name: "should get sprcified entity as select by ids",
			fields: fields{
				database: createDB(),
			},
			args: args{
				ids: []string{
					"queryConditionid0",
					"queryConditionid1",
				},
			},
			wantQueryConditions: domain.Conditions{
				createExpectedQueryCondition0Entity(),
				createExpectedQueryCondition1Entity(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			con := setUpQueryConditionTest()
			defer tearDownQueryConditionTest(con)
			setupTestData()

			q := &QueryConditionRepo{
				database: tt.fields.database,
			}
			gotQueryConditions, err := q.SelectByIDs(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryConditionRepo.SelectByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotQueryConditions, tt.wantQueryConditions) {
				t.Errorf("QueryConditionRepo.SelectByIDs() = %v, want %v", gotQueryConditions, tt.wantQueryConditions)
			}
		})
	}
}

func TestQueryConditionRepo_Insert(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		queryCondition domain.Condition
		operatorID     string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantEntity *domain.Condition
		wantErr    bool
	}{
		{
			name: "should insert to database as called insert",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryCondition: domain.Condition{
					PatternName:    "testPatternName",
					Category:       categories[0],
					IsDisclose:     true,
					DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
					ConditionData: domain.ConditionData{
						DisplayItemList: []domain.FieldAttr{},
						SearchConditionList: []domain.SearchConditionItem{
							{
								SearchField:    staffSearchConditionList[0],
								ConditionValue: "aaa",
								MatchType:      domain.QueryMatchTypeEnum.MATCH,
								Operator:       domain.QueryOperatorEnum.AND,
							},
							{
								SearchField:    staffSearchConditionList[1],
								ConditionValue: "bbb",
								MatchType:      domain.QueryMatchTypeEnum.UNMATCH,
								Operator:       domain.QueryOperatorEnum.AND,
							},
						},
						OrderConditionList: []domain.OrderConditionItem{},
					},
					Owner: createExpectedStaff1Entity(),
				},
				operatorID: createExpectedStaff1Entity().ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			con := setUpQueryConditionTest()
			defer tearDownQueryConditionTest(con)
			setupTestData()
			tt.wantEntity = &domain.Condition{
				PatternName:    tt.args.queryCondition.PatternName,
				Category:       tt.args.queryCondition.Category,
				IsDisclose:     tt.args.queryCondition.IsDisclose,
				ConditionData:  tt.args.queryCondition.ConditionData,
				DiscloseGroups: tt.args.queryCondition.DiscloseGroups,
				Owner:          tt.args.queryCondition.Owner,
			}

			q := &QueryConditionRepo{
				database: tt.fields.database,
			}
			gotID, err := q.Insert(&tt.args.queryCondition, tt.args.operatorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryConditionRepo.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.wantEntity.ID = gotID
			got, _ := sqlboiler.QueryConditions(qm.Where("id=?", tt.wantEntity.ID),
				qm.And("del != ?", true),
				qm.Load(qm.Rels(sqlboiler.QueryConditionRels.Owner, sqlboiler.StaffRels.StaffGroups), qm.Where("del != true")),
				qm.Load(sqlboiler.QueryConditionRels.QueryDisplayItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryDisplayItemColumns.RowOrder)),
				qm.Load(sqlboiler.QueryConditionRels.QueryOrderConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryOrderConditionItemColumns.RowOrder)),
				qm.Load(sqlboiler.QueryConditionRels.QuerySearchConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QuerySearchConditionItemColumns.RowOrder)),
				qm.Load(sqlboiler.QueryConditionRels.StaffGroups, qm.Where("del != true")),
			).One(context.Background(), con)
			resultEntity := QueryConditionObjectMap(got)

			if diff := cmp.Diff(resultEntity, tt.wantEntity); diff != "" {
				t.Errorf("differs = %s", diff)
			}
		})
	}
}

func TestQueryConditionRepo_Update(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		queryCondition *domain.Condition
		operatorID     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should update to database as called update",
			fields: fields{
				database: createDB(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpQueryConditionTest()
			defer tearDownQueryConditionTest(con)
			setupTestData()

			beforeQueryCondition := createExpectedQueryCondition0Entity()
			beforeQueryCondition.PatternName = "changed pattern name"
			beforeQueryCondition.ConditionData = domain.ConditionData{
				DisplayItemList: []domain.FieldAttr{},
				SearchConditionList: []domain.SearchConditionItem{
					{
						SearchField:    staffSearchConditionList[1],
						ConditionValue: "2",
						MatchType:      domain.QueryMatchTypeEnum.UNMATCH,
						Operator:       domain.QueryOperatorEnum.AND,
					},
				},
				OrderConditionList: []domain.OrderConditionItem{},
			}
			beforeQueryCondition.IsDisclose = true
			tt.args.queryCondition = beforeQueryCondition
			tt.args.operatorID = createExpectedStaff1Entity().ID

			q := &QueryConditionRepo{
				database: tt.fields.database,
			}
			if err := q.Update(tt.args.queryCondition, tt.args.operatorID); (err != nil) != tt.wantErr {
				t.Errorf("QueryConditionRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, _ := sqlboiler.QueryConditions(qm.Where("id=?", beforeQueryCondition.ID),
				qm.And("del != ?", true),
				qm.Load(qm.Rels(sqlboiler.QueryConditionRels.Owner, sqlboiler.StaffRels.StaffGroups), qm.Where("del != true")),
				qm.Load(sqlboiler.QueryConditionRels.QueryDisplayItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryDisplayItemColumns.RowOrder)),
				qm.Load(sqlboiler.QueryConditionRels.QueryOrderConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryOrderConditionItemColumns.RowOrder)),
				qm.Load(sqlboiler.QueryConditionRels.QuerySearchConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QuerySearchConditionItemColumns.RowOrder)),
				qm.Load(sqlboiler.QueryConditionRels.StaffGroups, qm.Where("del != true")),
			).One(context.Background(), con)
			resultEntity := QueryConditionObjectMap(got)

			if diff := cmp.Diff(resultEntity, beforeQueryCondition); diff != "" {
				t.Errorf("differs = %s", diff)
			}
		})
	}
}

func TestQueryConditionRepo_Dalete(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		id         string
		operatorID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should return nil as called delete",
			fields: fields{
				database: createDB(),
			},
			args: args{
				id:         "queryConditionid1",
				operatorID: createExpectedStaff1Entity().ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpQueryConditionTest()
			defer tearDownQueryConditionTest(con)
			setupTestData()

			q := &QueryConditionRepo{
				database: tt.fields.database,
			}
			if err := q.Delete(tt.args.id, tt.args.operatorID); (err != nil) != tt.wantErr {
				t.Errorf("QueryConditionRepo.Dalete() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := sqlboiler.QueryConditions(qm.Where("id=?", tt.args.id),
				qm.And("del != ?", true),
				qm.Load(qm.Rels(sqlboiler.QueryConditionRels.Owner, sqlboiler.StaffRels.StaffGroups), qm.Where("del != true")),
				qm.Load(sqlboiler.QueryConditionRels.QueryDisplayItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryDisplayItemColumns.RowOrder)),
				qm.Load(sqlboiler.QueryConditionRels.QueryOrderConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryOrderConditionItemColumns.RowOrder)),
				qm.Load(sqlboiler.QueryConditionRels.QuerySearchConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QuerySearchConditionItemColumns.RowOrder)),
				qm.Load(sqlboiler.QueryConditionRels.StaffGroups, qm.Where("del != true")),
			).One(context.Background(), con)
			if got != nil {
				t.Errorf("id: %v is not deleted. got is %v", tt.args.id, got.Del)
			}
		})
	}
}

func TestQueryConditionRepo_SelectQueryOperatorUsable(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		operatorID string
	}
	tests := []struct {
		name                      string
		fields                    fields
		args                      args
		wantResultQueryConditions domain.Conditions
		wantErr                   bool
	}{
		{
			name: "shuld return staffid1 usable condition list",
			fields: fields{
				database: createDB(),
			},
			args: args{
				operatorID: "staffid1",
			},
			wantResultQueryConditions: domain.Conditions{
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
				createExpectedQueryCondition10Entity(),
				createExpectedQueryCondition11Entity(),
			},
			wantErr: false,
		},
		{
			name: "shuld return staffid4 usable condition list",
			fields: fields{
				database: createDB(),
			},
			args: args{
				operatorID: "staffid4",
			},
			wantResultQueryConditions: domain.Conditions{
				createExpectedQueryCondition9Entity(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpQueryConditionTest()
			defer tearDownQueryConditionTest(con)
			setupTestData()

			q := &QueryConditionRepo{
				database: tt.fields.database,
			}
			gotResultQueryConditions, err := q.SelectQueryOperatorUsable(tt.args.operatorID)

			sort.Slice(gotResultQueryConditions, func(i, j int) bool {

				return gotResultQueryConditions[i].PatternName > gotResultQueryConditions[j].PatternName
			})

			sort.Slice(tt.wantResultQueryConditions, func(i, j int) bool {

				return tt.wantResultQueryConditions[i].PatternName > tt.wantResultQueryConditions[j].PatternName
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("QueryConditionRepo.SelectQueryOperatorUsable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(gotResultQueryConditions, tt.wantResultQueryConditions); diff != "" {
				t.Errorf("differs = %s", diff)
			}
		})
	}
}
