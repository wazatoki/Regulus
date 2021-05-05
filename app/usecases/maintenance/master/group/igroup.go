package group

import (
	"regulus/app/domain"
)

// Persistance API用のインターフェース
type Persistance interface {
	//Insert(*group.Group) (string, error)
	//Update(*group.Group) error
	//Delete(string) error
	//SelectByID(string) (*group.Group, error)
	SelectAll() ([]*domain.Group, error)
	//Select(string) ([]group.Group, error)
}
