package app

//推送门店状态变更
//请求方式：POST
//文档地址：https://open-shangou.meituan.com/home/docDetail/349
type AppPoiStatusRequest struct {
	AppPoiCode  string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，传商家中台系统里门店的编码，最长不超过128个字符。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	PoiStatus   int    `url:"poi_status"`   //是否必须：是；描述：门店当前状态，参考值：121-营业；120-休息；18-上线；19-下线。
	Reason      string `url:"reason"`       //是否必须：是；描述：操作原因
	OperateTime int64  `url:"operate_time"` //是否必须：是；描述：操作时间，传10位秒级的时间戳。
	OperateUser string `url:"operate_user"` //是否必须：是；描述：操作方，参考值：BUSINESS或者MEITUAN。
}
