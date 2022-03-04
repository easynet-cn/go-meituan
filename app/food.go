package app

type Food struct {
	AppFoodCode     string  `json:"app_food_code"`     //是否必须：是；描述：APP方商品id，即商家中台系统里商品的编码(spu_code值)，此app_food_code等同于app_medicine_code。(1)不同门店之间商品id可以重复，同一门店内商品id不允许重复。(2)字段信息限定长度不超过128个字符。
	AppMedicineCde  string  `json:"app_medicine_code"` //是否必须：是；描述：APP方商品id，即商家中台系统里商品的编码(spu_code值)(1)不同门店之间商品id可以重复，同一门店内商品id不允许重复。(2)字段信息限定长度不超过128个字符。此字段内容与app_food_code字段一致，平台建议商家对接此字段。
	FoodName        string  `json:"food_name"`         //是否必须：是；描述：退款商品名称
	SkuId           string  `json:"sku_id"`            //是否必须：是；描述：SKU码(商家的规格编码)，是商品sku唯一标识码，商品都是单规格。此字段信息与app_food_code相同，等同于药品app_medicine_code。
	Upc             string  `json:"upc"`               //是否必须：是；描述：商品的UPC码信息，即商品包装上的UPC/EAN码编号，长度一般8位或者13位。是商家同步商品信息时维护的UPC码，同一门店内，商品UPC码不允许重复。
	Spec            string  `json:"spec"`              //是否必须：是；描述：药品规格，展示美团标准库中此药品的规格信息
	FoodPrice       float64 `json:"food_price"`        //是否必须：是；描述：当前商品sku参加商品类活动优惠后的金额（单价），单位是元。
	Count           int     `json:"count"`             //是否必须：是；描述：本次退款的商品sku数量
	BoxNum          float64 `json:"box_num"`           //是否必须：是；描述：当前商品单件商品sku需使用包装盒的数量。药品类无商品维度打包盒，药品商家无需参考此字段信息。
	BoxPrice        float64 `json:"box_price"`         //是否必须：是；描述：当前商品sku需使用包装盒的单价，即单个包装盒的价格，单位是元。药品类无商品维度打包盒，药品商家无需参考此字段信息。
	OriginFoodPrice float64 `json:"origin_food_price"` //是否必须：是；描述：当前商品sku原价（单价），单位是元。
	RefundPrice     float64 `json:"refund_price"`      //是否必须：是；描述：当前商品单件可退金额（扣减商品所参加的所有营销活动优惠分摊后的金额），单位是元。
}
