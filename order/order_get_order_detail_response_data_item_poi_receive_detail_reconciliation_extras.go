package order

type OrderGetOrderDetailResposneDataItemPoiReceiveDetailReconciliationExtras struct {
	ChargeMode            int     `json:"chargeMode"`            //描述：订单服务费的费率模式，枚举：1-美配旧收费模式；2-美配新收费模式；6-闪购企客模式。
	PerformanceServiceFee float64 `json:"performanceServiceFee"` //描述：美配新收费模式&企客配送模式下的订单履约服务费金额，单位元。如订单无履约服务费，则返回0。
	TechnicalServiceFee   float64 `json:"technicalServiceFee"`   //描述：企客配送模式的“技术服务费”，单位为元；本参数仅在企客配送模式下推送。
	DistanceFee           float64 `json:"distanceFee"`           //描述：履约服务费里的“距离收费”，单位为元；本参数仅在企客配送模式下推送。
	SlaFee                float64 `json:"slaFee"`                //描述：履约服务费里的“时段增值收费”，单位为元；本参数仅在企客配送模式下推送。
	BaseShippingAmount    float64 `json:"baseShippingAmount"`    //描述：履约服务费里的“基础配送收费”，单位为元；本参数仅在企客配送模式下推送。
	CategoryChargeFee     float64 `json:"categoryChargeFee"`     //描述：履约服务费里的“品类收费”，单位为元；本参数仅在企客配送模式下推送。
	WeightChargeFee       float64 `json:"weightChargeFee"`       //描述：履约服务费里的“重量收费”，单位为元；本参数仅在企客配送模式下推送。
	HolidayChargeFee      float64 `json:"holidayChargeFee"`      //描述：履约服务费里的“节假日/大促收费"，单位为元；本参数仅在企客配送模式下推送。
}
