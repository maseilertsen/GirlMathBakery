package utils

type BakeReq struct {
	Item  string `json:"item"`
	Qty   int    `json:"qty"`
	Time  string `json:"when,omitempty"` // ISO8601 - "YYYY-MM-DD"
	User  string `json:"user,omitempty"`
	Token string `json:"token"` // Secret password
}
