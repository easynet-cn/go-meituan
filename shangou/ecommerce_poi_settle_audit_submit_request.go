package shangou

//批量提交门店审核
type EcommercePoiSettleAuditSubmitRequest struct {
	AppPoiCodes string `url:"app_poi_codes"` //是否必须：是；描述：提审门店列表，用逗号分割，最多支持50个
}
