package shangou

type ActFullDiscountShippingfeeListResponseDataItem struct {
	AppPoiCode string   `json:"app_poi_code"` //描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActId      int64    `json:"act_id"`       //描述：活动id
	StartTime  int      `json:"start_time"`   //描述：活动开始时间，传10位秒级的时间戳。系统已自动将传入的时间转换为所传日期的00:00:00。
	EndTime    int      `json:"end_time"`     //描述：活动结束时间，传10位秒级的时间戳。系统已自动将传入的时间转换为所传日期的23:59:59。
	WeeksTime  string   `json:"end_time"`     //描述：生效活动周期，1～7表示周一至周日，多个星期以英文逗号分隔。如创建活动时此字段未填写，则默认周期为"1,2,3,4,5,6,7"。
	Period     string   `json:"period"`       //描述：生效时间段，多个时段以英文逗号分隔；如创建活动时此字段未填写，则默认时段为"00:00:23:59"。
	ActDetail  string   `json:"act_detail"`   //描述：阶梯满减活动详情，json格式数组。
	MaxPrice   *float64 `json:"max_price"`    //描述：免配送费门槛，商品消费大于或等于该金额时免除全部配送费。若在创建时此字段未设置，则查询时值为null。
	ActStatus  string   `json:"actStatus"`    //描述：
}
