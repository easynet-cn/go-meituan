package order

type OrderGetOrderDetailResposneDataItemPoiReceiveDetailYuanReconciliationExtras struct {
	ChargeMode            int     `json:"chargeMode"`            //描述：费率模式 1:旧收费模式 2:新收费模式 6 闪购企客模式
	PerformanceServiceFee float64 `json:"performanceServiceFee"` //描述：订单履约服务费金额，单位元。如订单无履约服务费，则返回0。
	TechnicalServiceFee   float64 `json:"technicalServiceFee"`   //描述：技术服务费，单位为元，仅当费率模式为6时存在
	DistanceFee           float64 `json:"distanceFee"`           //描述：距离收费，单位为，仅当费率模式为6时存在
	SlaFee                float64 `json:"slaFee"`                //描述：时段增值收费，单位为元，仅当费率模式为6时存在
	BaseShippingAmount    float64 `json:"baseShippingAmount"`    //描述：基础配送费，单位为元，仅当费率模式为6时存在
	CategoryChargeFee     float64 `json:"categoryChargeFee"`     //描述：品类加价，单位为元，仅当费率模式为6时存在
	WeightChargeFee       float64 `json:"weightChargeFee"`       //描述：重量收费，单位为元，仅当费率模式为6时存在
	HolidayChargeFee      float64 `json:"holidayChargeFee"`      //描述：特殊日期收费，单位为元，仅当费率模式为6时存在
}
