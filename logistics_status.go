package meituan

var (
	LogisticsStatus = map[int]string{
		0:   "配送单发往配送",
		5:   "配送压单",
		10:  "配送单已确认",
		15:  "骑手已到店",
		20:  "骑手已取餐",
		40:  "骑手已送达",
		100: "配送单已取消",
	}
)
