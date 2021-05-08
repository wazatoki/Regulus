package repositories

import (
	"regulus/app/infrastructures/postgresql"

	"regulus/app/domain"

	"github.com/jmoiron/sqlx"
)

func createDB() *postgresql.Postgresql {
	return postgresql.NewPostgresql()
}

type db interface {
	WithDbContext(fn func(db *sqlx.DB) error) error
}

func comparisonOperator(matchType domain.QueryMatchType, val string) (string, string) {

	switch matchType.String() {
	case domain.QueryMatchTypeEnum.MATCH.String():
		return "=", val
	case domain.QueryMatchTypeEnum.UNMATCH.String():
		return "!=", val
	case domain.QueryMatchTypeEnum.PERTIALMATCH.String():
		return "like", "%" + val + "%"
	case domain.QueryMatchTypeEnum.GT.String():
		return ">", val
	case domain.QueryMatchTypeEnum.GE.String():
		return ">=", val
	case domain.QueryMatchTypeEnum.LT.String():
		return "<", val
	case domain.QueryMatchTypeEnum.LE.String():
		return "<=", val
	case domain.QueryMatchTypeEnum.IN.String():
		return "in", val
	default:
		return "=", val
	}
}
