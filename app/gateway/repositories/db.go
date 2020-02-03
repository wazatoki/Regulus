package repositories

import (
	"regulus/app/infrastructures/postgresql"

	"github.com/jmoiron/sqlx"
)

func createDB() *postgresql.Postgresql {
	return postgresql.NewPostgresql()
}

type db interface {
	WithDbContext(fn func(db *sqlx.DB) error) error
}

func comparisonOperator(matchType string, val string) (string, string) {

	switch matchType {
	case "match":
		return "=", val
	case "unmatch":
		return "!=", val
	case "pertialmatch":
		return "like", "%" + val + "%"
	case "gt":
		return ">", val
	case "ge":
		return "<=", val
	case "lt":
		return "<", val
	case "le":
		return "<=", val
	default:
		return "=", val
	}
}
