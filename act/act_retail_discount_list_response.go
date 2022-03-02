package act

type ActRetailDiscountListResponse struct {
	Data []ActRetailDiscountListResponseDataItem `json:"data"` //描述：折扣商品活动信息，json格式数组。 注意：app_food_code即为药品的app_medicine_code。
}
