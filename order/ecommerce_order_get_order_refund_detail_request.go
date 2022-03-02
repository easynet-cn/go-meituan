package order

//获取订单退款记录
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/ecommerce/order/getOrderRefundDetail
//文档地址：https://open-shangou.meituan.com/home/docDetail/322
type EcommerceOrderGetOrderRefundDetailRequest struct {
	WmOrderIdView int64 `url:"wm_order_id_view"` //是否必须：是；描述：订单号（同订单展示ID）
	RefundType    int   `url:"refund_type"`      //是否必须：否；描述：退款类型：1-全额退款；2-部分退款。如不传此字段代表查询全部类型。
}
