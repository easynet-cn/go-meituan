package act

type ActInstoreCouponBatchsaveRequestLimitTime struct {
	StartTime int `json:"start_time"` //是否必须：是；描述：start_time，活动开始时间，为10位秒级的时间戳；活动开始时间不得早于当前时间。
	EndTime   int `json:"end_time"`   //是否必须：是；描述：end_time，活动结束时间，为10位秒级的时间戳；活动结束时间不得早于当前时间和开始时间。
}
