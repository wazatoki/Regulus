package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"regulus/app/domain"
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
func (q *QueryConditionRepo) Update(queryCondition *domain.Condition, operatorID string) (err error) {
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
			MatchType:         s.MatchType.String(),
			Operator:          s.Operator.String(),
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
			OrderFieldKeyWord: o.OrderFieldKeyWord.String(),
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
func (q *QueryConditionRepo) Insert(queryCondition *domain.Condition, operatorID string) (id string, err error) {
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
			MatchType:         s.MatchType.String(),
			Operator:          s.Operator.String(),
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
			OrderFieldKeyWord: o.OrderFieldKeyWord.String(),
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
func (q *QueryConditionRepo) SelectByIDs(ids []string) (queryConditions domain.Conditions, err error) {
	queryConditions = domain.Conditions{}
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
func (q *QueryConditionRepo) SelectByID(id string) (queryCondition *domain.Condition, err error) {
	if id == "" {
		return nil, errors.New("id must be required")
	}

	err = q.database.WithDbContext(func(db *sqlx.DB) error {
		queries := q.createQueryModSlice()
		queries = append(
			queries,
			qm.And(sqlboiler.QueryConditionColumns.ID+" =?", id),
		)
		fetchedQueryCondition, err := sqlboiler.QueryConditions(queries...).One(context.Background(), db.DB)
		if err == nil {
			queryCondition = QueryConditionObjectMap(fetchedQueryCondition)
		}

		return err
	})

	return
}

// SelectAll select all query condition data without not del from database
func (q *QueryConditionRepo) SelectAll() (queryConditions domain.Conditions, err error) {
	queryConditions = domain.Conditions{}

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
func (q *QueryConditionRepo) Select(queryItems ...domain.SearchConditionItem) (resultQueryConditions domain.Conditions, err error) {

	err = q.database.WithDbContext(func(db *sqlx.DB) error {

		var args []interface{} = make([]interface{}, 0)
		ids := []string{}

		queryStr := "select distinct qc.id " +
			"from query_conditions qc " +
			"inner join staffs owner on qc.owner_id = owner.id " +
			"inner join join_query_conditions_staff_groups jqcsg on qc.id = jqcsg.query_conditions_id " +
			"inner join staff_groups sg on jqcsg.staff_groups_id = sg.id " +
			"where qc.del != true"

		// 条件構築
		searchConditionItems := []domain.SearchConditionItem{}

		for _, queryItem := range queryItems {

			// 検索条件がDB管理されていない場合の処理
			if queryItem.SearchField.ID == "category-view-value" {

				item := domain.SearchConditionItem{
					SearchField: domain.FieldAttr{},
				}
				item.Operator = queryItem.Operator
				item.MatchType = domain.QueryMatchTypeEnum.IN
				item.SearchField.ID = "category-name"
				categoryNames := domain.CategoryNameListByMatchType(queryItem.ConditionValue, queryItem.MatchType)
				tmpByte, _ := json.Marshal(categoryNames)
				item.ConditionValue = string(tmpByte)
				searchConditionItems = append(searchConditionItems, item)

			} else {
				searchConditionItems = append(searchConditionItems, queryItem)
			}
		}

		for _, searchConditionItem := range searchConditionItems {
			qu, pslice := q.createQueryModWhere(searchConditionItem)
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
		queries := q.createQueryModSlice()
		queries = append(
			queries,
			qm.And("query_conditions."+sqlboiler.QueryConditionColumns.Del+" != ?", true),
			qm.AndIn("query_conditions."+sqlboiler.QueryConditionColumns.ID+" in ?", convertedIDs...),
		)
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

func (q *QueryConditionRepo) createQueryModWhere(queryItem domain.SearchConditionItem) (string, []string) {

	mt, val := comparisonOperator(queryItem.MatchType, queryItem.ConditionValue)

	switch queryItem.SearchField.ID {
	case "pattern-name":
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return " or qc." + sqlboiler.QueryConditionColumns.PatternName + " " + mt + " ?", []string{val}
		}
		return " and qc." + sqlboiler.QueryConditionColumns.PatternName + " " + mt + " ?", []string{val}
	case "is-disclose":
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return " or qc." + sqlboiler.QueryConditionColumns.IsDisclose + " " + mt + " ?", []string{val}
		}
		return " and qc." + sqlboiler.QueryConditionColumns.IsDisclose + " " + mt + " ?", []string{val}
	case "disclose-groups":
		var ids []string
		json.Unmarshal([]byte(val), &ids)
		var convertedIDs []interface{} = make([]interface{}, len(ids))
		for i, d := range ids {
			convertedIDs[i] = d
		}
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			q, _, _ := sqlx.In(" or sg.id in (?)", convertedIDs)
			return q, ids
		}
		q, _, _ := sqlx.In(" and sg.id in (?)", convertedIDs)
		return q, ids
	case "owner":
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return " or owner." + sqlboiler.StaffColumns.Name + " " + mt + " ?", []string{val}
		}
		return " and owner." + sqlboiler.StaffColumns.Name + " " + mt + " ?", []string{val}
	case "category-name":
		var ids []string
		json.Unmarshal([]byte(val), &ids)
		var convertedIDs []interface{} = make([]interface{}, len(ids))
		for i, d := range ids {
			convertedIDs[i] = d
		}
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			q, _, _ := sqlx.In(" or qc."+sqlboiler.QueryConditionColumns.CategoryName+" in (?)", convertedIDs)
			return q, ids
		}
		q, _, _ := sqlx.In(" and qc."+sqlboiler.QueryConditionColumns.CategoryName+" in (?)", convertedIDs)
		return q, ids

	default:
		if queryItem.Operator.String() == domain.QueryOperatorEnum.OR.String() {
			return " or query_conditions." + sqlboiler.QueryConditionColumns.PatternName + " " + mt + " ?", []string{val}
		}
		return " and query_conditions." + sqlboiler.QueryConditionColumns.PatternName + " " + mt + " ?", []string{val}
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
func QueryConditionObjectMap(sqc *sqlboiler.QueryCondition) (eqc *domain.Condition) {
	var category *domain.Category
	var staffGroups domain.StaffGroups
	var displayItemList []domain.FieldAttr
	var searchConditionList []domain.SearchConditionItem
	var orderConditionList []domain.OrderConditionItem

	if sqc == nil {
		return nil
	}

	r := NewStaffGroupRepo()
	groups, _ := r.SelectAll()

	for _, category = range domain.CreateCategories(groups) {
		if category.Name == sqc.CategoryName {
			break
		}
	}

	for _, group := range sqc.R.StaffGroups {
		staffGroups = append(staffGroups, StaffGroupObjectMap(group))
	}

	displayItemList = []domain.FieldAttr{}
	for _, item := range sqc.R.QueryDisplayItems {
		var displayItem domain.FieldAttr
		for _, displayItem = range category.SearchItems.DisplayItemList {
			if item.DisplayFieldID == displayItem.ID {
				break
			}
		}
		displayItemList = append(displayItemList, displayItem)
	}

	searchConditionList = []domain.SearchConditionItem{}
	for _, item := range sqc.R.QuerySearchConditionItems {
		var searchField domain.FieldAttr
		for _, searchField = range category.SearchItems.SearchConditionList {
			if item.SearchFieldID == searchField.ID {
				break
			}
		}

		matchTypeEnum := domain.QueryMatchType{}
		operatorEnum := domain.QueryOperator{}

		searchConditionItem := domain.SearchConditionItem{
			SearchField:    searchField,
			ConditionValue: item.ConditionValue,
			MatchType:      matchTypeEnum.StrToEnum(item.MatchType),
			Operator:       operatorEnum.StrToEnum(item.Operator),
		}

		searchConditionList = append(searchConditionList, searchConditionItem)
	}

	orderConditionList = []domain.OrderConditionItem{}
	for _, item := range sqc.R.QueryOrderConditionItems {
		var orderField domain.FieldAttr
		for _, orderField = range category.SearchItems.OrderConditionList {
			if item.OrderFieldID == orderField.ID {
				break
			}
		}
		orderTypeEnum := domain.QueryOrderType{}

		orderConditionItem := domain.OrderConditionItem{
			OrderField:        orderField,
			OrderFieldKeyWord: orderTypeEnum.StrToEnum(item.OrderFieldKeyWord),
		}
		orderConditionList = append(orderConditionList, orderConditionItem)
	}
	eqc = &domain.Condition{
		ID:             sqc.ID,
		PatternName:    sqc.PatternName,
		Category:       category,
		IsDisclose:     sqc.IsDisclose,
		DiscloseGroups: staffGroups,
		ConditionData: domain.ConditionData{
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
