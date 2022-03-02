package act

type ActMarkupRepurchaseListResponseDataItem struct {
	ActId           int64   `json:"act_id"`            //描述：活动id
	ActName         string  `json:"act_name"`          //描述：活动名称，不超过12个字符。
	StartTime       int     `json:"start_time"`        //描述：活动开始时间，传10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的00:00:00。
	EndTime         int     `json:"end_time"`          //描述：活动结束时间，传10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的23:59:59。
	ActPrice        float64 `json:"act_price"`         //描述：换购门槛，表示可参加本次换购活动的订单金额门店(不包含运费)，单位元。须为1～999999.99之间的数字。
	ActRemark       string  `json:"actRemark"`         //描述：加价换购活动描述
	OrderLimitTotal int     `json:"order_limit_total"` //描述：每单可换购的商品总数，为大于0的整数或1，如为1则表示不限制。
}
