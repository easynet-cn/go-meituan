package app

type PoiReceiveDetailYuan struct {
	ActOrderChargeByMt              []ActOrderCharge `json:"actOrderChargeByMt"`              //是否必须：是；描述：美团承担明细
	ActOrderChargeByPoi             []ActOrderCharge `json:"actOrderChargeByPoi"`             //是否必须：是；描述：商家承担明细
	FoodShareFeeChargeByPoi         string           `json:"foodShareFeeChargeByPoi"`         //是否必须：是；描述：商品分成，即平台服务费，单位为元
	LogisticsFee                    string           `json:"logisticsFee"`                    //是否必须：是；描述：配送费，单位为元
	OnlinePayment                   string           `json:"onlinePayment"`                   //是否必须：是；描述：在线支付款，单位为元
	PoiReceive                      string           `json:"poiReceive"`                      //是否必须：是；描述：商家应收款，单位为元
	AgreementCommissionRebateAmount string           `json:"agreementCommissionRebateAmount"` //是否必须：是；描述：无门槛订单返利字段
	ReconciliationExtras            string           `json:"reconciliationExtras"`            //是否必须：是；描述：商家对账信息相关的额外信息
}
