package query

import (
	"encoding/json"
	"regulus/app/domain/query"
	"regulus/app/usecases/maintenance/master/group"
	"regulus/app/utils/log"
)

/*

Delete は検索条件削除時のユースケースです。

*/
func Delete(queryConditionIDs *[]string, queryRepo persistance, operatorID string) (result []query.Condition) {

	result = []query.Condition{}

	for _, id := range *queryConditionIDs {

		err := queryRepo.Delete(id, operatorID)

		if err != nil {
			log.Error("usecases:query:Update:message:" + err.Error())
			c, _ := queryRepo.SelectByID(id)
			result = append(result, *c)
		}
	}

	return
}

/*

Update は検索条件更新時のユースケースです。

*/
func Update(queryCondition *query.Condition, queryRepo persistance, operatorID string) error {

	err := queryRepo.Update(queryCondition, operatorID)

	if err != nil {
		log.Error("usecases:query:Update:message:" + err.Error())
	}

	return err
}

/*

Add は検索条件追加時のユースケースです。成功した場合は id を返却します。

*/
func Add(queryCondition *query.Condition, queryRepo persistance, operatorID string) (*query.Condition, error) {

	id, err := queryRepo.Insert(queryCondition, operatorID)

	if err != nil {
		log.Error("usecases:query:Add:message:" + err.Error())
		return nil, err
	}

	result, e := queryRepo.SelectByID(id)
	return result, e
}

/*

Find は検索時のユースケースです。条件指定がない場合は全件検索の結果を返します。

*/
func Find(conditionData *query.ConditionData, queryRepo persistance) ([]*query.Condition, error) {

	conditionList := []query.SearchConditionItem{}

	for _, queryItem := range conditionData.SearchConditionList {

		// 検索条件がDB管理されていない場合の処理
		if queryItem.SearchField.ID == "category-view-value" {

			item := query.SearchConditionItem{
				SearchField: query.FieldAttr{},
			}
			item.Operator = queryItem.Operator
			item.MatchType = query.In
			item.SearchField.ID = "category-name"
			categoryNames := query.CategoryNameListByMatchType(queryItem.ConditionValue, queryItem.MatchType)
			tmpByte, _ := json.Marshal(categoryNames)
			item.ConditionValue = string(tmpByte)
			conditionList = append(conditionList, item)

		} else {
			conditionList = append(conditionList, queryItem)
		}
	}

	items, err := queryRepo.Select(conditionData.SearchConditionList...)

	if err != nil {

		log.Error("usecases:query:Find:message:" + err.Error())
		return nil, err
	}

	items = query.Sort(items, conditionData.OrderConditionList...)
	return items, nil
}

/*

FetchDataInputFormItems は検索条件登録フォームを開く際に必要なデータを取得するユースケースです。

*/
func FetchDataInputFormItems(groupRepo group.Persistance) ([]*query.Category, error) {
	groups, err := groupRepo.SelectAll()

	if err != nil {
		log.Error("usecases:query:FetchDataInputFormItems:message:" + err.Error())
		return nil, err
	}
	return query.CreateCategories(groups), nil

}
