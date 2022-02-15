package shangou

//批量创建/更新单店响应
type EcommercePoiSettleSingleSaveResponse struct {
	ResultCode  int    `json:"result_code"` //1全部操作成功，如需获取具体成功信息可以查看success_list字段；2部分成功，成功的数据存储在success_list字段，失败的数据存在error_list字段中；3全部操作失败，失败的数据存在error_list字段中；4请求失败，一般为签名或限流问题，关注error字段中的具体描述即可
	Data        string `json:"data"`
	SuccessList string `json:"success_list"` //成功返回值集合
	ErrorList   string `json:"error_list"`   //失败返回值集合
}
