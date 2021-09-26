package repositories

import (
	"context"
	"errors"
	"regulus/app/domain"
	"regulus/app/infrastructures/sqlboiler"
	"regulus/app/utils"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Dalete delete data to database
func (g *StaffGroupRepo) Delete(id string, operatorID string) error {
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
func (g *StaffGroupRepo) Update(staffGroup *domain.StaffGroup, operatorID string) (err error) {
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
func (g *StaffGroupRepo) Insert(staffGroup *domain.StaffGroup, operatorID string) (id string, err error) {
	id = ""
	sqlStaffGroup := &sqlboiler.StaffGroup{
		ID:            utils.CreateID(),
		CreStaffID:    null.StringFrom(operatorID),
		UpdateStaffID: null.StringFrom(operatorID),
		Name:          staffGroup.Name,
	}

	err = g.database.WithDbContext(func(db *sqlx.DB) error {

		var err = sqlStaffGroup.Insert(context.Background(), db.DB, boil.Infer())
		return err
	})

	if err == nil {
		id = sqlStaffGroup.ID
	}

	return
}

// SelectByIDs select staff data by id list from database
func (g *StaffGroupRepo) SelectByIDs(ids []string) (staffGroups domain.StaffGroups, err error) {
	staffGroups = domain.StaffGroups{}
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
func (g *StaffGroupRepo) SelectByID(id string) (staffGroup *domain.StaffGroup, err error) {
	if id == "" {
		return &domain.StaffGroup{}, errors.New("id must be required")
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
func (g *StaffGroupRepo) SelectAll() (domain.StaffGroups, error) {
	geSlice := domain.StaffGroups{}

	err := g.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffGroupColumns.Del+"!=?", true),
		}

		groups, err := sqlboiler.StaffGroups(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, group := range groups {

				var ge = &domain.StaffGroup{}

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
func (g *StaffGroupRepo) Select(queryItems ...domain.SearchConditionItem) (result domain.StaffGroups, err error) {

	err = g.database.WithDbContext(func(db *sqlx.DB) error {
		var args []interface{} = make([]interface{}, 0)
		ids := []string{}

		queryStr := "select distinct sg.id " +
			"from staff_groups sg " +
			"left join join_staffs_staff_groups jsg on sg.id = jsg.staff_groups_id " +
			"left join staffs s on jsg.staffs_id = s.id " +
			"where sg.del != true"

		// 条件構築
		searchConditionItems := []domain.SearchConditionItem{}

		searchConditionItems = append(searchConditionItems, queryItems...)

		for _, searchConditionItem := range searchConditionItems {
			qu, pslice := g.createQueryModWhere(searchConditionItem)
			queryStr += qu

			for _, p := range pslice {
				args = append(args, p)
			}
		}

		// クエリをDBドライバに併せて再構築
		queryStr = db.Rebind(queryStr)

		// データ取得処理
		db.Select(&ids, queryStr, args...)
		var convertedIDs []interface{} = make([]interface{}, len(ids))
		for i, d := range ids {
			convertedIDs[i] = d
		}
		queries := g.createQueryModSlice()
		queries = append(
			queries,
			qm.And("staff_groups."+sqlboiler.StaffGroupColumns.Del+" != ?", true),
			qm.AndIn("staff_groups."+sqlboiler.StaffGroupColumns.ID+" in ?", convertedIDs...),
		)
		fetchedStaffGroups, err := sqlboiler.StaffGroups(queries...).All(context.Background(), db.DB)

		if err == nil {
			for _, fg := range fetchedStaffGroups {
				result = append(result, StaffGroupObjectMap(fg))
			}
		}
		return err
	})

	return
}

func (g *StaffGroupRepo) createQueryModWhere(queryItem domain.SearchConditionItem) (string, []string) {

	mt, val := comparisonOperator(queryItem.MatchType, queryItem.ConditionValue)

	switch queryItem.SearchField.ID {
	case "name":
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return " or sg." + sqlboiler.StaffGroupColumns.Name + " " + mt + " ?", []string{val}
		}
		return " and sg." + sqlboiler.StaffGroupColumns.Name + " " + mt + " ?", []string{val}
	case "staff-name":
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return " or s." + sqlboiler.StaffColumns.Name + " " + mt + " ?", []string{val}
		}
		return " and s." + sqlboiler.StaffColumns.Name + " " + mt + " ?", []string{val}
	case "staff-account-id":
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return " or s." + sqlboiler.StaffColumns.AccountID + " " + mt + " ?", []string{val}
		}
		return " and s." + sqlboiler.StaffColumns.AccountID + " " + mt + " ?", []string{val}
	default:
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return " or sg." + sqlboiler.StaffGroupColumns.Name + " " + mt + " ?", []string{val}
		}
		//queryItem.Operator = "and"
		return " and sg." + sqlboiler.StaffGroupColumns.Name + " " + mt + " ?", []string{val}
	}
}

func (g *StaffGroupRepo) createQueryModSlice() (qslice []qm.QueryMod) {
	qslice = []qm.QueryMod{}
	qslice = append(
		qslice,
		qm.Select("distinct staff_groups.*"),
		qm.Where("staff_groups."+sqlboiler.StaffGroupColumns.ID+" IS NOT NULL"),
		qm.Load(qm.Rels(sqlboiler.StaffGroupRels.Staffs), qm.Where("del != true")),
	)
	return
}

// StaffGroupObjectMap data mapper sqlboiler object to entities object
func StaffGroupObjectMap(sg *sqlboiler.StaffGroup) (eg *domain.StaffGroup) {

	if sg == nil {
		return nil
	}
	eg = &domain.StaffGroup{
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
