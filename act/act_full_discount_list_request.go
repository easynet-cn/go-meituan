package act

//【商品满减】批量查询指定商品满减活动或店内满减活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/full/discount/list
//文档地址：https://open-shangou.meituan.com/home/docDetail/319
type ActFullDiscountListRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActType    int    `url:"act_type"`     //是否必须：是；描述：活动类型，取值范围：0全店满减。药品类商家仅支持传0。
}
