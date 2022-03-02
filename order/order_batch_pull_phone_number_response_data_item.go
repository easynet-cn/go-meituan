package order

type OrderBatchPullPhoneNumberResponseDataItem struct {
	OrderId         int64  `json:"order_id"`          //描述：订单号
	AppPoiCode      string `json:"app_poi_code"`      //描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	WmOrderIdView   int64  `json:"wm_order_id_view"`  //描述：订单展示ID，与用户端、商家端订单详情中展示的订单号码一致。数据库中请用bigint(20)存储此字段。
	DaySeq          int    `json:"day_seq"`           //描述：订单的流水号，门店内每日已支付订单的流水号从1开始。
	RealPhoneNumber string `json:"real_phone_number"` //描述：订单收货人的真实手机号码
}
