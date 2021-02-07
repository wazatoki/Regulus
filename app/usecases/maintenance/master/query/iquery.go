package query

import (
	"regulus/app/domain/query"
)

// persistance API用のインターフェース
type persistance interface {
	Insert(*query.Condition, string) (string, error)
	Update(*query.Condition, string) error
	Delete(string, string) error
	SelectByID(string) (*query.Condition, error)
	SelectAll() ([]*query.Condition, error)
	Select(...query.SearchConditionItem) ([]*query.Condition, error)
}
