package repositories

import (
	"context"
	groupEntity "regulus/app/domain/entities"
	"regulus/app/infrastructures/sqlboiler"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// GroupRepo repository struct
type GroupRepo struct {
	database db
}

// SelectAll select all group data without not del from database
func (g *GroupRepo) SelectAll() ([]groupEntity.Group, error) {
	geSlice := []groupEntity.Group{}

	err := g.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffGroupColumns.Del+"!=?", true),
		}

		groups, err := sqlboiler.StaffGroups(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, group := range groups {
				var ge *groupEntity.Group
				ge = &groupEntity.Group{}

				ge.ID = group.ID
				ge.Name = group.Name

				geSlice = append(geSlice, *ge)
			}
		}

		return err
	})

	return geSlice, err
}

// NewGroupRepo constructor
func NewGroupRepo() *GroupRepo {
	return &GroupRepo{database: createDB()}
}
