package query

import (
	"regulus/app/domain/entities"
	"regulus/app/domain/query"
)

// persistance API用のインターフェース
type persistance interface {
	Insert(*entities.QueryCondition, string) (string, error)
	Update(*entities.QueryCondition, string) error
	Delete(string, string) error
	SelectByID(string) (*entities.QueryCondition, error)
	SelectAll() ([]*entities.QueryCondition, error)
	Select(...query.SearchConditionItem) ([]*entities.QueryCondition, error)
}
