package order

type OrderLogisticsStatusResponseData struct {
	Result           string  `json:"result"`
	OrderId          int64   `json:"order_id"`          //描述：订单号（同订单展示ID），数据库中请用bigint(20)存储此字段。
	LogisticsStatus  int     `json:"logistics_status"`  //描述：订单配送状态code。美团配送状态值有：0-配送单发往配送，5-配送侧压单，10-配送单已确认，15-骑手已到店，20-骑手已取货，40-骑手已送达，100-配送单已取消；医药b2c（快递配送）订单配送状态值有：0-初始，1-推送给商家，20-已发货，40-已送达（目前仅建仓门店的仓库相关订单才会有这个值，无仓门店的订单没有这个枚举）。
	LogisticsName    string  `json:"logistics_name"`    //描述：配送方名称，配送方id对应配送方名称：1-商家自配；2-趣活；3-达达；4-美团跑腿；5-加盟；6-自建；7-E代送；8-城市代理；100-角马；101-快送；102-混合送（专送+快送）；103-全城送；301-混合加盟；302-混合自建；303-混合快送
	SendTime         int     `json:"send_time"`         //描述：美团配送单发往配送的时间，此字段信息取最新一次发往配送的时间(logistics_status=0)，为10位秒级的时间戳。
	ConfirmTime      int     `json:"confirm_time"`      //描述：美团配送单确认时间，此字段信息取最新一次配送单确认的时间(logistics_status=10)，为10位秒级的时间戳。如当前无此状态，则字段值为0。
	CancelTime       int     `json:"cancel_time"`       //描述：美团配送单取消时间，此字段信息取最新一次配送单取消的时间(logistics_status=100)，为10位秒级的时间戳。如最新一次配送单最终未取消，则此字段值为0。
	FetchTime        int     `json:"fetch_time"`        //描述：美团骑手取单时间，此字段信息取最新一次骑手取单的时间(logistics_status=20)，为10位秒级的时间戳。如当前无此状态，则字段值为0。
	CompletedTime    int     `json:"completed_time"`    //描述：美团配送单完成时间，即骑手送达订单商品的时间(logistics_status=40)，为10位秒级的时间戳。如当前无此状态，则字段值为0。
	DispatcherName   string  `json:"dispatcher_name"`   //描述：美团配送骑手的姓名，取最新一次指派的骑手信息。
	DispatcherMobile string  `json:"dispatcher_mobile"` //描述：美团配送骑手的联系电话，取最新一次指派的骑手信息。订单取消、已完成或超过24小时该字段为空。
	PoiShippingFee   float64 `json:"poi_shipping_fee"`  //描述：商家承担的运费金额，单位是元。
	ShippingTips     string  `json:"shipping_tips"`     //描述：商家承担运费的具体说明
	TipAmount        float64 `json:"tip_amount"`        //描述：商家给定的配送小费金额，单位是元。
}
