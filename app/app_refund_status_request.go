package app

//推送退款状态（成功/失败）信息
//请求方式：POST
//文档地址：https://open-shangou.meituan.com/home/docDetail/502
type AppRefundRequest struct {
	OrderId      int64  `url:"order_id"`      //是否必须：是；描述：订单展示ID，与用户端、商家端订单详情中展示的订单号码一致。数据库中请用bigint(20)存储此字段。
	RefundId     string `url:"refund_id"`     //是否必须：是；描述：退款id，每次发起部分退款的退款id不同。
	RefundStatus string `url:"refund_status"` //是否必须：是；描述：退款状态：3-全部退款退款成功；4-全部退款失败；13-部分退款退款成功；14-部款退款失败
	Money        string `url:"money"`         //是否必须：是；描述：本次退款的合计金额，单位「元」。
	CTime        string `url:"ctime"`         //是否必须：是；描述：退款申请的发起时间。
	UTime        string `url:"utime"`         //是否必须：是；描述：退款状态的更新时间。
}
