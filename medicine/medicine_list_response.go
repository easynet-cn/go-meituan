package medicine

//查询门店药品列表响应
type MedicineListResponse struct {
	Data      []string  `json:"data"`
	ExtraInfo ExtraInfo `json:"extra_info"`
}
