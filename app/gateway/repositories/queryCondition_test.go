package repositories

import (
	"context"
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

func createExpectedQueryCondition1Entity() entities.QueryCondition {
	return entities.QueryCondition{
		ID:          "queryConditionid1",
		PatternName: "patternName1",
		Category: entities.Category{
			Name:      "staff",
			ViewValue: "利用者",
			SearchItems: entities.ComplexSearchItems{
				SearchConditionList: query.StaffSearchConditionList,
				DisplayItemList:     []query.FieldAttr{},
				OrderConditionList:  []query.FieldAttr{},
				Groups:              nil,
			},
		},
		ConditionData: query.ConditionData{
			DisplayItemList: []query.FieldAttr{},
			SearchConditionList: []query.SearchConditionItem{
				{
					SearchField:    query.StaffSearchConditionList[0],
					ConditionValue: "123",
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
