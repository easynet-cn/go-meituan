package poi

type EcommercePoiBoundListResponse struct {
	Data      []PoiBoundListItem `json:"data"`
	ExtraInfo string             `json:"extra_info"`
}
