package act

type ActGoodsCouponListRequestDataItemDiscountCouponInfo struct {
	FullPrice      int `json:"full_price"`       //描述：券使用门槛金额，范围是1~999的正整数，单位元。当discount等于0时，为商品兑换券活动，不支持设置门槛金额，full_price默认为0。
	Discount       int `json:"discount"`         //描述：折扣系数，范围010之间，最多一位小数；或为0，表示无门槛商品兑换券活动。
	MaxReducePrice int `json:"max_reduce_price"` //描述：券最高可减金额，单位元。(1)当discount大于0时，为折扣商品券活动，优惠信息最多三层。(2)当discount等于0时，为商品兑换券活动，不支持设置最高可减金额，max_reduce_price默认为0。
	Stock          int `json:"stock"`            //描述：券库存，范围199999的正整数。(1)当选择单门店券时(is_single_poi=1)，该库存指的是每门店的券库存。(2) 当选择门店通用券时(is_single_poi=0)，该库存指的是所有门店共享的券库存。
	UserType       int `json:"user_type"`        //描述：可领券的用户，取值范围：null或0新客&老客；1新客；2老客。
}
