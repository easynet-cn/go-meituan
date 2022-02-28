package shangou

//【商品券】批量查询商品券活动中的商品
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/act/goods/coupon/goodslist
//文档地址：https://open-shangou.meituan.com/home/docDetail/379
type ActGoodsCouponGoodslistRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActIds     string `url:"act_ids"`      //是否必须：是；描述：活动id，同一次调用中支持传入的活动id个数不能超过100个，多个之间用英文逗号分隔。
}
