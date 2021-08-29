package domain

// Staff 利用者グループを示す構造体
type Staff struct {
	ID                      string      `json:"id"`
	AccountID               string      `json:"staffID"`
	Password                string      `json:"password"`
	Name                    string      `json:"name"`
	StaffGroups             StaffGroups `json:"staffGroups"`
	OeratorUsableConditions Conditions  `json:"operatorUsableConditions"`
}

// Group 利用者グループを示す構造体
type StaffGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
