package shangou

//下发美团配送订单
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/logistics/push
//文档地址：https://open-shangou.meituan.com/home/docDetail/138
type OrderLogisticsPushRequest struct {
	OrderId int64 `url:"order_id"` //是否必须：是；描述：订单号（同订单展示ID）
}
