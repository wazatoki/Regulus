package query

import (
	"regulus/app/domain/entities"
	"regulus/app/domain/services"
	"regulus/app/domain/vo/query"
)

/*

Find は検索時のユースケースです。条件指定がない場合は全件検索の結果を返します。

*/
func Find(queryRepo persistance, conditionData *query.ConditionData) ([]entities.QueryCondition, error) {

	items, err := queryRepo.Select(conditionData.SearchConditionList...)

	if err != nil {
		items = services.Sort(items, conditionData.OrderConditionList...)

		return items, nil
	}

	return nil, err
}
