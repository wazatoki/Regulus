package entities

// Staff 利用者グループを示す構造体
type Staff struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Groups []Group `json:"groups"`
}
