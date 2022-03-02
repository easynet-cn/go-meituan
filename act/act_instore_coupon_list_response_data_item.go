package act

type ActInstoreCouponListResponseDataItem struct {
	StartTime  int    `json:"start_time"`   //描述：活动开始时间，为10位秒级的时间戳。
	ActData    string `json:"act_data"`     //描述：活动数据，json格式数组。
	AppPoiCode string `json:"app_poi_code"` //描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	EndTime    int    `json:"end_time"`     //描述：活动结束时间，为10位秒级的时间戳。
	ActId      int64  `json:"actId"`        //描述：活动id
	RemainDays int    `json:"remain_days"`  //描述：活动剩余天数，活动结束日期与当前日期的差。
	Status     int    `json:"status"`       //描述：活动当前的状态，参考值：4进行中；3待生效；5已结束；6提报冻结。
}
