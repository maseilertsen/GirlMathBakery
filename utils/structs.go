package utils

// Bake request
type BakeReq struct {
	Item  string `json:"item"`
	Qty   int    `json:"qty"`
	Time  string `json:"when,omitempty"` // ISO8601 - "YYYY-MM-DD"
	User  string `json:"user,omitempty"`
	Token string `json:"token"` // Secret password
}

/*DTOs*/

// ItemDTO A data transfer object
type ItemDTO struct {
	Item      string  `json:"item"`
	UnitCost  float64 `json:"unit_cost"`
	UnitStore float64 `json:"unit_store"`
	Unit      string  `json:"unit"`
}
