package act

//【商品券】创建商品券活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/goods/coupon/save
//文档地址：https://open-shangou.meituan.com/home/docDetail/376
type ActGoodsCouponSaveRequest struct {
	GoodsCouponData string `url:"goods_coupon_data"` //是否必须：是；描述：商品券活动信息，json格式字符串。 注意：app_food_code即为药品的app_medicine_code。
}
