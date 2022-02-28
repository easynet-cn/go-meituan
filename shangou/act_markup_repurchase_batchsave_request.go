package shangou

//【加价购】批量新增或者更新加价购活动接口
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/markup/repurchase/batchsave
//文档地址：https://open-shangou.meituan.com/home/docDetail/363
type ActMarkupRrepurchaseBatchsaveRequest struct {
	AppPoiCode string                                      `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActData    ActMarkupRrepurchaseBatchsaveRequestActData `url:"act_data"`     //是否必须：是；描述：活动信息，json格式字符串。门店内同一时间只能有一个生效中的加价购活动。门店生效中和待生效的加价换购活动最多不能超过5个。 注意：app_food_code即为药品的app_medicine_code。
}
