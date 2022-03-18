package act

type ActGoodsCouponSaveRequestGoodsCouponData struct {
	AppPoiCodes         string                                                           `json:"app_poi_codes"`          //是否必须：是；描述：可使用券的门店id列表，多个用英文逗号分隔。传APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	CouponName          string                                                           `json:"coupon_name"`            //是否必须：是；描述：券名称，不能超过12个字符，仅支持汉字、数字和字。
	IsSinglePoi         int                                                              `json:"is_single_poi"`          //是否必须：是；描述：券使用范围是否限制单门店，取值范围：0门店通用券；1单门店券。(1)门店通用券：如app_poi_codes字段传的门店ABC，则A门店领，ABC门店通用。(2)单门店券：A门店领，仅A门店用。
	AppFoodCodes        string                                                           `json:"app_food_codes"`         //是否必须：是；描述：参加活动的商品，传APP方商品id，即商家中台系统里商品的编码（spu_code值）。(1)如不同门店内同一商品的app_food_code不同，则视为不同商品，多个商品id用英文逗号分隔。(2)当创建按折扣的商品券活动且discount=0时，本参数仅支持传一个参加活动的商品。 注意：app_food_code即为药品的app_medicine_code。
	TakeCouponStartTime int                                                              `json:"take_coupon_start_time"` //是否必须：是；描述：领券开始时间，传10位秒级的时间戳。(1)系统会自动将传入的时间转换为所传日期的00:00:00。(2)领券开始时间不得早于当前时间。
	TakeCouponEndTime   int                                                              `json:"take_coupon_end_time"`   //是否必须：是；描述：领券结束时间，传10位秒级的时间戳。(1)系统会自动将传入的时间转换为所传日期的23:59:59。(2)领券结束时间不得早于当前时间和领券开始时间。
	UseCouponStartTime  int                                                              `json:"use_coupon_start_time"`  //是否必须：是；描述：用券开始时间，传10位秒级的时间戳。(1)系统会自动将传入的时间转换为所传日期的00:00:00。(2)用券开始时间不得早于领券开始时间，且不得晚于领券结束时间。注：用券结束时间默认为领券结束时间。
	CouponLimitCount    int                                                              `json:"coupon_limit_count"`     //是否必须：是；描述：活动期间每个用户领券的最大数量限制，范围1~99999999的整数。
	Type                int                                                              `json:"type"`                   //是否必须：是；描述：商品券类型，取值范围：1按活动价格；2按折扣。
	ActPriceCouponInfo  []ActGoodsCouponSaveRequestGoodsCouponDataActPriceCouponInfoItem `json:"act_price_coupon_info"`  //是否必须：是；描述：按活动价格的商品券信息，json格式数组。按活动价格的商品券活动最多支持设置三层优惠。
	DiscountCouponInfo  []ActGoodsCouponSaveRequestGoodsCouponDataDiscountCouponInfo     `json:"discount_coupon_info"`   //是否必须：是；描述：按折扣的商品券信息，json格式数组。(1)当discount大于0时，折扣商品券活动的优惠信息不能超过三层。(2)当discount等于0时，为商品兑换券活动，仅支持设置一层。
}
