package order

type OrderApplyPartRefundRequestFoodDataItem struct {
	AppFoodCode string  `json:"app_food_code"` //描述：APP方商品id，即商家中台系统里商品的编码（spu_code值），此app_food_code等同于app_medicine_code。字段信息限定长度不超过128个字符。
	SkuId       string  `json:"sku_id"`        //描述：SKU码(商家的规格编码)，是商品sku唯一标识码，商品都是单规格，此字段可不传。此字段信息与app_food_code相同，等同于药品app_medicine_code。
	Count       float64 `json:"count"`         //描述：本次退款的商品数量
	ItemId      string  `json:"item_id"`       //描述：订单内商品行维度的商品标识id。
}
