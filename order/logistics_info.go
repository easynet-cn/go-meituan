package order

type LogisticsInfo struct {
	ExpressCompany string `json:"expressCompany"` //描述：物流公司名称
	ExpressNumber  string `json:"expressNumber"`  //描述：物流订单号
	Reason         string `json:"reason"`         //描述：退货说明，用户寄回商品时填写的退货备注信息
}
