package poi

type PoiGetidsResponse struct {
	Data []string `json:"data"` //描述：APP方门店id，为商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
}
