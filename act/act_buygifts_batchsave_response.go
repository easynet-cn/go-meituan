package act

type ActBuygiftsBatchsaveResponse struct {
	Data       string                                       `json:"data"`
	SuccessMsg []ActBuygiftsBatchsaveResponseSuccessMsgItem `json:"success_msg"` //描述：创建成功的买赠活动基本信息集合，json格式数组。 注意：app_food_code即为药品的app_medicine_code。
}
