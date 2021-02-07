package maker

import (
	"regulus/app/domain/supplier"
)

// persistance API用のインターフェース
type persistance interface {
	Insert(*supplier.Maker) (string, error)
	Update(*supplier.Maker) error
	Delete(string) error
	SelectByID(string) (*supplier.Maker, error)
	SelectAll() ([]supplier.Maker, error)
	Select(string) ([]supplier.Maker, error)
}
