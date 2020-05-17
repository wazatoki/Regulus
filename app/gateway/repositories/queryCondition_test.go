package repositories

import (
	"context"
	"reflect"
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/sqlboiler"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func setUpQueryConditionTest() *sqlx.DB {
	db := createDB()
	con, _ := db.Open()

	return con
}

func tearDownQueryConditionTest(con *sqlx.DB) {
	con.Close()
}

func createExpectedQueryCondition0Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid0",
		PatternName: "patternName0",
		Category:    entities.Categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffSearchConditionList[0],
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

func createExpectedQueryCondition1Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid1",
		PatternName: "patternName1",
		Category:    entities.Categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffSearchConditionList[0],
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

func createExpectedQueryCondition2Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid2",
		PatternName: "patternName2",
		Category:    entities.Categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffSearchConditionList[0],
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

func createExpectedQueryCondition3Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid3",
		PatternName: "patternName3",
		Category:    entities.Categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffSearchConditionList[0],
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

func createExpectedQueryCondition4Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid4",
		PatternName: "patternName4",
		Category:    entities.Categories[0],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffSearchConditionList[0],
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

func createExpectedQueryCondition5Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid5",
		PatternName: "patternName5",
		Category:    entities.Categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffGroupSearchConditionList[0],
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

func createExpectedQueryCondition6Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid6",
		PatternName: "patternName6",
		Category:    entities.Categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffGroupSearchConditionList[0],
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

func createExpectedQueryCondition7Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid7",
		PatternName: "patternName7",
		Category:    entities.Categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffGroupSearchConditionList[0],
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

func createExpectedQueryCondition8Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid8",
		PatternName: "patternName8",
		Category:    entities.Categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffGroupSearchConditionList[0],
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

func createExpectedQueryCondition9Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid9",
		PatternName: "patternName9",
		Category:    entities.Categories[1],
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffGroupSearchConditionList[0],
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

func TestQueryConditionObjectMap(t *testing.T) {
	type args struct {
		sqc *sqlboiler.QueryCondition
	}
	tests := []struct {
		name    string
		args    args
		wantEqc entities.QueryCondition
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
		queryItems []*query.SearchConditionItem
	}
	tests := []struct {
		name                      string
		fields                    fields
		args                      args
		wantResultQueryConditions []entities.QueryCondition
		wantErr                   bool
	}{
		{
			name: "should get sprcified entity as select by condition",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []*query.SearchConditionItem{
					{
						SearchField:    query.QueryConditionSearchConditionList[0],
						ConditionValue: "5",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
					{
						SearchField:    query.QueryConditionSearchConditionList[1],
						ConditionValue: "利用者グ",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
				},
			},
			wantResultQueryConditions: []entities.QueryCondition{
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
		wantQueryCondition entities.QueryCondition
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
		wantQueryConditions []entities.QueryCondition
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
			wantQueryConditions: []entities.QueryCondition{
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
		queryCondition entities.QueryCondition
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantEntity entities.QueryCondition
		wantErr    bool
	}{
		{
			name: "should insert to database as called insert",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryCondition: entities.QueryCondition{
					PatternName:    "testPatternName",
					Category:       entities.Categories[0],
					IsDisclose:     true,
					DiscloseGroups: createExpectedStaffGroupEntity1Slice(),
					ConditionData: query.ConditionData{
						DisplayItemList: []query.FieldAttr{},
						SearchConditionList: []query.SearchConditionItem{
							{
								SearchField:    query.StaffSearchConditionList[0],
								ConditionValue: "aaa",
								MatchType:      query.Match,
								Operator:       query.And,
							},
							{
								SearchField:    query.StaffSearchConditionList[1],
								ConditionValue: "bbb",
								MatchType:      query.Unmatch,
								Operator:       query.And,
							},
						},
						OrderConditionList: []query.OrderConditionItem{},
					},
					Owner: createExpectedStaff1Entity(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			con := setUpStaffTest()
			defer tearDownStaffTest(con)
			setupTestData()
			tt.wantEntity = entities.QueryCondition{
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
			gotID, err := q.Insert(tt.args.queryCondition)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryConditionRepo.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.wantEntity.ID = gotID
			got, _ := sqlboiler.QueryConditions(qm.Where("id=?", tt.wantEntity.ID),
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
