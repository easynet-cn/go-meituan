package shangou

type ActNthDiscountSaveRequestActDataActProductsItem struct {
	AppFoodCode string `json:"app_food_code"` //是否必须：是；描述：1）字段描述：活动商品编码。 2）选择第N件X元时，所选商品原价必须大于活动价X。 3）仅支持单规格商品参加活动 4）商品原价乘活动件数必须大于活动价。 5）如填写了活动id，输入未在活动中存在的app_food_code，默认为在该活动中新增参加活动的商品。 6）如填写了活动id，输入已在活动中存在的app_food_code，则认为更新此活动中该商品库存信息。
	DayLimit    int    `json:"day_limit"`     //是否必须：是；描述：1）字段描述：活动商品当日库存。 2）格式：正整数，不大于999999999。 3）如当日库存为无限，则传-1。
}
