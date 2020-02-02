package repositories

import (
	"reflect"
	makerEntity "regulus/app/domain/entities/maker"
	"regulus/app/infrastructures/viper"
	"testing"

	"github.com/jmoiron/sqlx"
)

func connectDB() *sqlx.DB {
	db := createDB()
	con, _ := db.Open()
	con.Exec("delete from maker")
	return con
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
			con := connectDB()
			defer con.Close()

			m := &MakerRepo{
				database: tt.fields.database,
			}
			got, _ := m.Insert(tt.args.makerEntity)

			maker := makerEntity.Maker{}
			con.Get(&maker, "select id, name from maker limit 1")

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
			con := connectDB()
			defer con.Close()

			con.Exec("insert into maker (id, name) values('id1', 'testname1')")
			m := &MakerRepo{
				database: tt.fields.database,
			}

			m.Update(tt.args.makerEntity)

			maker := makerEntity.Maker{}
			con.Get(&maker, "select id, name from maker limit 1")

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
			con := connectDB()
			defer con.Close()

			con.Exec("insert into maker (id, name) values('id1', 'testname1'),('id2', 'testname2'),('id3', 'testname3')")

			m := &MakerRepo{
				database: tt.fields.database,
			}

			m.Dalete(tt.args.id)

			maker := makerEntity.Maker{}

			con.Get(&maker, "select id, name from maker where del = true")

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
			con := connectDB()
			defer con.Close()

			con.Exec("insert into maker (id, name) values('id1', 'testname1'),('id2', 'testname2'),('id3', 'testname3')")

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
