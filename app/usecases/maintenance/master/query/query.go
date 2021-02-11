package query

import (
	"encoding/json"
	"regulus/app/domain/query"
	"regulus/app/usecases/maintenance/master/group"
	"regulus/app/utils/log"
)

/*

UpdateCondition は検索条件更新時のユースケースです。

*/
func UpdateCondition(queryRepo persistance, queryCondition *query.Condition, operatorID string) error {

	err := queryRepo.Update(queryCondition, operatorID)

	if err != nil {
		log.Error("usecases:query:UpdateCondition:message:" + err.Error())
	}

	return err
}

/*

Find は検索時のユースケースです。条件指定がない場合は全件検索の結果を返します。

*/
func Find(conditionData *query.ConditionData, queryRepo persistance) ([]*query.Condition, error) {

	// 検索条件がDB管理されていない場合category-view-valueの処理

	conditionList := []query.SearchConditionItem{}

	for _, queryItem := range conditionData.SearchConditionList {

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

AddCondition は検索条件追加時のユースケースです。成功した場合は id を返却します。

*/
func AddCondition(queryRepo persistance, queryCondition *query.Condition, operatorID string) (string, error) {

	id, err := queryRepo.Insert(queryCondition, operatorID)

	if err != nil {
		log.Error("usecases:query:AddCondition:message:" + err.Error())
	}

	return id, err
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
