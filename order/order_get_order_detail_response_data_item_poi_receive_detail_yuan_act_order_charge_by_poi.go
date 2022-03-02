package order

type OrderGetOrderDetailResposneDataItemPoiReceiveDetailYuanActOrderChargeByPoiItem struct {
	Comment     string `json:"comment"`     //描述：备注
	FeeTypeDesc string `json:"feeTypeDesc"` //描述：明细费用类型描述
	FeeTypeId   int    `json:"feeTypeId"`   //描述：明细费用类型ID
	Money       string `json:"money"`       //描述：明细金额，单位为元
}
