package query

import (
	"regulus/app/domain"
	"regulus/app/usecases/maintenance/master/group"
	"regulus/app/utils/log"
)

/*

Delete は検索条件削除時のユースケースです。

*/
func Delete(queryConditionIDs *[]string, queryRepo persistance, operatorID string) (result []domain.Condition) {

	result = []domain.Condition{}

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
func Update(queryCondition *domain.Condition, queryRepo persistance, operatorID string) error {

	err := queryRepo.Update(queryCondition, operatorID)

	if err != nil {
		log.Error("usecases:query:Update:message:" + err.Error())
	}

	return err
}

/*

Add は検索条件追加時のユースケースです。成功した場合は id を返却します。

*/
func Add(queryCondition *domain.Condition, queryRepo persistance, operatorID string) (*domain.Condition, error) {

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
func Find(conditionData *domain.ConditionData, queryRepo persistance) ([]*domain.Condition, error) {

	items, err := queryRepo.Select(conditionData.SearchConditionList...)

	if err != nil {

		log.Error("usecases:query:Find:message:" + err.Error())
		return nil, err
	}

	items = domain.Sort(items, conditionData.OrderConditionList...)
	return items, nil
}

/*

FetchDataInputFormItems は検索条件登録フォームを開く際に必要なデータを取得するユースケースです。

*/
func FetchDataInputFormItems(groupRepo group.Persistance) ([]*domain.Category, error) {
	groups, err := groupRepo.SelectAll()

	if err != nil {
		log.Error("usecases:query:FetchDataInputFormItems:message:" + err.Error())
		return nil, err
	}
	return domain.CreateCategories(groups), nil

}
