package order

type OrderGetOrderDetailResposneDataItemPoiReceiveDetail struct {
	ActOrderChargeByMt              []OrderGetOrderDetailResposneDataItemPoiReceiveDetailActOrderChargeByMtItem  `json:"actOrderChargeByMt"`              //描述：美团承担优惠活动明细
	AgreementCommissionRebateAmount string                                                                       `json:"agreementCommissionRebateAmount"` //描述：无门槛订单返利字段
	ActOrderChargeByPoi             []OrderGetOrderDetailResposneDataItemPoiReceiveDetailActOrderChargeByPoiItem `json:"actOrderChargeByPoi"`             //描述：商家承担优惠活动明细
	FoodShareFeeChargeByPoi         int64                                                                        `json:"foodShareFeeChargeByPoi"`         //描述：商品分成，即平台服务费，单位是分。订单的平台服务费与协议保底佣金，取较大值；如对平台服务费计算方法或保底佣金有疑问，请商家咨询美团品牌经理。 注意： 美配新收费模式下，本参数是指技术服务费。 美配履约服务费请参考reconciliationExtras里的performanceServiceFee字段。
	ReconciliationExtras            string                                                                       `json:"reconciliationExtras"`            //描述：商家对账信息相关的额外信息
	LogisticsFee                    int64                                                                        `json:"logisticsFee"`                    //描述：门店配送费，单位是分。当前订单产生时该门店的配送费（商家自配送运费或美团配送运费），此字段数据为运费优惠前的原价。
	OnlinePayment                   int64                                                                        `json:"onlinePayment"`                   //描述：订单在线支付款，单位是分。此字段信息为用户实际支付的订单总金额。
	WmPoiReceiveCent                int64                                                                        `json:"wmPoiReceiveCent"`                //描述：商家应收款，单位是分。根据订单原始相关信息计算所得，计算公式：结算金额=商品原价合计+包装盒费合计+赠品原价合计+打包袋+配送费（仅限商家自配、美团跑腿）-商家活动支出-平台服务费-公益捐款。 注意：订单下单时的计算金额统计口径为商家协议配送方式下的商业费用，仅做营业参考数值，最终结算是按订单完成时的配送方式计算为准，切勿混淆。 关于商品分成/平台服务费、商家应收款计算公式，如疑问请商家咨询美团品牌经理或美团商服10105557。
}
