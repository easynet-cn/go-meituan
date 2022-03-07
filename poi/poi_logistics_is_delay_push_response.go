package poi

type PoiLogisticsIsDelayPushResponse struct {
	Data string `json:"data"` //描述：门店是否支持延迟发配送的信息，json格式字符串。
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
