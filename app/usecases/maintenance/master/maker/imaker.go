package maker

import (
	makerEntity "regulus/app/domain/entities"
)

// persistance API用のインターフェース
type persistance interface {
	Insert(*makerEntity.Maker) (string, error)
	Update(*makerEntity.Maker) error
	Delete(string) error
	SelectByID(string) (*makerEntity.Maker, error)
	SelectAll() ([]makerEntity.Maker, error)
	Select(string) ([]makerEntity.Maker, error)
}
