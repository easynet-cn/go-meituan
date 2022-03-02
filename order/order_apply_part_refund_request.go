package order

//发起部分退款
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/order/applyPartRefund
//文档地址：https://open-shangou.meituan.com/home/docDetail/165
type OrderApplyPartRefundRequest struct {
	OrderId  int64                                     `url:"order_id"`  //是否必须：是；描述：订单号（同订单展示ID）
	Reason   string                                    `url:"reason"`    //是否必须：是；描述：部分退款的原因，此字段信息传商家发起本次部分退款的原因。
	FoodData []OrderApplyPartRefundRequestFoodDataItem `url:"food_data"` //是否必须：是；描述：部分退款商品数据集合的json格式数组
}
