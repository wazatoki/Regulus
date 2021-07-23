package usecases

import (
	"errors"
	"regulus/app/domain"
	"regulus/app/utils/log"
)

/*

GroupFind は検索時のユースケースです。条件指定がない場合は全件検索の結果を返します。

*/
func GroupFind(conditionData *domain.ConditionData, groupRepo groupRepo) (domain.StaffGroups, error) {

	items, err := groupRepo.Select(conditionData.SearchConditionList...)

	if err != nil {

		log.Error("usecases:GroupFind:message:" + err.Error())
		return nil, err
	}

	items = items.Sort()
	return items, nil
}

/*

GroupDelete は検索条件削除時のユースケースです。

*/
func GroupDelete(staffGroupIDs *[]string, groupRepo groupRepo, operatorID string) (result []domain.StaffGroup) {

	result = []domain.StaffGroup{}

	for _, id := range *staffGroupIDs {

		err := groupRepo.Delete(id, operatorID)

		if err != nil {
			log.Error("usecases:GroupDelete:message:" + err.Error())
			sg, _ := groupRepo.SelectByID(id)
			result = append(result, *sg)
		}
	}

	return
}

/*

GroupUpdate は検索条件更新時のユースケースです。

*/
func GroupUpdate(staffGroup *domain.StaffGroup, groupRepo groupRepo, operatorID string) error {

	// nameの重複をチェック
	condition := &domain.SearchConditionItem{
		SearchField: domain.FieldAttr{
			ID: "name",
		},
		ConditionValue: staffGroup.Name,
		MatchType:      domain.QueryMatchTypeEnum.MATCH,
		Operator:       domain.QueryOperatorEnum.AND,
	}
	g, e := groupRepo.Select(*condition)

	if e != nil {
		log.Error("usecases:GroupAdd:message:" + e.Error())
		return e
	}

	if g != nil {
		log.Error("usecases:GroupAdd:message: duplicate staffGroup name")
		return errors.New("duplicate-staffGroup-name")
	}

	err := groupRepo.Update(staffGroup, operatorID)

	if err != nil {
		log.Error("usecases:GroupUpdate:message:" + err.Error())
	}

	return err
}

/*

GroupAdd は検索条件追加時のユースケースです。成功した場合は id を返却します。

*/
func GroupAdd(staffGroup *domain.StaffGroup, groupRepo groupRepo, operatorID string) (*domain.StaffGroup, error) {

	// nameの重複をチェック
	condition := &domain.SearchConditionItem{
		SearchField: domain.FieldAttr{
			ID: "name",
		},
		ConditionValue: staffGroup.Name,
		MatchType:      domain.QueryMatchTypeEnum.MATCH,
		Operator:       domain.QueryOperatorEnum.AND,
	}
	g, e := groupRepo.Select(*condition)

	if e != nil {
		log.Error("usecases:GroupAdd:message:" + e.Error())
		return nil, e
	}

	if g != nil {
		log.Error("usecases:GroupAdd:message: duplicate staffGroup name")
		return nil, errors.New("duplicate-staffGroup-name")
	}

	id, err := groupRepo.Insert(staffGroup, operatorID)

	if err != nil {
		log.Error("usecases:GroupAdd:message:" + err.Error())
		return nil, err
	}

	result, e := groupRepo.SelectByID(id)
	return result, e
}

// groupRepo API用のインターフェース
type groupRepo interface {
	Insert(*domain.StaffGroup, string) (string, error)
	Update(*domain.StaffGroup, string) error
	Delete(string, string) error
	SelectByID(string) (*domain.StaffGroup, error)
	SelectAll() (domain.StaffGroups, error)
	Select(...domain.SearchConditionItem) (domain.StaffGroups, error)
}
