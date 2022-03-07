package poi

type PoiLogisticsIsDelayPushInfo struct {
	SupportDelayPush int `json:"support_delay_push"` //描述：门店是否在延迟发配送白名单内，参考值：1-该门店在延迟发配送名单内；0-该门店不在白名单内。
	DelayPushSecond  int `json:"delay_push_second"`  //描述：如门店在延迟发配送白名单内，此字段返回延迟发配送时间，单位秒。不在白名单中则返回0。
}
