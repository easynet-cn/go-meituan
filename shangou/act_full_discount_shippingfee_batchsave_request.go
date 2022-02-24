package shangou

//【运费满减】批量创建阶梯满减配送费活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/full/discount/shippingfee/batchsave
//文档地址：https://open-shangou.meituan.com/home/docDetail/302
type ActFullDiscountShippingfeeBatchsaveRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActData    string `url:"act_data"`     //是否必须：是；描述：门店满减运费活动信息，json格式数组。(1)同一门店同一时间仅允许参加一个活动。(2)同一门店的生效中+待生效活动最多支持20组。
}
