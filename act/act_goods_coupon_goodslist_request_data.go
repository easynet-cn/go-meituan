package act

type ActGoodsCouponGoodslistRequestData struct {
	Data []string `json:"data"` //描述：商品券活动信息，json格式数组。 注意：app_food_code即为药品的app_medicine_code。
}
