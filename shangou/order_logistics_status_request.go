package shangou

//获取订单配送状态
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/logistics/status
//文档地址：https://open-shangou.meituan.com/home/docDetail/144
type OrderLogisticsStatusRequest struct {
	OrderId int64 `url:"order_id"` //是否必须：是；描述：订单号（同订单展示ID）
}
