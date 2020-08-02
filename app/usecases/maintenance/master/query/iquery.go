package query

import (
	"regulus/app/domain/entities"
	"regulus/app/domain/vo/query"
)

// persistance API用のインターフェース
type persistance interface {
	Insert(*entities.QueryCondition) (string, error)
	Update(*entities.QueryCondition) error
	Delete(string) error
	SelectByID(string) (entities.QueryCondition, error)
	SelectAll() ([]entities.QueryCondition, error)
	Select(...query.SearchConditionItem) ([]entities.QueryCondition, error)
}
