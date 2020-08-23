package login

import "regulus/app/domain/entities"

// persistance API用のインターフェース
type persistance interface {
	SelectByAccountID(string) (*entities.Staff, error)
}
