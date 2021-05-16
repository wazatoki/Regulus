package usecases

import (
	"regulus/app/domain"
	"regulus/app/utils/log"
)

/*

QueryFetchDataInputFormItems は検索条件登録フォームを開く際に必要なデータを取得するユースケースです。

*/
func QueryFetchDataInputFormItems(groupRepo groupRepo) ([]*domain.Category, error) {
	groups, err := groupRepo.SelectAll()

	if err != nil {
		log.Error("usecases:query:FetchDataInputFormItems:message:" + err.Error())
		return nil, err
	}
	return domain.CreateCategories(groups), nil

}

/*

QueryFind は検索時のユースケースです。条件指定がない場合は全件検索の結果を返します。

*/
func QueryFind(conditionData *domain.ConditionData, queryRepo queryRepo) (domain.Conditions, error) {

	items, err := queryRepo.Select(conditionData.SearchConditionList...)

	if err != nil {

		log.Error("usecases:QueryFind:message:" + err.Error())
		return nil, err
	}

	items = items.Sort(conditionData.OrderConditionList...)
	return items, nil
}

/*

QueryDelete は検索条件削除時のユースケースです。

*/
func QueryDelete(queryConditionIDs *[]string, queryRepo queryRepo, operatorID string) (result []domain.Condition) {

	result = []domain.Condition{}

	for _, id := range *queryConditionIDs {

		err := queryRepo.Delete(id, operatorID)

		if err != nil {
			log.Error("usecases:QueryDelete:message:" + err.Error())
			c, _ := queryRepo.SelectByID(id)
			result = append(result, *c)
		}
	}

	return
}

/*

QueryUpdate は検索条件更新時のユースケースです。

*/
func QueryUpdate(queryCondition *domain.Condition, queryRepo queryRepo, operatorID string) error {

	err := queryRepo.Update(queryCondition, operatorID)

	if err != nil {
		log.Error("usecases:QueryUpdate:message:" + err.Error())
	}

	return err
}

/*

QueryAdd は検索条件追加時のユースケースです。成功した場合は id を返却します。

*/
func QueryAdd(queryCondition *domain.Condition, queryRepo queryRepo, operatorID string) (*domain.Condition, error) {

	id, err := queryRepo.Insert(queryCondition, operatorID)

	if err != nil {
		log.Error("usecases:QueryAdd:message:" + err.Error())
		return nil, err
	}

	result, e := queryRepo.SelectByID(id)
	return result, e
}

// persistance API用のインターフェース
type queryRepo interface {
	Insert(*domain.Condition, string) (string, error)
	Update(*domain.Condition, string) error
	Delete(string, string) error
	SelectByID(string) (*domain.Condition, error)
	SelectAll() (domain.Conditions, error)
	Select(...domain.SearchConditionItem) (domain.Conditions, error)
}
