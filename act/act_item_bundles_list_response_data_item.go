package act

type ActItemBundlesListResposeDataItem struct {
	ActId     int64  `json:"act_id"`     //描述：活动id
	ActName   string `json:"act_name"`   //描述：活动名称
	StartTime int    `json:"start_time"` //描述：活动开始时间，为10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的00:00:00。
	EndTime   int    `json:"end_time"`   //描述：活动结束时间，为10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的23:59:59。
	ActRemark string `json:"act_remark"` //描述：活动文案
}
