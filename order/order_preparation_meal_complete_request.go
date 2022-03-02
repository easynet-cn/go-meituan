package order

//商家确认已完成拣货
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/order/preparationMealComplete
//文档地址：https://open-shangou.meituan.com/home/docDetail/217
type OrderPreparationMealCompleteRequest struct {
	OrderId int64 `url:"order_id"` //是否必须：是；描述：订单号（同订单展示ID）
}
