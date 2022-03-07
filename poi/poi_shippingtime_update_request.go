package poi

//更新门店营业时间
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/poi/shippingtime/update
//文档地址：https://open-shangou.meituan.com/home/docDetail/33
type PoiShippingtimeUpdateRequest struct {
	AppPoiCode   string `json:"app_poi_code"`  //描述：APP方门店id，传商家中台系统里门店的编码，最长不超过128个字符。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ShippingTime string `json:"shipping_time"` //描述：门店营业时间 (1)一天中的营业时间支持传多个时段,以英文逗号分隔。 (2)一周的营业时间，支持按周一至周日的顺序分别设置营业时间，以英文分号隔开，每天都需要上传时段。 (3)注意格式:每个时段之间不能存在交集。
}
