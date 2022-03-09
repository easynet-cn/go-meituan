package shipping

type ShippingFetchResponse struct {
	Data []ShippingFetchItem `json:"data"` //描述：门店配送信息，json格式。
}
