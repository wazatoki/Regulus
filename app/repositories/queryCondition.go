package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"regulus/app/domain/authentication"
	"regulus/app/domain/query"
	"regulus/app/infrastructures/sqlboiler"
	"regulus/app/utils"
	"regulus/app/utils/log"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Delete delete data to database
func (q *QueryConditionRepo) Delete(id string, operatorID string) error {
	if id == "" {
		return errors.New("id must be required")
	}

	err := q.database.WithDbContext(func(db *sqlx.DB) error {

		updateCols := map[string]interface{}{
			sqlboiler.QueryConditionColumns.Del:           true,
			sqlboiler.QueryConditionColumns.UpdateStaffID: null.StringFrom(operatorID),
		}
		query := qm.Where(sqlboiler.QueryConditionColumns.ID+" = ?", id)
		_, err := sqlboiler.QueryConditions(query).UpdateAll(context.Background(), db.DB, updateCols)

		return err
	})

	return err
}

// Update update data to database
func (q *QueryConditionRepo) Update(queryCondition *query.Condition, operatorID string) (err error) {
	if queryCondition.ID == "" {
		return errors.New("ID must be required")
	}

	sqlQueryCondition := &sqlboiler.QueryCondition{
		ID:            queryCondition.ID,
		UpdateStaffID: null.StringFrom(operatorID),
		CategoryName:  queryCondition.Category.Name,
		IsDisclose:    queryCondition.IsDisclose,
		OwnerID:       queryCondition.Owner.ID,
		PatternName:   queryCondition.PatternName,
	}
	sqlQueryDisplayItems := make([]*sqlboiler.QueryDisplayItem,
		len(queryCondition.ConditionData.DisplayItemList))
	sqlQuerySearchConditionItems := make([]*sqlboiler.QuerySearchConditionItem,
		len(queryCondition.ConditionData.SearchConditionList))
	sqlQueryOrderConditionItems := make([]*sqlboiler.QueryOrderConditionItem,
		len(queryCondition.ConditionData.OrderConditionList))
	sqlDiscloseGroups := make([]*sqlboiler.StaffGroup, len(queryCondition.DiscloseGroups))
	for i, d := range queryCondition.ConditionData.DisplayItemList {
		sqlQueryDisplayItems[i] = &sqlboiler.QueryDisplayItem{
			ID:                utils.CreateID(),
			CreStaffID:        null.StringFrom(operatorID),
			UpdateStaffID:     null.StringFrom(operatorID),
			QueryConditionsID: sqlQueryCondition.ID,
			DisplayFieldID:    d.ID,
			RowOrder:          i,
		}
	}

	for i, s := range queryCondition.ConditionData.SearchConditionList {
		sqlQuerySearchConditionItems[i] = &sqlboiler.QuerySearchConditionItem{
			ID:                utils.CreateID(),
			CreStaffID:        null.StringFrom(operatorID),
			UpdateStaffID:     null.StringFrom(operatorID),
			QueryConditionsID: sqlQueryCondition.ID,
			SearchFieldID:     s.SearchField.ID,
			ConditionValue:    s.ConditionValue,
			MatchType:         string(s.MatchType),
			Operator:          string(s.Operator),
			RowOrder:          i,
		}
	}

	for i, o := range queryCondition.ConditionData.OrderConditionList {
		sqlQueryOrderConditionItems[i] = &sqlboiler.QueryOrderConditionItem{
			ID:                utils.CreateID(),
			CreStaffID:        null.StringFrom(operatorID),
			UpdateStaffID:     null.StringFrom(operatorID),
			QueryConditionsID: sqlQueryCondition.ID,
			OrderFieldID:      o.OrderField.ID,
			OrderFieldKeyWord: string(o.OrderFieldKeyWord),
			RowOrder:          i,
		}
	}

	for i, g := range queryCondition.DiscloseGroups {
		sqlDiscloseGroups[i] = &sqlboiler.StaffGroup{
			ID: g.ID,
		}
	}

	err = q.database.WithDbContext(func(db *sqlx.DB) error {
		var err error

		/*
		 元の検索条件を論理削除
		*/
		displayUpdateCols := map[string]interface{}{
			sqlboiler.QueryDisplayItemColumns.Del:           true,
			sqlboiler.QueryDisplayItemColumns.UpdateStaffID: null.StringFrom(operatorID),
		}
		displayItemQuery := qm.Where(sqlboiler.QueryDisplayItemColumns.QueryConditionsID+" = ?", queryCondition.ID)
		_, err = sqlboiler.QueryDisplayItems(displayItemQuery).UpdateAll(context.Background(), db.DB, displayUpdateCols)

		searchUpdateCols := map[string]interface{}{
			sqlboiler.QuerySearchConditionItemColumns.Del:           true,
			sqlboiler.QuerySearchConditionItemColumns.UpdateStaffID: null.StringFrom(operatorID),
		}
		searchItemQuery := qm.Where(sqlboiler.QueryDisplayItemColumns.QueryConditionsID+" = ?", queryCondition.ID)
		_, err = sqlboiler.QuerySearchConditionItems(searchItemQuery).UpdateAll(context.Background(), db.DB, searchUpdateCols)

		orderUpdateCols := map[string]interface{}{
			sqlboiler.QueryOrderConditionItemColumns.Del:           true,
			sqlboiler.QueryOrderConditionItemColumns.UpdateStaffID: null.StringFrom(operatorID),
		}
		orderItemQuery := qm.Where(sqlboiler.QueryOrderConditionItemColumns.QueryConditionsID+" = ?", queryCondition.ID)
		_, err = sqlboiler.QueryOrderConditionItems(orderItemQuery).UpdateAll(context.Background(), db.DB, orderUpdateCols)

		// update
		_, err = sqlQueryCondition.Update(context.Background(), db.DB, boil.Infer())
		for _, d := range sqlQueryDisplayItems {
			err = d.Insert(context.Background(), db.DB, boil.Infer())
		}
		for _, s := range sqlQuerySearchConditionItems {
			err = s.Insert(context.Background(), db.DB, boil.Infer())
		}
		for _, o := range sqlQueryOrderConditionItems {
			err = o.Insert(context.Background(), db.DB, boil.Infer())
		}
		err = sqlQueryCondition.SetStaffGroups(context.Background(), db.DB, false, sqlDiscloseGroups...)
		return err

	})

	return err
}

// Insert insert data to database
func (q *QueryConditionRepo) Insert(queryCondition *query.Condition, operatorID string) (id string, err error) {
	id = ""
	sqlQueryCondition := &sqlboiler.QueryCondition{
		ID:            utils.CreateID(),
		CreStaffID:    null.StringFrom(operatorID),
		UpdateStaffID: null.StringFrom(operatorID),
		CategoryName:  queryCondition.Category.Name,
		IsDisclose:    queryCondition.IsDisclose,
		OwnerID:       operatorID,
		PatternName:   queryCondition.PatternName,
	}

	sqlQueryDisplayItems := make([]*sqlboiler.QueryDisplayItem,
		len(queryCondition.ConditionData.DisplayItemList))
	sqlQuerySearchConditionItems := make([]*sqlboiler.QuerySearchConditionItem,
		len(queryCondition.ConditionData.SearchConditionList))
	sqlQueryOrderConditionItems := make([]*sqlboiler.QueryOrderConditionItem,
		len(queryCondition.ConditionData.OrderConditionList))
	sqlDiscloseGroups := make([]*sqlboiler.StaffGroup, len(queryCondition.DiscloseGroups))

	for i, d := range queryCondition.ConditionData.DisplayItemList {
		sqlQueryDisplayItems[i] = &sqlboiler.QueryDisplayItem{
			ID:                utils.CreateID(),
			CreStaffID:        null.StringFrom(operatorID),
			UpdateStaffID:     null.StringFrom(operatorID),
			QueryConditionsID: sqlQueryCondition.ID,
			DisplayFieldID:    d.ID,
			RowOrder:          i,
		}
	}

	for i, s := range queryCondition.ConditionData.SearchConditionList {
		sqlQuerySearchConditionItems[i] = &sqlboiler.QuerySearchConditionItem{
			ID:                utils.CreateID(),
			CreStaffID:        null.StringFrom(operatorID),
			UpdateStaffID:     null.StringFrom(operatorID),
			QueryConditionsID: sqlQueryCondition.ID,
			SearchFieldID:     s.SearchField.ID,
			ConditionValue:    s.ConditionValue,
			MatchType:         string(s.MatchType),
			Operator:          string(s.Operator),
			RowOrder:          i,
		}
	}

	for i, o := range queryCondition.ConditionData.OrderConditionList {
		sqlQueryOrderConditionItems[i] = &sqlboiler.QueryOrderConditionItem{
			ID:                utils.CreateID(),
			CreStaffID:        null.StringFrom(operatorID),
			UpdateStaffID:     null.StringFrom(operatorID),
			QueryConditionsID: sqlQueryCondition.ID,
			OrderFieldID:      o.OrderField.ID,
			OrderFieldKeyWord: string(o.OrderFieldKeyWord),
			RowOrder:          i,
		}
	}

	for i, g := range queryCondition.DiscloseGroups {
		sqlDiscloseGroups[i] = &sqlboiler.StaffGroup{
			ID: g.ID,
		}
	}

	err = q.database.WithDbContext(func(db *sqlx.DB) error {
		var err error

		err = sqlQueryCondition.Insert(context.Background(), db.DB, boil.Infer())
		if err != nil {
			log.Error(err.Error())
			return err
		}
		for _, d := range sqlQueryDisplayItems {
			err = d.Insert(context.Background(), db.DB, boil.Infer())
			if err != nil {
				log.Error(err.Error())
				return err
			}
		}
		for _, s := range sqlQuerySearchConditionItems {
			err = s.Insert(context.Background(), db.DB, boil.Infer())
			if err != nil {
				log.Error(err.Error())

				return err
			}
		}
		for _, o := range sqlQueryOrderConditionItems {
			err = o.Insert(context.Background(), db.DB, boil.Infer())
			if err != nil {
				log.Error(err.Error())
				return err
			}
		}
		err = sqlQueryCondition.SetStaffGroups(context.Background(), db.DB, false, sqlDiscloseGroups...)

		if err != nil {
			log.Error(err.Error())
		}

		return err
	})

	if err != nil {
		return "", err
	}

	id = sqlQueryCondition.ID
	return
}

// SelectByIDs select staff data by id list from database
func (q *QueryConditionRepo) SelectByIDs(ids []string) (queryConditions []*query.Condition, err error) {
	queryConditions = []*query.Condition{}
	if len(ids) == 0 {
		return nil, errors.New("id list must be required")
	}

	var convertedIDs []interface{} = make([]interface{}, len(ids))
	for i, d := range ids {
		convertedIDs[i] = d
	}

	err = q.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.QueryConditionColumns.Del+" !=?", true),
			qm.AndIn(sqlboiler.QueryConditionColumns.ID+" in ?", convertedIDs...),
			qm.Load(qm.Rels(sqlboiler.QueryConditionRels.Owner, sqlboiler.StaffRels.StaffGroups), qm.Where("del != true")),
			qm.Load(sqlboiler.QueryConditionRels.QueryDisplayItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryDisplayItemColumns.RowOrder)),
			qm.Load(sqlboiler.QueryConditionRels.QueryOrderConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryOrderConditionItemColumns.RowOrder)),
			qm.Load(sqlboiler.QueryConditionRels.QuerySearchConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QuerySearchConditionItemColumns.RowOrder)),
			qm.Load(sqlboiler.QueryConditionRels.StaffGroups, qm.Where("del != true")),
		}

		fetchedQueryConditions, err := sqlboiler.QueryConditions(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, fq := range fetchedQueryConditions {
				queryConditions = append(queryConditions, QueryConditionObjectMap(fq))
			}
		}

		return err
	})

	return
}

// SelectByID select staaff data by id from database
func (q *QueryConditionRepo) SelectByID(id string) (queryCondition *query.Condition, err error) {
	if id == "" {
		return nil, errors.New("id must be required")
	}

	err = q.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.QueryConditionColumns.Del+" !=?", true),
			qm.And(sqlboiler.QueryConditionColumns.ID+" =?", id),
			qm.Load(qm.Rels(sqlboiler.QueryConditionRels.Owner, sqlboiler.StaffRels.StaffGroups), qm.Where("del != true")),
			qm.Load(sqlboiler.QueryConditionRels.QueryDisplayItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryDisplayItemColumns.RowOrder)),
			qm.Load(sqlboiler.QueryConditionRels.QueryOrderConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryOrderConditionItemColumns.RowOrder)),
			qm.Load(sqlboiler.QueryConditionRels.QuerySearchConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QuerySearchConditionItemColumns.RowOrder)),
			qm.Load(sqlboiler.QueryConditionRels.StaffGroups, qm.Where("del != true")),
		}
		fetchedQueryCondition, err := sqlboiler.QueryConditions(queries...).One(context.Background(), db.DB)
		if err == nil {
			queryCondition = QueryConditionObjectMap(fetchedQueryCondition)
		}

		return err
	})

	return
}

// SelectAll select all query condition data without not del from database
func (q *QueryConditionRepo) SelectAll() (queryConditions []*query.Condition, err error) {
	queryConditions = []*query.Condition{}

	err = q.database.WithDbContext(func(db *sqlx.DB) error {
		queries := q.createQueryModSlice()
		fetchedQueryConditions, err := sqlboiler.QueryConditions(queries...).All(context.Background(), db.DB)
		if err == nil {
			for _, fqc := range fetchedQueryConditions {
				queryConditions = append(queryConditions, QueryConditionObjectMap(fqc))
			}
		}

		return err
	})

	return
}

// Select select query condition data by condition from database
func (q *QueryConditionRepo) Select(queryItems ...query.SearchConditionItem) (resultQueryConditions []*query.Condition, err error) {

	err = q.database.WithDbContext(func(db *sqlx.DB) error {

		queries := q.createQueryModSlice()
		qmod := qm.And("query_conditions."+sqlboiler.QueryConditionColumns.Del+" != ?", true)
		queries = append(queries, qmod)

		// 条件構築
		for _, queryItem := range queryItems {

			qmod = q.createQueryModWhere(queryItem)
			queries = append(queries, qmod)

		}

		// データ取得処理
		fetchedQueryConditions, err := sqlboiler.QueryConditions(queries...).All(context.Background(), db.DB)

		if err == nil {
			for _, fc := range fetchedQueryConditions {
				resultQueryConditions = append(resultQueryConditions, QueryConditionObjectMap(fc))
			}
		}

		return err
	})

	return
}

func (q *QueryConditionRepo) createQueryModWhere(queryItem query.SearchConditionItem) qm.QueryMod {

	mt, val := comparisonOperator(queryItem.MatchType, queryItem.ConditionValue)

	switch queryItem.SearchField.ID {
	case "pattern-name":
		if queryItem.Operator == query.Or {
			return qm.Or("query_conditions."+sqlboiler.QueryConditionColumns.PatternName+" "+mt+" ?", val)
		}
		return qm.And("query_conditions."+sqlboiler.QueryConditionColumns.PatternName+" "+mt+" ?", val)
	case "is-disclose":
		if queryItem.Operator == query.Or {
			return qm.Or("query_conditions."+sqlboiler.QueryConditionColumns.IsDisclose+" "+mt+" ?", val)
		}
		return qm.And("query_conditions."+sqlboiler.QueryConditionColumns.IsDisclose+" "+mt+" ?", val)
	case "disclose-groups":
		var ids []interface{}
		json.Unmarshal([]byte(val), &ids)
		if queryItem.Operator == query.Or {
			return qm.OrIn("sg.id"+" "+mt+" ?", ids...)
		}
		return qm.AndIn("sg.id"+" "+mt+" ?", ids...)
	case "owner":
		if queryItem.Operator == query.Or {
			return qm.Or("owner."+sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
		}
		return qm.And("owner."+sqlboiler.StaffColumns.Name+" "+mt+" ?", val)
	case "category-name":
		var ids []interface{}
		json.Unmarshal([]byte(val), &ids)
		if queryItem.Operator == query.Or {
			return qm.OrIn("query_conditions."+sqlboiler.QueryConditionColumns.CategoryName+" "+mt+" ?", ids...)
		}
		return qm.AndIn("query_conditions."+sqlboiler.QueryConditionColumns.CategoryName+" "+mt+" ?", ids...)

	default:
		if queryItem.Operator == query.Or {
			return qm.Or("query_conditions."+sqlboiler.QueryConditionColumns.PatternName+" "+mt+" ?", val)
		}
		return qm.And("query_conditions."+sqlboiler.QueryConditionColumns.PatternName+" "+mt+" ?", val)
		// queryItem.Operator == and
	}
}

func (q *QueryConditionRepo) createQueryModSlice() (qslice []qm.QueryMod) {
	qslice = []qm.QueryMod{}
	qslice = append(
		qslice,
		qm.Select("distinct query_conditions.*"),
		qm.Where("query_conditions."+sqlboiler.QueryConditionColumns.ID+" IS NOT NULL"),
		qm.Load(qm.Rels(sqlboiler.QueryConditionRels.Owner, sqlboiler.StaffRels.StaffGroups), qm.Where("del != true")),
		qm.Load(sqlboiler.QueryConditionRels.QueryDisplayItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryDisplayItemColumns.RowOrder)),
		qm.Load(sqlboiler.QueryConditionRels.QueryOrderConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryOrderConditionItemColumns.RowOrder)),
		qm.Load(sqlboiler.QueryConditionRels.QuerySearchConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QuerySearchConditionItemColumns.RowOrder)),
		qm.Load(sqlboiler.QueryConditionRels.StaffGroups, qm.Where("del != true")),
	)
	return
}

// QueryConditionObjectMap data mapper sqlboiler object to entities object
func QueryConditionObjectMap(sqc *sqlboiler.QueryCondition) (eqc *query.Condition) {
	var category *query.Category
	var staffGroups []*authentication.Group
	var displayItemList []query.FieldAttr
	var searchConditionList []query.SearchConditionItem
	var orderConditionList []query.OrderConditionItem

	if sqc == nil {
		return nil
	}

	r := NewStaffGroupRepo()
	groups, _ := r.SelectAll()

	for _, category = range query.CreateCategories(groups) {
		if category.Name == sqc.CategoryName {
			break
		}
	}

	for _, group := range sqc.R.StaffGroups {
		staffGroups = append(staffGroups, StaffGroupObjectMap(group))
	}

	displayItemList = []query.FieldAttr{}
	for _, item := range sqc.R.QueryDisplayItems {
		var displayItem query.FieldAttr
		for _, displayItem = range category.SearchItems.DisplayItemList {
			if item.DisplayFieldID == displayItem.ID {
				break
			}
		}
		displayItemList = append(displayItemList, displayItem)
	}

	searchConditionList = []query.SearchConditionItem{}
	for _, item := range sqc.R.QuerySearchConditionItems {
		var searchField query.FieldAttr
		for _, searchField = range category.SearchItems.SearchConditionList {
			if item.SearchFieldID == searchField.ID {
				break
			}
		}

		var matchTypeEnum query.MatchTypeEnum
		var operatorEnum query.OperatorEnum

		searchConditionItem := query.SearchConditionItem{
			SearchField:    searchField,
			ConditionValue: item.ConditionValue,
			MatchType:      matchTypeEnum.StrToEnum(item.MatchType),
			Operator:       operatorEnum.StrToEnum(item.Operator),
		}

		searchConditionList = append(searchConditionList, searchConditionItem)
	}

	orderConditionList = []query.OrderConditionItem{}
	for _, item := range sqc.R.QueryOrderConditionItems {
		var orderField query.FieldAttr
		for _, orderField = range category.SearchItems.OrderConditionList {
			if item.OrderFieldID == orderField.ID {
				break
			}
		}
		var orderTypeEnum query.OrderTypeEnum

		orderConditionItem := query.OrderConditionItem{
			OrderField:        orderField,
			OrderFieldKeyWord: orderTypeEnum.StrToEnum(item.OrderFieldKeyWord),
		}
		orderConditionList = append(orderConditionList, orderConditionItem)
	}
	eqc = &query.Condition{
		ID:             sqc.ID,
		PatternName:    sqc.PatternName,
		Category:       category,
		IsDisclose:     sqc.IsDisclose,
		DiscloseGroups: staffGroups,
		ConditionData: query.ConditionData{
			DisplayItemList:     displayItemList,
			SearchConditionList: searchConditionList,
			OrderConditionList:  orderConditionList,
		},
		Owner: StaffObjectMap(sqc.R.Owner),
	}
	return
}

// NewQueryConditionRepo constructor
func NewQueryConditionRepo() *QueryConditionRepo {
	return &QueryConditionRepo{database: createDB()}
}

// QueryConditionRepo repository struct
type QueryConditionRepo struct {
	database db
}
