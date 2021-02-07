package login

import "regulus/app/domain/authentication"

// persistance API用のインターフェース
type persistance interface {
	SelectByAccountID(string) (*authentication.Staff, error)
}
