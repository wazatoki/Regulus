package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
	"regulus/app/infrastructures/sqlboiler"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// SelectByID select staaff data by id from database
func (q *QueryConditionRepo) SelectByID(id string) (queryCondition entities.QueryCondition, err error) {
	if id == "" {
		return entities.QueryCondition{}, errors.New("id must be required")
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
func (q *QueryConditionRepo) SelectAll() (queryConditions []entities.QueryCondition, err error) {
	queryConditions = []entities.QueryCondition{}

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
func (q *QueryConditionRepo) Select(queryItems ...*query.SearchConditionItem) (resultQueryConditions []entities.QueryCondition, err error) {
	tempQueryConditions := []entities.QueryCondition{}
	addQueryConditions := []entities.QueryCondition{}
	var allQueryConditions []entities.QueryCondition
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
				addQueryConditions = []entities.QueryCondition{}
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
				addQueryConditions = []entities.QueryCondition{}
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

func (q *QueryConditionRepo) isDependentDB(queryItem *query.SearchConditionItem) bool {
	switch queryItem.SearchField.ID {
	case "category-view-value":
		return false
	default:
		return true
	}
}

func (q *QueryConditionRepo) createQueryModWhere(queryItem *query.SearchConditionItem) qm.QueryMod {

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
func QueryConditionObjectMap(sqc *sqlboiler.QueryCondition) (eqc entities.QueryCondition) {
	var category entities.Category
	var staffGroups []entities.StaffGroup
	var displayItemList []query.FieldAttr
	var searchConditionList []query.SearchConditionItem
	var orderConditionList []query.OrderConditionItem

	for _, category = range entities.Categories {
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
	eqc = entities.QueryCondition{
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
