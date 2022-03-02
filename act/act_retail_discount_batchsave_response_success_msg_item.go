package act

type ActRetailDiscountBatchsaveResponseMsgSuccessItem struct {
	AppFoodCode string `json:"app_food_code"` //描述：活动商品，传APP方商品id，即商家中台系统里商品的编码（spu_code值）。 注意：app_food_code即为药品的app_medicine_code。
	EndTime     int    `json:"end_time"`      //描述：活动结束时间，传10位秒级的时间戳。系统已自动将传入的时间转换为所传日期的23:59:59。
	ActId       int64  `json:"act_id"`        //描述：活动id，不同的商品折扣活动对应的活动id不同。
	Period      string `json:"period"`        //描述：生效时间段，多个时段以英文逗号分隔；如创建活动时此字段未填写，则默认时段为"00:00:23:59"。
	StartTime   int    `json:"start_time"`    //描述：活动开始时间，传10位秒级的时间戳。系统已自动将传入的时间转换为所传日期的00:00:00。
	WeeksTime   string `json:"weeks_time"`    //描述：生效活动周期，1～7表示周一至周日，多个星期以英文逗号分隔。如创建活动时此字段未填写，则默认周期为"1,2,3,4,5,6,7"。
}
