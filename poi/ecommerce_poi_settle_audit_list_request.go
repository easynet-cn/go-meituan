package poi

//批量查询门店审核状态
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/ecommerce/poi/settle/audit/list
//文档地址：https://open-shangou.meituan.com/home/docDetail/531
type EcommercePoiSettleAuditListRequest struct {
	AppPoiCodes string `url:"app_poi_codes"` //是否必须：是；描述：查询门店列表，最多支持50个
	AuditType   *int   `url:"audit_type"`    //是否必须：否；描述：审核类型：1-门店入驻；2-门店logo(头图)。不传默认为1，查门店入驻审核结果。
}
