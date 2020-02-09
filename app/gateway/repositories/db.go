package repositories

import (
	"regulus/app/infrastructures/postgresql"

	"regulus/app/domain/vo/query"

	"github.com/jmoiron/sqlx"
)

func createDB() *postgresql.Postgresql {
	return postgresql.NewPostgresql()
}

type db interface {
	WithDbContext(fn func(db *sqlx.DB) error) error
}

func comparisonOperator(matchType query.MatchTypeEnum, val string) (string, string) {

	switch matchType {
	case query.Match:
		return "=", val
	case query.Unmatch:
		return "!=", val
	case query.Pertialmatch:
		return "like", "%" + val + "%"
	case query.Gt:
		return ">", val
	case query.Ge:
		return ">=", val
	case query.Lt:
		return "<", val
	case query.Le:
		return "<=", val
	default:
		return "=", val
	}
}
