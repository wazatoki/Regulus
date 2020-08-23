package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"regulus/app/domain/entities"
	"regulus/app/domain/services"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/sqlboiler"
	"regulus/app/utils"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
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
func (q *QueryConditionRepo) Update(queryCondition *entities.QueryCondition, operatorID string) (err error) {
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
func (q *QueryConditionRepo) Insert(queryCondition *entities.QueryCondition, operatorID string) (id string, err error) {
	id = ""
	sqlQueryCondition := &sqlboiler.QueryCondition{
		ID:            utils.CreateID(),
		CreStaffID:    null.StringFrom(operatorID),
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

		err = sqlQueryCondition.Insert(context.Background(), db.DB, boil.Infer())
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

	if err == nil {
		id = sqlQueryCondition.ID
	}

	return
}

// SelectByIDs select staff data by id list from database
func (q *QueryConditionRepo) SelectByIDs(ids []string) (queryConditions []*entities.QueryCondition, err error) {
	queryConditions = []*entities.QueryCondition{}
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
func (q *QueryConditionRepo) SelectByID(id string) (queryCondition *entities.QueryCondition, err error) {
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
func (q *QueryConditionRepo) SelectAll() (queryConditions []*entities.QueryCondition, err error) {
	queryConditions = []*entities.QueryCondition{}

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
func (q *QueryConditionRepo) Select(queryItems ...query.SearchConditionItem) (resultQueryConditions []*entities.QueryCondition, err error) {
	tempQueryConditions := []*entities.QueryCondition{}
	addQueryConditions := []*entities.QueryCondition{}
	var allQueryConditions []*entities.QueryCondition
	var queries []qm.QueryMod
	var qmod qm.QueryMod

	err = q.database.WithDbContext(func(db *sqlx.DB) error {

		var selectAllerr error
		allQueryConditions, selectAllerr = q.SelectAll()
		if selectAllerr != nil {
			return selectAllerr
		}

		resultQueryConditions = allQueryConditions

		for i, queryItem := range queryItems {

			operator := queryItem.Operator

			if q.isDependentDB(queryItem) {
				queries = q.createQueryModSlice()
				qmod = qm.And("query_conditions."+sqlboiler.QueryConditionColumns.Del+" != ?", true)

				// DBに依存する条件が続く場合はまとめてクエリを作成する
				j := i
				for ; j < len(queryItems); j++ {
					qmod = qm.Expr(qmod, q.createQueryModWhere(queryItems[j]))
					if j < len(queryItems)-1 && !q.isDependentDB(queryItems[j+1]) {
						break
					}
				}
				i = j

				queries = append(queries, qmod)
				fetchedStaffs, err := sqlboiler.QueryConditions(queries...).All(context.Background(), db.DB)
				if err == nil {
					for _, fs := range fetchedStaffs {
						tempQueryConditions = append(tempQueryConditions, QueryConditionObjectMap(fs))
					}
				}
			} else {
				for _, item := range allQueryConditions {
					if queryItem.SearchField.ID == "category-view-value" {
						switch queryItem.MatchType {
						case query.Match:
							if item.Category.Name == queryItem.ConditionValue {
								tempQueryConditions = append(tempQueryConditions, item)
							}
						case query.Unmatch:
							if item.Category.Name != queryItem.ConditionValue {
								tempQueryConditions = append(tempQueryConditions, item)
							}
						default: // query.Pertialmatch
							if strings.Contains(item.Category.Name, queryItem.ConditionValue) {
								tempQueryConditions = append(tempQueryConditions, item)
							}
						}
					}
				}
			}

			if operator == query.Or {
				addQueryConditions = []*entities.QueryCondition{}
				for _, tempItem := range tempQueryConditions {
					isMatchResult := false
					for _, resultItem := range resultQueryConditions {
						if resultItem.ID == tempItem.ID {
							isMatchResult = true
							break
						}
					}
					if !isMatchResult {
						// 最新の結果と一致しなかったものの集合を作る
						addQueryConditions = append(addQueryConditions, tempItem)
					}
				}
				// 最終結果に反映
				resultQueryConditions = append(resultQueryConditions, addQueryConditions...)
			} else { // operator == And
				addQueryConditions = []*entities.QueryCondition{}
				for _, resultItem := range resultQueryConditions {
					isMatchResult := false
					for _, tempItem := range tempQueryConditions {
						if resultItem.ID == tempItem.ID {
							isMatchResult = true
							break
						}
					}
					if isMatchResult {
						// 最新の結果と一致したものだけの集合を作る
						addQueryConditions = append(addQueryConditions, resultItem)
					}
				}
				// 最終結果に反映
				resultQueryConditions = addQueryConditions
			}

		}
		return err
	})

	return
}

func (q *QueryConditionRepo) isDependentDB(queryItem query.SearchConditionItem) bool {
	switch queryItem.SearchField.ID {
	case "category-view-value":
		return false
	default:
		return true
	}
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
		qm.InnerJoin("join_query_conditions_staff_groups jqcsg on query_conditions.id = jqcsg.query_conditions_id"),
		qm.InnerJoin("staff_groups sg on jqcsg.staff_groups_id = sg.id"),
		qm.InnerJoin("staffs owner on query_conditions.owner_id = owner.id"),
		qm.Where("query_conditions."+sqlboiler.QueryConditionColumns.ID+" IS NOT NULL"),
		qm.Load(qm.Rels(sqlboiler.QueryConditionRels.Owner, sqlboiler.StaffRels.StaffGroups), qm.Where("del != true")),
		qm.Load(sqlboiler.QueryConditionRels.QueryDisplayItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryDisplayItemColumns.RowOrder)),
		qm.Load(sqlboiler.QueryConditionRels.QueryOrderConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QueryOrderConditionItemColumns.RowOrder)),
		qm.Load(sqlboiler.QueryConditionRels.QuerySearchConditionItems, qm.Where("del != true"), qm.OrderBy(sqlboiler.QuerySearchConditionItemColumns.RowOrder)),
		qm.Load(sqlboiler.QueryConditionRels.StaffGroups, qm.Where("del != true")),
		//qm.Load(sqlboiler.StaffRels.StaffGroups, qm.Where("del != true")),
	)
	return
}

// QueryConditionObjectMap data mapper sqlboiler object to entities object
func QueryConditionObjectMap(sqc *sqlboiler.QueryCondition) (eqc *entities.QueryCondition) {
	var category entities.Category
	var staffGroups []*entities.StaffGroup
	var displayItemList []query.FieldAttr
	var searchConditionList []query.SearchConditionItem
	var orderConditionList []query.OrderConditionItem

	if sqc == nil {
		return nil
	}

	r := NewStaffGroupRepo()
	groups, _ := r.SelectAll()

	for _, category = range services.CreateCategories(groups) {
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
	eqc = &entities.QueryCondition{
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
