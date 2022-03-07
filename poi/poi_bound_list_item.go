package poi

type PoiBoundListItem struct {
	WmPoiId    int64  `json:"wm_poi_id"`    //描述：美团门店ID
	AppPoiCode string `json:"app_poi_code"` //描述：三方门店ID
	PoiName    string `json:"poi_name"`     //描述：美团门店名称
	IsOnline   int    `json:"is_online"`    //描述：门店上下线状态，参考值：0-下线；1-上线；2-上单中；3-审核通过可上线。
}
