package shangou

type OrderGetOrderDetailResposneDataItemPoiReceiveDetailActOrderChargeByMtItem struct {
	Comment     string `json:"comment"`     //描述：活动内容备注，如无优惠活动则此字段信息默认为“活动款”。
	FeeTypeDesc string `json:"feeTypeDesc"` //描述：明细费用类型描述，目前默认为“活动款”。
	FeeTypeId   int    `json:"feeTypeId"`   //描述：明细费用类型ID，目前默认为“10019”。
	MoneyCent   int64  `json:"moneyCent"`   //描述：本活动优惠的总金额，单位是分。
}
