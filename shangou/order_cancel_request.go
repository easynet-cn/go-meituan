package shangou

//商家取消订单
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/cancel
//文档地址：https://open-shangou.meituan.com/home/docDetail/114
type OrderCancelRequest struct {
	OrderId    int64  `url:"order_id"`    //是否必须：是；描述：订单号（同订单展示ID），商家收到新订单后至订单完成前，均可调用接口操作取消此订单。
	Reason     string `url:"reason"`      //是否必须：否；描述：此字段传取商家消订单的原因
	ReasonCode int    `url:"reason_code"` //是否必须：否；描述：此字段传美团规范化取消订单的原因code，可参考【附录】文档的“取消订单原因code列表”；字符长度不能超过9个字符。
}
