package act

type ActInstoreCouponBatchsaveResponse struct {
	Data       string                                            `json:"data"`
	SuccessMsg []ActInstoreCouponBatchsaveResponseSuccessMsgItem `json:"success_msg"` //描述：活动信息，json格式字符串。
}
