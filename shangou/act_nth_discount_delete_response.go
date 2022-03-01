package shangou

type ActNthDiscountDeleteResponse struct {
	Data  string `json:"data"`  //描述：创建活动成功，则为生成的活动id 创建活动失败，则为 ng
	Error string `json:"error"` //描述：失败信息
}
