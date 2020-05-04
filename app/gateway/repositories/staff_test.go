package repositories

import (
	"context"
	"reflect"
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/sqlboiler"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func setUpStaffTest() *sqlx.DB {
	db := createDB()
	con, _ := db.Open()
	con.Exec("delete from staffs")
	con.Exec("delete from staff_groups")
	con.Exec("delete from join_staffs_staff_groups")
	return con
}

func tearDownStaffTest(con *sqlx.DB) {
	con.Close()
}

func createExpectedStaff1Entity() entities.Staff {
	return entities.Staff{
		ID:        "staffid1",
		AccountID: "12345",
		Name:      "name 1",
		Password:  "password 1",
		Groups:    createExpectedStaffGroupEntity1Slice(),
	}
}

func createExpectedStaff2Entity() entities.Staff {
	return entities.Staff{
		ID:        "staffid2",
		AccountID: "22345",
		Name:      "name 2",
		Password:  "password 2",
		Groups: []entities.StaffGroup{
			createExpectedStaffGroup1Entity(),
		},
	}
}

func createExpectedStaff3Entity() entities.Staff {
	return entities.Staff{
		ID:        "staffid3",
		AccountID: "32345",
		Name:      "name 3",
		Password:  "password 3",
		Groups: []entities.StaffGroup{
			createExpectedStaffGroup2Entity(),
		},
	}
}

func createExpectedStaffEntity1Slice() []entities.Staff {
	return []entities.Staff{
		createExpectedStaff1Entity(),
		createExpectedStaff3Entity(),
	}
}

func insertStaffTestData(con *sqlx.DB) {
	staffstr := "insert into staffs (id, account_id, name, password) "
	staffstr += "values('staffid1', '12345', 'name 1', 'password 1'), "
	staffstr += "('staffid2', '22345', 'name 2', 'password 2'), "
	staffstr += "('staffid3', '32345', 'name 3', 'password 3')"
	con.Exec(staffstr)
	insertStaffGroupTestData(con)
	groupstr := "insert into join_staffs_staff_groups (staffs_id, staff_groups_id) "
	groupstr += "values('staffid1', 'staffgroupid1'), "
	groupstr += "('staffid1', 'staffgroupid2'), "
	groupstr += "('staffid2', 'staffgroupid1'), "
	groupstr += "('staffid3', 'staffgroupid2')"
	con.Exec(groupstr)
}

func TestStaffObjectMap(t *testing.T) {

	type args struct {
		ss *sqlboiler.Staff
	}
	tests := []struct {
		name   string
		args   args
		wantEs entities.Staff
	}{
		{
			name:   "convert sqlboiler.staff to entities.staff",
			args:   args{},
			wantEs: createExpectedStaff1Entity(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffTest()
			defer tearDownStaffTest(con)
			insertStaffTestData(con)

			tt.args.ss, _ = sqlboiler.Staffs(qm.Where("id=?", "staffid1"), qm.Load(sqlboiler.StaffRels.StaffGroups)).One(context.Background(), con)

			if gotEs := StaffObjectMap(tt.args.ss); !reflect.DeepEqual(gotEs, tt.wantEs) {
				t.Errorf("StaffObjectMap() = %v, want %v", gotEs, tt.wantEs)
			}
		})
	}
}

func TestStaffRepo_Select(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		queryItems []*query.SearchConditionItem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Staff
		wantErr bool
	}{
		{
			name: "select with condition",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []*query.SearchConditionItem{
					{
						SearchField: query.FieldAttr{
							ID:        "name",
							ViewValue: "利用者名称",
							FieldType: query.STRING,
						},
						ConditionValue: "name",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
					{
						SearchField: query.FieldAttr{
							ID:        "groups",
							ViewValue: "所属グループ",
							FieldType: query.ARRAY,
						},
						ConditionValue: "[\"staffgroupid2\"]",
						MatchType:      query.In,
						Operator:       query.And,
					},
				},
			},
			want:    createExpectedStaffEntity1Slice(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffTest()
			defer tearDownStaffTest(con)
			insertStaffTestData(con)

			s := &StaffRepo{
				database: tt.fields.database,
			}
			got, err := s.Select(tt.args.queryItems...)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffRepo.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StaffRepo.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
