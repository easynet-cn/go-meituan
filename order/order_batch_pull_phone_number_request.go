package order

//拉取用户真实手机号
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/order/batchPullPhoneNumber
//文档地址：https://open-shangou.meituan.com/home/docDetail/223
type OrderBatchPullPhoneNumberRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：非必填；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	Offset     int    `url:"offset"`       //是否必须：是；描述：分页查询的偏移量，表示本次查询从第几条数据开始查，0是第一条。
	Limit      int    `url:"limit"`        //是否必须：是；描述：分页查询每页展示数量，最大不能超过1000。
}
