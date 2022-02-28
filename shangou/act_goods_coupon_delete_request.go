package shangou

//【商品券】删除商品券活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/goods/coupon/delete
//文档地址：https://open-shangou.meituan.com/home/docDetail/377
type ActGoodsCouponDeleteRequest struct {
	ActIds string `url:"act_ids"` //是否必须：是；描述：活动id列表，传需要删除的活动id。同一次调用中，可上传需删除的活动id不能超过100个，多个之间用英文逗号分隔；不支持部分删除成功。
}
