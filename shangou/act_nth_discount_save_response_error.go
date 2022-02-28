package shangou

type ActNthDiscountSaveResponseError struct {
	Msg  string `json:"msg"`  //描述：失败说明
	Code int    `json:"code"` //描述：错误状态码
}
