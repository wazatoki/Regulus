package query

import (
	"regulus/app/domain/entities"
	"regulus/app/domain/services"
	"regulus/app/domain/vo/query"
	"regulus/app/usecases/maintenance/master/group"
	"regulus/app/utils/log"
)

/*

UpdateCondition は検索条件更新時のユースケースです。

*/
func UpdateCondition(queryRepo persistance, queryCondition *entities.QueryCondition, operatorID string) error {

	err := queryRepo.Update(queryCondition, operatorID)

	if err != nil {
		log.Error("usecases:query:UpdateCondition:message:" + err.Error())
	}

	return err
}

/*

Find は検索時のユースケースです。条件指定がない場合は全件検索の結果を返します。

*/
func Find(queryRepo persistance, conditionData *query.ConditionData) ([]*entities.QueryCondition, error) {

	items, err := queryRepo.SelectAll()
	if err != nil {

		log.Error("usecases:query:Find:message:" + err.Error())
		return nil, err
	}

	items = items.FindBySearchStrings(conditionData.SearchStrings...)
	items = items.FindBySearchConditionItem(conditionData.SearchConditionList...)
	items = items.Sort(conditionData.OrderConditionList...)

	return items, nil
}

/*

AddCondition は検索条件追加時のユースケースです。成功した場合は id を返却します。

*/
func AddCondition(queryRepo persistance, queryCondition *entities.QueryCondition, operatorID string) (string, error) {

	id, err := queryRepo.Insert(queryCondition, operatorID)

	if err != nil {
		log.Error("usecases:query:AddCondition:message:" + err.Error())
	}

	return id, err
}

/*

FetchDataInputFormItems は検索条件登録フォームを開く際に必要なデータを取得するユースケースです。

*/
func FetchDataInputFormItems(groupRepo group.Persistance) ([]*entities.Category, error) {
	groups, err := groupRepo.SelectAll()

	if err != nil {
		log.Error("usecases:query:FetchDataInputFormItems:message:" + err.Error())
		return nil, err
	}
	return services.CreateCategories(groups), nil

}
