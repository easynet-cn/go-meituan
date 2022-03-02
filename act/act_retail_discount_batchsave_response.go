package act

type ActRetailDiscountBatchsaveResponse struct {
	Data       string                                             `json:"data"`
	SuccessMsg []ActRetailDiscountBatchsaveResponseMsgSuccessItem `json:"success_msg"` //描述：创建成功的商品折扣活动基本信息，json格式。
	ExtraInfo  ActRetailDiscountBatchsaveResponseExtraInfo        `json:"extra_info"`
}
