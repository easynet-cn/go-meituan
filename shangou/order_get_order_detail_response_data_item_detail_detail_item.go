package shangou

type OrderGetOrderDetailResposneDataItemDetailDetailItem struct {
	ActualPrice     string  `json:"actual_price"`      //描述：商品现价，单位是元
	OriginalPrice   string  `json:"original_price"`    //描述：商品原价，单位是元
	AppFoodCode     string  `json:"app_food_code"`     //描述：APP方商品id，即商家中台系统里商品的编码(spu_code值)，此app_food_code等同于app_medicine_code。(1)不同门店之间商品id可以重复，同一门店内商品id不允许重复。(2)字段信息限定长度不超过128个字符。(3)如此字段信息推送的是商品名称或信息为空，则表示商家没有维护商品编码，请商家自行维护。
	AppMedicineCode string  `json:"app_medicine_code"` //描述：APP方商品id，即商家中台系统里商品的编码(spu_code值)，内容等同于app_food_code编码，平台建议对接此字段。
	FoodName        string  `json:"food_name"`         //描述：商品名称
	SkuId           string  `json:"sku_id"`            //描述：SKU码(商家的规格编码)，是商品sku唯一标识码，药品都属于单规格。此字段信息与app_food_code相同，等同于药品app_medicine_code。
	Upc             string  `json:"upc"`               //描述：商品的UPC码信息，即商品包装上的UPC/EAN码编号，长度一般8位或者13位。是商家同步商品信息时维护的UPC码，同一门店内，商品UPC码不允许重复。
	Quantity        int     `json:"quantity"`          //描述：订单中此商品sku的购买数量
	Price           float64 `json:"price"`             //描述：商品单价，单位是元。此字段信息默认返回活动折扣后价格，商家如有需求将价格替换为原价，可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“18 替换菜品折扣价格为原价”。
	BoxNum          float64 `json:"box_num"`           //描述：订单中当前商品sku需使用包装盒的总数量（单件商品sku需使用包装盒数量*商品售卖数量）。单件商品sku包装盒数量是商家同步商品时维护的信息。药品类无商品维度包装盒，药品商家无需参考此字段信息。
	BoxPrice        float64 `json:"box_price"`         //描述：包装盒单价，单位是元。为订单中当前商品sku单个包装盒的价格，是商家同步商品时维护的信息。药品类无商品维度包装盒，药品商家无需参考此字段信息。
	Unit            string  `json:"unit"`              //描述：商品售卖单位
	FoodDiscount    float64 `json:"food_discount"`     //描述：商品折扣系数，目前此字段默认为1，商家无需参考此字段的信息。
	FoodProperty    string  `json:"food_property"`     //描述：商品属性，多个属性用半角逗号隔开。该信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“23 菜品属性”。
	CartId          int     `json:"cart_id"`           //描述：商品所在的口袋，为用户在购物车中自行设置商品分开包装的口袋号。0为1号口袋，1为2号口袋，以此类推。此信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“26 商品口袋”。此字段适用于餐饮类订单，目前闪购品类订单此字段信息默认为0。
	Weight          int64   `json:"weight"`            //描述：商品重量：单位是克/g
	WarehouseInfo   string  `json:"warehouse_info"`    //描述：订单有分仓信息时，返回仓库编码与该仓库对应的商品数量。
}
