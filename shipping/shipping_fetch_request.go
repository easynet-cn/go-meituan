package shipping

//获取门店配送范围（混合送）
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/shipping/fetch
//文档地址：https://open-shangou.meituan.com/home/docDetail/51
type ShippingFetchRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，传商家中台系统里门店的编码，最长不超过128个字符。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
}
