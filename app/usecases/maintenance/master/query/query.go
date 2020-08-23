package query

import (
	"regulus/app/domain/entities"
	"regulus/app/domain/services"
	"regulus/app/domain/vo/query"
)

/*

Find は検索時のユースケースです。条件指定がない場合は全件検索の結果を返します。

*/
func Find(queryRepo persistance, conditionData *query.ConditionData) ([]*entities.QueryCondition, error) {

	items, err := queryRepo.Select(conditionData.SearchConditionList...)

	if err != nil {
		items = services.Sort(items, conditionData.OrderConditionList...)

		return items, nil
	}

	return nil, err
}

/*

Find は検索条件追加時のユースケースです。成功した場合は id を返却します。

*/
func AddCondition(queryRepo persistance, queryCondition *entities.QueryCondition, operatorID string) (string, error) {

	id, err := queryRepo.Insert(queryCondition, operatorID)

	return id, err
}
