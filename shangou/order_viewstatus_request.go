package shangou

//查询订单状态
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/viewstatus
//文档地址：https://open-shangou.meituan.com/home/docDetail/129
type OrderViewstatusRequest struct {
	OrderId int64 `url:"order_id"` //是否必须：是；描述：订单号（同订单展示ID），根据订单号查询订单当前的状态。
}
