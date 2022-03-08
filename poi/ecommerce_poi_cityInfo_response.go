package poi

type EcommercePoiCityInfoResponse struct {
	Data       []CityInfo `json:"data"`
	ResultCode int        `json:"result_code"` //描述：1-全部操作成功，如需获取具体成功信息可以查看success_list字段；2-部分成功，成功的数据存储在success_list字段，失败的数据存在error_list字段中；3-全部操作失败，失败的数据存在error_list字段中；4-请求失败，一般为签名或限流问题，关注error字段中的具体描述即可
}
