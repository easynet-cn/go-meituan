package order

//订单确认退款请求
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/refund/agree
//文档地址：https://open-shangou.meituan.com/home/docDetail/123
type OrderRefundAgreeRequest struct {
	OrderId int64  `url:"order_id"` //是否必须：是；描述：订单号（同订单展示ID），当商家收到用户的退款申请，如同意退款，商家可调用此接口操作确认此订单的退款申请。
	Reason  string `url:"reason"`   //是否必须：是；描述：此字段传确认退款的原因
}
