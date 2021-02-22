package repositories

import (
	"context"
	"reflect"
	"regulus/app/domain/query"
	"regulus/app/infrastructures/sqlboiler"
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

var categories = []*query.Category{
	{
		Name:      "staff",
		ViewValue: "利用者",
		SearchItems: query.ComplexSearchItems{
			SearchConditionList: staffSearchConditionList,
			DisplayItemList:     []query.FieldAttr{},
			OrderConditionList:  []query.FieldAttr{},
			Groups:              createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "staff-group",
		ViewValue: "利用者グループ",
		SearchItems: query.ComplexSearchItems{
			SearchConditionList: staffGroupSearchConditionList,
			DisplayItemList:     []query.FieldAttr{},
			OrderConditionList:  []query.FieldAttr{},
			Groups:              createExpectedStaffGroupEntity2Slice(),
		},
	},
	{
		Name:      "query-condition",
		ViewValue: "検索条件管理",
		SearchItems: query.ComplexSearchItems{
			SearchConditionList: queryConditionSearchConditionList,
			DisplayItemList:     []query.FieldAttr{},
			OrderConditionList:  []query.FieldAttr{},
			Groups:              nil,
		},
	},
}

func createExpectedQueryCondition0Entity() *query.Condition {
	return &query.Condition{
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

func createExpectedQueryCondition1Entity() *query.Condition {
	return &query.Condition{
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

func createExpectedQueryCondition2Entity() *query.Condition {
	return &query.Condition{
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

func createExpectedQueryCondition3Entity() *query.Condition {
	return &query.Condition{
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

func createExpectedQueryCondition4Entity() *query.Condition {
	return &query.Condition{
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
		DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff1Entity(),
	}
}

func createExpectedQueryCondition5Entity() *query.Condition {
	return &query.Condition{
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

func createExpectedQueryCondition6Entity() *query.Condition {
	return &query.Condition{
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
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition7Entity() *query.Condition {
	return &query.Condition{
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

func createExpectedQueryCondition8Entity() *query.Condition {
	return &query.Condition{
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
		IsDisclose:     true,
		Owner:          createExpectedStaff2Entity(),
	}
}

func createExpectedQueryCondition9Entity() *query.Condition {
	return &query.Condition{
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
		DiscloseGroups: createExpectedStaffGroupEntity2Slice(),
		IsDisclose:     false,
		Owner:          createExpectedStaff2Entity(),
	}
}

func TestQueryConditionObjectMap(t *testing.T) {
	type args struct {
		sqc *sqlboiler.QueryCondition
	}
	tests := []struct {
		name    string
		args    args
		wantEqc *query.Condition
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
			want:    10,
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
		queryItems []query.SearchConditionItem
	}
	tests := []struct {
		name                      string
		fields                    fields
		args                      args
		wantResultQueryConditions []*query.Condition
		wantErr                   bool
	}{
		{
			name: "should get sprcified entity as select by condition",
			fields: fields{
				database: createDB(),
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
						ConditionValue: "利用者",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
				},
			},
			wantResultQueryConditions: []*query.Condition{
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
				queryItems: []query.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[2],
						ConditionValue: "true",
						MatchType:      query.Match,
						Operator:       query.And,
					},
				},
			},
			wantResultQueryConditions: []*query.Condition{
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
				queryItems: []query.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[3],
						ConditionValue: "[\"staffgroupid3\"]",
						MatchType:      query.Match,
						Operator:       query.And,
					},
				},
			},
			wantResultQueryConditions: []*query.Condition{
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
				queryItems: []query.SearchConditionItem{
					{
						SearchField:    queryConditionSearchConditionList[4],
						ConditionValue: "2",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
					{
						SearchField:    queryConditionSearchConditionList[0],
						ConditionValue: "5",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
				},
			},
			wantResultQueryConditions: []*query.Condition{
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
		wantQueryCondition *query.Condition
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
		wantQueryConditions []*query.Condition
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
			wantQueryConditions: []*query.Condition{
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
		queryCondition query.Condition
		operatorID     string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantEntity *query.Condition
		wantErr    bool
	}{
		{
			name: "should insert to database as called insert",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryCondition: query.Condition{
					PatternName:    "testPatternName",
					Category:       categories[0],
					IsDisclose:     true,
					DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
					ConditionData: query.ConditionData{
						DisplayItemList: []query.FieldAttr{},
						SearchConditionList: []query.SearchConditionItem{
							{
								SearchField:    staffSearchConditionList[0],
								ConditionValue: "aaa",
								MatchType:      query.Match,
								Operator:       query.And,
							},
							{
								SearchField:    staffSearchConditionList[1],
								ConditionValue: "bbb",
								MatchType:      query.Unmatch,
								Operator:       query.And,
							},
						},
						OrderConditionList: []query.OrderConditionItem{},
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
			tt.wantEntity = &query.Condition{
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
		queryCondition *query.Condition
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
			beforeQueryCondition.ConditionData = query.ConditionData{
				DisplayItemList: []query.FieldAttr{},
				SearchConditionList: []query.SearchConditionItem{
					{
						SearchField:    staffSearchConditionList[1],
						ConditionValue: "2",
						MatchType:      query.Unmatch,
						Operator:       query.And,
					},
				},
				OrderConditionList: []query.OrderConditionItem{},
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
