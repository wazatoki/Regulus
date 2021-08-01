package repositories

import (
	"context"
	"reflect"
	"regulus/app/domain"
	"regulus/app/infrastructures/sqlboiler"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func setUpStaffTest() *sqlx.DB {
	db := createDB()
	con, _ := db.Open()

	return con
}

func tearDownStaffTest(con *sqlx.DB) {
	con.Close()
}

func createExpectedStaff1Entity() *domain.Staff {
	return &domain.Staff{
		ID:                 "staffid1",
		AccountID:          "12345",
		Name:               "name 1",
		Password:           "password 1",
		StaffGroups:        createExpectedStaffGroupEntity1Slice(),
		FavoriteConditions: domain.Conditions{},
	}
}

func createExpectedStaff2Entity() *domain.Staff {
	return &domain.Staff{
		ID:        "staffid2",
		AccountID: "22345",
		Name:      "name 2",
		Password:  "password 2",
		StaffGroups: domain.StaffGroups{
			createExpectedStaffGroup1Entity(),
		},
		FavoriteConditions: domain.Conditions{},
	}
}

func createExpectedStaff3Entity() *domain.Staff {
	return &domain.Staff{
		ID:        "staffid3",
		AccountID: "32345",
		Name:      "name 3",
		Password:  "password 3",
		StaffGroups: domain.StaffGroups{
			createExpectedStaffGroup2Entity(),
		},
		FavoriteConditions: domain.Conditions{},
	}
}

func createExpectedStaff4Entity() *domain.Staff {
	return &domain.Staff{
		ID:        "staffid4",
		AccountID: "42345",
		Name:      "name 4",
		Password:  "password 4",
		StaffGroups: domain.StaffGroups{
			createExpectedStaffGroup3Entity(),
		},
		FavoriteConditions: domain.Conditions{},
	}
}

func createExpectedStaff5Entity() *domain.Staff {
	return &domain.Staff{
		ID:        "staffid5",
		AccountID: "52345",
		Name:      "name 5",
		Password:  "password 5",
		StaffGroups: domain.StaffGroups{
			createExpectedStaffGroup1Entity(),
		},
		FavoriteConditions: domain.Conditions{},
	}
}

func createExpectedStaffEntity1Slice() []*domain.Staff {
	return []*domain.Staff{
		createExpectedStaff1Entity(),
		createExpectedStaff3Entity(),
		createExpectedStaff4Entity(),
	}
}

func createExpectedStaffEntity2Slice() []*domain.Staff {
	return []*domain.Staff{
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
		wantEs *domain.Staff
	}{
		{
			name:   "convert sqlboiler.staff to domain.Staff",
			args:   args{},
			wantEs: createExpectedStaff1Entity(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffTest()
			defer tearDownStaffTest(con)
			setupTestData()

			tt.args.ss, _ = sqlboiler.Staffs(qm.Where("id=?", "staffid1"), qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true"))).One(context.Background(), con)

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
		queryItems []*domain.SearchConditionItem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Staff
		wantErr bool
	}{
		{
			name: "select staff with condition",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []*domain.SearchConditionItem{
					{
						SearchField: domain.FieldAttr{
							ID:        "name",
							ViewValue: "利用者名称",
							FieldType: domain.QueryValueTypeEnum.STRING,
						},
						ConditionValue: "name",
						MatchType:      domain.QueryMatchTypeEnum.PERTIALMATCH,
						Operator:       domain.QueryOperatorEnum.AND,
					},
					{
						SearchField: domain.FieldAttr{
							ID:        "groups",
							ViewValue: "所属グループ",
							FieldType: domain.QueryValueTypeEnum.ARRAY,
						},
						ConditionValue: "[\"staffgroupid2\"]",
						MatchType:      domain.QueryMatchTypeEnum.IN,
						Operator:       domain.QueryOperatorEnum.AND,
					},
					{
						SearchField: domain.FieldAttr{
							ID:        "groups",
							ViewValue: "所属グループ",
							FieldType: domain.QueryValueTypeEnum.ARRAY,
						},
						ConditionValue: "[\"staffgroupid3\"]",
						MatchType:      domain.QueryMatchTypeEnum.IN,
						Operator:       domain.QueryOperatorEnum.OR,
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
		wantStaffs []*domain.Staff
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
		want    *domain.Staff
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
		wantStaffs []*domain.Staff
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
		staff      *domain.Staff
		operatorID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should insert to database as called insert",
			fields: fields{
				database: createDB(),
			},
			args: args{
				staff: &domain.Staff{
					ID:        "",
					AccountID: "62345",
					Name:      "name 6",
					Password:  "password 6",
					StaffGroups: domain.StaffGroups{
						createExpectedStaffGroup1Entity(),
						createExpectedStaffGroup2Entity(),
					},
				},
				operatorID: createExpectedStaff1Entity().ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffTest()
			defer tearDownStaffTest(con)
			setupTestData()
			want := &domain.Staff{
				ID:        "",
				AccountID: "62345",
				Name:      "name 6",
				Password:  "password 6",
				StaffGroups: domain.StaffGroups{
					createExpectedStaffGroup1Entity(),
					createExpectedStaffGroup2Entity(),
				},
				FavoriteConditions: domain.Conditions{},
			}

			s := &StaffRepo{
				database: tt.fields.database,
			}
			gotID, err := s.Insert(tt.args.staff, tt.args.operatorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffRepo.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want.ID = gotID

			got, _ := sqlboiler.Staffs(qm.Where("id=?", want.ID), qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true"))).One(context.Background(), con)
			resultEntity := StaffObjectMap(got)

			if !reflect.DeepEqual(resultEntity, want) {
				t.Errorf("StaffRepo.SelectByID() = %v, want %v", resultEntity, want)
			}

		})
	}
}

func TestStaffRepo_Update(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		staff      *domain.Staff
		operatorID string
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
			con := setUpStaffTest()
			defer tearDownStaffTest(con)
			setupTestData()
			beforeStaff := createExpectedStaff1Entity()
			tt.args.staff = beforeStaff
			tt.args.staff.AccountID = "1234512345"
			tt.args.staff.Name = "name 1name 1"
			tt.args.staff.Password = "password 1password 1"
			tt.args.staff.StaffGroups = domain.StaffGroups{
				createExpectedStaffGroup2Entity(),
				createExpectedStaffGroup3Entity(),
			}
			tt.args.operatorID = createExpectedStaff1Entity().ID

			s := &StaffRepo{
				database: tt.fields.database,
			}
			if err := s.Update(tt.args.staff, tt.args.operatorID); (err != nil) != tt.wantErr {
				t.Errorf("StaffRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := sqlboiler.Staffs(qm.Where("id=?", beforeStaff.ID), qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true"))).One(context.Background(), con)
			resultEntity := StaffObjectMap(got)

			if !reflect.DeepEqual(resultEntity, tt.args.staff) {
				t.Errorf("StaffRepo.SelectByID() = %v, want %v", resultEntity, tt.args.staff)
			}

		})
	}
}

func TestStaffRepo_Dalete(t *testing.T) {
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
				id:         "staffid2",
				operatorID: createExpectedStaff1Entity().ID,
			},
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
			if err := s.Dalete(tt.args.id, tt.args.operatorID); (err != nil) != tt.wantErr {
				t.Errorf("StaffRepo.Dalete() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, _ := sqlboiler.Staffs(
				qm.Where("id=?", "staffid2"),
				qm.And("del != true"),
				qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true")),
			).One(context.Background(), con)

			if got != nil {
				t.Errorf("StaffRepo.Dalete() = %v, want %v", got, nil)
			}
		})
	}
}

func TestStaffRepo_SelectByAccountID(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantStaff *domain.Staff
		wantErr   bool
	}{
		{
			name: "it should get specified entity as select by accountID",
			fields: fields{
				database: createDB(),
			},
			args: args{
				id: "12345",
			},
			wantStaff: createExpectedStaff1Entity(),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StaffRepo{
				database: tt.fields.database,
			}
			gotStaff, err := s.SelectByAccountID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffRepo.SelectByAccountID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStaff, tt.wantStaff) {
				t.Errorf("StaffRepo.SelectByAccountID() = %v, want %v", gotStaff, tt.wantStaff)
			}
		})
	}
}
