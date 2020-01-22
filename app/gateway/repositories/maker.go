package repositories

import (
	"context"
	"errors"
	makerEntity "regulus/app/domain/entities/maker"
	"regulus/app/infrastructures/sqlboiler"
	"regulus/app/utils"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"
	// . "github.com/volatiletech/sqlboiler/queries/qm"
)

// MakerRepo repository struct
type MakerRepo struct {
	database db
}

// Update insert data to database
func (m *MakerRepo) Update(makerEntity *makerEntity.Maker) error {
	if makerEntity.ID == "" {
		return errors.New("ID must be required")
	}

	err := m.database.WithDbContext(func(db *sqlx.DB) error {
		maker, _ := sqlboiler.FindMaker(context.Background(), db.DB, makerEntity.ID)
		maker.Name = makerEntity.Name
		var err error
		_, err = maker.Update(context.Background(), db.DB, boil.Infer())
		return err
	})

	return err
}

// Insert insert data to database
func (m *MakerRepo) Insert(makerEntity *makerEntity.Maker) (string, error) {
	id := ""
	maker := &sqlboiler.Maker{}
	maker.ID = utils.CreateID()
	maker.Name = makerEntity.Name

	err := m.database.WithDbContext(func(db *sqlx.DB) error {
		var err error
		err = maker.Insert(context.Background(), db.DB, boil.Infer())
		return err
	})

	if err == nil {
		id = maker.ID
	}

	return id, err
}

// NewMakerRepo constructor
func NewMakerRepo() *MakerRepo {
	return &MakerRepo{database: createDB()}
}
