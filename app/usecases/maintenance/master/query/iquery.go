package query

import (
	"regulus/app/domain"
)

// persistance API用のインターフェース
type persistance interface {
	Insert(*domain.Condition, string) (string, error)
	Update(*domain.Condition, string) error
	Delete(string, string) error
	SelectByID(string) (*domain.Condition, error)
	SelectAll() ([]*domain.Condition, error)
	Select(...domain.SearchConditionItem) ([]*domain.Condition, error)
}
