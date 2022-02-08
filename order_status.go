package meituan

var (
	OrderStatus = map[int]string{
		1: "用户已提交订单",
		2: "向商家推送订单",
		4: "商家已确认",
		8: "订单已完成",
		9: "订单已取消",
	}
)
