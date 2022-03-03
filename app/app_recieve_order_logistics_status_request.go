package app

//推送美配/企客订单配送状态
//请求方式：POST
//文档地址：https://open-shangou.meituan.com/home/docDetail/171
type AppRecieveOrderLogisticsStatusRequest struct {
	OrderId          int64  `url:"order_id"`          //是否必须：是；描述：订单号，数据库中请用bigint(20)存储此字段。
	WmOrderIdView    int64  `url:"wm_order_id_view"`  //是否必须：是；描述：订单展示ID，与用户端、商家端订单详情中展示的订单号码一致。数据库中请用bigint(20)存储此字段。
	LogisticsStatus  int    `url:"logistics_status"`  //是否必须：是；描述：美团配送订单状态code，目前美团配送状态值有：0-配送单发往配送，5-配送侧压单，10-配送单已确认，15-骑手已到店，20-骑手已取货，40-骑手已送达，100-配送单已取消。
	Time             int    `url:"time"`              //是否必须：是；描述：订单配送状态变更为当前状态的时间，推送10位秒级的时间戳。
	DispatcherName   string `url:"dispatcher_name"`   //是否必须：是；描述：美团配送骑手的姓名，取最新一次指派的骑手信息。
	DispatcherMobile string `url:"dispatcher_mobile"` //是否必须：是；描述：美团配送骑手的联系电话，取最新一次指派的骑手信息。 骑手手机号会以隐私号形式推送，请兼容13812345678和13812345678_123456两种号码格式，最多不超过20位，以便对接隐私号订单。
}
