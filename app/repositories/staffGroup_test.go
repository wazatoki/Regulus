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

func setUpStaffGroupTest() *sqlx.DB {
	db := createDB()
	con, _ := db.Open()

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

func createExpectedStaffGroup3Entity() entities.StaffGroup {
	return entities.StaffGroup{
		ID:   "staffgroupid3",
		Name: "staff group name 3",
	}
}

func createExpectedStaffGroupEntity1Slice() []entities.StaffGroup {
	return []entities.StaffGroup{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
	}
}

func createExpectedStaffGroupEntity2Slice() []entities.StaffGroup {
	return []entities.StaffGroup{
		createExpectedStaffGroup1Entity(),
		createExpectedStaffGroup2Entity(),
		createExpectedStaffGroup3Entity(),
	}
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
			setupTestData()

			tt.args.sg, _ = sqlboiler.StaffGroups(qm.Where("id=?", "staffgroupid1")).One(context.Background(), con)

			if gotEg := StaffGroupObjectMap(tt.args.sg); !reflect.DeepEqual(gotEg, tt.wantEg) {
				t.Errorf("StaffGroupObjectMap() = %v, want %v", gotEg, tt.wantEg)
			}
		})
	}
}

func TestStaffGroupRepo_Select(t *testing.T) {
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
		want    []entities.StaffGroup
		wantErr bool
	}{
		{
			name: "select staff group with condition",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []*query.SearchConditionItem{
					{
						SearchField: query.FieldAttr{
							ID:        "name",
							ViewValue: "グループ名称",
							FieldType: query.STRING,
						},
						ConditionValue: "name 1",
						MatchType:      query.Pertialmatch,
						Operator:       query.And,
					},
				},
			},
			want: []entities.StaffGroup{
				createExpectedStaffGroup1Entity(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffGroupTest()
			defer tearDownStaffGroupTest(con)
			setupTestData()

			g := &StaffGroupRepo{
				database: tt.fields.database,
			}
			got, err := g.Select(tt.args.queryItems...)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffGroupRepo.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StaffGroupRepo.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStaffGroupRepo_SelectByID(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		id string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStaffGroup entities.StaffGroup
		wantErr        bool
	}{
		{
			name: "it should get specified entity as select by id",
			fields: fields{
				database: createDB(),
			},
			args: args{
				id: "staffgroupid1",
			},
			wantStaffGroup: createExpectedStaffGroup1Entity(),
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffGroupTest()
			defer tearDownStaffGroupTest(con)
			setupTestData()

			g := &StaffGroupRepo{
				database: tt.fields.database,
			}
			gotStaffGroup, err := g.SelectByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffGroupRepo.SelectByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStaffGroup, tt.wantStaffGroup) {
				t.Errorf("StaffGroupRepo.SelectByID() = %v, want %v", gotStaffGroup, tt.wantStaffGroup)
			}
		})
	}
}

func TestStaffGroupRepo_SelectByIDs(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		ids []string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantStaffGroups []entities.StaffGroup
		wantErr         bool
	}{
		{
			name: "it should get specified entities as select by ids",
			fields: fields{
				database: createDB(),
			},
			args: args{
				ids: []string{
					"staffgroupid1",
					"staffgroupid2",
				},
			},
			wantStaffGroups: []entities.StaffGroup{
				createExpectedStaffGroup1Entity(),
				createExpectedStaffGroup2Entity(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffGroupTest()
			defer tearDownStaffGroupTest(con)
			setupTestData()

			g := &StaffGroupRepo{
				database: tt.fields.database,
			}
			gotStaffGroups, err := g.SelectByIDs(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffGroupRepo.SelectByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStaffGroups, tt.wantStaffGroups) {
				t.Errorf("StaffGroupRepo.SelectByIDs() = %v, want %v", gotStaffGroups, tt.wantStaffGroups)
			}
		})
	}
}

func TestStaffGroupRepo_Insert(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		staffGroup *entities.StaffGroup
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
				&entities.StaffGroup{
					ID:   "staffgroupid3",
					Name: "staff group name 5",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffGroupTest()
			defer tearDownStaffGroupTest(con)
			setupTestData()
			want := entities.StaffGroup{
				ID:   "staffgroupid3",
				Name: "staff group name 5",
			}
			g := &StaffGroupRepo{
				database: tt.fields.database,
			}
			gotID, err := g.Insert(tt.args.staffGroup)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffGroupRepo.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want.ID = gotID

			got, _ := sqlboiler.StaffGroups(qm.Where("id=?", want.ID)).One(context.Background(), con)
			resultEntity := StaffGroupObjectMap(got)

			if !reflect.DeepEqual(resultEntity, want) {
				t.Errorf("StaffRepo.SelectByID() = %v, want %v", resultEntity, want)
			}
		})
	}
}

func TestStaffGroupRepo_Update(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		staffGroup *entities.StaffGroup
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
			con := setUpStaffGroupTest()
			defer tearDownStaffGroupTest(con)
			setupTestData()
			beforeStaffGroup := createExpectedStaffGroup1Entity()
			tt.args.staffGroup = &beforeStaffGroup
			tt.args.staffGroup.Name = "staff group name 5"
			g := &StaffGroupRepo{
				database: tt.fields.database,
			}
			if err := g.Update(tt.args.staffGroup); (err != nil) != tt.wantErr {
				t.Errorf("StaffGroupRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, _ := sqlboiler.StaffGroups(qm.Where("id=?", beforeStaffGroup.ID)).One(context.Background(), con)
			resultEntity := StaffGroupObjectMap(got)

			if !reflect.DeepEqual(resultEntity, *tt.args.staffGroup) {
				t.Errorf("StaffRepo.SelectByID() = %v, want %v", resultEntity, tt.args.staffGroup)
			}
		})
	}
}

func TestStaffGroupRepo_Dalete(t *testing.T) {
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
		wantErr bool
	}{
		{
			name: "should return nil as called delete",
			fields: fields{
				database: createDB(),
			},
			args: args{
				id: "staffgroupid2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpStaffGroupTest()
			defer tearDownStaffGroupTest(con)
			setupTestData()

			g := &StaffGroupRepo{
				database: tt.fields.database,
			}
			if err := g.Dalete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("StaffGroupRepo.Dalete() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := sqlboiler.StaffGroups(
				qm.Where("id=?", "staffgroupid2"),
				qm.And("del != true"),
			).One(context.Background(), con)

			if got != nil {
				t.Errorf("StaffRepo.Dalete() = %v, want %v", got, nil)
			}
		})
	}
}
