package poi

//批量提交门店审核
//请求地址：https://waimaiopen.meituan.com/api/v1/ecommerce/poi/settle/audit/submit
type EcommercePoiSettleAuditSubmitRequest struct {
	AppPoiCodes string `url:"app_poi_codes"` //是否必须：是；描述：提审门店列表，用逗号分割，最多支持50个
}
