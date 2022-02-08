package meituan

var (
	AppCancelOrderReason = map[int]string{
		2001: "APP方商家超时接单",
		2002: "APP方非顾客原因修改订单",
		2003: "APP方非顾客原因取消订单",
		2004: "APP方配送延迟",
		2005: "APP方售后投诉",
		2006: "APP方用户要求取消",
		2007: "APP方其他原因取消（未传code，默认为此原因）",
		2008: "店铺太忙",
		2009: "商品已售完",
		2010: "地址无法配送",
		2011: "店铺已打烊",
		2012: "联系不上用户",
		2013: "重复订单",
		2014: "配送员取餐慢",
		2015: "配送员送餐慢",
		2016: "配送员丢餐、少餐、餐洒",
	}
)
