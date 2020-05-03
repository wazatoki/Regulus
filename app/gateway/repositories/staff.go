package repositories

import (
	"context"
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/sqlboiler"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// Select select maker data by condition from database
func (s *StaffRepo) Select(queryItems ...*query.SearchConditionItem) ([]*entities.Staff, error) {
	staffs := []*entities.Staff{}
	queries := s.createQueryModSlice()
	var q qm.QueryMod

	err := s.database.WithDbContext(func(db *sqlx.DB) error {
		q = qm.Where(sqlboiler.StaffColumns.ID + " IS NOT NULL")

		for _, queryItem := range queryItems {
			q = qm.Expr(q, s.createQueryModWhere(queryItem))
		}

		q = qm.Expr(q, qm.And(sqlboiler.StaffColumns.Del+" != ?", true))

		queries = append(queries, q)

		fetchedStaffs, err := sqlboiler.Staffs(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, fs := range fetchedStaffs {
				staffs = append(staffs, StaffObjectMap(fs))
			}
		}

		return err
	})

	return staffs, err
}

func (s *StaffRepo) createQueryModSlice() (qslice []qm.QueryMod) {
	qslice = []qm.QueryMod{}
	qslice = append(
		qslice,
		qm.InnerJoin("join_staffs_staff_groups jsg on staffs.id = jsg.staffs_id"),
		qm.InnerJoin("staff_groups sg on jsg.staff_groups_id = sg.id"),
	)
	return
}

func (s *StaffRepo) createQueryModWhere(queryItem *query.SearchConditionItem) qm.QueryMod {

	mt, val := comparisonOperator(queryItem.MatchType, queryItem.ConditionValue)

	switch queryItem.SearchField.ID {
	case "staff-account-id":
		if queryItem.Operator == query.Or {
			return qm.Or(sqlboiler.StaffColumns.AccountID+" "+mt+" ?", val)
		}
		return qm.And(sqlboiler.StaffColumns.AccountID+" "+mt+" ?", val)
	case "staff-name":
		if queryItem.Operator == query.Or {
			return qm.Or(sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
		}
		return qm.And(sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
	case "staff-groups":
		var ids []string
		json.Unmarshal([]byte(val), &ids)
		if queryItem.Operator == query.Or {
			return qm.OrIn("sg.id"+" "+mt+" ?", ids)
		}
		return qm.AndIn("sg.id"+" "+mt+" ?", ids)
	case "staff-group-names":
		if queryItem.Operator == query.Or {
			return qm.Or("sg.name"+" "+mt+" ?", val)
		}
		return qm.And("sg.name"+" "+mt+" ?", val)
	default:
		if queryItem.Operator == "or" {
			return qm.Or(sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
		}
		//queryItem.Operator = "and"
		return qm.And(sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
	}
}

// StaffObjectMap data mapper sqlboiler object to entities object
func StaffObjectMap(ss *sqlboiler.Staff) (es *entities.Staff) {
	groups := []entities.StaffGroup{}
	for _, group := range ss.R.StaffGroups {
		groups = append(groups, StaffGroupObjectMap(group))
	}
	es = &entities.Staff{
		ID:        ss.ID,
		AccountID: ss.AccountID,
		Name:      ss.Name,
		Password:  ss.Password,
		Groups:    groups,
	}
	return
}

// NewStaffRepo constructor
func NewStaffRepo() *StaffRepo {
	return &StaffRepo{database: createDB()}
}

// StaffRepo repository struct
type StaffRepo struct {
	database db
}
