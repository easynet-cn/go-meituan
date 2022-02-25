package shangou

type ActFullDiscountBatchsaveResponse struct {
	Data       string                                    `json:"data"`
	SuccessMsg string                                    `json:"success_msg"` //描述：活动id
	ExtraInfo  ActFullDiscountBatchsaveResponseExtraInfo `json:"extra_info"`
}
