package shangou

type ActInstoreCouponBatchsaveResponseSuccessMsgItem struct {
	ActId     int64 `json:"act_id"`     //描述：活动id
	StartTime int   `json:"start_time"` //描述：活动开始时间，为10位秒级的时间戳。
	EndTime   int   `json:"end_time"`   //描述：活动结束时间，为10位秒级的时间戳。
}
