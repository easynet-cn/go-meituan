package act

type ActMarkupRrepurchaseBatchsaveRequestActDataAppFoodsItem struct {
	AppFoodCode *string  `json:"app_food_code"` //是否必须：否，创建时必填；描述：参加本次换购活动的商品，传APP方商品id，即商家中台系统里商品的编码（spu_code值）。 注意：app_food_code即为药品的app_medicine_code。
	ActPrice    *float64 `json:"act_price"`     //是否必须：否，创建时必填；描述：活动价(单价)，表示换购一件此商品需额外支付的价格，单位元；活动价不能大于商品原价。
	PrderLimit  *int     `json:"order_limit"`   //是否必须：否，创建时必填；描述：每单限购，表示达到换购门槛后，可以加价换购此商品的最大数量。须为大于0的整数或1，传1表示不限制。
	DayLimit    *int     `json:"day_limit"`     //是否必须：否，创建时必填；描述：商品当日活动库存，须为大于0的整数或1，如传1则表示不限制。
}
