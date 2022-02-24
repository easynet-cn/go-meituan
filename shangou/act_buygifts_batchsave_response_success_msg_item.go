package shangou

type ActBuygiftsBatchsaveResponseSuccessMsgItem struct {
	ItemId      int64  `json:"item_id"`       //描述：活动id，不同商品的买赠活动的活动id不同。
	AppFoodCode string `json:"app_food_code"` //描述：APP方商品id，即商家中台系统里商品的编码（spu_code值）。 注意：app_food_code即为药品的app_medicine_code。
	StartTime   int    `json:"start_time"`    //描述：活动开始时间，为10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的00:00:00。
	EndTime     int    `json:"end_time"`      //描述：活动结束时间，为10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的23:59:59。
}
