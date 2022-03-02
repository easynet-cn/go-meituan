package order

//驳回订单退款申请
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/refund/reject
//文档地址：https://open-shangou.meituan.com/home/docDetail/126
type OrderRefundRejectRequest struct {
	OrderId int64  `url:"order_id"` //是否必须：是；描述：订单号（同订单展示ID），当商家收到用户的退款申请，如拒绝申请，商家可调用此接口操作驳回此订单的退款申请。
	Reason  string `url:"reason"`   //是否必须：是；描述：此字段传驳回退款的原因
}
