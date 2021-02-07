package repositories

import (
	"context"
	"errors"
	makerEntity "regulus/app/domain/entities"
	"regulus/app/infrastructures/sqlboiler"
	"regulus/app/utils"
	"sort"

	"regulus/app/domain/query"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// MakerRepo repository struct
type MakerRepo struct {
	database db
}

// Select select maker data by condition from database
func (m *MakerRepo) Select(queryItems ...*query.SearchConditionItem) ([]makerEntity.Maker, error) {
	meSlice := []makerEntity.Maker{}
	queries := []qm.QueryMod{}
	var q qm.QueryMod

	err := m.database.WithDbContext(func(db *sqlx.DB) error {
		q = qm.Where(sqlboiler.MakerColumns.ID + " IS NOT NULL")

		for _, queryItem := range queryItems {
			q = qm.Expr(q, m.createQueryMod(queryItem))
		}

		q = qm.Expr(q, qm.And(sqlboiler.MakerColumns.Del+"!=?", true))

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

// SelectByIDs select maker data by id list from database
func (m *MakerRepo) SelectByIDs(ids []string) ([]makerEntity.Maker, error) {
	if len(ids) == 0 {
		return nil, errors.New("id list must be required")
	}

	meSlice := []makerEntity.Maker{}

	var convertedIDs []interface{} = make([]interface{}, len(ids))
	for i, d := range ids {
		convertedIDs[i] = d
	}

	err := m.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.MakerColumns.Del+" !=?", true),
		}

		q := qm.AndIn(sqlboiler.MakerColumns.ID+" in ?", convertedIDs...)

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
		maker.Del = true
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

func (m *MakerRepo) createQueryMod(queryItem *query.SearchConditionItem) qm.QueryMod {

	mt, val := comparisonOperator(queryItem.MatchType, queryItem.ConditionValue)

	switch queryItem.SearchField.ID {
	default:
		if queryItem.Operator == "or" {
			return qm.Or(m.columnName(queryItem.SearchField.ID)+" "+mt+" ?", val)
		}
		//queryItem.Operator = "and"
		return qm.And(m.columnName(queryItem.SearchField.ID)+" "+mt+" ?", val)
	}
}

/*
Sort is sort maker slice by orderItems
*/
func Sort(makers []makerEntity.Maker, orderItems []query.OrderConditionItem) []makerEntity.Maker {
	sort.Slice(makers, func(i int, j int) bool {
		return compare(makers[i], makers[j], orderItems, 0)
	})
	return makers
}

func compare(maker1 makerEntity.Maker, maker2 makerEntity.Maker, orderItems []query.OrderConditionItem, orderIndex int) bool {

	if len(orderItems) <= orderIndex {
		return false
	}

	switch orderItems[orderIndex].OrderField.ID {
	case "maker-name": // 基本的にはこれしか選択されてない
		if maker1.Name == maker2.Name {
			orderIndex++
			return compare(maker1, maker2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return maker1.Name > maker2.Name
		}
		return maker1.Name < maker2.Name

	default:
		if maker1.Name == maker2.Name {
			orderIndex++
			return compare(maker1, maker2, orderItems, orderIndex)
		}
		if orderItems[orderIndex].OrderFieldKeyWord == query.Desc {
			return maker1.Name > maker2.Name
		}
		return maker1.Name < maker2.Name

	}
}

// NewMakerRepo constructor
func NewMakerRepo() *MakerRepo {
	return &MakerRepo{database: createDB()}
}
