package act

type ActGoodsCouponGoodslistRequestDataItemActPriceCouponInfo struct {
	FullPrice   int `json:"full_price"`   //描述：券使用门槛金额，范围是1~999的正整数，单位元。
	ReducePrice int `json:"reduce_price"` //描述：券金额，范围是1~999的正整数，单位元；表示达到使用门槛后可以减的金额。注意：券立减金额需随门槛增加而增加。
	Stock       int `json:"stock"`        //描述：券库存，范围199999的正整数。(1)当选择单门店券时(is_single_poi=1)，该库存指的是每门店的券库存。(2) 当选择门店通用券时(is_single_poi=0)，该库存指的是所有门店共享的券库存。
	UserType    int `json:"user_type"`    //描述：可领券的用户，参考值：null或0新客&老客；1新客；2老客。
}
