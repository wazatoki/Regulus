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
	var beforeQueryItem *query.Item

	switch len(queryItems) {
	case 0:
		return m.SelectAll()
	case 1: // todo
		queries = append(queries, qm.Where(sqlboiler.MakerColumns.Del+"!=?", true))
	default:

	}

	for _, queryItem := range queryItems {
		if beforeQueryItem != nil {

		}

		beforeQueryItem = queryItem
	}
	return meSlice, nil
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

func createQueryMod(queryItem *query.Item) qm.QueryMod {
	mt := queryItem.MatchType

	switch mt {
	case "match":
		return sqlboiler.MakerWhere.Name.EQ(queryItem.StringValue)
	case "unmatch":
		return sqlboiler.MakerWhere.Name.NEQ(queryItem.StringValue)
	default:
		return qm.Where(sqlboiler.MakerColumns.Name+"like", "%"+queryItem.StringValue+"%")
	}
}

// NewMakerRepo constructor
func NewMakerRepo() *MakerRepo {
	return &MakerRepo{database: createDB()}
}
