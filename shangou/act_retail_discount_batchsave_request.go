package shangou

//【折扣】批量创建或更新药品折扣商品和爆品商品活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/retail/discount/batchsave
//文档地址：https://open-shangou.meituan.com/home/docDetail/305
type ActRetailDiscountBatchsaveRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActType    string `url:"act_type"`     //是否必须：否；描述：活动类型可取值： 1001折扣活动 56爆品活动 创建时如不传本参数，则默认为创建折扣活动，按折扣活动的规则进行各项校验。
	ActData    string `url:"act_data"`     //是否必须：是；描述：折扣活动信息，json格式数组。(1)同一次调用中，此字段支持上传的活动信息不能超过50组。(2)同一次调用中，支持创建活动和更新活动两种方式同时存在。(3)支持折扣活动部分创建成功。 注意：app_food_code即为药品的app_medicine_code。
}
