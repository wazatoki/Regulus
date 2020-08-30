package entities

// Staff 利用者グループを示す構造体
type Staff struct {
	ID        string        `json:"id"`
	AccountID string        `json:"staffID"`
	Password  string        `json:"password"`
	Name      string        `json:"name"`
	Groups    []*StaffGroup `json:"groups"`
}
