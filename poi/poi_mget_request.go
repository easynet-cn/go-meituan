package poi

//批量获取门店详细信息
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/poi/mget
//文档地址：https://open-shangou.meituan.com/home/docDetail/9
type PoiMgetRequest struct {
	AppPoiCodes string `url:"app_poi_codes"` //是否必须：是；描述：APP方门店id，传商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。 支持传多个门店id批量查询，一次调用可上传200个门店id，多个之间以英文逗号分隔；支持部分门店查询成功。 仅支持返回门店id正确的门店信息。
}
