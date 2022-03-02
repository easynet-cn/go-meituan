package shipping

type TimeMarkupFactor struct {
	TimeRange string  `json:"timeRange"` //是否必须：是；描述：加价时段， 1.精确到分钟，格式为HH:MMHH:MM，2.多个规则之间，不能有重叠时间，3.不允许跨天，开始时间须小于结束时间
	MarkupNum float64 `json:"markupNum"` //是否必须：是；描述：加价金额， 1.单位元 2.正数且最多支持1位小数 3.最大支持传10000
}
