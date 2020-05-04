package repositories

import (
	"reflect"
	makerEntity "regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/viper"
	"testing"

	"github.com/jmoiron/sqlx"
)

func setUpMakerTest() *sqlx.DB {
	db := createDB()
	con, _ := db.Open()
	con.Exec("delete from makers")
	return con
}

func tearDownMakerTest(con *sqlx.DB) {
	con.Close()
}

func TestMakerRepo_Insert(t *testing.T) {
	viper.SetupTestConfig()
	type fields struct {
		database db
	}
	type args struct {
		makerEntity *makerEntity.Maker
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "normal insert",
			fields: fields{
				database: createDB(),
			},
			args: args{
				makerEntity: &makerEntity.Maker{
					ID:   "",
					Name: "testname",
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpMakerTest()
			defer tearDownMakerTest(con)

			m := &MakerRepo{
				database: tt.fields.database,
			}
			got, _ := m.Insert(tt.args.makerEntity)

			maker := makerEntity.Maker{}
			con.Get(&maker, "select id, name from makers limit 1")

			if got == "" {
				t.Errorf("MakerRepo.Insert() = %v, want not ''", got)
			}

			if maker.Name != "testname" {
				t.Errorf("MakerRepo.Insert() = %v, want %v", maker.Name, "testname")
			}
		})
	}
}

func TestMakerRepo_Update(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		makerEntity *makerEntity.Maker
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal update",
			fields: fields{
				database: createDB(),
			},
			args: args{
				makerEntity: &makerEntity.Maker{
					ID:   "id1",
					Name: "testname2",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpMakerTest()
			defer tearDownMakerTest(con)

			con.Exec("insert into makers (id, name) values('id1', 'testname1')")
			m := &MakerRepo{
				database: tt.fields.database,
			}

			m.Update(tt.args.makerEntity)

			maker := makerEntity.Maker{}
			con.Get(&maker, "select id, name from makers limit 1")

			if maker.Name != tt.args.makerEntity.Name {
				t.Errorf("MakerRepo.Update() id = %v, name = %v, wantName %v", maker.ID, maker.Name, tt.args.makerEntity.Name)
			}
		})
	}
}

func TestMakerRepo_Dalete(t *testing.T) {
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
			name: "normal delete",
			fields: fields{
				database: createDB(),
			},
			args: args{
				id: "id1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpMakerTest()
			defer tearDownMakerTest(con)

			con.Exec("insert into makers (id, name) values('id1', 'testname1'),('id2', 'testname2'),('id3', 'testname3')")

			m := &MakerRepo{
				database: tt.fields.database,
			}

			m.Dalete(tt.args.id)

			maker := makerEntity.Maker{}

			con.Get(&maker, "select id, name from makers where del = true")

			if maker.ID != "id1" {
				t.Errorf("MakerRepo.Dalete() deleted id = %v, wantErr %v", maker.ID, "id1")
			}
		})
	}
}

func TestMakerRepo_SelectByID(t *testing.T) {
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
		want    *makerEntity.Maker
		wantErr bool
	}{
		{
			name: "normal select by id",
			fields: fields{
				database: createDB(),
			},
			args: args{
				id: "id1",
			},
			want: &makerEntity.Maker{
				ID:   "id1",
				Name: "testname1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpMakerTest()
			defer tearDownMakerTest(con)

			con.Exec("insert into makers (id, name) values('id1', 'testname1'),('id2', 'testname2'),('id3', 'testname3')")

			m := &MakerRepo{
				database: tt.fields.database,
			}
			got, _ := m.SelectByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakerRepo.SelectByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakerRepo_SelectAll(t *testing.T) {
	type fields struct {
		database db
	}
	tests := []struct {
		name    string
		fields  fields
		want    []makerEntity.Maker
		wantErr bool
	}{
		{
			name: "normal select all",
			fields: fields{
				database: createDB(),
			},
			want: []makerEntity.Maker{
				{
					ID:   "id1",
					Name: "testname1",
				},
				{
					ID:   "id2",
					Name: "testname2",
				},
				{
					ID:   "id3",
					Name: "testname3",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := setUpMakerTest()
			defer tearDownMakerTest(con)

			con.Exec("insert into makers (id, name) values('id1', 'testname1'),('id2', 'testname2'),('id3', 'testname3')")

			m := &MakerRepo{
				database: tt.fields.database,
			}
			got, _ := m.SelectAll()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakerRepo.SelectAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakerRepo_Select(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		queryItems []*query.SearchConditionItem
	}
	qi1 := query.SearchConditionItem{
		SearchField: query.FieldAttr{
			ID:        "id1",
			ViewValue: "MakerName",
			FieldType: query.STRING,
		},
		ConditionValue: "1",
		MatchType:      "pertialmatch",
		Operator:       "and",
	}
	qi2 := query.SearchConditionItem{
		SearchField: query.FieldAttr{
			ID:        "id1",
			ViewValue: "MakerName",
			FieldType: query.STRING,
		},
		ConditionValue: "2",
		MatchType:      "pertialmatch",
		Operator:       "or",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []makerEntity.Maker
		wantErr bool
	}{
		{
			name: "normal select with no query",
			fields: fields{
				database: createDB(),
			},
			args: args{},
			want: []makerEntity.Maker{
				{
					ID:   "id1",
					Name: "testname1",
				},
				{
					ID:   "id3",
					Name: "testname3",
				},
			},
			wantErr: true,
		},
		{
			name: "normal select with a query",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []*query.SearchConditionItem{
					&qi1,
				},
			},
			want: []makerEntity.Maker{
				{
					ID:   "id1",
					Name: "testname1",
				},
			},
			wantErr: true,
		},
		{
			name: "normal select with two queries",
			fields: fields{
				database: createDB(),
			},
			args: args{
				queryItems: []*query.SearchConditionItem{
					&qi1,
					&qi2,
				},
			},
			want: []makerEntity.Maker{
				{
					ID:   "id1",
					Name: "testname1",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		con := setUpMakerTest()
		defer tearDownMakerTest(con)

		con.Exec("insert into makers (id, name, del) values('id1', 'testname1', false),('id2', 'testname2', true),('id3', 'testname3', false)")

		t.Run(tt.name, func(t *testing.T) {
			m := &MakerRepo{
				database: tt.fields.database,
			}
			got, _ := m.Select(tt.args.queryItems...)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("MakerRepo.Select() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakerRepo.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakerRepo_SelectByIDs(t *testing.T) {
	type fields struct {
		database db
	}
	type args struct {
		ids []string
	}
	arg := args{
		ids: []string{
			"id1", "id2",
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []makerEntity.Maker
		wantErr bool
	}{
		{
			name: "normal select by ids",
			fields: fields{
				database: createDB(),
			},
			args: arg,
			want: []makerEntity.Maker{
				{
					ID:   "id1",
					Name: "testname1",
				},
				{
					ID:   "id2",
					Name: "testname2",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		con := setUpMakerTest()
		defer tearDownMakerTest(con)

		con.Exec("insert into makers (id, name, del) values('id1', 'testname1', false),('id2', 'testname2', false),('id3', 'testname3', false)")

		t.Run(tt.name, func(t *testing.T) {
			m := &MakerRepo{
				database: tt.fields.database,
			}

			got, _ := m.SelectByIDs(tt.args.ids)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("MakerRepo.SelectByIDs() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakerRepo.SelectByIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		makers     []makerEntity.Maker
		orderItems []query.OrderConditionItem
	}
	unsortMakers := []makerEntity.Maker{
		{
			ID:   "id1",
			Name: "testname1",
		},
		{
			ID:   "id2",
			Name: "testname2",
		},
		{
			ID:   "id3",
			Name: "testname3",
		},
	}
	tests := []struct {
		name string
		args args
		want []makerEntity.Maker
	}{
		{
			name: "sort by asc",
			args: args{
				makers: unsortMakers,
				orderItems: []query.OrderConditionItem{
					{
						OrderField: query.FieldAttr{
							ID:        "id1",
							FieldType: query.STRING,
							ViewValue: "makerName",
						},
						OrderFieldKeyWord: query.Asc,
					},
				},
			},
			want: []makerEntity.Maker{
				{
					ID:   "id1",
					Name: "testname1",
				},
				{
					ID:   "id2",
					Name: "testname2",
				},
				{
					ID:   "id3",
					Name: "testname3",
				},
			},
		},
		{
			name: "sort by desc",
			args: args{
				makers: unsortMakers,
				orderItems: []query.OrderConditionItem{
					{
						OrderField: query.FieldAttr{
							ID:        "id1",
							FieldType: query.STRING,
							ViewValue: "makerName",
						},
						OrderFieldKeyWord: query.Desc,
					},
				},
			},
			want: []makerEntity.Maker{
				{
					ID:   "id3",
					Name: "testname3",
				},
				{
					ID:   "id2",
					Name: "testname2",
				},
				{
					ID:   "id1",
					Name: "testname1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.makers, tt.args.orderItems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
