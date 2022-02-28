package shangou

type ActNthDiscountSaveResponse struct {
	Data  string `json:"data"`  //描述：创建活动成功，则为生成的活动id 创建活动失败，则为 ng
	Msg   string `json:"msg"`   //描述：成功提示信息
	Error string `json:"error"` //描述：失败信息
}
