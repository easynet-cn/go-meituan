package shipping

type ShippingListResponse struct {
	Data []ShippingListItem `json:"data"` //描述：门店配送范围信息，json格式数组。
}
