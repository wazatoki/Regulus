package repositories

import (
	"context"
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/sqlboiler"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// SelectAll select all group data without not del from database
func (g *StaffGroupRepo) SelectAll() ([]entities.StaffGroup, error) {
	geSlice := []entities.StaffGroup{}

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

				geSlice = append(geSlice, *ge)
			}
		}

		return err
	})

	return geSlice, err
}

// Select select maker data by condition from database
func (g *StaffGroupRepo) Select(queryItems ...*query.SearchConditionItem) ([]entities.StaffGroup, error) {
	staffGroups := []entities.StaffGroup{}
	queries := g.createQueryModSlice()
	var q qm.QueryMod

	err := g.database.WithDbContext(func(db *sqlx.DB) error {
		q = qm.Where(sqlboiler.StaffGroupColumns.ID + " IS NOT NULL")

		for _, queryItem := range queryItems {
			q = qm.Expr(q, g.createQueryModWhere(queryItem))
		}

		q = qm.Expr(q, qm.And(sqlboiler.StaffGroupColumns.Del+" != ?", true))

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
			return qm.Or(sqlboiler.StaffGroupColumns.Name+" "+mt+" ?", val)
		}
		return qm.And(sqlboiler.StaffGroupColumns.Name+" "+mt+" ?", val)
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
			return qm.Or(sqlboiler.StaffGroupColumns.Name+" "+mt+" ?", val)
		}
		//queryItem.Operator = "and"
		return qm.And(sqlboiler.StaffGroupColumns.Name+" "+mt+" ?", val)
	}
}

func (g *StaffGroupRepo) createQueryModSlice() (qslice []qm.QueryMod) {
	qslice = []qm.QueryMod{}
	qslice = append(
		qslice,
		qm.InnerJoin("join_staffs_staff_groups jsg on staff_groups.id = jsg.staff_groups_id"),
		qm.InnerJoin("staffs s on jsg.staffs_id = s.id"),
	)
	return
}

// StaffGroupObjectMap data mapper sqlboiler object to entities object
func StaffGroupObjectMap(sg *sqlboiler.StaffGroup) (eg entities.StaffGroup) {
	eg = entities.StaffGroup{
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
