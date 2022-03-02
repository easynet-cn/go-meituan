package act

type ActMarkupRepurchaseFoodListResponseDataAppFoodsItem struct {
	AppFoodCode string  `json:"app_food_code"` //描述：参加本次换购活动的商品，为APP方商品id，即商家中台系统里商品的编码（spu_code值）。
	ActPrice    float64 `json:"act_price"`     //描述：活动价(单价)，表示换购一件此商品需额外支付的价格，单位元；活动价不能大于商品原价。
	OrderLimit  int     `json:"order_limit"`   //描述：每单限购，表示达到换购门槛后，可以加价换购此商品的最大数量。为大于0的整数或1，如为1表示不限制。
	DayLimit    int     `json:"day_limit"`     //描述：商品当日活动库存，为大于0的整数或1，如为1则表示不限制。
	PriceRemark string  `json:"price_remark"`  //描述：加价换购活动描述
}
