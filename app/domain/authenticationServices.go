package domain

import (
	"sort"
)

type StaffGroups []*StaffGroup

/*
Sort is sort staffGroup slice by name
*/
func (g *StaffGroups) Sort() StaffGroups {
	staffGroups := *g
	sort.Slice(staffGroups, func(i int, j int) bool {

		return staffGroups[i].Name < staffGroups[j].Name

	})
	return staffGroups
}
