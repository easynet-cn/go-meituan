package act

type ActItemBundlesSaveRequestActDataAppFoodsItem struct {
	AppFoodCode string `json:"app_food_code"` //是否必须：是；描述：活动商品，传APP方商品id，即商家中台系统里商品的编码（spu_code值）。 注意：app_food_code即为药品的app_medicine_code。
	DayLimit    int    `json:"day_limit"`     //是否必须：是；描述：商品当日活动库存，须传大于0的整数或1，如传1则表示不限制商品活动库存。
}
