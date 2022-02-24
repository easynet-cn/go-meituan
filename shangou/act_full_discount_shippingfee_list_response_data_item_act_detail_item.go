package shangou

type ActFullDiscountShippingfeeListResponseDataItemActDetailItem struct {
	LimitPrice    float64 `json:"limit_price"`    //描述：满减门槛，单位元，消费大于等于该金额时减配送费。(1)字段信息必须为大于0的数字，且不能超过2位小数。(2)不能比同一阶梯中的减配送费金额低。
	DiscountPrice float64 `json:"discount_price"` //描述：减配送费金额，单位元，享受活动时减免的配送费金额。(1)字段信息必须为大于0的数字，且不能超过2位小数。(2)不能比同一阶梯中的满减门槛金额高。(3)减免的金额必须随满减门槛增加而增加。
	PoiCharge     float64 `json:"poi_charge"`     //描述：商家承担活动费用，单位是元。
	MtCharge      float64 `json:"mt_charge"`      //描述：美团承担活动费用，单位是元。
}
