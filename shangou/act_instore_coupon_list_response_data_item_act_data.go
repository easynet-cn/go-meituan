package shangou

type ActInstoreCouponListResponseDataItemActData struct {
	use_count     int     `json:"use_count"`     //描述：劵已使用数量
	user_type     int     `json:"user_type"`     //描述：可领券的用户类型，参考值：1新客；2新客&老客；3老客。
	coupon_price  float64 `json:"coupon_price"`  //描述：优惠券金额，单位元，表示此优惠券可优惠的金额。
	validity_days int     `json:"validity_days"` //描述：券使用有效期，单位天。表示用户领用了此商家券后在领用日起多久内可有效使用。参考值为07，分别表示当天内有效、1日内有效、2日内有效、...、7日内有效。
	stock         int     `json:"stock"`         //描述：劵库存
	limit_price   float64 `json:"limit_price"`   //描述：使用门槛金额，单位元，表示订单金额满足此门槛即可享受本活动。
	sent_count    int     `json:"sent_count"`    //描述：劵已发数量
}
