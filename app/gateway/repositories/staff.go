package repositories

import (
	"context"
	"errors"
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/sqlboiler"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// SelectByIDs select staff data by id list from database
func (s *StaffRepo) SelectByIDs(ids []string) (staffs []entities.Staff, err error) {
	if len(ids) == 0 {
		return nil, errors.New("id list must be required")
	}

	var convertedIDs []interface{} = make([]interface{}, len(ids))
	for i, d := range ids {
		convertedIDs[i] = d
	}

	err = s.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffColumns.Del+" !=?", true),
			qm.AndIn(sqlboiler.MakerColumns.ID+" in ?", convertedIDs...),
			qm.Load(sqlboiler.StaffRels.StaffGroups),
		}

		fetchedStaffs, err := sqlboiler.Staffs(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, fs := range fetchedStaffs {
				staffs = append(staffs, StaffObjectMap(fs))
			}
		}

		return err
	})

	return
}

// SelectByID select staaff data by id from database
func (s *StaffRepo) SelectByID(id string) (staff entities.Staff, err error) {
	if id == "" {
		return entities.Staff{}, errors.New("id must be required")
	}

	err = s.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffColumns.Del+" !=?", true),
			qm.And(sqlboiler.StaffColumns.ID+" =?", id),
			qm.Load(sqlboiler.StaffRels.StaffGroups),
		}
		fetchedStaff, err := sqlboiler.Staffs(queries...).One(context.Background(), db.DB)
		if err == nil {
			staff = StaffObjectMap(fetchedStaff)
		}

		return err
	})

	return
}

// SelectAll select all staff data without not del from database
func (s *StaffRepo) SelectAll() (staffs []entities.Staff, err error) {

	err = s.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffColumns.Del+"!=?", true),
			qm.Load(sqlboiler.StaffRels.StaffGroups),
		}

		fetchedStaffs, err := sqlboiler.Staffs(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, fs := range fetchedStaffs {
				staffs = append(staffs, StaffObjectMap(fs))
			}
		}

		return err
	})

	return
}

// Select select staff data by condition from database
func (s *StaffRepo) Select(queryItems ...*query.SearchConditionItem) (staffs []entities.Staff, err error) {
	queries := s.createQueryModSlice()
	var q qm.QueryMod

	err = s.database.WithDbContext(func(db *sqlx.DB) error {
		q = qm.Where("staffs." + sqlboiler.StaffColumns.ID + " IS NOT NULL")

		for _, queryItem := range queryItems {
			q = qm.Expr(q, s.createQueryModWhere(queryItem))
		}

		q = qm.Expr(q, qm.And("staffs."+sqlboiler.StaffColumns.Del+" != ?", true))

		queries = append(queries, q, qm.Load(sqlboiler.StaffRels.StaffGroups))

		fetchedStaffs, err := sqlboiler.Staffs(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, fs := range fetchedStaffs {
				staffs = append(staffs, StaffObjectMap(fs))
			}
		}

		return err
	})

	return
}

func (s *StaffRepo) createQueryModWhere(queryItem *query.SearchConditionItem) qm.QueryMod {

	mt, val := comparisonOperator(queryItem.MatchType, queryItem.ConditionValue)

	switch queryItem.SearchField.ID {
	case "account-id":
		if queryItem.Operator == query.Or {
			return qm.Or("staffs."+sqlboiler.StaffColumns.AccountID+" "+mt+" ?", val)
		}
		return qm.And("staffs."+sqlboiler.StaffColumns.AccountID+" "+mt+" ?", val)
	case "name":
		if queryItem.Operator == query.Or {
			return qm.Or("staffs."+sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
		}
		return qm.And("staffs."+sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
	case "groups":
		var ids []interface{}
		json.Unmarshal([]byte(val), &ids)
		if queryItem.Operator == query.Or {
			return qm.OrIn("sg.id"+" "+mt+" ?", ids...)
		}
		return qm.AndIn("sg.id"+" "+mt+" ?", ids...)
	case "group-name":
		if queryItem.Operator == query.Or {
			return qm.Or("sg.name"+" "+mt+" ?", val)
		}
		return qm.And("sg.name"+" "+mt+" ?", val)
	default:
		if queryItem.Operator == query.Or {
			return qm.Or("staffs."+sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
		}
		// queryItem.Operator == and
		return qm.And("staffs."+sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
	}
}

func (s *StaffRepo) createQueryModSlice() (qslice []qm.QueryMod) {
	qslice = []qm.QueryMod{}
	qslice = append(
		qslice,
		qm.Select("distinct staffs.*"),
		qm.InnerJoin("join_staffs_staff_groups jsg on staffs.id = jsg.staffs_id"),
		qm.InnerJoin("staff_groups sg on jsg.staff_groups_id = sg.id"),
	)
	return
}

// StaffObjectMap data mapper sqlboiler object to entities object
func StaffObjectMap(ss *sqlboiler.Staff) (es entities.Staff) {
	groups := []entities.StaffGroup{}
	for _, group := range ss.R.StaffGroups {
		groups = append(groups, StaffGroupObjectMap(group))
	}
	es = entities.Staff{
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
