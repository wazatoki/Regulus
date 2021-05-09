package usecases

import (
	"regulus/app/domain"
)

// groupRepo API用のインターフェース
type groupRepo interface {
	SelectAll() (domain.Groups, error)
}
