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

func createExpectedStaff4Entity() entities.Staff {
	return entities.Staff{
		ID:        "staffid4",
		AccountID: "42345",
		Name:      "name 4",
		Password:  "password 4",
		Groups: []entities.StaffGroup{
			createExpectedStaffGroup3Entity(),
		},
	}
}

func createExpectedStaff5Entity() entities.Staff {
	return entities.Staff{
		ID:        "staffid5",
		AccountID: "52345",
		Name:      "name 5",
		Password:  "password 5",
		Groups: []entities.StaffGroup{
			createExpectedStaffGroup1Entity(),
		},
	}
}

func createExpectedStaffEntity1Slice() []entities.Staff {
	return []entities.Staff{
		createExpectedStaff1Entity(),
		createExpectedStaff3Entity(),
		createExpectedStaff4Entity(),
	}
}

func createExpectedStaffEntity2Slice() []entities.Staff {
	return []entities.Staff{
		createExpectedStaff1Entity(),
		createExpectedStaff2Entity(),
		createExpectedStaff3Entity(),
		createExpectedStaff4Entity(),
		createExpectedStaff5Entity(),
	}
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
			setupTestData()

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
			name: "select staff with condition",
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
					{
						SearchField: query.FieldAttr{
							ID:        "groups",
							ViewValue: "所属グループ",
							FieldType: query.ARRAY,
						},
						ConditionValue: "[\"staffgroupid3\"]",
						MatchType:      query.In,
						Operator:       query.Or,
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
			setupTestData()

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

func TestStaffRepo_SelectAll(t *testing.T) {
	type fields struct {
		database db
	}
	tests := []struct {
		name       string
		fields     fields
		wantStaffs []entities.Staff
		wantErr    bool
	}{
		{
			name: "it should get 5entities as select all ",
			fields: fields{
				database: createDB(),
			},
			wantStaffs: createExpectedStaffEntity2Slice(),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffTest()
			defer tearDownStaffTest(con)
			setupTestData()

			s := &StaffRepo{
				database: tt.fields.database,
			}
			gotStaffs, err := s.SelectAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffRepo.SelectAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStaffs, tt.wantStaffs) {
				t.Errorf("StaffRepo.SelectAll() = %v, want %v", gotStaffs, tt.wantStaffs)
			}
		})
	}
}

func TestStaffRepo_SelectByID(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Staff
		wantErr bool
	}{
		{
			name: "it should get specified entity as select by id",
			fields: fields{
				database: createDB(),
			},
			args: args{
				id: "staffid1",
			},
			want:    createExpectedStaff1Entity(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffTest()
			defer tearDownStaffTest(con)
			setupTestData()

			s := &StaffRepo{
				database: tt.fields.database,
			}
			got, err := s.SelectByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffRepo.SelectByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StaffRepo.SelectByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStaffRepo_SelectByIDs(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		ids []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStaffs []entities.Staff
		wantErr    bool
	}{
		{
			name: "it should get specified entities as select by ids",
			fields: fields{
				database: createDB(),
			},
			args: args{
				ids: []string{
					"staffid1",
					"staffid3",
					"staffid4",
				},
			},
			wantStaffs: createExpectedStaffEntity1Slice(),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffTest()
			defer tearDownStaffTest(con)
			setupTestData()

			s := &StaffRepo{
				database: tt.fields.database,
			}
			gotStaffs, err := s.SelectByIDs(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffRepo.SelectByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStaffs, tt.wantStaffs) {
				t.Errorf("StaffRepo.SelectByIDs() = %v, want %v", gotStaffs, tt.wantStaffs)
			}
		})
	}
}

func TestStaffRepo_Insert(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		staff *entities.Staff
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  string
		wantErr bool
	}{
		{
			name: "should insert to database as called insert",
			fields: fields{
				database: createDB(),
			},
			args: args{
				staff: &entities.Staff{
					ID:        "",
					AccountID: "62345",
					Name:      "name 6",
					Password:  "password 6",
					Groups: []entities.StaffGroup{
						createExpectedStaffGroup1Entity(),
						createExpectedStaffGroup2Entity(),
					},
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
			want := entities.Staff{
				ID:        "",
				AccountID: "62345",
				Name:      "name 6",
				Password:  "password 6",
				Groups: []entities.StaffGroup{
					createExpectedStaffGroup1Entity(),
					createExpectedStaffGroup2Entity(),
				},
			}

			s := &StaffRepo{
				database: tt.fields.database,
			}
			gotID, err := s.Insert(tt.args.staff)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffRepo.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want.ID = gotID

			got, _ := sqlboiler.Staffs(qm.Where("id=?", want.ID), qm.Load(sqlboiler.StaffRels.StaffGroups)).One(context.Background(), con)
			resultEntity := StaffObjectMap(got)

			if !reflect.DeepEqual(resultEntity, want) {
				t.Errorf("StaffRepo.SelectByID() = %v, want %v", resultEntity, want)
			}

		})
	}
}
