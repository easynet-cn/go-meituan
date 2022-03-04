package app

//推送用户或客服取消订单
//请求方式：POST
//文档地址：https://open-shangou.meituan.com/home/docDetail/201
type AppCanceledOrderRequest struct {
	OrderId    int64  `url:"order_id"`     //是否必须：是；描述：订单号（同订单展示ID），数据库中请用bigint(20)存储此字段。
	ReasonCode int    `url:"reason_code"`  //是否必须：是；描述：订单取消的原因code，请商家查看【附录】文档里“取消订单原因code列表”。
	Reason     string `url:"reason"`       //是否必须：是；描述：订单取消的原因
	DealOpType string `url:"deal_op_type"` //是否必须：是；描述：当前订单取消操作人类型， 1-用户 2-商家端 3-客服 4-BD 5-系统 6-开放平台
}
