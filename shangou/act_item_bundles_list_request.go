package shangou

//【X元M件】查询X元M件活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/item/bundles/list
//文档地址：https://open-shangou.meituan.com/home/docDetail/330
type ActItemBundlesListRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
}
