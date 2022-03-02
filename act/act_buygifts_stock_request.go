package act

//【买赠】批量修改买赠活动当日活动库存
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/buygifts/stock
//文档地址：https://open-shangou.meituan.com/home/docDetail/250
type ActBuygiftsStockRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActData    string `url:"act_data"`     //是否必须：是；描述：买赠活动的当日活动赠品库存信息，json格式数组。(1)参数包括：item_id(必填项)，为活动id。gifts_day_limit(必填项)，为当日活动赠品库存；仅支持上传大于0的正整数或1，如传1时则表示不限制活动赠品库存。(2)同一次调用可上传的活动数据最多不能超过200组；支持部分修改成功。
}
