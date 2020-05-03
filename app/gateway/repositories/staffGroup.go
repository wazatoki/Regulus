package repositories

import (
	"context"
	"regulus/app/domain/entities"
	"regulus/app/infrastructures/sqlboiler"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// SelectAll select all group data without not del from database
func (g *StaffGroupRepo) SelectAll() ([]entities.StaffGroup, error) {
	geSlice := []entities.StaffGroup{}

	err := g.database.WithDbContext(func(db *sqlx.DB) error {
		queries := []qm.QueryMod{
			qm.Where(sqlboiler.StaffGroupColumns.Del+"!=?", true),
		}

		groups, err := sqlboiler.StaffGroups(queries...).All(context.Background(), db.DB)

		if err == nil {

			for _, group := range groups {
				var ge *entities.StaffGroup
				ge = &entities.StaffGroup{}

				ge.ID = group.ID
				ge.Name = group.Name

				geSlice = append(geSlice, *ge)
			}
		}

		return err
	})

	return geSlice, err
}

// StaffGroupObjectMap data mapper sqlboiler object to entities object
func StaffGroupObjectMap(sg *sqlboiler.StaffGroup) (eg entities.StaffGroup) {
	eg = entities.StaffGroup{
		ID:   sg.ID,
		Name: sg.Name,
	}
	return
}

// NewStaffGroupRepo constructor
func NewStaffGroupRepo() *StaffGroupRepo {
	return &StaffGroupRepo{database: createDB()}
}

// StaffGroupRepo repository struct
type StaffGroupRepo struct {
	database db
}
