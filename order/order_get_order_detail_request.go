package order

//获取订单详细信息
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/getOrderDetail
//文档地址：https://open-shangou.meituan.com/home/docDetail/135
type OrderGetOrderDetailRequest struct {
	OrderId       int64 `url:"order_id"`        //是否必须：是；描述：订单号（同订单展示ID），商家可根据订单号查询订单当前的详细信息。
	IsMtLogistics *int  `url:"is_mt_logistics"` //是否必须：否；描述：是否为美团配送。当需要查询美团配送的详细信息时，此字段需上传且取值为1；医药b2c（快递配送）此字段上传0。
}
