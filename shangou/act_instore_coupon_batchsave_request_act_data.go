package shangou

type ActInstoreCouponBatchsaveRequestActData struct {
	LimitPrice   int `json:"limit_price"`   //是否必须：是；描述：limit_price(必填项)：使用门槛金额，单位元，须为1～999间的整数。表示订单金额满足此门槛即可享受本活动。
	CouponPrice  int `json:"coupon_price"`  //是否必须：是；描述：coupon_price(必填项)：优惠券金额，单位元，须为1～999间的整数。表示此优惠券可优惠的金额。券金额需大于使用门槛的3%，不可超过使用门槛的50%。
	UserType     int `json:"user_type"`     //是否必须：是；描述：user_type(必填项)：可领券的用户类型，取值范围：1新客；2新客&老客；3老客。 新客、老客可分别设置活动。
	ValidityDays int `json:"validity_days"` //是否必须：是；描述：validity_days(必填项)：券使用有效期，表示用户领用了此商家券后在领用日起多久内可有效使用。取值范围为07，分别表示当天内有效、1日内有效、2日内有效、...、7日内有效。
	Stock        int `json:"stock"`         //是否必须：是；描述：stock(必填项)，券的库存，须为1~99999间整数。
}
