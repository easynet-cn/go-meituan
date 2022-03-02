package shangou

type OrderViewstatusResponseData struct {
	Status int `json:"status"` //描述：订单状态，返回订单当前的状态。目前平台的订单状态参考值有：1-用户已提交订单；2-向商家推送订单；4-商家已确认；8-订单已完成；9-订单已取消。
}
