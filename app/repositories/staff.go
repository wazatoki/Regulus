package repositories

import (
	"context"
	"errors"
	"regulus/app/domain"
	"regulus/app/infrastructures/sqlboiler"
	"regulus/app/utils"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Dalete delete data to database
func (s *StaffRepo) Dalete(id string, operatorID string) error {
	if id == "" {
		return errors.New("id must be required")
	}

	err := s.database.WithDbContext(func(db *sqlx.DB) error {
		sqlStaff, _ := sqlboiler.FindStaff(context.Background(), db.DB, id)
		sqlStaff.Del = true
		sqlStaff.UpdateStaffID = null.StringFrom(operatorID)
		var err error
		_, err = sqlStaff.Update(context.Background(), db.DB, boil.Infer())
		return err
	})

	return err
}

// Update update data to database
func (s *StaffRepo) Update(staff *domain.Staff, operatorID string) (err error) {
	if staff.ID == "" {
		return errors.New("ID must be required")
	}

	sqlStaff := &sqlboiler.Staff{
		ID:            staff.ID,
		UpdateStaffID: null.StringFrom(operatorID),
		AccountID:     staff.AccountID,
		Name:          staff.Name,
		Password:      staff.Password,
	}

	sqlStaffGroups := make([]*sqlboiler.StaffGroup, len(staff.StaffGroups))
	for i, g := range staff.StaffGroups {
		sqlStaffGroups[i] = &sqlboiler.StaffGroup{
			ID: g.ID,
		}
	}

	err = s.database.WithDbContext(func(db *sqlx.DB) error {
		var err error
		_, err = sqlStaff.Update(context.Background(), db.DB, boil.Infer())
		sqlStaff.SetStaffGroups(context.Background(), db.DB, false, sqlStaffGroups...)
		return err
	})

	return err
}

// Insert insert data to database
func (s *StaffRepo) Insert(staff *domain.Staff, operatorID string) (id string, err error) {
	id = ""
	sqlStaff := &sqlboiler.Staff{
		ID:            utils.CreateID(),
		CreStaffID:    null.StringFrom(operatorID),
		UpdateStaffID: null.StringFrom(operatorID),
		AccountID:     staff.AccountID,
		Name:          staff.Name,
		Password:      staff.Password,
	}

	sqlStaffGroups := make([]*sqlboiler.StaffGroup, len(staff.StaffGroups))
	for i, g := range staff.StaffGroups {
		sqlStaffGroups[i] = &sqlboiler.StaffGroup{
			ID: g.ID,
		}
	}

	err = s.database.WithDbContext(func(db *sqlx.DB) error {

		var err = sqlStaff.Insert(context.Background(), db.DB, boil.Infer())
		sqlStaff.SetStaffGroups(context.Background(), db.DB, false, sqlStaffGroups...)
		return err
	})

	if err == nil {
		id = sqlStaff.ID
	}

	return
}

// SelectByIDs select staff data by id list from database
func (s *StaffRepo) SelectByIDs(ids []string) (staffs []*domain.Staff, err error) {
	staffs = []*domain.Staff{}
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
			qm.AndIn(sqlboiler.StaffColumns.ID+" in ?", convertedIDs...),
			qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true")),
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
func (s *StaffRepo) SelectByID(id string) (staff *domain.Staff, err error) {
	if id == "" {
		return nil, errors.New("id must be required")
	}

	err = s.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffColumns.Del+" !=?", true),
			qm.And(sqlboiler.StaffColumns.ID+" =?", id),
			qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true")),
		}
		fetchedStaff, err := sqlboiler.Staffs(queries...).One(context.Background(), db.DB)
		if err == nil {
			staff = StaffObjectMap(fetchedStaff)
		}

		return err
	})

	return
}

// SelectByAccountID select staaff data by accountID from database
func (s *StaffRepo) SelectByAccountID(id string) (staff *domain.Staff, err error) {
	if id == "" {
		return nil, errors.New("accountID must be required")
	}

	err = s.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffColumns.Del+" !=?", true),
			qm.And(sqlboiler.StaffColumns.AccountID+" =?", id),
			qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true")),
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
func (s *StaffRepo) SelectAll() (staffs []*domain.Staff, err error) {
	staffs = []*domain.Staff{}
	err = s.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffColumns.Del+"!=?", true),
			qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true")),
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
func (s *StaffRepo) Select(queryItems ...*domain.SearchConditionItem) (staffs []*domain.Staff, err error) {
	staffs = []*domain.Staff{}
	queries := s.createQueryModSlice()
	var q qm.QueryMod

	err = s.database.WithDbContext(func(db *sqlx.DB) error {
		q = qm.Where("staffs." + sqlboiler.StaffColumns.ID + " IS NOT NULL")

		for _, queryItem := range queryItems {
			q = qm.Expr(q, s.createQueryModWhere(queryItem))
		}

		q = qm.Expr(q, qm.And("staffs."+sqlboiler.StaffColumns.Del+" != ?", true))

		queries = append(queries, q, qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true")))

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

func (s *StaffRepo) createQueryModWhere(queryItem *domain.SearchConditionItem) qm.QueryMod {

	mt, val := comparisonOperator(queryItem.MatchType, queryItem.ConditionValue)

	switch queryItem.SearchField.ID {
	case "account-id":
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return qm.Or("staffs."+sqlboiler.StaffColumns.AccountID+" "+mt+" ?", val)
		}
		return qm.And("staffs."+sqlboiler.StaffColumns.AccountID+" "+mt+" ?", val)
	case "name":
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return qm.Or("staffs."+sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
		}
		return qm.And("staffs."+sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
	case "groups":
		var ids []interface{}
		json.Unmarshal([]byte(val), &ids)
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return qm.OrIn("sg.id"+" "+mt+" ?", ids...)
		}
		return qm.AndIn("sg.id"+" "+mt+" ?", ids...)
	case "group-name":
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return qm.Or("sg.name"+" "+mt+" ?", val)
		}
		return qm.And("sg.name"+" "+mt+" ?", val)
	default:
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
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
func StaffObjectMap(ss *sqlboiler.Staff) (es *domain.Staff) {

	if ss == nil {
		return nil
	}

	groups := domain.StaffGroups{}
	for _, group := range ss.R.StaffGroups {
		groups = append(groups, StaffGroupObjectMap(group))
	}
	conditions := domain.Conditions{}
	es = &domain.Staff{
		ID:                      ss.ID,
		AccountID:               ss.AccountID,
		Name:                    ss.Name,
		Password:                ss.Password,
		StaffGroups:             groups,
		OeratorUsableConditions: conditions,
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
