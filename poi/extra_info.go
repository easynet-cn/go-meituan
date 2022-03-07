package poi

type ExtraInfo struct {
	TotalCount int `json:"total_count"` //描述：按app查询时返回app下已绑定的门店总数量；查单个门店时，本参数返回1。
}
