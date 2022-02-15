package shangou

//批量创建/更新单店请求
type EcommercePoiSettleSingleSaveRequest struct {
	Type      int    //是否必须：是；描述：0为创建，1为更新
	ApplyInfo string //是否必须：是；描述：创建/更新门店信息集合，json格式数组，最多支持传入50个门店信息 （1）单店更新时仅支持更新被驳回的模块信息，其他模块数据不会生效且没有错误信息。
}
