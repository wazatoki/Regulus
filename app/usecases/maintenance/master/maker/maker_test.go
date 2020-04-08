/*

maker はメーカーマスターメンテナンスのためのパッケージを提供します。

*/

package maker

import (
	"errors"
	"reflect"
	makerEntity "regulus/app/domain/entities/maker"
	"testing"
)

type persistanceStb struct{}

func (p *persistanceStb) Insert(*makerEntity.Maker) (string, error) {
	return "sampleId1", nil
}
func (p *persistanceStb) Update(*makerEntity.Maker) error {
	return nil
}
func (p *persistanceStb) Delete(string) error {
	return nil
}
func (p *persistanceStb) SelectByID(string) (*makerEntity.Maker, error) {
	return &makerEntity.Maker{
		ID:   "sampleid1",
		Name: "samplename1",
	}, nil
}
func (p *persistanceStb) SelectAll() ([]makerEntity.Maker, error) {
	var results []makerEntity.Maker = []makerEntity.Maker{}
	return results, nil
}
func (p *persistanceStb) Select(string) ([]makerEntity.Maker, error) {
	var results []makerEntity.Maker = []makerEntity.Maker{}
	return results, nil
}

type persistanceStbErr struct{}

func (p *persistanceStbErr) Insert(*makerEntity.Maker) (string, error) {
	return "", errors.New("error occurred")
}
func (p *persistanceStbErr) Update(*makerEntity.Maker) error {
	return errors.New("error occurred")
}
func (p *persistanceStbErr) Delete(string) error {
	return errors.New("error occurred")
}
func (p *persistanceStbErr) SelectByID(string) (*makerEntity.Maker, error) {
	return &makerEntity.Maker{
		ID:   "sampleid1",
		Name: "samplename1",
	}, nil
}
func (p *persistanceStbErr) SelectAll() ([]makerEntity.Maker, error) {
	var results []makerEntity.Maker = []makerEntity.Maker{}
	return results, nil
}
func (p *persistanceStbErr) Select(string) ([]makerEntity.Maker, error) {
	var results []makerEntity.Maker = []makerEntity.Maker{}
	return results, nil
}

type persistanceStbErr1 struct{}

func (p *persistanceStbErr1) Insert(*makerEntity.Maker) (string, error) {
	return "", errors.New("error occurred")
}
func (p *persistanceStbErr1) Update(*makerEntity.Maker) error {
	return errors.New("error occurred")
}
func (p *persistanceStbErr1) Delete(string) error {
	return errors.New("error occurred")
}
func (p *persistanceStbErr1) SelectByID(string) (*makerEntity.Maker, error) {
	return &makerEntity.Maker{}, errors.New("error occurred")
}
func (p *persistanceStbErr1) SelectAll() ([]makerEntity.Maker, error) {
	var results []makerEntity.Maker = []makerEntity.Maker{}
	return results, nil
}
func (p *persistanceStbErr1) Select(string) ([]makerEntity.Maker, error) {
	var results []makerEntity.Maker = []makerEntity.Maker{}
	return results, nil
}

func TestCreate(t *testing.T) {
	type args struct {
		makerEn makerEntity.Maker
		p       persistance
	}

	mk := makerEntity.Maker{
		ID:   "",
		Name: "name1",
	}

	expectMk := makerEntity.Maker{
		ID:   "sampleId1",
		Name: "name1",
	}

	var pstb *persistanceStb = &persistanceStb{}
	a := args{
		makerEn: mk,
		p:       pstb,
	}
	var pstbErr *persistanceStbErr = &persistanceStbErr{}
	b := args{
		makerEn: mk,
		p:       pstbErr,
	}

	tests := []struct {
		name string
		args args
		want *makerEntity.Maker
	}{
		{
			name: "create successfull",
			args: a,
			want: &expectMk,
		},
		{
			name: "create unsuccessfull",
			args: b,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.makerEn, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		makerEn makerEntity.Maker
		p       persistance
	}

	mk := makerEntity.Maker{
		ID:   "sampleId1",
		Name: "name1",
	}

	var pstb *persistanceStb = &persistanceStb{}
	a := args{
		makerEn: mk,
		p:       pstb,
	}
	var pstbErr *persistanceStbErr = &persistanceStbErr{}
	b := args{
		makerEn: mk,
		p:       pstbErr,
	}

	tests := []struct {
		name string
		args args
		want *makerEntity.Maker
	}{
		{
			name: "update successfull",
			args: a,
			want: &mk,
		},
		{
			name: "update unsuccessfull",
			args: b,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Update(tt.args.makerEn, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		idList []string
		p      persistance
	}

	idList := []string{"aaa", "bbb"}
	var pstb *persistanceStb = &persistanceStb{}
	a := args{
		idList: idList,
		p:      pstb,
	}
	var pstbErr *persistanceStbErr = &persistanceStbErr{}
	b := args{
		idList: idList,
		p:      pstbErr,
	}

	tests := []struct {
		name string
		args args
		want *[]makerEntity.Maker
	}{
		{
			name: "delete successfull",
			args: a,
			want: &[]makerEntity.Maker{},
		},
		{
			name: "delete unsuccessfull",
			args: b,
			want: &[]makerEntity.Maker{
				{
					ID:   "sampleid1",
					Name: "samplename1",
				},
				{
					ID:   "sampleid1",
					Name: "samplename1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args.idList, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindByID(t *testing.T) {
	type args struct {
		id string
		p  persistance
	}

	id := "aaa"
	var pstb *persistanceStb = &persistanceStb{}
	a := args{
		id: id,
		p:  pstb,
	}
	var pstbErr *persistanceStbErr1 = &persistanceStbErr1{}
	b := args{
		id: id,
		p:  pstbErr,
	}

	tests := []struct {
		name string
		args args
		want *makerEntity.Maker
	}{
		{
			name: "FindByID successful",
			args: a,
			want: &makerEntity.Maker{
				ID:   "sampleid1",
				Name: "samplename1",
			},
		},
		{
			name: "FindByID unsuccessful",
			args: b,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindByID(tt.args.id, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
