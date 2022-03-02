package order

//自配订单已送达（不支持同步骑手信息）
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/arrived
//文档地址：https://open-shangou.meituan.com/home/docDetail/120
type OrderArrivedRequest struct {
	OrderId int64 `url:"order_id"` //是否必须：是；描述：订单号（同订单展示ID），门店配送方式仅为自配送的商家在接单后，如商品已完成配送，可调用此接口同步此订单配送状态“商品已送达”至用户端。
}
