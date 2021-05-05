package login

import "regulus/app/domain"

// persistance API用のインターフェース
type persistance interface {
	SelectByAccountID(string) (*domain.Staff, error)
}
