package shangou

//取消美团配送订单
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/logistics/cancel
//文档地址：https://open-shangou.meituan.com/home/docDetail/141
type OrderLogisticsCancelRequest struct {
	OrderId int64 `url:"app_poi_code"` //是否必须：是；描述：订单号（同订单展示ID）
}
