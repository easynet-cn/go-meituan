package order

//设订单为商家已收到
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/poi_received
//文档地址：https://open-shangou.meituan.com/home/docDetail/108
type OrderPoiReceivedRequest struct {
	OrderId int64 `url:"order_id"` //是否必须：是；描述：订单号（同订单展示ID），当商家收到新订单推送消息后，可调用此接口向美团同步商家系统已收到此订单信息的消息。
}
