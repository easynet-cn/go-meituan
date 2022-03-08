package app

//门店审核结果推送
//请求方式：POST
//文档地址：https://open-shangou.meituan.com/home/docDetail/535
type AppPoiAuditResultRequest struct {
	AppId       string `url:"app_id"`       //是否必须：是；描述：APP方id
	AppPoiCode  string `url:"app_poi_code"` //是否必须：是；描述：门店编号
	PoiType     string `url:"poi_type"`     //是否必须：是；描述：门店类型，1:单店，2:多店
	AuditDetail string `url:"audit_detail"` //是否必须：是；描述：审核详情，json数组
}
