package repositories

import (
	"context"
	"reflect"
	"regulus/app/domain/entities"
	"regulus/app/infrastructures/sqlboiler"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func setUpStaffGroupTest() *sqlx.DB {
	db := createDB()
	con, _ := db.Open()
	con.Exec("delete from staff_groups")
	return con
}

func tearDownStaffGroupTest(con *sqlx.DB) {
	con.Close()
}

func createExpectedStaffGroup1Entity() entities.StaffGroup {
	return entities.StaffGroup{
		ID:   "staffgroupid1",
		Name: "staff group name 1",
	}
}

func createExpectedStaffGroup2Entity() entities.StaffGroup {
	return entities.StaffGroup{
		ID:   "staffgroupid2",
		Name: "staff group name 2",
	}
}

func createExpectedStaffGroupEntity1Slice() []entities.StaffGroup {
	return []entities.StaffGroup{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
	}
}

func insertStaffGroupTestData(con *sqlx.DB) {
	con.Exec("insert into staff_groups (id, name) values('staffgroupid1', 'staff group name 1'), ('staffgroupid2', 'staff group name 2')")
}

func TestStaffGroupObjectMap(t *testing.T) {
	type args struct {
		sg *sqlboiler.StaffGroup
	}
	tests := []struct {
		name   string
		args   args
		wantEg entities.StaffGroup
	}{
		{
			name:   "convert sqlboiler.staffGroup to entities.staffGroup",
			args:   args{},
			wantEg: createExpectedStaffGroup1Entity(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffGroupTest()
			defer tearDownStaffGroupTest(con)

			insertStaffGroupTestData(con)

			tt.args.sg, _ = sqlboiler.StaffGroups(qm.Where("id=?", "staffgroupid1")).One(context.Background(), con)

			if gotEg := StaffGroupObjectMap(tt.args.sg); !reflect.DeepEqual(gotEg, tt.wantEg) {
				t.Errorf("StaffGroupObjectMap() = %v, want %v", gotEg, tt.wantEg)
			}
		})
	}
}
