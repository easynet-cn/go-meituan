package shangou

//商家确认订单
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/order/confirm
//文档地址：https://open-shangou.meituan.com/home/docDetail/111
type OrderConfirmRequest struct {
	OrderId int64 `json:"order_id"` //描述：订单号（同订单展示ID），当商家收到新订单推送消息后，可调用接口操作接单。
}
