package shangou

type ActRetailDiscountListResponseDataItem struct {
	AppFoodCode         string   `json:"app_food_code"`        //描述：活动商品，为APP方商品id，即商家中台系统里商品的编码（spu_code值）。 注意：app_food_code即为药品的app_medicine_code。
	UserType            int      `json:"user_type"`            //描述：允许参与活动的用户类型，参考值：0不限；1门店新客。 在爆品活动类型下，参考值有3种：0不限；1门店新用户；4闪购新客。
	StartTime           int      `json:"start_time"`           //描述：活动开始时间，为10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的00:00:00。
	EndTime             int      `json:"end_time"`             //描述：活动结束时间，为10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的23:59:59。
	OrderLimit          int      `json:"order_limit"`          //描述：每单限购，表示同一订单中可享折扣价的数量。(1)此字段信息为大于0的整数或1，为1则表示不限购。(2)如为大于0的整数，则最大值为10。
	DayLimit            int      `json:"day_limit"`            //描述：折扣商品当日活动库存，此字段信息为大于0的整数或1，如为1则表示不限制商品活动库存。
	Period              string   `json:"period"`               //描述：生效时间段：(1)同一个活动最多可以有3个生效时段，多个时段以英文逗号分隔。(2)生效时间最短为1小时。(3)此字段未填写则默认时段为"00:00:23:59"。
	WeeksTime           string   `json:"weeks_time"`           //描述：生效活动周期，1～7表示周一至周日，多个星期以英文逗号分隔。此字段未填写，则默认周期为"1,2,3,4,5,6,7"。
	SettingType         int      `json:"setting_type"`         //描述：活动开展类型，参考值：0按折扣系数开展活动；1按折扣价格开展活动。
	ActPrice            float64  `json:"act_price"`            //描述：折扣后价格，单位元。
	DiscountCoefficient *float64 `json:"discount_coefficient"` //描述：折扣系数，当setting_type=1时此字段信息为null。
	ItemId              int64    `json:"item_id"`              //描述：活动id
	OriginPrice         float64  `json:"origin_price"`         //描述：商品原价，单位元。
	Stock               int      `json:"stock"`                //描述：当日剩余活动商品数量。只有当发起查询时间处于活动生效时段内时(start_time、end_time、period、weeks_time均需满足)，此字段信息才代表当前实际剩余的活动商品数量，否则表示的是创建活动时设置的商品当日活动库存。
	Status              int      `json:"status"`               //描述：活动当前的状态，参考值：0已过期；1已生效；2待生效。
	Name                string   `json:"name"`                 //描述：商品名称
	Sequence            int      `json:"sequence"`             //描述：商品在折扣活动内的排序序号，值越小排位越靠前。如未填写此字段，则默认按创建顺序排序，创建时间越新的排位越靠前。
}
