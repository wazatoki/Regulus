package repositories

import (
	"context"
	"errors"
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/sqlboiler"
	"regulus/app/utils"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// Dalete delete data to database
func (g *StaffGroupRepo) Dalete(id string, operatorID string) error {
	if id == "" {
		return errors.New("id must be required")
	}

	err := g.database.WithDbContext(func(db *sqlx.DB) error {
		sqlStaffGroup, _ := sqlboiler.FindStaffGroup(context.Background(), db.DB, id)
		sqlStaffGroup.Del = true
		sqlStaffGroup.UpdateStaffID = null.StringFrom(operatorID)
		var err error
		_, err = sqlStaffGroup.Update(context.Background(), db.DB, boil.Infer())
		return err
	})

	return err
}

// Update update data to database
func (g *StaffGroupRepo) Update(staffGroup *entities.StaffGroup, operatorID string) (err error) {
	if staffGroup.ID == "" {
		return errors.New("ID must be required")
	}

	sqlStaffGroup := &sqlboiler.StaffGroup{
		ID:            staffGroup.ID,
		UpdateStaffID: null.StringFrom(operatorID),
		Name:          staffGroup.Name,
	}

	err = g.database.WithDbContext(func(db *sqlx.DB) error {
		var err error
		_, err = sqlStaffGroup.Update(context.Background(), db.DB, boil.Infer())
		return err
	})

	return err
}

// Insert insert data to database
func (g *StaffGroupRepo) Insert(staffGroup *entities.StaffGroup, operatorID string) (id string, err error) {
	id = ""
	sqlStaffGroup := &sqlboiler.StaffGroup{
		ID:            utils.CreateID(),
		CreStaffID:    null.StringFrom(operatorID),
		UpdateStaffID: null.StringFrom(operatorID),
		Name:          staffGroup.Name,
	}

	err = g.database.WithDbContext(func(db *sqlx.DB) error {
		var err error

		err = sqlStaffGroup.Insert(context.Background(), db.DB, boil.Infer())
		return err
	})

	if err == nil {
		id = sqlStaffGroup.ID
	}

	return
}

// SelectByIDs select staff data by id list from database
func (g *StaffGroupRepo) SelectByIDs(ids []string) (staffGroups []*entities.StaffGroup, err error) {
	staffGroups = []*entities.StaffGroup{}
	if len(ids) == 0 {
		return nil, errors.New("id list must be required")
	}

	var convertedIDs []interface{} = make([]interface{}, len(ids))
	for i, d := range ids {
		convertedIDs[i] = d
	}

	err = g.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffGroupColumns.Del+" !=?", true),
			qm.AndIn(sqlboiler.StaffGroupColumns.ID+" in ?", convertedIDs...),
		}

		fetchedStaffGroups, err := sqlboiler.StaffGroups(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, fs := range fetchedStaffGroups {
				staffGroups = append(staffGroups, StaffGroupObjectMap(fs))
			}
		}

		return err
	})

	return
}

// SelectByID select staaffGroup data by id from database
func (g *StaffGroupRepo) SelectByID(id string) (staffGroup *entities.StaffGroup, err error) {
	if id == "" {
		return &entities.StaffGroup{}, errors.New("id must be required")
	}

	err = g.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffGroupColumns.Del+" !=?", true),
			qm.And(sqlboiler.StaffGroupColumns.ID+" =?", id),
		}
		fetchedStaffGroup, err := sqlboiler.StaffGroups(queries...).One(context.Background(), db.DB)
		if err == nil {
			staffGroup = StaffGroupObjectMap(fetchedStaffGroup)
		}

		return err
	})

	return
}

// SelectAll select all group data without not del from database
func (g *StaffGroupRepo) SelectAll() ([]*entities.StaffGroup, error) {
	geSlice := []*entities.StaffGroup{}

	err := g.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffGroupColumns.Del+"!=?", true),
		}

		groups, err := sqlboiler.StaffGroups(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, group := range groups {
				var ge *entities.StaffGroup
				ge = &entities.StaffGroup{}

				ge.ID = group.ID
				ge.Name = group.Name

				geSlice = append(geSlice, ge)
			}
		}

		return err
	})

	return geSlice, err
}

// Select select staffGroup data by condition from database
func (g *StaffGroupRepo) Select(queryItems ...*query.SearchConditionItem) ([]*entities.StaffGroup, error) {
	staffGroups := []*entities.StaffGroup{}
	queries := g.createQueryModSlice()
	var q qm.QueryMod

	err := g.database.WithDbContext(func(db *sqlx.DB) error {
		q = qm.Where("staff_groups.id IS NOT NULL")

		for _, queryItem := range queryItems {
			q = qm.Expr(q, g.createQueryModWhere(queryItem))
		}

		q = qm.Expr(q, qm.And("staff_groups.del != ?", true))

		queries = append(queries, q)

		fetchedStaffGroups, err := sqlboiler.StaffGroups(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, fsg := range fetchedStaffGroups {
				staffGroups = append(staffGroups, StaffGroupObjectMap(fsg))
			}
		}

		return err
	})

	return staffGroups, err
}

func (g *StaffGroupRepo) createQueryModWhere(queryItem *query.SearchConditionItem) qm.QueryMod {

	mt, val := comparisonOperator(queryItem.MatchType, queryItem.ConditionValue)

	switch queryItem.SearchField.ID {
	case "name":
		if queryItem.Operator == query.Or {
			return qm.Or("staff_groups.name "+mt+" ?", val)
		}
		return qm.And("staff_groups.name "+mt+" ?", val)
	case "staff-name":
		if queryItem.Operator == query.Or {
			return qm.Or("s.name "+mt+" ?", val)
		}
		return qm.And("s.name "+mt+" ?", val)
	case "staff-account-id":
		if queryItem.Operator == query.Or {
			return qm.Or("s.account_id "+mt+" ?", val)
		}
		return qm.And("s.account_id "+mt+" ?", val)
	default:
		if queryItem.Operator == "or" {
			return qm.Or("staff_groups.name "+mt+" ?", val)
		}
		//queryItem.Operator = "and"
		return qm.And("staff_groups.name "+mt+" ?", val)
	}
}

func (g *StaffGroupRepo) createQueryModSlice() (qslice []qm.QueryMod) {
	qslice = []qm.QueryMod{}
	qslice = append(
		qslice,
		qm.Select("distinct staff_groups.*"),
		qm.InnerJoin("join_staffs_staff_groups jsg on staff_groups.id = jsg.staff_groups_id"),
		qm.InnerJoin("staffs s on jsg.staffs_id = s.id"),
	)
	return
}

// StaffGroupObjectMap data mapper sqlboiler object to entities object
func StaffGroupObjectMap(sg *sqlboiler.StaffGroup) (eg *entities.StaffGroup) {

	if sg == nil {
		return nil
	}
	eg = &entities.StaffGroup{
		ID:   sg.ID,
		Name: sg.Name,
	}
	return
}

// NewStaffGroupRepo constructor
func NewStaffGroupRepo() *StaffGroupRepo {
	return &StaffGroupRepo{database: createDB()}
}

// StaffGroupRepo repository struct
type StaffGroupRepo struct {
	database db
}
