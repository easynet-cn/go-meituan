package shangou

//【商家券】批量删除商家券（店内发券）活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/instore/coupon/delete
//文档地址：https://open-shangou.meituan.com/home/docDetail/316
type ActInstoreCouponDeleteRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActIds     string `url:"act_ids"`      //是否必须：是；描述：活动id，上传多个需删除的折扣活动id时，用英文逗号隔开；同一次调用最多可上传的活动id数量不能超过100个。注意：同一次调用中，不支持部分活动删除成功。
}
