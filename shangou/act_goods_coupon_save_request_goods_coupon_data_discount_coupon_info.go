package shangou

type ActGoodsCouponSaveRequestGoodsCouponDataDiscountCouponInfo struct {
	FullPrice      int     `json:"full_price"`       //是否必须：是；描述：券使用门槛金额，范围是1~999的正整数，单位元。(1)当discount大于0时，此字段必传，折扣商品券活动的优惠信息不能超过三层。(2)当discount等于0时，为商品兑换券活动，不支持设置门槛金额，full_price默认为0。
	Discount       float64 `json:"discount"`         //是否必须：是；描述：折扣系数，需大于0且小于10，最多一位小数；或传0，表示兑换券。当discount=0时，则此活动为无门槛商品兑换券活动；此时app_food_codes字段上传的商品数不能大于1个，且该商品不可再参加其他优惠活动。
	MaxReducePrice *int    `json:"max_reduce_price"` //是否必须：否；描述：券最高可减金额，单位元。(1)当discount大于0时，此字段必传，折扣商品券活动的优惠信息不能超过三层。(2)当discount等于0时，为商品兑换券活动，不支持设置最高可减金额，max_reduce_price默认为0。
	Stock          int     `json:"stock"`            //是否必须：是；描述：券库存，范围199999的正整数。(1)当选择单门店券时(is_single_poi=1)，该库存指的是每门店的券库存。(2) 当选择门店通用券时(is_single_poi=0)，该库存指的是所有门店共享的券库存。
	UserType       *int    `json:"user_type"`        //是否必须：否；描述：可领券的用户，取值范围：0新客&老客；1新客；2老客。如不传此字段，则默认为0。
}
