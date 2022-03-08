package poi

type EcommercePoiSettleCategoryListResponse struct {
	ResultCode int    `json:"result_code"` //描述：1-全部操作成功，查询到的数据在data字段中；2-部分成功，成功的数据存储在data字段中，失败的数据存在error_list字段中；3-全部操作失败，失败的数据存在error_list字段中；4-请求失败，一般为签名或限流问题，关注error字段中的具体描述即可
	Data       string `json:"data"`
}
