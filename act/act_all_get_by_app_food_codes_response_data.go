package act

type ActAllGetByAppFoodCodesResponseData struct {
	TotalCount int                                              `json:"total_count"` //描述：当前传入的商品所参加的活动的总量
	ActList    []ActAllGetByAppFoodCodesResponseDataActListItem `json:"act_list"`    //描述：json数组
}
