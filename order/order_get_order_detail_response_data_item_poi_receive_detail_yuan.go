package order

type OrderGetOrderDetailResposneDataItemPoiReceiveDetailYuan struct {
	ActOrderChargeByMt      []OrderGetOrderDetailResposneDataItemPoiReceiveDetailYuanActOrderChargeByMtItem  `json:"actOrderChargeByMt"`      //描述：美团承担明细
	ActOrderChargeByPoi     []OrderGetOrderDetailResposneDataItemPoiReceiveDetailYuanActOrderChargeByPoiItem `json:"actOrderChargeByPoi"`     //描述：商家承担明细
	FoodShareFeeChargeByPoi string                                                                           `json:"foodShareFeeChargeByPoi"` //描述：商品分成，即平台服务费，单位为元
	ReconciliationExtras    string                                                                           `json:"reconciliationExtras"`    //描述：商家对账信息相关的额外信息
	LogisticsFee            string                                                                           `json:"logisticsFee"`            //描述：配送费，单位为元
	OnlinePayment           string                                                                           `json:"onlinePayment"`           //描述：在线支付款，单位为元
	PoiReceive              string                                                                           `json:"poiReceive"`              //描述：商家应收款，单位为元
}
