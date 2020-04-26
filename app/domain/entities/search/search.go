package search

import (
	"regulus/app/domain/vo/query"
)

/*
SaveData is save query save data
*/
type SaveData struct {
	ID             string              `json:"id"`
	PatternName    string              `json:"patternName"`
	Category       string              `json:"category"`
	IsDisclose     bool                `json:"isDisclose"`
	DiscloseGroups []string            `json:"discloseGroups"`
	OwnerID        string              `json:"ownerID"`
	ConditionData  query.ConditionData `json:"conditionData"`
}
