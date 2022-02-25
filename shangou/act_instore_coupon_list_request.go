package shangou

//【商家券】查询商家券（店内发券）活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/instore/coupon/list
//文档地址：https://open-shangou.meituan.com/home/docDetail/317
type ActInstoreCouponListRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店IDAPP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	PageNum    int    `url:"page_num"`     //是否必须：是；描述：分页查询偏移量，此字段为分页查询时页码偏移量，表示从第几页开始查，1为第一页。 例如：offset=2，limit=3时， (1)如活动总数有8条，则会分为13，46，78，此时查询结果返回46的数据； (2)如活动商品总数有5，则会分为13，45，此时查询结果返回45的数据。
	PageSize   int    `url:"page_size"`    //是否必须：是；描述：分页查询时每页展示的条数，此字段信息最大不超过200。 例如：offset=2，limit=3时， (1)如活动总数有8条，则会分为13，46，78，此时查询结果返回46的数据； (2)如活动商品总数有5，则会分为13，45，此时查询结果返回45的数据。
}
