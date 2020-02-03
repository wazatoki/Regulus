package repositories

import (
	"context"
	"errors"
	makerEntity "regulus/app/domain/entities/maker"
	"regulus/app/infrastructures/sqlboiler"
	"regulus/app/utils"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"regulus/app/domain/vo/query"
)

// MakerRepo repository struct
type MakerRepo struct {
	database db
}

// Select select maker data by condition from database
func (m *MakerRepo) Select(queryItems ...*query.Item) ([]makerEntity.Maker, error) {
	meSlice := []makerEntity.Maker{}
	queries := []qm.QueryMod{}
	var q qm.QueryMod

	err := m.database.WithDbContext(func(db *sqlx.DB) error {
		q = qm.Where(sqlboiler.MakerColumns.Del+"!=?", true)

		for _, queryItem := range queryItems {
			q = qm.Expr(q, m.createQueryMod(queryItem))
		}

		queries = append(queries, q)

		makers, err := sqlboiler.Makers(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, maker := range makers {
				var me *makerEntity.Maker
				me = &makerEntity.Maker{}

				me.ID = maker.ID
				me.Name = maker.Name

				meSlice = append(meSlice, *me)
			}
		}

		return err
	})

	return meSlice, err
}

// SelectAll select all maker data without not del from database
func (m *MakerRepo) SelectAll() ([]makerEntity.Maker, error) {
	meSlice := []makerEntity.Maker{}

	err := m.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.MakerColumns.Del+"!=?", true),
		}

		makers, err := sqlboiler.Makers(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, maker := range makers {
				var me *makerEntity.Maker
				me = &makerEntity.Maker{}

				me.ID = maker.ID
				me.Name = maker.Name

				meSlice = append(meSlice, *me)
			}
		}

		return err
	})

	return meSlice, err
}

// SelectByID select maker data by id from database
func (m *MakerRepo) SelectByID(id string) (*makerEntity.Maker, error) {
	if id == "" {
		return nil, errors.New("id must be required")
	}

	var me *makerEntity.Maker
	me = &makerEntity.Maker{}

	err := m.database.WithDbContext(func(db *sqlx.DB) error {
		maker, err := sqlboiler.FindMaker(context.Background(), db.DB, id)

		if err == nil {
			me.ID = maker.ID
			me.Name = maker.Name
		}

		return err
	})

	return me, err
}

// Dalete delete data to database
func (m *MakerRepo) Dalete(id string) error {
	if id == "" {
		return errors.New("id must be required")
	}

	err := m.database.WithDbContext(func(db *sqlx.DB) error {
		maker, _ := sqlboiler.FindMaker(context.Background(), db.DB, id)
		maker.Del = null.BoolFrom(true)
		var err error
		_, err = maker.Update(context.Background(), db.DB, boil.Infer())
		return err
	})

	return err
}

// Update update data to database
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

func (m *MakerRepo) columnName(fieldName string) string {
	switch fieldName {
	case "Name":
		return "name"
	default:
		return "name"
	}
}

func (m *MakerRepo) createQueryMod(queryItem *query.Item) qm.QueryMod {

	mt, val := comparisonOperator(queryItem.MatchType, queryItem.Value)

	switch queryItem.EntityName {
	default:
		if queryItem.Operator == "or" {
			return qm.Or(m.columnName(queryItem.FieldName)+" "+mt+" ?", val)
		}
		//queryItem.Operator = "and"
		return qm.And(m.columnName(queryItem.FieldName)+" "+mt+" ?", val)
	}
}

// NewMakerRepo constructor
func NewMakerRepo() *MakerRepo {
	return &MakerRepo{database: createDB()}
}
