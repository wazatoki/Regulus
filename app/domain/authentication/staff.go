package authentication

// Staff 利用者グループを示す構造体
type Staff struct {
	ID        string   `json:"id"`
	AccountID string   `json:"staffID"`
	Password  string   `json:"password"`
	Name      string   `json:"name"`
	Groups    []*Group `json:"groups"`
}

// Group 利用者グループを示す構造体
type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
