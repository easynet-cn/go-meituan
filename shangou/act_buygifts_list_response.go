package shangou

type ActBuygiftsListResponse struct {
	Data []ActBuygiftsListResponseDataItem `json:"data"` //描述：买赠活动信息集合，json格式数组。 注意：app_food_code即为药品的app_medicine_code。
}
